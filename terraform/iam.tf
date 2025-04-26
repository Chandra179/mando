resource "aws_iam_role" "ecs_task_execution_role" {
  name = "${local.base_name}-ecs-task-execution-role"

  # Trust policy allowing ECS tasks to assume this role
  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "ecs-tasks.amazonaws.com"
        }
      }
    ]
  })

  tags = {
    Name = "${local.base_name}-ecs-task-execution-role"
  }
}

# Attach the AWS managed policy granting necessary permissions
resource "aws_iam_role_policy_attachment" "ecs_task_execution_role_policy" {
  role       = aws_iam_role.ecs_task_execution_role.name
  # This policy allows pulling ECR images, writing CloudWatch logs, etc.
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"
}

# Optional: Define a Task Role if your *application containers* need to call AWS APIs
# resource "aws_iam_role" "ecs_task_role" { ... }
# resource "aws_iam_role_policy_attachment" "ecs_task_role_policy" { ... }