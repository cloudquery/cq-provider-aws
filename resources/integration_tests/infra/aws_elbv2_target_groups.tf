resource "aws_lb_target_group" "elbv2_target_groups_tg" {
  name = substr("${var.test_prefix}${var.test_suffix}",0,31)
  port = 80
  protocol = "HTTP"
  vpc_id = aws_vpc.elbv2_target_groups_vpc.id
  tags = {
    test = "test"
  }
}

resource "aws_vpc" "elbv2_target_groups_vpc" {
  cidr_block = "10.0.0.0/16"
}