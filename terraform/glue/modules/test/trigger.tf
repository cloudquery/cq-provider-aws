resource "aws_glue_trigger" "example" {
  name = "${var.prefix}-glue-trigger"
  type = "ON_DEMAND"
  description = "test trigger"

  actions {
    job_name = aws_glue_job.example.name
  }
}
