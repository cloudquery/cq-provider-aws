resource "aws_vpc" "fsx_test_vpc" {
  cidr_block = "10.0.0.0/16"
}

resource "aws_subnet" "fsx_test_subnet" {
  vpc_id = aws_vpc.fsx_test_vpc.id
  cidr_block = "10.0.1.0/24"

  tags = {
    Name = "test_vpc"
  }
}

resource "aws_fsx_backup" "test_fsx_backup" {
  file_system_id = aws_fsx_lustre_file_system.test_fsx.id
}

resource "aws_fsx_lustre_file_system" "test_fsx" {
  storage_capacity = 1200
  storage_type = "HDD"
  deployment_type = "PERSISTENT_1"
  per_unit_storage_throughput = 50
  subnet_ids = [aws_subnet.fsx_test_subnet.id]
}

