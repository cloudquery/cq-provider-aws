resource "aws_ebs_volume" "aws_ebs_volume_example" {
  availability_zone = "us-west-2a"
  size              = 40

  tags = {
    Name = "ec2-ebs-snapshot-${var.test_prefix}${var.test_suffix}"
  }
}

resource "aws_ebs_snapshot" "aws_ebs_snapshot_example" {
  volume_id = aws_ebs_volume.example.id

  tags = {
    Name = "cq-provider-aws ${var.test_prefix}-${var.test_suffix}"
  }
}