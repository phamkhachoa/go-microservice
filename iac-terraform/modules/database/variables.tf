variable "region" {
  description = "AWS region to deploy resources"
  type        = string
}

variable "vpc_id" {
  description = "CIDR block for the VPC"
  type        = string
}

variable "db_name" {
  description = "Name of the database"
  type        = string
  default     = "mydb"
}

variable "subnet_ids" {
  description = "Subnet IDs"
  type        = list(string)
}

variable "publicly_accessible" {
  description = "Whether the DB should be publicly accessible"
  type        = bool
  default     = true
}

variable "db_allocated_storage" {
  description = "Allocated storage for the database in GB"
  type        = number
  default     = 20
}

variable "db_instance_class" {
  description = "Instance class for the database"
  type        = string
  default     = "db.t3.micro"
}

variable "db_engine_version" {
  description = "MySQL engine version"
  type        = string
  default     = "8.0"
}

variable "backup_retention_period" {
  description = "Number of days to retain backups"
  type        = number
  default     = 7
}