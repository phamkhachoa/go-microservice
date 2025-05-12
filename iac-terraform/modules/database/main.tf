provider "aws" {
  region = var.region
}

resource "aws_db_subnet_group" "mysql" {
  name        = "${var.environment}-${var.name}-subnet-group"
  description = "DB subnet group for ${var.name}"
  subnet_ids  = var.subnet_ids

  tags = {
    Name        = "${var.environment}-${var.name}-subnet-group"
    Environment = var.environment
  }
}

resource "aws_security_group" "mysql" {
  name        = "${var.environment}-${var.name}-sg"
  description = "Security group for ${var.name} RDS cluster"
  vpc_id      = var.vpc_id

  ingress {
    from_port   = 3306
    to_port     = 3306
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name        = "${var.environment}-${var.name}-sg"
    Environment = var.environment
  }
}

resource "aws_rds_cluster" "mysql" {
  cluster_identifier        = "${var.environment}-${var.name}"
  engine                    = "aurora-mysql"
  engine_version            = var.engine_version
  database_name             = var.database_name
  master_username           = var.master_username
  master_password           = var.master_password
  backup_retention_period   = var.backup_retention_period
  preferred_backup_window   = var.preferred_backup_window
  skip_final_snapshot       = var.skip_final_snapshot
  final_snapshot_identifier = var.skip_final_snapshot ? null : "${var.environment}-${var.name}-final-snapshot"
  db_subnet_group_name      = aws_db_subnet_group.mysql.name
  vpc_security_group_ids    = [aws_security_group.mysql.id]

  tags = {
    Name        = "${var.environment}-${var.name}"
    Environment = var.environment
  }
}

resource "aws_rds_cluster_instance" "mysql" {
  count                = var.instance_count
  identifier           = "${var.environment}-${var.name}-${count.index}"
  cluster_identifier   = aws_rds_cluster.mysql.id
  instance_class       = var.instance_class
  engine               = "aurora-mysql"
  engine_version       = var.engine_version
  db_subnet_group_name = aws_db_subnet_group.mysql.name

  tags = {
    Name        = "${var.environment}-${var.name}-${count.index}"
    Environment = var.environment
  }
}
