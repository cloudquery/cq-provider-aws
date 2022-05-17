service          = "aws"
output_directory = "."
add_generate     = true

description_modifier "remove_read_only" {
  words = ["  This member is required."]
}


resource "aws" "iam" "groups" {
  path = "github.com/aws/aws-sdk-go-v2/service/iam/types.Group"
  ignoreError "IgnoreAccessDenied" {
    path = "github.com/cloudquery/cq-provider-aws/client.IgnoreAccessDeniedServiceDisabled"
  }
  multiplex "AwsAccountRegion" {
    path = "github.com/cloudquery/cq-provider-aws/client.AccountMultiplex"
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
  options {
    primary_keys = ["account_id", "id"]
  }

  column "group_id" {
    rename = "id"
  }

  column "group_name" {
    rename = "name"
  }

  userDefinedColumn "policies" {
    type              = "json"
    generate_resolver = true
    description       = "List of policies attached to group."
  }

  user_relation "aws" "iam" "policies" {
    path = "github.com/aws/aws-sdk-go-v2/service/iam.GetGroupPolicyOutput"
    userDefinedColumn "group_id" {
      type        = "string"
      description = "Group ID the policy belongs too."
      resolver "resolveAWSAccount" {
        path   = "github.com/cloudquery/cq-provider-sdk/provider/schema.ParentResourceFieldResolver"
        params = ["id"]
      }
    }

    userDefinedColumn "account_id" {
      type        = "string"
      description = "The AWS Account ID of the resource."
      resolver "resolveAWSAccount" {
        path = "github.com/cloudquery/cq-provider-aws/client.ResolveAWSAccount"
      }
    }
    options {
      primary_keys = ["group_cq_id", "policy_name"]
    }

    column "policy_document" {
      type              = "json"
      generate_resolver = true
    }
  }

  user_relation "aws" "iam" "accessed_details" {
    path = "github.com/cloudquery/cq-provider-aws/resources/services/iam.AccessedDetails"
    column "service_last_accessed" {
      skip_prefix = true
    }


    relation "aws" "iam" "entities" {
      path = "github.com/aws/aws-sdk-go-v2/service/iam/types.EntityDetails"

      column "entity_info" {
        skip_prefix = true
      }
    }
  }
}
