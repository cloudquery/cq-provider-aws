service          = "aws"
output_directory = "."
add_generate     = true

description_modifier "remove_read_only" {
  words = ["  This member is required."]
}

resource "aws" "inspector" "findings" {
  path = "github.com/aws/aws-sdk-go-v2/service/inspector2/types.Finding"
  ignoreError "IgnoreCommonErrors" {
    path = "github.com/cloudquery/cq-provider-aws/client.IgnoreCommonErrors"
  }
  multiplex "AwsAccountRegion" {
    path   = "github.com/cloudquery/cq-provider-aws/client.ServiceAccountRegionMultiplexer"
    params = ["inspector2"]
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cq-provider-aws/client.DeleteAccountRegionFilter"
  }

  options {
    primary_keys = [
      "arn"
    ]
  }

  column "finding_arn" {
    rename = "arn"
  }

  column "aws_account_id" {
    rename = "account_id"
  }

  userDefinedColumn "region" {
    type        = "string"
    description = "The AWS Region of the resource"
    resolver "resolveAWSRegion" {
      path = "github.com/cloudquery/cq-provider-aws/client.ResolveAWSRegion"
    }
  }

  ignore_columns_in_tests = [
    "inspector_score_details",
    "network_reachability_details",
    "package_vulnerability_details",
  ]

  column "inspector_score_details" {
    type = "json"
  }

  column "network_reachability_details" {
    type = "json"
  }

  column "package_vulnerability_details" {
    type = "json"
  }

  relation "aws" "inspector" "resources" {
    ignore_columns_in_tests = [
      "details_aws_ec2_instance",
      "details_aws_ecr_container_image",
    ]

    column "details_aws_ec2_instance" {
      rename = "aws_ec2_instance"
      type   = "json"
    }

    column "details_aws_ecr_container_image" {
      rename = "aws_ecr_container_image"
      type   = "json"
    }
  }
}