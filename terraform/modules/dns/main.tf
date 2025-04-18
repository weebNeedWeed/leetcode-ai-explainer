terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

data "aws_route53_zone" "dns_zone" {
  name         = var.domain_name
  private_zone = false
}

resource "aws_acm_certificate" "cert" {
  domain_name       = "leetcode.${var.domain_name}"
  validation_method = "DNS"
}

resource "aws_route53_record" "rec" {
  for_each = {
    for dvo in aws_acm_certificate.cert.domain_validation_options : dvo.domain_name => {
      name   = dvo.resource_record_name
      record = dvo.resource_record_value
      type   = dvo.resource_record_type
    }
  }

  allow_overwrite = true
  name            = each.value.name
  records         = [each.value.record]
  ttl             = 60
  type            = each.value.type
  zone_id         = data.aws_route53_zone.dns_zone.zone_id
}

resource "aws_acm_certificate_validation" "cert_validation" {
  certificate_arn         = aws_acm_certificate.cert.arn
  validation_record_fqdns = [for record in aws_route53_record.rec : record.fqdn]
}

resource "aws_route53_record" "www" {
  zone_id = data.aws_route53_zone.dns_zone.zone_id
  name    = "leetcode.${var.domain_name}"
  type    = "A"

  alias {
    name                   = var.lb_alias.dns_name
    zone_id                = var.lb_alias.zone_id
    evaluate_target_health = true
  }
}
