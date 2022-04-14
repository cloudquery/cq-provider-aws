service = "aws"

output_directory = "."

resource "aws" "ssm" "instance_patches" {
  path = "github.com/aws/aws-sdk-go-v2/service/ssm/types.PatchComplianceData"
  ignoreError "IgnoreAccessDenied" {
    path = "github.com/cloudquery/cq-provider-aws/client.IgnoreAccessDeniedServiceDisabled"
  }
}
