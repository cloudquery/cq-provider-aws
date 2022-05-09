resource "aws_xray_group" "xray-group" {
  group_name        = "${var.prefix}-xray-group"
  filter_expression = "responsetime > 5"

  insights_configuration {
    insights_enabled      = true
    notifications_enabled = true
  }

  tags = merge(
    { Name = "${var.prefix}-xray-group" },
    var.tags
  )
}