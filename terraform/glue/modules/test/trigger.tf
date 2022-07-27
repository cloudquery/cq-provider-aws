resource "aws_glue_trigger" "example" {
  name = "${var.prefix}-glue-trigger"
  type = "CONDITIONAL"

  actions {
    job_name = aws_glue_job.aws_glue_job1.name
  }

  predicate {
    conditions {
      job_name = aws_glue_job.aws_glue_job2.name
      state    = "SUCCEEDED"
    }
  }
}