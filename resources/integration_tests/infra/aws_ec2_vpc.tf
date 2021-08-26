resource "aws_vpc_peering_connection" "aws_ec2_vpc_peering_connections" {
  peer_vpc_id = aws_vpc.aws_vpc.id
  vpc_id = aws_vpc.aws_ec2_vpc_peering_vpc.id

  accepter {
    allow_remote_vpc_dns_resolution = true
  }

  requester {
    allow_remote_vpc_dns_resolution = true
  }
  auto_accept = true

}

resource "aws_vpc" "aws_ec2_vpc_peering_vpc" {
  cidr_block = "10.1.0.0/16"
  tags = {
    Name = "vpc${var.test_prefix}-${var.test_suffix}"
  }
  enable_dns_hostnames = true
}

resource "aws_vpc_endpoint" "aws_ec2_vpc_endpoint" {
  vpc_id       = aws_vpc.aws_ec2_vpc_peering_vpc.id
  service_name = "com.amazonaws.${data.aws_region.current.name}.s3"

  tags = {
    Environment = "test"
  }
}

