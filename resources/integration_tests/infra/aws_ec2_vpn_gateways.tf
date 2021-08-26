resource "aws_vpn_gateway" "aws_ec2_vpn_gateway" {
  vpc_id = aws_vpc.aws_vpc.id

  tags = {
    Name = "vg_${var.test_prefix}${var.test_suffix}"
  }
}

