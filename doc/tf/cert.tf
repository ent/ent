provider "aws" {
  region  = "us-east-1"
  version = "~> 2.0"
  alias   = "us-east-1"
}

resource "aws_acm_certificate" "cert" {
  domain_name       = local.domain_name
  validation_method = "DNS"
  provider          = aws.us-east-1

  subject_alternative_names = [
    format("www.%s", local.domain_name),
  ]

  tags = {
    Name = local.domain_name,
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
  provider        = aws.us-east-1
}

resource "aws_acm_certificate_validation" "cert" {
  certificate_arn         = aws_acm_certificate.cert.arn
  validation_record_fqdns = aws_route53_record.cert_validation.*.fqdn
  provider                = aws.us-east-1
}
