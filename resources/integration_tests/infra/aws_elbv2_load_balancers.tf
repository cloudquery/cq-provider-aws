resource "aws_lb" "aws_elbv2_load_balancer" {
  name = "elbv2-${var.test_suffix}"
  load_balancer_type = "network"

  subnets = [
    aws_subnet.elbv2-subnet.id]
  tags = {
    test = "test"
  }
}

//todo might be moved to some general file
resource "aws_vpc" "elbv2-vpc" {
  cidr_block = "172.16.0.0/16"

  tags = {
    Name = "tf-example"
  }
}

resource "aws_internet_gateway" "elbv2-gw" {
  vpc_id = aws_vpc.elbv2-vpc.id
}

resource "aws_subnet" "elbv2-subnet" {

  vpc_id = aws_vpc.elbv2-vpc.id
  cidr_block = "172.16.1.0/24"
  availability_zone = "us-east-1a"

  tags = {
    Name = "tf-example"
  }
}