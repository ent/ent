locals {
  allowed_cidrs = length(var.allowed_cidrs) == 0 ? data.terraform_remote_state.current.outputs.allowed_cidrs : var.allowed_cidrs
}

resource "aws_waf_ipset" "ipset" {
  name = "whitelistset"

  dynamic "ip_set_descriptors" {
    for_each = local.allowed_cidrs

    content {
      type  = "IPV4"
      value = ip_set_descriptors.value
    }
  }
}

resource "aws_waf_rule" "rule" {
  metric_name = "whitelistrule"
  name        = "whitelistrule"

  predicates {
    data_id = aws_waf_ipset.ipset.id
    negated = false
    type    = "IPMatch"
  }
}

resource "aws_waf_web_acl" "acl" {
  metric_name = "whitelistacl"
  name        = "whitelistacl"

  default_action {
    type = "BLOCK"
  }

  rules {
    action {
      type = "ALLOW"
    }

    priority = 1
    rule_id  = aws_waf_rule.rule.id
    type     = "REGULAR"
  }
}
