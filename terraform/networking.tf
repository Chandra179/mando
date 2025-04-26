# Consistent naming using variables
locals {
  base_name = "${var.app_name}-${var.environment}"
}

resource "aws_vpc" "main" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_support   = true
  enable_dns_hostnames = true

  tags = {
    Name = "${local.base_name}-vpc"
  }
}

# Need at least one subnet for ECS Fargate tasks
resource "aws_subnet" "public" {
  vpc_id                  = aws_vpc.main.id
  cidr_block              = "10.0.1.0/24"
  map_public_ip_on_launch = true # Needed for Fargate tasks to pull images via IGW
  availability_zone       = "${var.region}a" # Adjust AZ if needed for your LocalStack

  tags = {
    Name = "${local.base_name}-public-subnet"
  }
}

resource "aws_internet_gateway" "gw" {
  vpc_id = aws_vpc.main.id

  tags = {
    Name = "${local.base_name}-igw"
  }
}

resource "aws_route_table" "public" {
  vpc_id = aws_vpc.main.id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.gw.id
  }

  tags = {
    Name = "${local.base_name}-public-rt"
  }
}

resource "aws_route_table_association" "public" {
  subnet_id      = aws_subnet.public.id
  route_table_id = aws_route_table.public.id
}

# Security Group for all ECS tasks
resource "aws_security_group" "ecs_tasks" {
  name        = "${local.base_name}-ecs-tasks-sg"
  description = "Allow traffic between ECS tasks and specific inbound ports"
  vpc_id      = aws_vpc.main.id

  # Allow all internal traffic within this security group (Best Practice)
  ingress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    self        = true
    description = "Allow internal communication"
  }

  # Allow external access to Go App (port 8080)
  ingress {
    from_port   = 8080
    to_port     = 8080
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"] # For local dev; restrict in production
    description = "Allow HTTP traffic to Go App"
  }

  # Allow external access to Mongo Express (port 8081)
  ingress {
    from_port   = 8081
    to_port     = 8081
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"] # For local dev; restrict in production
    description = "Allow HTTP traffic to Mongo Express"
  }

   # Allow external access to MongoDB (port 27017) - Optional for direct debugging
   # Primarily needed for internal access (covered by 'self = true' rule)
   # ingress {
   #   from_port   = 27017
   #   to_port     = 27017
   #   protocol    = "tcp"
   #   cidr_blocks = ["0.0.0.0/0"] # Restrict this severely!
   #   description = "Allow direct access to MongoDB (Use with caution!)"
   # }


  # Allow all outbound traffic (Needed for pulling images, etc.)
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "${local.base_name}-ecs-tasks-sg"
  }
}