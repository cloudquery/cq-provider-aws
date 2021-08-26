resource "aws_waf_ipset" "waf_ipset_2" {
  name = "waf_ipset_2"

  ip_set_descriptors {
    type  = "IPV4"
    value = "192.0.7.0/24"
  }
}

resource "aws_waf_rule" "waf_rule_3" {
  depends_on  = [aws_waf_ipset.waf_ipset_2]
  name        = "waf_rule_3"
  metric_name = "wafrule3"

  predicates {
    data_id = aws_waf_ipset.waf_ipset_2.id
    negated = false
    type    = "IPMatch"
  }
}

resource "aws_waf_web_acl" "waf_web_acl_1" {
  depends_on = [
    aws_waf_ipset.waf_ipset_2,
    aws_waf_rule.waf_rule_3,
  ]
  name        = "waf_web_acl_1"
  metric_name = "wafwebacl1"

  default_action {
    type = "ALLOW"
  }

  rules {
    action {
      type = "BLOCK"
    }

    priority = 1
    rule_id  = aws_waf_rule.waf_rule_3.id
    type     = "REGULAR"
  }
}
