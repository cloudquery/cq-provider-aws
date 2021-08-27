resource "aws_ebs_volume" "example" {
  availability_zone = data.aws_availability_zones.available.names[0]
  size = 5

  tags = {
    Name = "ebs-${var.test_prefix}${var.test_suffix}"
  }
}

data "aws_availability_zones" "available" {
  state = "available"
}