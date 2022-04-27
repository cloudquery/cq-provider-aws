resource "aws_s3_bucket" "athenabucket" {
  bucket        = "${var.prefix}athenabucket"
  force_destroy = true
}

resource "aws_athena_database" "aws_athena_database" {
  name   = "${var.prefix}athenadatabase"
  bucket = aws_s3_bucket.athenabucket.bucket
}