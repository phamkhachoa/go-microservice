output "cluster_endpoint" {
  description = "EKS cluster endpoint"
  value       = aws_eks_cluster.main.endpoint
}

output "cluster_name" {
  description = "EKS cluster name"
  value       = aws_eks_cluster.main.name
}


output "cluster_ca_certificate" {
  description = "EKS cluster name"
  value       = aws_eks_cluster.main.certificate_authority
}

# output "oidc_provider_arn" {
#   description = "EKS oidc_provider_arn"
#   value       = aws_eks_cluster.main.oidc_provider_arn
# }

output "eks_cluster" {
  description = "EKS cluster name"
  value       = aws_eks_cluster.main
}

# Output the OIDC Provider ARN for reference
output "oidc_provider_arn" {
  description = "The ARN of the IAM OIDC provider for the EKS cluster."
  value       = aws_iam_openid_connect_provider.my_cluster_oidc_provider.arn
}

output "cluster_certificate_authority_data" {
  description = "Base64 encoded certificate data required to communicate with the cluster"
  value       = aws_eks_cluster.main.certificate_authority[0].data
}

output "oidc_provider_url" {
  description = "URL of the OIDC Provider from the EKS cluster"
  value       = aws_eks_cluster.main.identity[0].oidc[0].issuer
}

# output "kubeconfig" {
#   description = "kubectl config that can be used to connect to the cluster"
#   value       = local.kubeconfig
#   sensitive   = true
# }

output "cluster_security_group_id" {
  description = "Security group ID attached to the EKS cluster"
  value       = aws_eks_cluster.main.vpc_config[0].cluster_security_group_id
}
