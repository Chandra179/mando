resource "aws_ecr_repository" "app" {
  # Name should match what you use in docker build/tag/push
  name                 = var.app_name
  image_tag_mutability = "MUTABLE" # Allow tag overwrites (e.g., 'latest') for dev

  image_scanning_configuration {
    scan_on_push = false # Typically false for LocalStack
  }

  tags = {
    Name        = "${local.base_name}-app-repo"
    Environment = var.environment
  }
}

# Output the repository URL (useful for build scripts/CI)
output "ecr_repository_url" {
  description = "URL of the ECR repository for the application"
  value       = aws_ecr_repository.app.repository_url
}