service = "aws"
output_directory = "."
add_generate = true

resource "aws" "backup" "global_settings" {
  path = "github.com/aws/aws-sdk-go-v2/service/backup.DescribeGlobalSettingsOutput"
  ignoreError "IgnoreAccessDenied" {
    path = "github.com/cloudquery/cq-provider-aws/client.IgnoreAccessDeniedServiceDisabled"
  }
  deleteFilter "AccountFilter" {
    path = "github.com/cloudquery/cq-provider-aws/client.DeleteAccountFilter"
  }
  multiplex "AwsAccount" {
    path = "github.com/cloudquery/cq-provider-aws/client.AccountMultiplex"
  }


  options {
    primary_keys = ["account_id"]
  }

  userDefinedColumn "account_id" {
    description = "The AWS Account ID of the resource."
    type        = "string"
    resolver "resolveAWSAccount" {
      path = "github.com/cloudquery/cq-provider-aws/client.ResolveAWSAccount"
    }
  }

  column "result_metadata" {
    skip = true
  }
}
