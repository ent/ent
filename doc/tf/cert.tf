
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
  name            = aws_acm_certificate.cert.domain_validation_options.0.resource_record_name
  type            = aws_acm_certificate.cert.domain_validation_options.0.resource_record_type
  zone_id         = aws_route53_zone.zone.id
  records         = [aws_acm_certificate.cert.domain_validation_options.0.resource_record_value]
  ttl             = 60
  allow_overwrite = true
}

resource "aws_acm_certificate_validation" "cert" {
  certificate_arn         = aws_acm_certificate.cert.arn
  validation_record_fqdns = [aws_route53_record.cert_validation.fqdn]
}
