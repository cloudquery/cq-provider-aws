resource "aws_config_configuration_recorder" "config_recorder_1" {
  name     = "config_recorder"
  role_arn = aws_iam_role.config_recorder_iam_role.arn
}

resource "aws_iam_role" "config_recorder_iam_role" {
  name = "config_recorder_iam_role"

  assume_role_policy = <<POLICY
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "config.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
POLICY
}
