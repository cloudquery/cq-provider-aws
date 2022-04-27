service          = "aws"
output_directory = "."
add_generate     = true

description_modifier "remove_read_only" {
  words = ["  This member is required."]
}

resource "aws" "athena" "data_catalogs" {
  path = "github.com/aws/aws-sdk-go-v2/service/athena/types.DataCatalog"

  ignoreError "IgnoreAccessDenied" {
    path = "github.com/cloudquery/cq-provider-aws/client.IgnoreAccessDeniedServiceDisabled"
  }
  multiplex "AwsAccountRegion" {
    path   = "github.com/cloudquery/cq-provider-aws/client.ServiceAccountRegionMultiplexer"
    params = ["athena"]
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cq-provider-aws/client.DeleteAccountRegionFilter"
  }

  options {
    primary_keys = [
      "account_id",
      "region",
      "name"
    ]
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

  user_relation "aws" "athena" "databases" {
    path = "github.com/aws/aws-sdk-go-v2/service/athena/types.Database"
    user_relation "aws" "athena" "tables" {
      path = "github.com/aws/aws-sdk-go-v2/service/athena/types.TableMetadata"
    }
  }

  userDefinedColumn "tags" {
    type              = "json"
    generate_resolver = true
  }
}