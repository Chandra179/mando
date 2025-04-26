terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0" # Use a recent version
    }
  }
}

provider "aws" {
  region                      = var.region
  access_key                  = "test"      # Dummy credentials for LocalStack
  secret_key                  = "test"
  skip_credentials_validation = true
  skip_metadata_api_check     = true
  skip_requesting_account_id  = true

  # Point Terraform to LocalStack endpoints (Best Practice: Explicitly list used services)
  endpoints {
    ec2        = "http://localhost:4566"
    ecr        = "http://localhost:4566"
    ecs        = "http://localhost:4566"
    efs        = "http://localhost:4566"
    iam        = "http://localhost:4566"
    logs       = "http://localhost:4566"
    sts        = "http://localhost:4566"
    # Add others if needed by specific resources (e.g., cloudwatch for alarms)
  }
}