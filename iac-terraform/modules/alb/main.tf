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

data "tls_certificate" "my_cluster_issuer" {
  url = var.eks_cluster.identity[0].oidc[0].issuer
}

# Create an IAM OIDC Identity Provider that trusts the EKS cluster's issuer URL
resource "aws_iam_openid_connect_provider" "eks" {
  url =var.eks_cluster.identity[0].oidc[0].issuer

  client_id_list = [
    "sts.amazonaws.com",
  ]

  thumbprint_list = [
    data.tls_certificate.my_cluster_issuer.certificates[0].sha1_fingerprint,
  ]
}

#it allows IAM roles to trust and authenticate using the OpenID Connect (OIDC) protocol
data "aws_iam_policy_document" "aws_load_balancer_controller_assume_role_policy" {
  statement {
    actions = ["sts:AssumeRoleWithWebIdentity"]
    effect  = "Allow"

    condition {
      test     = "StringEquals"
      variable = "${replace(aws_iam_openid_connect_provider.eks.url, "https://", "")}:sub"
      values   = ["system:serviceaccount:kube-system:aws-load-balancer-controller"]
    }

    principals {
      identifiers = [aws_iam_openid_connect_provider.eks.arn]
      type        = "Federated"
    }
  }
}

# Create IAM role for the ALB controller
resource "aws_iam_role" "aws-load-balancer-controller" {
  assume_role_policy = data.aws_iam_policy_document.aws_load_balancer_controller_assume_role_policy.json
  name               = "aws-load-balancer-controller"
}

# Create IAM policy for ALB controller
resource "aws_iam_policy" "aws-load-balancer-controller" {
  name   = "AWSLoadBalancerController"
  # The policy document is in a separate file in the root module
  policy = file("${path.root}/iam_policy.json")
}

# Attach ALB controller policy to the role
resource "aws_iam_role_policy_attachment" "aws_load_balancer_controller_attach" {
  role       = aws_iam_role.aws-load-balancer-controller.name
  policy_arn = aws_iam_policy.aws-load-balancer-controller.arn
}

# Install the AWS Load Balancer Controller using Helm
resource "helm_release" "aws-load-balancer-controller" {
  name             = "aws-load-balancer-controller"
  repository       = "https://aws.github.io/eks-charts"
  chart            = "aws-load-balancer-controller"
  namespace        = "kube-system"
  version          = "1.4.1"
  create_namespace = true

  values = [
    file("./ingress.yaml")
  ]
  set {
    name  = "serviceAccount.annotations.eks\\.amazonaws\\.com/role-arn"
    value = aws_iam_role.aws-load-balancer-controller.arn
  }
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