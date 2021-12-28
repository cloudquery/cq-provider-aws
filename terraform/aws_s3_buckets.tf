resource "aws_s3_bucket" "s3_bucket" {
  bucket        = "bucket-test"
  acl           = "private"
  force_destroy = true
}