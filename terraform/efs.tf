resource "aws_efs_file_system" "mongo_data" {
  creation_token = "${local.base_name}-mongo-data"
  # LocalStack Note: Encryption in transit might not be fully supported/necessary
  # encrypted = true
  # kms_key_id = aws_kms_key.efs.arn # Requires defining a KMS key

  tags = {
    Name = "${local.base_name}-mongo-data-efs"
  }
}

# Mount target allows the EFS to be accessed from the VPC subnet
resource "aws_efs_mount_target" "mongo_data" {
  file_system_id  = aws_efs_file_system.mongo_data.id
  subnet_id       = aws_subnet.public.id
  security_groups = [aws_security_group.ecs_tasks.id] # Allow access from ECS tasks
}

# Best Practice: Use an Access Point for controlled access within EFS
resource "aws_efs_access_point" "mongo_data" {
  file_system_id = aws_efs_file_system.mongo_data.id

  # Define a directory within EFS for this specific volume mount
  root_directory {
    path = "/data" # This path is *within* the EFS volume
    creation_info {
      # Match permissions/ownership needed by the MongoDB container user (often non-root)
      # Check MongoDB image docs; often runs as uid/gid 999 or similar. Use '1000' as a common default.
      owner_gid   = 1000
      owner_uid   = 1000
      permissions = "0755" # Or 0770/0777 if needed, but start stricter
    }
  }

  # Optional: Define POSIX user if different from creation_info
  # posix_user {
  #   gid = 1000
  #   uid = 1000
  # }

  tags = {
    Name = "${local.base_name}-mongo-data-ap"
  }

  # Ensure the mount target is ready before creating the AP
  depends_on = [aws_efs_mount_target.mongo_data]
}

output "efs_file_system_id" {
  description = "ID of the EFS file system for MongoDB data"
  value       = aws_efs_file_system.mongo_data.id
}

output "efs_access_point_id" {
  description = "ID of the EFS access point for MongoDB data"
  value       = aws_efs_access_point.mongo_data.id
}