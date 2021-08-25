resource "aws_apigatewayv2_vpc_link" "aws_apigatewayv2_vpc_links_link" {
  name = "link-${var.test_prefix}-${var.test_suffix}"
  security_group_ids = [
    aws_security_group.aws_apigatewayv2_vpc_links_sg.id]
  subnet_ids = [
    aws_subnet.aws_apigatewayv2_vpc_links_subnet.id]
}

resource "aws_security_group" "aws_apigatewayv2_vpc_links_sg" {
  name = "${var.test_prefix}-${var.test_suffix}-sg"
  vpc_id = aws_vpc.aws_apigatewayv2_vpc_links_vpc.id
  ingress {
    from_port = 22
    to_port = 22
    protocol = "tcp"
    cidr_blocks = [
      "0.0.0.0/0"]
  }
}

resource "aws_subnet" "aws_apigatewayv2_vpc_links_subnet" {
  vpc_id = aws_vpc.aws_apigatewayv2_vpc_links_vpc.id
  cidr_block = "172.16.10.0/24"
  availability_zone = "us-east-1a"
}

//todo might be moved to some general file
resource "aws_vpc" "aws_apigatewayv2_vpc_links_vpc" {
  cidr_block = "172.16.0.0/16"

  tags = {
    Name = "vpc-${var.test_prefix}-${var.test_suffix}"
  }
}
