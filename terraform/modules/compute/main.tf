terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

resource "aws_iam_role" "ecs_execution_role" {
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
}

resource "aws_iam_role_policy_attachment" "exec_role_attach" {
  for_each = toset([
    "arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryFullAccess",
    "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy",
    "arn:aws:iam::aws:policy/CloudWatchLogsFullAccess"
  ])
  role       = aws_iam_role.ecs_execution_role.name
  policy_arn = each.value
}

resource "aws_iam_role" "task_role" {
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
}

resource "aws_iam_role_policy_attachment" "task_role_attach" {
  for_each = toset([
    "arn:aws:iam::aws:policy/AmazonDynamoDBFullAccess",
    "arn:aws:iam::aws:policy/CloudWatchLogsFullAccess"
  ])
  role       = aws_iam_role.task_role.name
  policy_arn = each.value
}

resource "aws_iam_role_policy_attachment" "exec_role_attach_ecr" {
  role       = aws_iam_role.ecs_execution_role.name
  policy_arn = "arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryFullAccess"
}

resource "aws_iam_role_policy_attachment" "exec_role_attach_execrole" {
  role       = aws_iam_role.ecs_execution_role.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"
}

resource "aws_iam_role_policy_attachment" "exec_role_attach_cw" {
  role       = aws_iam_role.ecs_execution_role.name
  policy_arn = "arn:aws:iam::aws:policy/CloudWatchLogsFullAccess"
}

resource "aws_service_discovery_http_namespace" "namespace" {
  name = "leetcode"
}

resource "aws_ecs_cluster" "main" {
  name = "leetcode-cluster"

  setting {
    name  = "containerInsights"
    value = "enabled"
  }

  service_connect_defaults {
    namespace = aws_service_discovery_http_namespace.namespace.arn
  }
}

resource "aws_ecs_cluster_capacity_providers" "cp" {
  cluster_name       = aws_ecs_cluster.main.name
  capacity_providers = ["FARGATE"]
}

resource "aws_ecs_task_definition" "api_td" {
  family             = "api-go"
  execution_role_arn = aws_iam_role.ecs_execution_role.arn
  task_role_arn      = aws_iam_role.task_role.arn
  cpu                = 256
  memory             = 512

  runtime_platform {
    operating_system_family = "LINUX"
    cpu_architecture        = "X86_64"
  }

  container_definitions = jsonencode([
    {
      name      = "api-go"
      image     = "public.ecr.aws/f9n5f1l7/dgs:latest"
      cpu       = 128
      memory    = 256
      essential = true
      portMappings = [
        {
          containerPort = 9090
          hostPort      = 8080
          protocol      = "tcp"
          name          = "http"
        }
      ]
    }
  ])
}

resource "aws_ecs_service" "api_go_ecs_service" {
  name            = "api-go"
  cluster         = aws_ecs_cluster.main.id
  task_definition = aws_ecs_task_definition.api_td.arn
  desired_count   = 1

  network_configuration {
    subnets         = var.subnet_ids
    security_groups = [var.sg_ecs_task_id]
  }

  force_new_deployment = true

  triggers = {
    redeployment = plantimestamp()
  }

  force_delete = true

  service_connect_configuration {
    enabled   = true
    namespace = aws_service_discovery_http_namespace.namespace.id
    service {
      client_alias {
        dns_name = "api-go"
        port     = "8080"
      }
      discovery_name = "api-go"
      port_name      = "http"
    }
  }
}
