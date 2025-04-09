variable "domain_name" {
  type    = string
  default = "harley.id.vn"
}

variable "lb_alias" {
  type = object({
    dns_name = string
    zone_id  = string
  })
}
