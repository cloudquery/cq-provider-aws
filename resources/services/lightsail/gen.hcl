service          = "aws"
output_directory = "."
add_generate     = true

description_modifier "remove_read_only" {
  words = ["  This member is required."]
}

resource "aws" "lightsail" "instances" {
  path = "github.com/aws/aws-sdk-go-v2/service/lightsail/types.Instance"
  ignoreError "IgnoreAccessDenied" {
    path = "github.com/cloudquery/cq-provider-aws/client.IgnoreAccessDeniedServiceDisabled"
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cq-provider-aws/client.DeleteAccountRegionFilter"
  }
  multiplex "AwsAccountRegion" {
    path   = "github.com/cloudquery/cq-provider-aws/client.ServiceAccountRegionMultiplexer"
    params = ["lightsail"]
  }

  userDefinedColumn "account_id" {
    description = "The AWS Account ID of the resource."
    type        = "string"
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

  column "tags" {
    type              = "json"
    generate_resolver = true
  }

  relation "aws" "lightsail" "add_ons" {
    ignore_in_tests = true // see https://github.com/hashicorp/terraform-provider-aws/issues/23688
  }

  // disks attached to the instance
  relation "aws" "lightsail" "hardware_disks" {
    path = "github.com/aws/aws-sdk-go-v2/service/lightsail/types.Disk"

    column "tags" {
      type              = "json"
      generate_resolver = true
    }
    column "gb_in_use" {
      ignore_in_tests = true
    }
    column "attachment_state" {
      skip = true
    }

    relation "aws" "lightsail" "add_ons" {
      ignore_in_tests = true // see https://github.com/hashicorp/terraform-provider-aws/issues/23688
    }
  }

  // todo maybe skip original ports column
  user_relation "aws" "lightsail" "port_states" {
    path = "github.com/aws/aws-sdk-go-v2/service/lightsail/types.InstancePortState"
  }


  userDefinedColumn "access_details" {
    type              = "json"
    generate_resolver = true
  }
}

resource "aws" "lightsail" "buckets" {
  path = "github.com/aws/aws-sdk-go-v2/service/lightsail/types.Bucket"
  ignoreError "IgnoreAccessDenied" {
    path = "github.com/cloudquery/cq-provider-aws/client.IgnoreAccessDeniedServiceDisabled"
  }
  multiplex "AwsAccountRegion" {
    path   = "github.com/cloudquery/cq-provider-aws/client.ServiceAccountRegionMultiplexer"
    params = ["lightsail"]
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cq-provider-aws/client.DeleteAccountRegionFilter"
  }

  options {
    primary_keys = [
      "arn"
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


  column "tags" {
    type              = "json"
    generate_resolver = true
  }

  column "resources_receiving_access" {
    type = "json"
  }

  user_relation "aws" "lightsail" "access_keys" {
    path = "github.com/aws/aws-sdk-go-v2/service/lightsail/types.AccessKey"
  }
}


resource "aws" "lightsail" "disks" {
  path = "github.com/aws/aws-sdk-go-v2/service/lightsail/types.Disk"
  ignoreError "IgnoreAccessDenied" {
    path = "github.com/cloudquery/cq-provider-aws/client.IgnoreAccessDeniedServiceDisabled"
  }
  multiplex "AwsAccountRegion" {
    path   = "github.com/cloudquery/cq-provider-aws/client.ServiceAccountRegionMultiplexer"
    params = ["lightsail"]
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cq-provider-aws/client.DeleteAccountRegionFilter"
  }

  options {
    primary_keys = [
      "arn"
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

  column "tags" {
    type              = "json"
    generate_resolver = true
  }

  user_relation "aws" "lightsail" "disk_snapshot" {
    path = "github.com/aws/aws-sdk-go-v2/service/lightsail/types.DiskSnapshot"

    column "tags" {
      type              = "json"
      generate_resolver = true
    }
  }
}


resource "aws" "lightsail" "alarms" {
  path = "github.com/aws/aws-sdk-go-v2/service/lightsail/types.Alarm"
  ignoreError "IgnoreAccessDenied" {
    path = "github.com/cloudquery/cq-provider-aws/client.IgnoreAccessDeniedServiceDisabled"
  }
  multiplex "AwsAccountRegion" {
    path   = "github.com/cloudquery/cq-provider-aws/client.ServiceAccountRegionMultiplexer"
    params = ["lightsail"]
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cq-provider-aws/client.DeleteAccountRegionFilter"
  }

  options {
    primary_keys = [
      "arn"
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

  column "location" {
    skip_prefix = true
  }
  column "region_name" {
    skip = true
  }
  column "monitored_resource_info_name" {
    rename = "monitored_resource_name"
  }
  column "monitored_resource_info_resource_type" {
    rename = "monitored_resource_resource_type"
  }
}


resource "aws" "lightsail" "instance_snapshots" {
  path = "github.com/aws/aws-sdk-go-v2/service/lightsail/types.InstanceSnapshot"
  ignoreError "IgnoreAccessDenied" {
    path = "github.com/cloudquery/cq-provider-aws/client.IgnoreAccessDeniedServiceDisabled"
  }
  multiplex "AwsAccountRegion" {
    path   = "github.com/cloudquery/cq-provider-aws/client.ServiceAccountRegionMultiplexer"
    params = ["lightsail"]
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cq-provider-aws/client.DeleteAccountRegionFilter"
  }

  options {
    primary_keys = [
      "arn"
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

  column "location" {
    skip_prefix = true
  }
  column "region_name" {
    skip = true
  }

  column "tags" {
    type              = "json"
    generate_resolver = true
  }

  relation "aws" "lightsail" "from_attached_disks" {
    column "tags" {
      type              = "json"
      generate_resolver = true
    }
  }
}