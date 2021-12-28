resource "aws_ebs_volume" "aws_ebs_volume_example" {
  availability_zone = "us-east-1a"
  size              = 40

  tags = {
    Name = "ec2-ebs-volume-test"
  }
}

resource "aws_ebs_snapshot" "aws_ebs_snapshot_example" {
  volume_id = aws_ebs_volume.aws_ebs_volume_example.id

  tags = {
    Name = "ec2-ebs-snapshot-test"
  }
}