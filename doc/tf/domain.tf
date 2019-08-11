locals {
  domain_name = "entgo.io"
}

resource "aws_route53_zone" "zone" {
  name = local.domain_name
}

resource "aws_route53_record" "ns" {
  name    = aws_route53_zone.zone.name
  type    = "NS"
  zone_id = aws_route53_zone.zone.id
  ttl     = 300
  records = aws_route53_zone.zone.name_servers
}
