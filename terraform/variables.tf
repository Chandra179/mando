variable "region" {
  description = "AWS region for LocalStack"
  type        = string
  default     = "us-east-1" # Common LocalStack default
}

variable "app_image_tag" {
  description = "Tag for the mando-app Docker image"
  type        = string
  default     = "latest"
}

variable "mongo_root_user" {
  description = "MongoDB root username"
  type        = string
  default     = "root"
  # sensitive = true # Good practice, but less critical for known local dev defaults
}

variable "mongo_root_password" {
  description = "MongoDB root password"
  type        = string
  default     = "root"
  sensitive   = true # Mark passwords as sensitive
}

variable "mongo_db_name" {
  description = "MongoDB database name"
  type        = string
  default     = "mando"
}

variable "mongo_express_user" {
  description = "Mongo Express basic auth username"
  type        = string
  default     = "root"
  # sensitive = true
}

variable "mongo_express_password" {
  description = "Mongo Express basic auth password"
  type        = string
  default     = "root"
  sensitive   = true
}

variable "app_name" {
  description = "Base name for application resources"
  type        = string
  default     = "mando"
}

variable "environment" {
  description = "Deployment environment (e.g., local, dev)"
  type        = string
  default     = "localstack"
}