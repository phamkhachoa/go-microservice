# Output the database endpoint and password
output "rds_endpoint" {
  value = aws_db_instance.mysql.endpoint
}