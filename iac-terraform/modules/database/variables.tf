variable "region" {
  description = "AWS region to deploy the RDS cluster"
  type        = string
}

variable "environment" {
  description = "Environment name (e.g. dev, staging, prod)"
  type        = string
}

variable "name" {
  description = "Name prefix for the RDS resources"
  type        = string
  default     = "mysql-cluster"
}

variable "vpc_id" {
  description = "VPC ID where the RDS cluster will be deployed"
  type        = string
}

variable "subnet_ids" {
  description = "List of subnet IDs for the DB subnet group"
  type        = list(string)
}

variable "allowed_cidr_blocks" {
  description = "List of CIDR blocks allowed to connect to the RDS cluster"
  type        = list(string)
  default     = ["10.0.0.0/16", "0.0.0.0/0"]
}

variable "database_name" {
  description = "Name of the database to create"
  type        = string
  default     = "mydb"
}

variable "master_username" {
  description = "Master username for the RDS cluster"
  type        = string
  default     = "admin"
}

variable "master_password" {
  description = "Master password for the RDS cluster"
  type        = string
  sensitive   = true
  default     = "root1234"
}

variable "engine_version" {
  description = "MySQL engine version"
  type        = string
  default     = "5.7.mysql_aurora.2.11.1"
}

variable "instance_class" {
  description = "Instance class for RDS cluster instances"
  type        = string
  default     = "db.t3.micro"
}

variable "instance_count" {
  description = "Number of instances in the RDS cluster"
  type        = number
  default     = 1
}

variable "backup_retention_period" {
  description = "Backup retention period in days"
  type        = number
  default     = 7
}

variable "preferred_backup_window" {
  description = "Preferred backup window"
  type        = string
  default     = "03:00-05:00"
}

variable "skip_final_snapshot" {
  description = "Whether to skip the final snapshot when deleting the cluster"
  type        = bool
  default     = true
}
