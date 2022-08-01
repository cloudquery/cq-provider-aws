description_modifier "remove_read_only" {
  words = ["  This member is required."]
}

resource "aws" "kinesis" "streams" {
  path = "github.com/aws/aws-sdk-go-v2/service/kinesis/types.StreamDescriptionSummary"
  ignoreError "IgnoreAccessDenied" {
    path = "github.com/cloudquery/cq-provider-aws/client.IgnoreAccessDeniedServiceDisabled"
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cq-provider-aws/client.DeleteAccountRegionFilter"
  }
  multiplex "AwsAccountRegion" {
    path   = "github.com/cloudquery/cq-provider-aws/client.ServiceAccountRegionMultiplexer"
    params = ["kinesis"]
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
  ignore_columns_in_tests = ["kms_key_id", "retention_in_days"]

  options {
    primary_keys = ["arn"]
  }

  userDefinedColumn "arn" {
    type = "string"
    resolver "resolveStreamArn" {
      path          = "github.com/cloudquery/cq-provider-sdk/provider/schema.PathResolver"
      path_resolver = true
      // TODO: require manual changing from ARN -> StreamARN for the path resolver as its not supported by cq-gen yet
      params = ["stream_arn"]
    }
  }

  userDefinedColumn "tags" {
    type              = "json"
    generate_resolver = true
  }
}

resource "aws" "kinesis" "firehoses" {

  path = "github.com/aws/aws-sdk-go-v2/service/firehose/types.DeliveryStreamDescription"
  ignoreError "IgnoreAccessDenied" {
    path = "github.com/cloudquery/cq-provider-aws/client.IgnoreAccessDeniedServiceDisabled"
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cq-provider-aws/client.DeleteAccountRegionFilter"
  }
  multiplex "AwsAccountRegion" {
    path   = "github.com/cloudquery/cq-provider-aws/client.ServiceAccountRegionMultiplexer"
    params = ["firehose"]
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


  userDefinedColumn "tags" {
    type              = "json"
    generate_resolver = true
  }


  column "destinations" {
    skip = true
  }

  column "has_more_destinations" {
    skip = true
  }

  userDefinedColumn "arn" {
    type        = "string"
    description = "The Amazon Resource Name (ARN) of the delivery stream"

    resolver "resolveStreamArn" {
      path          = "github.com/cloudquery/cq-provider-sdk/provider/schema.PathResolver"
      path_resolver = true
      params        = ["DeliveryStreamARN"]
    }
  }

  column "delivery_stream_encryption_configuration" {
    // skip_prefix = true
    rename = "encryption_config"
  }

  column "delivery_stream_encryption_configuration" {
    // skip_prefix = true
    rename = "encryption_config"
  }

  column "delivery_stream_encryption_configuration_failure_description_details" {
    skip = false
  }
  column "delivery_stream_encryption_configuration_failure_description_type" {
    skip = false
  }
  column "source" {
    rename = "source"
  }
  column "source_kinesis_stream_source_description" {
    rename = "_kinesis_stream"
    // skip_prefix = true
  }
  user_relation "aws" "kinesis" "open_search_destination" {
    path = "github.com/aws/aws-sdk-go-v2/service/firehose/types.AmazonopensearchserviceDestinationDescription"
    resolver "resolveTable" {
      path          = "github.com/cloudquery/cq-provider-sdk/provider/schema.PathTableResolver"
      path_resolver = true
      params        = ["Destinations.AmazonopensearchserviceDestinationDescription"]
    }
    column "s3_destination_description"{
      rename = "s3_destination"
    }
       column "s3_destination_encryption_configuration" {
        skip_prefix = true
      }

       column "s3_destination_cloud_watch_logging_options_" {
        skip_prefix = true
      }

    // column "s3_destination_description_encryption_configuration_kms_encryption_config" {
    //   skip = true
    // }
    // column "s3_destination_description_encryption_configuration_no_encryption_config" {
    //   skip = true
    // }
    // column "s3_destination_description_cloud_watch_logging_options_log_group_name" {
    //   skip = true
    // }
    // column "s3_destination_description_cloud_watch_logging_options_log_stream_name" {
    //   skip = true
    // }
  }
  // user_relation "aws" "kinesis" "elasticsearch_destination" {
  //   path = "github.com/aws/aws-sdk-go-v2/service/firehose/types.ElasticsearchDestinationDescription"
  //   resolver "resolveTable" {
  //     path          = "github.com/cloudquery/cq-provider-sdk/provider/schema.PathTableResolver"
  //     path_resolver = true
  //     params        = ["Destinations.ElasticsearchDestinationDescription"]
  //   }
  //   column "s3_destination_description_encryption_configuration_kms_encryption_config" {
  //     skip = true
  //   }
  //   column "s3_destination_description_encryption_configuration_no_encryption_config" {
  //     skip = true
  //   }
  //   column "s3_destination_description_cloud_watch_logging_options_log_group_name" {
  //     skip = true
  //   }
  //   column "s3_destination_description_cloud_watch_logging_options_log_stream_name" {
  //     skip = true
  //   }
  //   column "output_format_configuration_serializer_orc_ser_de_bloom_filter_false_positive_probability" {
  //     skip = true
  //   }
  // }

  user_relation "aws" "kinesis" "extended_s3_destination" {
    path = "github.com/aws/aws-sdk-go-v2/service/firehose/types.ExtendedS3DestinationDescription"
    resolver "resolveTable" {
      path          = "github.com/cloudquery/cq-provider-sdk/provider/schema.PathTableResolver"
      path_resolver = true
      params        = ["Destinations.ExtendedS3DestinationDescription"]
    }
    column "data_format_conversion_configuration" {
      skip_prefix = true
    }
    column "input_format_configuration" {
      skip_prefix = true
    }

    column "deserializer_open_x_json_ser_de_convert_dots_in_json_keys_to_underscores" {
      skip = true
    }
    column "dynamic_partitioning_configuration_retry_options_duration_in_seconds" {
      skip = true
    }

    column "output_format_configuration" {
      skip_prefix = true
    }
    column "s3_backup_description" {
      rename = "s3_backup"
    }

   column "s3_backup_encryption_configuration" {
      skip_prefix = true
    }
    
    // column "s3_backup_description_encryption_configuration_kms_encryption_config" {
    //   member_trusted_principal = ""
    // }

  }

  // user_relation "aws" "kinesis" "http_destination" {
  //   path = "github.com/aws/aws-sdk-go-v2/service/firehose/types.HttpEndpointDestinationDescription"
  //   resolver "resolveTable" {
  //     path          = "github.com/cloudquery/cq-provider-sdk/provider/schema.PathTableResolver"
  //     path_resolver = true
  //     params        = ["Destinations.HttpEndpointDestinationDescription"]
  //   }
  //   column "s3_destination_description" {
  //     skip_prefix = true
  //   }
  // }
  
  // user_relation "aws" "kinesis" "redshift_destination" {
  //   path = "github.com/aws/aws-sdk-go-v2/service/firehose/types.RedshiftDestinationDescription"
  //   resolver "resolveTable" {
  //     path          = "github.com/cloudquery/cq-provider-sdk/provider/schema.PathTableResolver"
  //     path_resolver = true
  //     params        = ["Destinations.RedshiftDestinationDescription"]
  //   }
  //   column "encryption_configuration" {
  //     skip_prefix = true
  //   }
  //   column "encryption_configuration_kms_encryption" {
  //     skip_prefix = true
  //   }
  //   column "s3_destination_description" {
  //     skip_prefix = true
  //   }
  //   column "s3_backup_description" {
  //     skip_prefix = true
  //   }
  // }
  // user_relation "aws" "kinesis" "splunk_destination" {
  //   path = "github.com/aws/aws-sdk-go-v2/service/firehose/types.SplunkDestinationDescription"
  //   resolver "resolveTable" {
  //     path          = "github.com/cloudquery/cq-provider-sdk/provider/schema.PathTableResolver"
  //     path_resolver = true
  //     params        = ["Destinations.SplunkDestinationDescription"]
  //   }
  //   column "s3_destination_description" {
  //     skip_prefix = true
  //   }

  // }
}
