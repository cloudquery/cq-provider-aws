resource "aws_config_conformance_pack" "aws_config_cp" {
  name = "config-cp-"

  input_parameter {
    parameter_name  = "AccessKeysRotatedParameterMaxAccessKeyAge"
    parameter_value = "90"
  }

  template_body = <<EOT
Parameters:
  AccessKeysRotatedParameterMaxAccessKeyAge:
    Type: String
Resources:
  IAMPasswordPolicy:
    Properties:
      ConfigRuleName: IAMPasswordPolicy
      Source:
        Owner: AWS
        SourceIdentifier: IAM_PASSWORD_POLICY
    Type: AWS::Config::ConfigRule
EOT

  depends_on = [aws_config_configuration_recorder.aws_config_cr]
}

resource "aws_config_configuration_recorder" "aws_config_cr" {
  name     = "config-cr-"
  role_arn = aws_iam_role.aws_config_ir.arn
}

resource "aws_iam_role" "aws_config_ir" {
  name = "config-role-"

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