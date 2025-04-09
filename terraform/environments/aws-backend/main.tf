terraform {
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

resource "aws_s3_bucket" "s3_state_bucket" {
  bucket        = "s3-state-bucket-04092025"
  force_destroy = true
}
