terraform {
  backend "s3" {
    bucket = "s3-state-bucket-04092025"
    region = "ap-southeast-1"
    key    = "states/leetcode/terraform.tfstate"
  }

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

variable "github_token" {
  type      = string
  sensitive = true
}

variable "gemini_apikey" {
  type      = string
  sensitive = true
}

provider "aws" {
  region = "ap-southeast-1"
}

module "dns" {
  source = "../../modules/dns"
  lb_alias = {
    dns_name = module.networking.lb_alias.dns_name
    zone_id  = module.networking.lb_alias.zone_id
  }
}

module "networking" {
  source = "../../modules/networking"

  certificate_arn = module.dns.certificate_arn
}

module "compute" {
  source = "../../modules/compute"

  subnet_ids         = module.networking.private_subnet_ids
  sg_ecs_task_id     = module.networking.sg_ecs_task_id
  ecr_repository_url = module.cicd.ecr_repository_url
  target_group_arn   = module.networking.target_group_arn
  gemini_apikey      = var.gemini_apikey
  github_token       = var.github_token
}

module "cicd" {
  source = "../../modules/cicd"
}
