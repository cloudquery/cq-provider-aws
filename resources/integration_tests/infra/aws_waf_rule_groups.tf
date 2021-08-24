resource "aws_waf_rule" "waf_rule_2" {
  name        = "waf_rule_2"
  metric_name = "wafrule2"
}

resource "aws_waf_rule_group" "waf_rule_group_1" {
  name        = "waf_rule_group_1"
  metric_name = "wafrulegroup1"

  activated_rule {
    action {
      type = "COUNT"
    }

    priority = 50
    rule_id  = aws_waf_rule.waf_rule_2.id
  }
}
