service          = "aws"
output_directory = "."
add_generate     = true

resource "aws" "shield" "protections" {
  path = "github.com/aws/aws-sdk-go-v2/service/shield/types.Protection"
  ignoreError "IgnoreAccessDenied" {
    path = "github.com/cloudquery/cq-provider-aws/client.IgnoreAccessDeniedServiceDisabled"
  }
  multiplex "AwsAccount" {
    path   = "github.com/cloudquery/cq-provider-aws/client.AccountMultiplex,"
    params = ["shield"]
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cq-provider-aws/client.DeleteAccountFilter"
  }
  userDefinedColumn "account_id" {
    type        = "string"
    description = "The AWS Account ID of the resource."
    resolver "resolveAWSAccount" {
      path = "github.com/cloudquery/cq-provider-aws/client.ResolveAWSAccount"
    }
  }
  userDefinedColumn "region" {
    type        = "string"
    description = "The AWS Region of the resource."
    resolver "resolveAWSRegion" {
      path = "github.com/cloudquery/cq-provider-aws/client.ResolveAWSRegion"
    }
  }

  options {
    primary_keys = ["arn"]
  }

  column "protection_arn" {
    rename = "arn"
  }

  userDefinedColumn "tags" {
    type              = "json"
    description       = "The AWS tags of the resource."
    generate_resolver = true
  }
}