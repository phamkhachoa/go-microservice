variable "regison" {
  description = "AWS region to deploy resources"
  type        = string
}

variable "vpc_id" {
  description = "ID of the existing VPC"
  type        = string
}

variable "vpc_cidr" {
  description = "Subnet IDs"
  type        = list(string)
}

variable "subnet_ids" {
  description = "List of subnet IDs for the Redis cluster"
  type        = list(string)
}

variable "cluster_id" {
  description = "ID for the Redis cluster"
  type        = string
  default     = "redis-free-tier"
  validation {
    condition     = can(regex("^[a-z0-9\\-]+$", var.cluster_id))
    error_message = "The cluster ID must be lowercase, contain only alphanumeric characters and hyphens."
  }
}