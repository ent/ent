terraform {
  backend "s3" {
    bucket         = "entgo.tfstate"
    region         = "eu-central-1"
    key            = "terraform.tfstate"
    dynamodb_table = "entgo.terraform.lock"
  }

  required_version = "> 0.12"
}

provider "aws" {
  region  = var.region
  version = "~> 2.0"
}

resource "aws_route53_zone" "zone" {
  name = "entgo.io"
}

resource "aws_route53_record" "ns" {
  name    = aws_route53_zone.zone.name
  type    = "NS"
  zone_id = aws_route53_zone.zone.id
  ttl     = 300
  records = aws_route53_zone.zone.name_servers
}

resource "aws_acm_certificate" "cert" {
  domain_name       = aws_route53_zone.zone.name
  validation_method = "DNS"

  subject_alternative_names = [
    "*.${aws_route53_zone.zone.name}"
  ]

  tags = {
    Name = aws_route53_zone.zone.name
  }

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_route53_record" "cert_validation" {
  name            = aws_acm_certificate.cert.domain_validation_options[count.index]["resource_record_name"]
  type            = aws_acm_certificate.cert.domain_validation_options[count.index]["resource_record_type"]
  zone_id         = aws_route53_zone.zone.id
  records         = [aws_acm_certificate.cert.domain_validation_options[count.index]["resource_record_value"]]
  count           = length(aws_acm_certificate.cert.subject_alternative_names) + 1
  ttl             = 60
  allow_overwrite = true
}

resource "aws_acm_certificate_validation" "cert" {
  certificate_arn         = aws_acm_certificate.cert.arn
  validation_record_fqdns = aws_route53_record.cert_validation.*.fqdn
}

resource "aws_iam_user" "deployer" {
  name = format("%s.deployer", var.name)
  path = format("/%s/", var.name)
}

module "website" {
  source             = "git::https://github.com/cloudposse/terraform-aws-s3-website.git?ref=tags/0.8.0"
  name               = var.name
  hostname           = "entgo.io"
  versioning_enabled = true
  parent_zone_id     = aws_route53_zone.zone.id
  deployment_arns    = { tostring(aws_iam_user.deployer.arn) = "" }
}
