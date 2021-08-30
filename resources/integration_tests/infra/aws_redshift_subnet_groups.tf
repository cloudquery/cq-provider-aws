resource "aws_vpc" "redshift_vpc_1" {
  cidr_block = "10.173.0.0/16"
}

resource "aws_subnet" "redshift_subnet_1" {
  cidr_block        = "10.173.1.0/24"
  vpc_id            = aws_vpc.redshift_vpc_1.id

  tags = {
    Name = "tf-dbsubnet-test-1"
  }
}

resource "aws_subnet" "redshift_subnet_2" {
  cidr_block        = "10.173.2.0/24"
  vpc_id            = aws_vpc.redshift_vpc_1.id

  tags = {
    Name = "tf-dbsubnet-test-2"
  }
}

resource "aws_redshift_subnet_group" "redshift_subnet_group_1" {
  name       = "redshift-subnet-group-1"
  subnet_ids = [aws_subnet.redshift_subnet_1.id, aws_subnet.redshift_subnet_2.id]
  description = "my test description"

  tags = {
    environment = "Production"
  }
}
