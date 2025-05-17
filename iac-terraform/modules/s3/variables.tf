variable "region" {
  description = "AWS region to deploy resources"
  type        = string
}

variable "bucket_name" {
  description = "Name of the S3 bucket to create for website hosting"
  type        = string
}

variable "environment" {
  description = "Environment name (e.g., dev, staging, prod)"
  type        = string
  default     = "dev"
}

variable "domain_name" {
  description = "Domain name for the website (optional)"
  type        = string
  default     = ""
}

variable "acm_certificate_arn" {
  description = "ARN of ACM certificate for the domain (optional)"
  type        = string
  default     = ""
}

variable "alb_domain_name" {
  description = "Domain name for the website (optional)"
  type        = string
  default     = "k8s-default-gomicros-863a68ddbf-776386155.ap-northeast-1.elb.amazonaws.com"
}