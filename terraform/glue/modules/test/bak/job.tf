resource "aws_glue_job" "aws_glue_job1" {
  name     = "${var.prefix}-glue-job1"
  role_arn = aws_iam_role.aws_iam_role.arn

  command {
    script_location = "s3://${aws_s3_bucket.aws_s3_bucket.bucket}/scripts/example.py"
  }
}

resource "aws_glue_job" "aws_glue_job2" {
  name     = "${var.prefix}-glue-job2"
  role_arn = aws_iam_role.aws_iam_role.arn

  command {
    script_location = "s3://${aws_s3_bucket.aws_s3_bucket.bucket}/scripts/example.py"
  }
}