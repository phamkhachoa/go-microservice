# modules/alb/outputs.tf

output "alb_security_group_id" {
  description = "Security Group ID for ALB"
  value       = aws_security_group.alb_sg.id
}

output "alb_controller_role_arn" {
  description = "IAM Role ARN for ALB Controller"
  value       = aws_iam_role.lb_controller.arn
}

output "alb_controller_policy_arn" {
  description = "IAM Policy ARN for ALB Controller"
  value       = aws_iam_policy.alb_controller_policy.arn
}