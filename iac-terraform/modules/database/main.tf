provider "aws" {
  region = var.region
}

# resource "aws_db_subnet_group" "mysql" {
#   name        = "${var.environment}-${var.name}-subnet-group"
#   description = "DB subnet group for ${var.name}"
#   subnet_ids  = var.subnet_ids
#
#   tags = {
#     Name        = "${var.environment}-${var.name}-subnet-group"
#     Environment = var.environment
#   }
# }

# Subnet group for RDS
resource "aws_db_subnet_group" "rds_subnet_group" {
  name       = "rds-subnet-group"
  subnet_ids = var.subnet_ids
  tags = {
    Name = "RDS Subnet Group"
  }
}

# Security group for RDS
resource "aws_security_group" "rds_sg" {
  name        = "rds-security-group"
  description = "Allow MySQL traffic"
  vpc_id      = var.vpc_id

  # MySQL port
  ingress {
    from_port   = 3306
    to_port     = 3306
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"] # In production, restrict this to specific IPs
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "rds-sg"
  }
}

# resource "aws_security_group" "mysql" {
#   name        = "${var.environment}-${var.name}-sg"
#   description = "Security group for ${var.name} RDS cluster"
#   vpc_id      = var.vpc_id
#
#   ingress {
#     from_port   = 3306
#     to_port     = 3306
#     protocol    = "tcp"
#     cidr_blocks = ["0.0.0.0/0"]
#   }
#
#   egress {
#     from_port   = 0
#     to_port     = 0
#     protocol    = "-1"
#     cidr_blocks = ["0.0.0.0/0"]
#   }
#
#   tags = {
#     Name        = "${var.environment}-${var.name}-sg"
#     Environment = var.environment
#   }
# }

# resource "aws_rds_cluster" "mysql" {
#   cluster_identifier        = "${var.environment}-${var.name}"
#   engine                    = "aurora-mysql"
#   engine_version            = var.engine_version
#   database_name             = var.database_name
#   master_username           = var.master_username
#   master_password           = var.master_password
#   backup_retention_period   = var.backup_retention_period
#   preferred_backup_window   = var.preferred_backup_window
#   skip_final_snapshot       = var.skip_final_snapshot
#   final_snapshot_identifier = var.skip_final_snapshot ? null : "${var.environment}-${var.name}-final-snapshot"
#   db_subnet_group_name      = aws_db_subnet_group.mysql.name
#   vpc_security_group_ids    = [aws_security_group.mysql.id]
#
#   tags = {
#     Name        = "${var.environment}-${var.name}"
#     Environment = var.environment
#   }
# }

# RDS MySQL instance - Free Tier eligible
resource "aws_db_instance" "mysql" {
  identifier             = "mysql-instance"
  allocated_storage      = 20                           # Free tier offers 20GB
  storage_type           = "gp2"
  engine                 = "mysql"
  engine_version         = "8.0"                        # Choose your version
  instance_class         = "db.t3.micro"                # Free tier eligible
  db_name                = "mydb"
  username               = "admin"
  password               = "root1234"
  parameter_group_name   = "default.mysql8.0"
  skip_final_snapshot    = true
  publicly_accessible    = true                         # Set to false for production
  vpc_security_group_ids = [aws_security_group.rds_sg.id]
  db_subnet_group_name   = aws_db_subnet_group.rds_subnet_group.name
  multi_az               = false                        # Free tier is single-AZ only
  storage_encrypted      = false                        # Free tier doesn't support encryption

  # Performance Insights is not Free Tier eligible
  performance_insights_enabled = false

  # Automated backups
  backup_retention_period = 7
  backup_window           = "03:00-04:00"
  maintenance_window      = "Mon:04:00-Mon:05:00"

  tags = {
    Name = "MySQL Free Tier"
  }
}
