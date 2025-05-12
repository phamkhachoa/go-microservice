# modules/alb/main.tf

# Provider configuration
provider "aws" {
  region = var.region
  alias  = "alb"
}

provider "helm" {
  kubernetes {
    host                   = var.cluster_endpoint
    cluster_ca_certificate = base64decode(var.cluster_certificate_authority_data)
    exec {
      api_version = "client.authentication.k8s.io/v1beta1"
      args        = ["eks", "get-token", "--cluster-name", var.cluster_name]
      command     = "aws"
    }
  }
}

data "aws_caller_identity" "current" {}

# Get OIDC provider URL without https:// prefix for IAM role
locals {
  oidc_provider = replace(var.oidc_provider_url, "https://", "")
}

# Create IAM policy for ALB controller
resource "aws_iam_policy" "alb_controller_policy" {
  provider    = aws.alb
  name        = "AWSLoadBalancerControllerIAMPolicy-${var.cluster_name}"
  description = "Policy for ALB Ingress Controller for cluster ${var.cluster_name}"

  # The policy document is in a separate file in the root module
  policy = file("${path.root}/iam_policy.json")
}

# Create IAM role for the ALB controller
resource "aws_iam_role" "lb_controller" {
  provider = aws.alb
  name     = "eks-alb-ingress-controller-${var.cluster_name}"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Principal = {
          Federated = "arn:aws:iam::${data.aws_caller_identity.current.account_id}:oidc-provider/${local.oidc_provider}"
        }
        Action = "sts:AssumeRoleWithWebIdentity"
        Condition = {
          StringEquals = {
            "${local.oidc_provider}:sub" = "system:serviceaccount:kube-system:aws-load-balancer-controller"
          }
        }
      }
    ]
  })
}

# Attach ALB controller policy to the role
resource "aws_iam_role_policy_attachment" "lb_controller" {
  provider   = aws.alb
  policy_arn = aws_iam_policy.alb_controller_policy.arn
  role       = aws_iam_role.lb_controller.name
}

# Install the AWS Load Balancer Controller using Helm
resource "helm_release" "lb_controller" {
  name       = "aws-load-balancer-controller"
  repository = "https://aws.github.io/eks-charts"
  chart      = "aws-load-balancer-controller"
  namespace  = "kube-system"
  version    = "1.6.0" # Check for the latest version

  set {
    name  = "clusterName"
    value = var.cluster_name
  }

  set {
    name  = "serviceAccount.create"
    value = "true"
  }

  set {
    name  = "serviceAccount.name"
    value = "aws-load-balancer-controller"
  }

  set {
    name  = "serviceAccount.annotations.eks\\.amazonaws\\.com/role-arn"
    value = aws_iam_role.lb_controller.arn
  }

  # Wait for the load balancer controller to be fully deployed
  wait = true
}

# Create security group for the ALB
resource "aws_security_group" "alb_sg" {
  provider    = aws.alb
  name        = "alb-security-group-${var.cluster_name}"
  description = "Security group for ALB in cluster ${var.cluster_name}"
  vpc_id      = var.vpc_id

  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
    description = "Allow HTTP traffic"
  }

  ingress {
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
    description = "Allow HTTPS traffic"
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
    description = "Allow all outbound traffic"
  }

  tags = {
    Name = "alb-security-group-${var.cluster_name}"
  }
}