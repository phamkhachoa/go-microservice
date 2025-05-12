
output "cluster_id" {
  description = "The ID of the RDS cluster"
  value       = aws_rds_cluster.mysql.id
}

output "cluster_endpoint" {
  description = "The cluster endpoint"
  value       = aws_rds_cluster.mysql.endpoint
}

output "reader_endpoint" {
  description = "The cluster reader endpoint"
  value       = aws_rds_cluster.mysql.reader_endpoint
}

output "database_name" {
  description = "The database name"
  value       = aws_rds_cluster.mysql.database_name
}

output "port" {
  description = "The database port"
  value       = aws_rds_cluster.mysql.port
}

output "security_group_id" {
  description = "The ID of the security group created for the RDS cluster"
  value       = aws_security_group.mysql.id
}

output "subnet_group_name" {
  description = "The name of the subnet group"
  value       = aws_db_subnet_group.mysql.name
}

output "instance_ids" {
  description = "List of instance IDs in the cluster"
  value       = aws_rds_cluster_instance.mysql[*].id
}
