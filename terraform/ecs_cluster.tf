 resource "aws_ecs_cluster" "main" {
  name = "${local.base_name}-cluster"

  # Optional: Default capacity provider strategy if mixing Fargate/EC2
  # configuration {
  #   execute_command_configuration {
  #     logging = "DEFAULT"
  #   }
  # }

  tags = {
    Name = "${local.base_name}-cluster"
  }
}

# Log group for all container logs from this cluster/app
resource "aws_cloudwatch_log_group" "ecs_logs" {
  name = "/ecs/${local.base_name}"
  # Retention is good practice, though LocalStack might not enforce it strictly
  retention_in_days = 7

  tags = {
    Name        = "${local.base_name}-ecs-logs"
    Environment = var.environment
  }
}

output "ecs_cluster_name" {
  description = "Name of the ECS Cluster"
  value       = aws_ecs_cluster.main.name
}

output "cloudwatch_log_group_name" {
  description = "Name of the CloudWatch Log Group for ECS tasks"
  value       = aws_cloudwatch_log_group.ecs_logs.name
}