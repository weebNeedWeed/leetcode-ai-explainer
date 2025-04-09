variable "availability_zones" {
  type    = list(string)
  default = ["ap-southeast-1a", "ap-southeast-1b"]
}

variable "vpc_cidr" {
  type    = string
  default = "10.0.0.0/16"
}

variable "certificate_arn" {
  type        = string
  description = "ARN of the certificate that is being validated."
}
