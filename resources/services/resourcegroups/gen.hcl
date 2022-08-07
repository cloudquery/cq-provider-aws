service          = "aws"
output_directory = "."
add_generate     = true

description_modifier "remove_read_only" {
  words = ["  This member is required."]
}

resource "aws" "resourcegroups" "resource_groups" {
  path = "github.com/aws/aws-sdk-go-v2/service/resourcegroups/types.GroupIdentifier"

  ignoreError "IgnoreCommonErrors" {
    path = "github.com/cloudquery/cq-provider-aws/client.IgnoreCommonErrors"
  }
  multiplex "AwsAccountRegion" {
    path   = "github.com/cloudquery/cq-provider-aws/client.ServiceAccountRegionMultiplexer"
    params = ["resource-groups"]
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cq-provider-aws/client.DeleteAccountRegionFilter"
  }

  options {
    primary_keys = ["arn"]
  }

  column "group_arn" {
    rename = "arn"
  }
  column "group_name" {
    rename = "name"
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
  userDefinedColumn "tags" {
    type              = "json"
    generate_resolver = true
  }

  postResourceResolver "resolveGroupQuery" {
    path     = "github.com/cloudquery/cq-provider-sdk/provider/schema.RowResolver"
    generate = true
  }
  userDefinedColumn "resource_query_type" {
    description = "The type of the query."
    type        = "string"
  }
  userDefinedColumn "resource_query" {
    description = "The query that defines a group or a search."
    type        = "string"
  }
}
