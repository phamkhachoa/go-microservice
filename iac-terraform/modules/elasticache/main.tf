provider "aws" {
  region = var.regison
}

# Security group for ElastiCache Redis
resource "aws_security_group" "redis_sg" {
  name        = "redis-security-group"
  description = "Allow Redis traffic"
  vpc_id      = var.vpc_id

  # Redis port
  ingress {
    from_port   = 6379
    to_port     = 6379
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
    description = "Allow Redis traffic from VPC"
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
    description = "Allow all outbound traffic"
  }

  tags = {
    Name = "redis-sg"
  }
}

# Create ElastiCache subnet group with the three existing subnets
resource "aws_elasticache_subnet_group" "redis_subnet_group" {
  name       = "redis-subnet-group"
  subnet_ids = var.subnet_ids # This should be a list of your 3 subnet IDs

  tags = {
    Name = "Redis Subnet Group"
  }
}

# Create Redis parameter group
resource "aws_elasticache_parameter_group" "redis_parameter_group" {
  name   = "redis-params"
  family = "redis6.x"

  # Optional Redis parameters
  parameter {
    name  = "maxmemory-policy"
    value = "volatile-lru"
  }

  parameter {
    name  = "notify-keyspace-events"
    value = "Ex"
  }

  tags = {
    Name = "Redis Parameters"
  }
}

# Create ElastiCache Redis instance (free tier)
resource "aws_elasticache_cluster" "redis" {
  cluster_id           = var.cluster_id
  engine               = "redis"
  node_type            = "cache.t2.micro" # Free tier eligible
  num_cache_nodes      = 1                # Free tier allows only 1 node
  parameter_group_name = aws_elasticache_parameter_group.redis_parameter_group.name
  subnet_group_name    = aws_elasticache_subnet_group.redis_subnet_group.name
  security_group_ids   = [aws_security_group.redis_sg.id]

  engine_version       = "6.2"
  port                 = 6379

  # Free tier settings
  snapshot_retention_limit = 0 # No snapshots for free tier
  maintenance_window       = "sun:05:00-sun:06:00"

  tags = {
    Name = "redis-free-tier"
  }
}

# Output the Redis endpoint
output "redis_endpoint" {
  value = "${aws_elasticache_cluster.redis.cache_nodes.0.address}:${aws_elasticache_cluster.redis.cache_nodes.0.port}"
  description = "The endpoint of the Redis instance"
}

output "redis_security_group_id" {
  value = aws_security_group.redis_sg.id
  description = "The ID of the Redis security group"
}