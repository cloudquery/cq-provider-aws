//check-for-changes
service          = "aws"
output_directory = "."
add_generate     = true

description_modifier "remove_read_only" {
  words = ["  This member is required."]
}

resource "aws" "eventbridge" "event_buses" {
  path = "github.com/aws/aws-sdk-go-v2/service/eventbridge/types.EventBus"

  ignoreError "IgnoreCommonErrors" {
    path = "github.com/cloudquery/cq-provider-aws/client.IgnoreCommonErrors"
  }
  multiplex "AwsAccountRegion" {
    path   = "github.com/cloudquery/cq-provider-aws/client.ServiceAccountRegionMultiplexer"
    params = ["events"]
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cq-provider-aws/client.DeleteAccountRegionFilter"
  }

  options {
    primary_keys = ["arn"]
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

  user_relation "aws" "eventbridge" "rules" {
    path = "github.com/aws/aws-sdk-go-v2/service/eventbridge/types.Rule"
    options {
      primary_keys = [
        "event_bus_cq_id",
        "arn"
      ]
    }

    userDefinedColumn "tags" {
      type              = "json"
      generate_resolver = true
    }
  }
}
