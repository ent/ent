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
  region  = "eu-central-1"
  version = "~> 2.0"
}

locals {
  name   = "entgo"
  domain = "entgo.io"
}

resource "aws_route53_zone" "zone" {
  name = local.domain
}

resource "aws_route53_record" "ns" {
  name    = aws_route53_zone.zone.name
  type    = "NS"
  zone_id = aws_route53_zone.zone.id
  ttl     = 300
  records = aws_route53_zone.zone.name_servers
}

provider "aws" {
  region  = "us-east-1"
  version = "~> 2.0"
  alias   = "us-east-1"
}

module "certificate" {
  source                      = "git::https://github.com/cloudposse/terraform-aws-acm-request-certificate.git?ref=tags/0.4.0"
  domain_name                 = local.domain
  zone_name                   = local.domain
  subject_alternative_names   = [format("*.%s", local.domain)]
  wait_for_certificate_issued = true
  providers                   = { aws = aws.us-east-1 }
}
resource "aws_iam_user" "deployer" {
  name = format("%s.deployer", local.name)
  path = format("/%s/", local.name)
}

module "website" {
  source             = "git::https://github.com/cloudposse/terraform-aws-s3-website.git?ref=tags/0.8.0"
  name               = local.name
  hostname           = local.domain
  versioning_enabled = true
  parent_zone_id     = aws_route53_zone.zone.id
  deployment_arns    = { tostring(aws_iam_user.deployer.arn) = "" }
}
