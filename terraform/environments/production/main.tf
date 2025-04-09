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

provider "aws" {
  region = "ap-southeast-1"
}

# module "dns" {
#   source = "../../modules/dns"
#   lb_alias = {
#     dns_name = module.networking.lb_alias.dns_name
#     zone_id  = module.networking.lb_alias.zone_id
#   }
# }

# module "networking" {
#   source = "../../modules/networking"

#   certificate_arn = module.dns.certificate_arn
# }

# module "compute" {
#   source = "../../modules/compute"

#   subnet_ids     = module.networking.private_subnet_ids
#   sg_ecs_task_id = module.networking.sg_ecs_task_id
# }

module "cicd" {
  source = "../../modules/cicd"
}
