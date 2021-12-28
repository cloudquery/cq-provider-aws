resource "aws_efs_file_system" "aws_efs_filesystems_system" {
  creation_token = "efs-test"

  lifecycle_policy {
    transition_to_ia = "AFTER_30_DAYS"
  }

  tags = {
    Name = "efs-test"
  }
}