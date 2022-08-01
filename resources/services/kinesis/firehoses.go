package kinesis

import (
	"context"

	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func Firehoses() *schema.Table {
	return &schema.Table{
		Name:         "aws_kinesis_firehoses",
		Description:  "Contains information about a delivery stream.",
		Resolver:     fetchKinesisFirehoses,
		Multiplex:    client.ServiceAccountRegionMultiplexer("firehose"),
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: ResolveKinesisFirehoseTags,
			},
			{
				Name:        "delivery_stream_arn",
				Description: "The Amazon Resource Name (ARN) of the delivery stream",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DeliveryStreamARN"),
			},
			{
				Name:        "delivery_stream_name",
				Description: "The name of the delivery stream.  This member is required.",
				Type:        schema.TypeString,
			},
			{
				Name:        "delivery_stream_status",
				Description: "The status of the delivery stream",
				Type:        schema.TypeString,
			},
			{
				Name:        "delivery_stream_type",
				Description: "The delivery stream type",
				Type:        schema.TypeString,
			},
			{
				Name:        "has_more_destinations",
				Description: "Indicates whether there are more destinations available to list.  This member is required.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "version_id",
				Description: "Each time the destination is updated for a delivery stream, the version ID is changed, and the current version ID is required when updating the destination. This is so that the service knows it is applying the changes to the correct version of the delivery stream.  This member is required.",
				Type:        schema.TypeString,
			},
			{
				Name:        "create_timestamp",
				Description: "The date and time that the delivery stream was created.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "failure_description_details",
				Description: "A message providing details about the error that caused the failure.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DeliveryStreamEncryptionConfiguration.FailureDescription.Details"),
			},
			{
				Name:        "failure_description_type",
				Description: "The type of error that caused the failure.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DeliveryStreamEncryptionConfiguration.FailureDescription.Type"),
			},
			{
				Name:        "key_arn",
				Description: "If KeyType is CUSTOMER_MANAGED_CMK, this field contains the ARN of the customer managed CMK",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DeliveryStreamEncryptionConfiguration.KeyARN"),
			},
			{
				Name:        "key_type",
				Description: "Indicates the type of customer master key (CMK) that is used for encryption",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DeliveryStreamEncryptionConfiguration.KeyType"),
			},
			{
				Name:        "status",
				Description: "This is the server-side encryption (SSE) status for the delivery stream",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DeliveryStreamEncryptionConfiguration.Status"),
			},
			{
				Name:        "failure_description_details",
				Description: "A message providing details about the error that caused the failure.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("FailureDescription.Details"),
			},
			{
				Name:        "failure_description_type",
				Description: "The type of error that caused the failure.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("FailureDescription.Type"),
			},
			{
				Name:        "last_update_timestamp",
				Description: "The date and time that the delivery stream was last updated.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "source_kinesis_stream_source_description_kinesis_stream_arn",
				Description: "The Amazon Resource Name (ARN) of the source Kinesis data stream",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Source.KinesisStreamSourceDescription.KinesisStreamARN"),
			},
			{
				Name:        "source_kinesis_stream_source_description_role_arn",
				Description: "The ARN of the role used by the source Kinesis data stream",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Source.KinesisStreamSourceDescription.RoleARN"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "aws_kinesis_firehose_open_search_destination",
				Resolver: fetchKinesisFirehoseOpenSearchDestinations,
				Columns: []schema.Column{
					{
						Name:        "firehose_cq_id",
						Description: "Unique CloudQuery ID of aws_kinesis_firehoses table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:     "buffering_hints_interval_in_seconds",
						Type:     schema.TypeBigInt,
						Resolver: schema.PathResolver("BufferingHints.IntervalInSeconds"),
					},
					{
						Name:     "buffering_hints_size_in_mb_s",
						Type:     schema.TypeBigInt,
						Resolver: schema.PathResolver("BufferingHints.SizeInMBs"),
					},
					{
						Name:        "cloud_watch_logging_options_enabled",
						Description: "Enables or disables CloudWatch logging.",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("CloudWatchLoggingOptions.Enabled"),
					},
					{
						Name:        "cloud_watch_logging_options_log_group_name",
						Description: "The CloudWatch group name for logging",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("CloudWatchLoggingOptions.LogGroupName"),
					},
					{
						Name:        "cloud_watch_logging_options_log_stream_name",
						Description: "The CloudWatch log stream name for logging",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("CloudWatchLoggingOptions.LogStreamName"),
					},
					{
						Name: "cluster_endpoint",
						Type: schema.TypeString,
					},
					{
						Name:     "domain_arn",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("DomainARN"),
					},
					{
						Name: "index_name",
						Type: schema.TypeString,
					},
					{
						Name: "index_rotation_period",
						Type: schema.TypeString,
					},
					{
						Name:        "processing_configuration_enabled",
						Description: "Enables or disables data processing.",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("ProcessingConfiguration.Enabled"),
					},
					{
						Name:     "retry_options_duration_in_seconds",
						Type:     schema.TypeBigInt,
						Resolver: schema.PathResolver("RetryOptions.DurationInSeconds"),
					},
					{
						Name:     "role_arn",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("RoleARN"),
					},
					{
						Name: "s3_backup_mode",
						Type: schema.TypeString,
					},
					{
						Name:        "s3_destination_description_bucket_arn",
						Description: "The ARN of the S3 bucket",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("S3DestinationDescription.BucketARN"),
					},
					{
						Name:        "s3_destination_description_buffering_hints_interval_in_seconds",
						Description: "Buffer incoming data for the specified period of time, in seconds, before delivering it to the destination",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("S3DestinationDescription.BufferingHints.IntervalInSeconds"),
					},
					{
						Name:        "s3_destination_description_buffering_hints_size_in_mb_s",
						Description: "Buffer incoming data to the specified size, in MiBs, before delivering it to the destination",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("S3DestinationDescription.BufferingHints.SizeInMBs"),
					},
					{
						Name:        "s3_destination_description_compression_format",
						Description: "The compression format",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("S3DestinationDescription.CompressionFormat"),
					},
					{
						Name:        "s3_destination_description_role_arn",
						Description: "The Amazon Resource Name (ARN) of the AWS credentials",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("S3DestinationDescription.RoleARN"),
					},
					{
						Name:        "s3_destination_description_cloud_watch_logging_options_enabled",
						Description: "Enables or disables CloudWatch logging.",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("S3DestinationDescription.CloudWatchLoggingOptions.Enabled"),
					},
					{
						Name:        "s3_destination_description_error_output_prefix",
						Description: "A prefix that Kinesis Data Firehose evaluates and adds to failed records before writing them to S3",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("S3DestinationDescription.ErrorOutputPrefix"),
					},
					{
						Name:        "s3_destination_description_prefix",
						Description: "The \"YYYY/MM/DD/HH\" time format prefix is automatically used for delivered Amazon S3 files",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("S3DestinationDescription.Prefix"),
					},
					{
						Name: "type_name",
						Type: schema.TypeString,
					},
					{
						Name:        "vpc_configuration_description_role_arn",
						Description: "The ARN of the IAM role that the delivery stream uses to create endpoints in the destination VPC",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VpcConfigurationDescription.RoleARN"),
					},
					{
						Name:        "vpc_configuration_description_security_group_ids",
						Description: "The IDs of the security groups that Kinesis Data Firehose uses when it creates ENIs in the VPC of the Amazon ES destination",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("VpcConfigurationDescription.SecurityGroupIds"),
					},
					{
						Name:        "vpc_configuration_description_subnet_ids",
						Description: "The IDs of the subnets that Kinesis Data Firehose uses to create ENIs in the VPC of the Amazon ES destination",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("VpcConfigurationDescription.SubnetIds"),
					},
					{
						Name:        "vpc_configuration_description_vpc_id",
						Description: "The ID of the Amazon ES destination's VPC.  This member is required.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VpcConfigurationDescription.VpcId"),
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "aws_kinesis_firehose_open_search_destination_processing_configuration_processors",
						Description: "Describes a data processor.",
						Resolver:    fetchKinesisFirehoseOpenSearchDestinationProcessingConfigurationProcessors,
						Columns: []schema.Column{
							{
								Name:        "firehose_open_search_destination_cq_id",
								Description: "Unique CloudQuery ID of aws_kinesis_firehose_open_search_destination table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "type",
								Description: "The type of processor.  This member is required.",
								Type:        schema.TypeString,
							},
						},
						Relations: []*schema.Table{
							{
								Name:        "aws_kinesis_firehose_open_search_destination_processing_configuration_processor_parameters",
								Description: "Describes the processor parameter.",
								Resolver:    fetchKinesisFirehoseOpenSearchDestinationProcessingConfigurationProcessorParameters,
								Columns: []schema.Column{
									{
										Name:        "firehose_open_search_destination_processing_configuration_processor_cq_id",
										Description: "Unique CloudQuery ID of aws_kinesis_firehose_open_search_destination_processing_configuration_processors table (FK)",
										Type:        schema.TypeUUID,
										Resolver:    schema.ParentIdResolver,
									},
									{
										Name:        "parameter_name",
										Description: "The name of the parameter.  This member is required.",
										Type:        schema.TypeString,
									},
									{
										Name:        "parameter_value",
										Description: "The parameter value.  This member is required.",
										Type:        schema.TypeString,
									},
								},
							},
						},
					},
				},
			},
			{
				Name:        "aws_kinesis_firehose_elasticsearch_destination",
				Description: "The destination description in Amazon ES.",
				Resolver:    fetchKinesisFirehoseElasticsearchDestinations,
				Columns: []schema.Column{
					{
						Name:        "firehose_cq_id",
						Description: "Unique CloudQuery ID of aws_kinesis_firehoses table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "buffering_hints_interval_in_seconds",
						Description: "Buffer incoming data for the specified period of time, in seconds, before delivering it to the destination",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("BufferingHints.IntervalInSeconds"),
					},
					{
						Name:        "buffering_hints_size_in_mb_s",
						Description: "Buffer incoming data to the specified size, in MBs, before delivering it to the destination",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("BufferingHints.SizeInMBs"),
					},
					{
						Name:        "cloud_watch_logging_options_enabled",
						Description: "Enables or disables CloudWatch logging.",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("CloudWatchLoggingOptions.Enabled"),
					},
					{
						Name:        "cloud_watch_logging_options_log_group_name",
						Description: "The CloudWatch group name for logging",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("CloudWatchLoggingOptions.LogGroupName"),
					},
					{
						Name:        "cloud_watch_logging_options_log_stream_name",
						Description: "The CloudWatch log stream name for logging",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("CloudWatchLoggingOptions.LogStreamName"),
					},
					{
						Name:        "cluster_endpoint",
						Description: "The endpoint to use when communicating with the cluster",
						Type:        schema.TypeString,
					},
					{
						Name:        "domain_arn",
						Description: "The ARN of the Amazon ES domain",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DomainARN"),
					},
					{
						Name:        "index_name",
						Description: "The Elasticsearch index name.",
						Type:        schema.TypeString,
					},
					{
						Name:        "index_rotation_period",
						Description: "The Elasticsearch index rotation period",
						Type:        schema.TypeString,
					},
					{
						Name:        "processing_configuration_enabled",
						Description: "Enables or disables data processing.",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("ProcessingConfiguration.Enabled"),
					},
					{
						Name:        "retry_options_duration_in_seconds",
						Description: "After an initial failure to deliver to Amazon ES, the total amount of time during which Kinesis Data Firehose retries delivery (including the first attempt)",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("RetryOptions.DurationInSeconds"),
					},
					{
						Name:        "role_arn",
						Description: "The Amazon Resource Name (ARN) of the AWS credentials",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("RoleARN"),
					},
					{
						Name:        "s3_backup_mode",
						Description: "The Amazon S3 backup mode.",
						Type:        schema.TypeString,
					},
					{
						Name:        "s3_destination_description_bucket_arn",
						Description: "The ARN of the S3 bucket",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("S3DestinationDescription.BucketARN"),
					},
					{
						Name:        "s3_destination_description_buffering_hints_interval_in_seconds",
						Description: "Buffer incoming data for the specified period of time, in seconds, before delivering it to the destination",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("S3DestinationDescription.BufferingHints.IntervalInSeconds"),
					},
					{
						Name:        "s3_destination_description_buffering_hints_size_in_mb_s",
						Description: "Buffer incoming data to the specified size, in MiBs, before delivering it to the destination",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("S3DestinationDescription.BufferingHints.SizeInMBs"),
					},
					{
						Name:        "s3_destination_description_compression_format",
						Description: "The compression format",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("S3DestinationDescription.CompressionFormat"),
					},
					{
						Name:        "s3_destination_description_role_arn",
						Description: "The Amazon Resource Name (ARN) of the AWS credentials",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("S3DestinationDescription.RoleARN"),
					},
					{
						Name:        "s3_destination_description_cloud_watch_logging_options_enabled",
						Description: "Enables or disables CloudWatch logging.",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("S3DestinationDescription.CloudWatchLoggingOptions.Enabled"),
					},
					{
						Name:        "s3_destination_description_error_output_prefix",
						Description: "A prefix that Kinesis Data Firehose evaluates and adds to failed records before writing them to S3",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("S3DestinationDescription.ErrorOutputPrefix"),
					},
					{
						Name:        "s3_destination_description_prefix",
						Description: "The \"YYYY/MM/DD/HH\" time format prefix is automatically used for delivered Amazon S3 files",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("S3DestinationDescription.Prefix"),
					},
					{
						Name:        "type_name",
						Description: "The Elasticsearch type name",
						Type:        schema.TypeString,
					},
					{
						Name:        "vpc_configuration_description_role_arn",
						Description: "The ARN of the IAM role that the delivery stream uses to create endpoints in the destination VPC",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VpcConfigurationDescription.RoleARN"),
					},
					{
						Name:        "vpc_configuration_description_security_group_ids",
						Description: "The IDs of the security groups that Kinesis Data Firehose uses when it creates ENIs in the VPC of the Amazon ES destination",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("VpcConfigurationDescription.SecurityGroupIds"),
					},
					{
						Name:        "vpc_configuration_description_subnet_ids",
						Description: "The IDs of the subnets that Kinesis Data Firehose uses to create ENIs in the VPC of the Amazon ES destination",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("VpcConfigurationDescription.SubnetIds"),
					},
					{
						Name:        "vpc_configuration_description_vpc_id",
						Description: "The ID of the Amazon ES destination's VPC.  This member is required.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VpcConfigurationDescription.VpcId"),
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "aws_kinesis_firehose_elasticsearch_destination_processing_configuration_processors",
						Description: "Describes a data processor.",
						Resolver:    fetchKinesisFirehoseElasticsearchDestinationProcessingConfigurationProcessors,
						Columns: []schema.Column{
							{
								Name:        "firehose_elasticsearch_destination_cq_id",
								Description: "Unique CloudQuery ID of aws_kinesis_firehose_elasticsearch_destination table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "type",
								Description: "The type of processor.  This member is required.",
								Type:        schema.TypeString,
							},
						},
						Relations: []*schema.Table{
							{
								Name:        "aws_kinesis_firehose_elasticsearch_destination_processing_configuration_processor_parameters",
								Description: "Describes the processor parameter.",
								Resolver:    fetchKinesisFirehoseElasticsearchDestinationProcessingConfigurationProcessorParameters,
								Columns: []schema.Column{
									{
										Name:        "firehose_elasticsearch_destination_processing_configuration_processor_cq_id",
										Description: "Unique CloudQuery ID of aws_kinesis_firehose_elasticsearch_destination_processing_configuration_processors table (FK)",
										Type:        schema.TypeUUID,
										Resolver:    schema.ParentIdResolver,
									},
									{
										Name:        "parameter_name",
										Description: "The name of the parameter.  This member is required.",
										Type:        schema.TypeString,
									},
									{
										Name:        "parameter_value",
										Description: "The parameter value.  This member is required.",
										Type:        schema.TypeString,
									},
								},
							},
						},
					},
				},
			},
			{
				Name:        "aws_kinesis_firehose_extended_s3_destination",
				Description: "Describes a destination in Amazon S3.",
				Resolver:    fetchKinesisFirehoseExtendedS3Destinations,
				Columns: []schema.Column{
					{
						Name:        "firehose_cq_id",
						Description: "Unique CloudQuery ID of aws_kinesis_firehoses table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "bucket_arn",
						Description: "The ARN of the S3 bucket",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("BucketARN"),
					},
					{
						Name:        "buffering_hints_interval_in_seconds",
						Description: "Buffer incoming data for the specified period of time, in seconds, before delivering it to the destination",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("BufferingHints.IntervalInSeconds"),
					},
					{
						Name:        "buffering_hints_size_in_mb_s",
						Description: "Buffer incoming data to the specified size, in MiBs, before delivering it to the destination",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("BufferingHints.SizeInMBs"),
					},
					{
						Name:        "compression_format",
						Description: "The compression format",
						Type:        schema.TypeString,
					},
					{
						Name:        "encryption_configuration_kms_encryption_config_aws_kms_key_arn",
						Description: "The Amazon Resource Name (ARN) of the encryption key",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EncryptionConfiguration.KMSEncryptionConfig.AWSKMSKeyARN"),
					},
					{
						Name:        "encryption_configuration_no_encryption_config",
						Description: "Specifically override existing encryption information to ensure that no encryption is used.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EncryptionConfiguration.NoEncryptionConfig"),
					},
					{
						Name:        "role_arn",
						Description: "The Amazon Resource Name (ARN) of the AWS credentials",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("RoleARN"),
					},
					{
						Name:        "cloud_watch_logging_options_enabled",
						Description: "Enables or disables CloudWatch logging.",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("CloudWatchLoggingOptions.Enabled"),
					},
					{
						Name:        "cloud_watch_logging_options_log_group_name",
						Description: "The CloudWatch group name for logging",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("CloudWatchLoggingOptions.LogGroupName"),
					},
					{
						Name:        "cloud_watch_logging_options_log_stream_name",
						Description: "The CloudWatch log stream name for logging",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("CloudWatchLoggingOptions.LogStreamName"),
					},
					{
						Name:        "enabled",
						Description: "Defaults to true",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("DataFormatConversionConfiguration.Enabled"),
					},
					{
						Name:        "deserializer_hive_json_ser_de_timestamp_formats",
						Description: "Indicates how you want Kinesis Data Firehose to parse the date and timestamps that may be present in your input data JSON",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("DataFormatConversionConfiguration.InputFormatConfiguration.Deserializer.HiveJsonSerDe.TimestampFormats"),
					},
					{
						Name:        "deserializer_open_x_json_ser_de_case_insensitive",
						Description: "When set to true, which is the default, Kinesis Data Firehose converts JSON keys to lowercase before deserializing them.",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("DataFormatConversionConfiguration.InputFormatConfiguration.Deserializer.OpenXJsonSerDe.CaseInsensitive"),
					},
					{
						Name:        "deserializer_open_x_json_ser_de_column_to_json_key_mappings",
						Description: "Maps column names to JSON keys that aren't identical to the column names",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("DataFormatConversionConfiguration.InputFormatConfiguration.Deserializer.OpenXJsonSerDe.ColumnToJsonKeyMappings"),
					},
					{
						Name:        "serializer_orc_ser_de_block_size_bytes",
						Description: "The Hadoop Distributed File System (HDFS) block size",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("DataFormatConversionConfiguration.OutputFormatConfiguration.Serializer.OrcSerDe.BlockSizeBytes"),
					},
					{
						Name:        "serializer_orc_ser_de_bloom_filter_columns",
						Description: "The column names for which you want Kinesis Data Firehose to create bloom filters",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("DataFormatConversionConfiguration.OutputFormatConfiguration.Serializer.OrcSerDe.BloomFilterColumns"),
					},
					{
						Name:        "serializer_orc_ser_de_bloom_filter_false_positive_probability",
						Description: "The Bloom filter false positive probability (FPP)",
						Type:        schema.TypeFloat,
						Resolver:    schema.PathResolver("DataFormatConversionConfiguration.OutputFormatConfiguration.Serializer.OrcSerDe.BloomFilterFalsePositiveProbability"),
					},
					{
						Name:        "serializer_orc_ser_de_compression",
						Description: "The compression code to use over data blocks",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DataFormatConversionConfiguration.OutputFormatConfiguration.Serializer.OrcSerDe.Compression"),
					},
					{
						Name:        "serializer_orc_ser_de_dictionary_key_threshold",
						Description: "Represents the fraction of the total number of non-null rows",
						Type:        schema.TypeFloat,
						Resolver:    schema.PathResolver("DataFormatConversionConfiguration.OutputFormatConfiguration.Serializer.OrcSerDe.DictionaryKeyThreshold"),
					},
					{
						Name:        "serializer_orc_ser_de_enable_padding",
						Description: "Set this to true to indicate that you want stripes to be padded to the HDFS block boundaries",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("DataFormatConversionConfiguration.OutputFormatConfiguration.Serializer.OrcSerDe.EnablePadding"),
					},
					{
						Name:        "serializer_orc_ser_de_format_version",
						Description: "The version of the file to write",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DataFormatConversionConfiguration.OutputFormatConfiguration.Serializer.OrcSerDe.FormatVersion"),
					},
					{
						Name:        "serializer_orc_ser_de_padding_tolerance",
						Description: "A number between 0 and 1 that defines the tolerance for block padding as a decimal fraction of stripe size",
						Type:        schema.TypeFloat,
						Resolver:    schema.PathResolver("DataFormatConversionConfiguration.OutputFormatConfiguration.Serializer.OrcSerDe.PaddingTolerance"),
					},
					{
						Name:        "serializer_orc_ser_de_row_index_stride",
						Description: "The number of rows between index entries",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("DataFormatConversionConfiguration.OutputFormatConfiguration.Serializer.OrcSerDe.RowIndexStride"),
					},
					{
						Name:        "serializer_orc_ser_de_stripe_size_bytes",
						Description: "The number of bytes in each stripe",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("DataFormatConversionConfiguration.OutputFormatConfiguration.Serializer.OrcSerDe.StripeSizeBytes"),
					},
					{
						Name:        "serializer_parquet_ser_de_block_size_bytes",
						Description: "The Hadoop Distributed File System (HDFS) block size",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("DataFormatConversionConfiguration.OutputFormatConfiguration.Serializer.ParquetSerDe.BlockSizeBytes"),
					},
					{
						Name:        "serializer_parquet_ser_de_compression",
						Description: "The compression code to use over data blocks",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DataFormatConversionConfiguration.OutputFormatConfiguration.Serializer.ParquetSerDe.Compression"),
					},
					{
						Name:        "serializer_parquet_ser_de_enable_dictionary_compression",
						Description: "Indicates whether to enable dictionary compression.",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("DataFormatConversionConfiguration.OutputFormatConfiguration.Serializer.ParquetSerDe.EnableDictionaryCompression"),
					},
					{
						Name:        "serializer_parquet_ser_de_max_padding_bytes",
						Description: "The maximum amount of padding to apply",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("DataFormatConversionConfiguration.OutputFormatConfiguration.Serializer.ParquetSerDe.MaxPaddingBytes"),
					},
					{
						Name:        "serializer_parquet_ser_de_page_size_bytes",
						Description: "The Parquet page size",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("DataFormatConversionConfiguration.OutputFormatConfiguration.Serializer.ParquetSerDe.PageSizeBytes"),
					},
					{
						Name:        "serializer_parquet_ser_de_writer_version",
						Description: "Indicates the version of row format to output",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DataFormatConversionConfiguration.OutputFormatConfiguration.Serializer.ParquetSerDe.WriterVersion"),
					},
					{
						Name:        "schema_configuration_catalog_id",
						Description: "The ID of the AWS Glue Data Catalog",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DataFormatConversionConfiguration.SchemaConfiguration.CatalogId"),
					},
					{
						Name:        "schema_configuration_database_name",
						Description: "Specifies the name of the AWS Glue database that contains the schema for the output data",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DataFormatConversionConfiguration.SchemaConfiguration.DatabaseName"),
					},
					{
						Name:        "schema_configuration_region",
						Description: "If you don't specify an AWS Region, the default is the current Region.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DataFormatConversionConfiguration.SchemaConfiguration.Region"),
					},
					{
						Name:        "schema_configuration_role_arn",
						Description: "The role that Kinesis Data Firehose can use to access AWS Glue",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DataFormatConversionConfiguration.SchemaConfiguration.RoleARN"),
					},
					{
						Name:        "schema_configuration_table_name",
						Description: "Specifies the AWS Glue table that contains the column information that constitutes your data schema",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DataFormatConversionConfiguration.SchemaConfiguration.TableName"),
					},
					{
						Name:        "schema_configuration_version_id",
						Description: "Specifies the table version for the output data schema",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DataFormatConversionConfiguration.SchemaConfiguration.VersionId"),
					},
					{
						Name:        "dynamic_partitioning_configuration_enabled",
						Description: "Specifies that the dynamic partitioning is enabled for this Kinesis Data Firehose delivery stream.",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("DynamicPartitioningConfiguration.Enabled"),
					},
					{
						Name:        "error_output_prefix",
						Description: "A prefix that Kinesis Data Firehose evaluates and adds to failed records before writing them to S3",
						Type:        schema.TypeString,
					},
					{
						Name:        "prefix",
						Description: "The \"YYYY/MM/DD/HH\" time format prefix is automatically used for delivered Amazon S3 files",
						Type:        schema.TypeString,
					},
					{
						Name:        "processing_configuration_enabled",
						Description: "Enables or disables data processing.",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("ProcessingConfiguration.Enabled"),
					},
					{
						Name:        "bucket_arn",
						Description: "The ARN of the S3 bucket",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("S3BackupDescription.BucketARN"),
					},
					{
						Name:        "buffering_hints_interval_in_seconds",
						Description: "Buffer incoming data for the specified period of time, in seconds, before delivering it to the destination",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("S3BackupDescription.BufferingHints.IntervalInSeconds"),
					},
					{
						Name:        "buffering_hints_size_in_mb_s",
						Description: "Buffer incoming data to the specified size, in MiBs, before delivering it to the destination",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("S3BackupDescription.BufferingHints.SizeInMBs"),
					},
					{
						Name:        "compression_format",
						Description: "The compression format",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("S3BackupDescription.CompressionFormat"),
					},
					{
						Name:        "encryption_configuration_kms_encryption_config_aws_kms_key_arn",
						Description: "The Amazon Resource Name (ARN) of the encryption key",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("S3BackupDescription.EncryptionConfiguration.KMSEncryptionConfig.AWSKMSKeyARN"),
					},
					{
						Name:        "encryption_configuration_no_encryption_config",
						Description: "Specifically override existing encryption information to ensure that no encryption is used.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("S3BackupDescription.EncryptionConfiguration.NoEncryptionConfig"),
					},
					{
						Name:        "role_arn",
						Description: "The Amazon Resource Name (ARN) of the AWS credentials",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("S3BackupDescription.RoleARN"),
					},
					{
						Name:        "cloud_watch_logging_options_enabled",
						Description: "Enables or disables CloudWatch logging.",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("S3BackupDescription.CloudWatchLoggingOptions.Enabled"),
					},
					{
						Name:        "cloud_watch_logging_options_log_group_name",
						Description: "The CloudWatch group name for logging",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("S3BackupDescription.CloudWatchLoggingOptions.LogGroupName"),
					},
					{
						Name:        "cloud_watch_logging_options_log_stream_name",
						Description: "The CloudWatch log stream name for logging",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("S3BackupDescription.CloudWatchLoggingOptions.LogStreamName"),
					},
					{
						Name:        "error_output_prefix",
						Description: "A prefix that Kinesis Data Firehose evaluates and adds to failed records before writing them to S3",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("S3BackupDescription.ErrorOutputPrefix"),
					},
					{
						Name:        "prefix",
						Description: "The \"YYYY/MM/DD/HH\" time format prefix is automatically used for delivered Amazon S3 files",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("S3BackupDescription.Prefix"),
					},
					{
						Name:        "s3_backup_mode",
						Description: "The Amazon S3 backup mode.",
						Type:        schema.TypeString,
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "aws_kinesis_firehose_extended_s3_destination_processing_configuration_processors",
						Description: "Describes a data processor.",
						Resolver:    fetchKinesisFirehoseExtendedS3DestinationProcessingConfigurationProcessors,
						Columns: []schema.Column{
							{
								Name:        "firehose_extended_s3_destination_cq_id",
								Description: "Unique CloudQuery ID of aws_kinesis_firehose_extended_s3_destination table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "type",
								Description: "The type of processor.  This member is required.",
								Type:        schema.TypeString,
							},
						},
						Relations: []*schema.Table{
							{
								Name:        "aws_kinesis_firehose_extended_s3_destination_processing_configuration_processor_parameters",
								Description: "Describes the processor parameter.",
								Resolver:    fetchKinesisFirehoseExtendedS3DestinationProcessingConfigurationProcessorParameters,
								Columns: []schema.Column{
									{
										Name:        "firehose_extended_s3_destination_processing_configuration_processor_cq_id",
										Description: "Unique CloudQuery ID of aws_kinesis_firehose_extended_s3_destination_processing_configuration_processors table (FK)",
										Type:        schema.TypeUUID,
										Resolver:    schema.ParentIdResolver,
									},
									{
										Name:        "parameter_name",
										Description: "The name of the parameter.  This member is required.",
										Type:        schema.TypeString,
									},
									{
										Name:        "parameter_value",
										Description: "The parameter value.  This member is required.",
										Type:        schema.TypeString,
									},
								},
							},
						},
					},
				},
			},
			{
				Name:        "aws_kinesis_firehose_http_destination",
				Description: "Describes the HTTP endpoint destination.",
				Resolver:    fetchKinesisFirehoseHttpDestinations,
				Columns: []schema.Column{
					{
						Name:        "firehose_cq_id",
						Description: "Unique CloudQuery ID of aws_kinesis_firehoses table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "buffering_hints_interval_in_seconds",
						Description: "Buffer incoming data for the specified period of time, in seconds, before delivering it to the destination",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("BufferingHints.IntervalInSeconds"),
					},
					{
						Name:        "buffering_hints_size_in_mb_s",
						Description: "Buffer incoming data to the specified size, in MBs, before delivering it to the destination",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("BufferingHints.SizeInMBs"),
					},
					{
						Name:        "cloud_watch_logging_options_enabled",
						Description: "Enables or disables CloudWatch logging.",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("CloudWatchLoggingOptions.Enabled"),
					},
					{
						Name:        "cloud_watch_logging_options_log_group_name",
						Description: "The CloudWatch group name for logging",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("CloudWatchLoggingOptions.LogGroupName"),
					},
					{
						Name:        "cloud_watch_logging_options_log_stream_name",
						Description: "The CloudWatch log stream name for logging",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("CloudWatchLoggingOptions.LogStreamName"),
					},
					{
						Name:        "endpoint_configuration_name",
						Description: "The name of the HTTP endpoint selected as the destination.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EndpointConfiguration.Name"),
					},
					{
						Name:        "endpoint_configuration_url",
						Description: "The URL of the HTTP endpoint selected as the destination.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EndpointConfiguration.Url"),
					},
					{
						Name:        "processing_configuration_enabled",
						Description: "Enables or disables data processing.",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("ProcessingConfiguration.Enabled"),
					},
					{
						Name:        "request_configuration_content_encoding",
						Description: "Kinesis Data Firehose uses the content encoding to compress the body of a request before sending the request to the destination",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("RequestConfiguration.ContentEncoding"),
					},
					{
						Name:        "retry_options_duration_in_seconds",
						Description: "The total amount of time that Kinesis Data Firehose spends on retries",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("RetryOptions.DurationInSeconds"),
					},
					{
						Name:        "role_arn",
						Description: "Kinesis Data Firehose uses this IAM role for all the permissions that the delivery stream needs.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("RoleARN"),
					},
					{
						Name:        "s3_backup_mode",
						Description: "Describes the S3 bucket backup options for the data that Kinesis Firehose delivers to the HTTP endpoint destination",
						Type:        schema.TypeString,
					},
					{
						Name:        "bucket_arn",
						Description: "The ARN of the S3 bucket",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("S3DestinationDescription.BucketARN"),
					},
					{
						Name:        "buffering_hints_interval_in_seconds",
						Description: "Buffer incoming data for the specified period of time, in seconds, before delivering it to the destination",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("S3DestinationDescription.BufferingHints.IntervalInSeconds"),
					},
					{
						Name:        "buffering_hints_size_in_mb_s",
						Description: "Buffer incoming data to the specified size, in MiBs, before delivering it to the destination",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("S3DestinationDescription.BufferingHints.SizeInMBs"),
					},
					{
						Name:        "compression_format",
						Description: "The compression format",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("S3DestinationDescription.CompressionFormat"),
					},
					{
						Name:        "encryption_configuration_kms_encryption_config_aws_kms_key_arn",
						Description: "The Amazon Resource Name (ARN) of the encryption key",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("S3DestinationDescription.EncryptionConfiguration.KMSEncryptionConfig.AWSKMSKeyARN"),
					},
					{
						Name:        "encryption_configuration_no_encryption_config",
						Description: "Specifically override existing encryption information to ensure that no encryption is used.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("S3DestinationDescription.EncryptionConfiguration.NoEncryptionConfig"),
					},
					{
						Name:        "role_arn",
						Description: "The Amazon Resource Name (ARN) of the AWS credentials",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("S3DestinationDescription.RoleARN"),
					},
					{
						Name:        "cloud_watch_logging_options_enabled",
						Description: "Enables or disables CloudWatch logging.",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("S3DestinationDescription.CloudWatchLoggingOptions.Enabled"),
					},
					{
						Name:        "cloud_watch_logging_options_log_group_name",
						Description: "The CloudWatch group name for logging",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("S3DestinationDescription.CloudWatchLoggingOptions.LogGroupName"),
					},
					{
						Name:        "cloud_watch_logging_options_log_stream_name",
						Description: "The CloudWatch log stream name for logging",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("S3DestinationDescription.CloudWatchLoggingOptions.LogStreamName"),
					},
					{
						Name:        "error_output_prefix",
						Description: "A prefix that Kinesis Data Firehose evaluates and adds to failed records before writing them to S3",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("S3DestinationDescription.ErrorOutputPrefix"),
					},
					{
						Name:        "prefix",
						Description: "The \"YYYY/MM/DD/HH\" time format prefix is automatically used for delivered Amazon S3 files",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("S3DestinationDescription.Prefix"),
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "aws_kinesis_firehose_http_destination_processing_configuration_processors",
						Description: "Describes a data processor.",
						Resolver:    fetchKinesisFirehoseHttpDestinationProcessingConfigurationProcessors,
						Columns: []schema.Column{
							{
								Name:        "firehose_http_destination_cq_id",
								Description: "Unique CloudQuery ID of aws_kinesis_firehose_http_destination table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "type",
								Description: "The type of processor.  This member is required.",
								Type:        schema.TypeString,
							},
						},
						Relations: []*schema.Table{
							{
								Name:        "aws_kinesis_firehose_http_destination_processing_configuration_processor_parameters",
								Description: "Describes the processor parameter.",
								Resolver:    fetchKinesisFirehoseHttpDestinationProcessingConfigurationProcessorParameters,
								Columns: []schema.Column{
									{
										Name:        "firehose_http_destination_processing_configuration_processor_cq_id",
										Description: "Unique CloudQuery ID of aws_kinesis_firehose_http_destination_processing_configuration_processors table (FK)",
										Type:        schema.TypeUUID,
										Resolver:    schema.ParentIdResolver,
									},
									{
										Name:        "parameter_name",
										Description: "The name of the parameter.  This member is required.",
										Type:        schema.TypeString,
									},
									{
										Name:        "parameter_value",
										Description: "The parameter value.  This member is required.",
										Type:        schema.TypeString,
									},
								},
							},
						},
					},
					{
						Name:        "aws_kinesis_firehose_http_destination_request_configuration_common_attributes",
						Description: "Describes the metadata that's delivered to the specified HTTP endpoint destination.",
						Resolver:    fetchKinesisFirehoseHttpDestinationRequestConfigurationCommonAttributes,
						Columns: []schema.Column{
							{
								Name:        "firehose_http_destination_cq_id",
								Description: "Unique CloudQuery ID of aws_kinesis_firehose_http_destination table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "attribute_name",
								Description: "The name of the HTTP endpoint common attribute.  This member is required.",
								Type:        schema.TypeString,
							},
							{
								Name:        "attribute_value",
								Description: "The value of the HTTP endpoint common attribute.  This member is required.",
								Type:        schema.TypeString,
							},
						},
					},
				},
			},
			{
				Name:        "aws_kinesis_firehose_redshift_destination",
				Description: "Describes a destination in Amazon Redshift.",
				Resolver:    fetchKinesisFirehoseRedshiftDestinations,
				Columns: []schema.Column{
					{
						Name:        "firehose_cq_id",
						Description: "Unique CloudQuery ID of aws_kinesis_firehoses table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "cluster_j_db_c_url",
						Description: "The database connection string.  This member is required.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ClusterJDBCURL"),
					},
					{
						Name:        "copy_command_data_table_name",
						Description: "The name of the target table",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("CopyCommand.DataTableName"),
					},
					{
						Name:        "copy_command_copy_options",
						Description: "Optional parameters to use with the Amazon Redshift COPY command",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("CopyCommand.CopyOptions"),
					},
					{
						Name:        "copy_command_data_table_columns",
						Description: "A comma-separated list of column names.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("CopyCommand.DataTableColumns"),
					},
					{
						Name:        "role_arn",
						Description: "The Amazon Resource Name (ARN) of the AWS credentials",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("RoleARN"),
					},
					{
						Name:        "bucket_arn",
						Description: "The ARN of the S3 bucket",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("S3DestinationDescription.BucketARN"),
					},
					{
						Name:        "buffering_hints_interval_in_seconds",
						Description: "Buffer incoming data for the specified period of time, in seconds, before delivering it to the destination",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("S3DestinationDescription.BufferingHints.IntervalInSeconds"),
					},
					{
						Name:        "buffering_hints_size_in_mb_s",
						Description: "Buffer incoming data to the specified size, in MiBs, before delivering it to the destination",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("S3DestinationDescription.BufferingHints.SizeInMBs"),
					},
					{
						Name:        "compression_format",
						Description: "The compression format",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("S3DestinationDescription.CompressionFormat"),
					},
					{
						Name:        "kms_encryption_config_aws_kms_key_arn",
						Description: "The Amazon Resource Name (ARN) of the encryption key",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("S3DestinationDescription.EncryptionConfiguration.KMSEncryptionConfig.AWSKMSKeyARN"),
					},
					{
						Name:        "no_encryption_config",
						Description: "Specifically override existing encryption information to ensure that no encryption is used.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("S3DestinationDescription.EncryptionConfiguration.NoEncryptionConfig"),
					},
					{
						Name:        "role_arn",
						Description: "The Amazon Resource Name (ARN) of the AWS credentials",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("S3DestinationDescription.RoleARN"),
					},
					{
						Name:        "cloud_watch_logging_options_enabled",
						Description: "Enables or disables CloudWatch logging.",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("S3DestinationDescription.CloudWatchLoggingOptions.Enabled"),
					},
					{
						Name:        "cloud_watch_logging_options_log_group_name",
						Description: "The CloudWatch group name for logging",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("S3DestinationDescription.CloudWatchLoggingOptions.LogGroupName"),
					},
					{
						Name:        "cloud_watch_logging_options_log_stream_name",
						Description: "The CloudWatch log stream name for logging",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("S3DestinationDescription.CloudWatchLoggingOptions.LogStreamName"),
					},
					{
						Name:        "error_output_prefix",
						Description: "A prefix that Kinesis Data Firehose evaluates and adds to failed records before writing them to S3",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("S3DestinationDescription.ErrorOutputPrefix"),
					},
					{
						Name:        "prefix",
						Description: "The \"YYYY/MM/DD/HH\" time format prefix is automatically used for delivered Amazon S3 files",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("S3DestinationDescription.Prefix"),
					},
					{
						Name:        "username",
						Description: "The name of the user.  This member is required.",
						Type:        schema.TypeString,
					},
					{
						Name:        "cloud_watch_logging_options_enabled",
						Description: "Enables or disables CloudWatch logging.",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("CloudWatchLoggingOptions.Enabled"),
					},
					{
						Name:        "cloud_watch_logging_options_log_group_name",
						Description: "The CloudWatch group name for logging",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("CloudWatchLoggingOptions.LogGroupName"),
					},
					{
						Name:        "cloud_watch_logging_options_log_stream_name",
						Description: "The CloudWatch log stream name for logging",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("CloudWatchLoggingOptions.LogStreamName"),
					},
					{
						Name:        "processing_configuration_enabled",
						Description: "Enables or disables data processing.",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("ProcessingConfiguration.Enabled"),
					},
					{
						Name:        "retry_options_duration_in_seconds",
						Description: "The length of time during which Kinesis Data Firehose retries delivery after a failure, starting from the initial request and including the first attempt",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("RetryOptions.DurationInSeconds"),
					},
					{
						Name:        "bucket_arn",
						Description: "The ARN of the S3 bucket",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("S3BackupDescription.BucketARN"),
					},
					{
						Name:        "buffering_hints_interval_in_seconds",
						Description: "Buffer incoming data for the specified period of time, in seconds, before delivering it to the destination",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("S3BackupDescription.BufferingHints.IntervalInSeconds"),
					},
					{
						Name:        "buffering_hints_size_in_mb_s",
						Description: "Buffer incoming data to the specified size, in MiBs, before delivering it to the destination",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("S3BackupDescription.BufferingHints.SizeInMBs"),
					},
					{
						Name:        "compression_format",
						Description: "The compression format",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("S3BackupDescription.CompressionFormat"),
					},
					{
						Name:        "kms_encryption_config_aws_kms_key_arn",
						Description: "The Amazon Resource Name (ARN) of the encryption key",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("S3BackupDescription.EncryptionConfiguration.KMSEncryptionConfig.AWSKMSKeyARN"),
					},
					{
						Name:        "no_encryption_config",
						Description: "Specifically override existing encryption information to ensure that no encryption is used.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("S3BackupDescription.EncryptionConfiguration.NoEncryptionConfig"),
					},
					{
						Name:        "role_arn",
						Description: "The Amazon Resource Name (ARN) of the AWS credentials",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("S3BackupDescription.RoleARN"),
					},
					{
						Name:        "cloud_watch_logging_options_enabled",
						Description: "Enables or disables CloudWatch logging.",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("S3BackupDescription.CloudWatchLoggingOptions.Enabled"),
					},
					{
						Name:        "cloud_watch_logging_options_log_group_name",
						Description: "The CloudWatch group name for logging",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("S3BackupDescription.CloudWatchLoggingOptions.LogGroupName"),
					},
					{
						Name:        "cloud_watch_logging_options_log_stream_name",
						Description: "The CloudWatch log stream name for logging",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("S3BackupDescription.CloudWatchLoggingOptions.LogStreamName"),
					},
					{
						Name:        "error_output_prefix",
						Description: "A prefix that Kinesis Data Firehose evaluates and adds to failed records before writing them to S3",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("S3BackupDescription.ErrorOutputPrefix"),
					},
					{
						Name:        "prefix",
						Description: "The \"YYYY/MM/DD/HH\" time format prefix is automatically used for delivered Amazon S3 files",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("S3BackupDescription.Prefix"),
					},
					{
						Name:        "s3_backup_mode",
						Description: "The Amazon S3 backup mode.",
						Type:        schema.TypeString,
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "aws_kinesis_firehose_redshift_destination_processing_configuration_processors",
						Description: "Describes a data processor.",
						Resolver:    fetchKinesisFirehoseRedshiftDestinationProcessingConfigurationProcessors,
						Columns: []schema.Column{
							{
								Name:        "firehose_redshift_destination_cq_id",
								Description: "Unique CloudQuery ID of aws_kinesis_firehose_redshift_destination table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "type",
								Description: "The type of processor.  This member is required.",
								Type:        schema.TypeString,
							},
						},
						Relations: []*schema.Table{
							{
								Name:        "aws_kinesis_firehose_redshift_destination_processing_configuration_processor_parameters",
								Description: "Describes the processor parameter.",
								Resolver:    fetchKinesisFirehoseRedshiftDestinationProcessingConfigurationProcessorParameters,
								Columns: []schema.Column{
									{
										Name:        "firehose_redshift_destination_processing_configuration_processor_cq_id",
										Description: "Unique CloudQuery ID of aws_kinesis_firehose_redshift_destination_processing_configuration_processors table (FK)",
										Type:        schema.TypeUUID,
										Resolver:    schema.ParentIdResolver,
									},
									{
										Name:        "parameter_name",
										Description: "The name of the parameter.  This member is required.",
										Type:        schema.TypeString,
									},
									{
										Name:        "parameter_value",
										Description: "The parameter value.  This member is required.",
										Type:        schema.TypeString,
									},
								},
							},
						},
					},
				},
			},
			{
				Name:        "aws_kinesis_firehose_splunk_destination",
				Description: "Describes a destination in Splunk.",
				Resolver:    fetchKinesisFirehoseSplunkDestinations,
				Columns: []schema.Column{
					{
						Name:        "firehose_cq_id",
						Description: "Unique CloudQuery ID of aws_kinesis_firehoses table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "cloud_watch_logging_options_enabled",
						Description: "Enables or disables CloudWatch logging.",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("CloudWatchLoggingOptions.Enabled"),
					},
					{
						Name:        "cloud_watch_logging_options_log_group_name",
						Description: "The CloudWatch group name for logging",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("CloudWatchLoggingOptions.LogGroupName"),
					},
					{
						Name:        "cloud_watch_logging_options_log_stream_name",
						Description: "The CloudWatch log stream name for logging",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("CloudWatchLoggingOptions.LogStreamName"),
					},
					{
						Name:        "h_e_ca_cknowledgment_timeout_in_seconds",
						Description: "The amount of time that Kinesis Data Firehose waits to receive an acknowledgment from Splunk after it sends it data",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("HECAcknowledgmentTimeoutInSeconds"),
					},
					{
						Name:        "h_e_c_endpoint",
						Description: "The HTTP Event Collector (HEC) endpoint to which Kinesis Data Firehose sends your data.",
						Type:        schema.TypeString,
					},
					{
						Name:        "h_e_c_endpoint_type",
						Description: "This type can be either \"Raw\" or \"Event.\"",
						Type:        schema.TypeString,
					},
					{
						Name:        "h_e_c_token",
						Description: "A GUID you obtain from your Splunk cluster when you create a new HEC endpoint.",
						Type:        schema.TypeString,
					},
					{
						Name:        "processing_configuration_enabled",
						Description: "Enables or disables data processing.",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("ProcessingConfiguration.Enabled"),
					},
					{
						Name:        "retry_options_duration_in_seconds",
						Description: "The total amount of time that Kinesis Data Firehose spends on retries",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("RetryOptions.DurationInSeconds"),
					},
					{
						Name:        "s3_backup_mode",
						Description: "Defines how documents should be delivered to Amazon S3",
						Type:        schema.TypeString,
					},
					{
						Name:        "bucket_arn",
						Description: "The ARN of the S3 bucket",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("S3DestinationDescription.BucketARN"),
					},
					{
						Name:        "buffering_hints_interval_in_seconds",
						Description: "Buffer incoming data for the specified period of time, in seconds, before delivering it to the destination",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("S3DestinationDescription.BufferingHints.IntervalInSeconds"),
					},
					{
						Name:        "buffering_hints_size_in_mb_s",
						Description: "Buffer incoming data to the specified size, in MiBs, before delivering it to the destination",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("S3DestinationDescription.BufferingHints.SizeInMBs"),
					},
					{
						Name:        "compression_format",
						Description: "The compression format",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("S3DestinationDescription.CompressionFormat"),
					},
					{
						Name:        "encryption_configuration_kms_encryption_config_aws_kms_key_arn",
						Description: "The Amazon Resource Name (ARN) of the encryption key",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("S3DestinationDescription.EncryptionConfiguration.KMSEncryptionConfig.AWSKMSKeyARN"),
					},
					{
						Name:        "encryption_configuration_no_encryption_config",
						Description: "Specifically override existing encryption information to ensure that no encryption is used.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("S3DestinationDescription.EncryptionConfiguration.NoEncryptionConfig"),
					},
					{
						Name:        "role_arn",
						Description: "The Amazon Resource Name (ARN) of the AWS credentials",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("S3DestinationDescription.RoleARN"),
					},
					{
						Name:        "cloud_watch_logging_options_enabled",
						Description: "Enables or disables CloudWatch logging.",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("S3DestinationDescription.CloudWatchLoggingOptions.Enabled"),
					},
					{
						Name:        "cloud_watch_logging_options_log_group_name",
						Description: "The CloudWatch group name for logging",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("S3DestinationDescription.CloudWatchLoggingOptions.LogGroupName"),
					},
					{
						Name:        "cloud_watch_logging_options_log_stream_name",
						Description: "The CloudWatch log stream name for logging",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("S3DestinationDescription.CloudWatchLoggingOptions.LogStreamName"),
					},
					{
						Name:        "error_output_prefix",
						Description: "A prefix that Kinesis Data Firehose evaluates and adds to failed records before writing them to S3",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("S3DestinationDescription.ErrorOutputPrefix"),
					},
					{
						Name:        "prefix",
						Description: "The \"YYYY/MM/DD/HH\" time format prefix is automatically used for delivered Amazon S3 files",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("S3DestinationDescription.Prefix"),
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "aws_kinesis_firehose_splunk_destination_processing_configuration_processors",
						Description: "Describes a data processor.",
						Resolver:    fetchKinesisFirehoseSplunkDestinationProcessingConfigurationProcessors,
						Columns: []schema.Column{
							{
								Name:        "firehose_splunk_destination_cq_id",
								Description: "Unique CloudQuery ID of aws_kinesis_firehose_splunk_destination table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "type",
								Description: "The type of processor.  This member is required.",
								Type:        schema.TypeString,
							},
						},
						Relations: []*schema.Table{
							{
								Name:        "aws_kinesis_firehose_splunk_destination_processing_configuration_processor_parameters",
								Description: "Describes the processor parameter.",
								Resolver:    fetchKinesisFirehoseSplunkDestinationProcessingConfigurationProcessorParameters,
								Columns: []schema.Column{
									{
										Name:        "firehose_splunk_destination_processing_configuration_processor_cq_id",
										Description: "Unique CloudQuery ID of aws_kinesis_firehose_splunk_destination_processing_configuration_processors table (FK)",
										Type:        schema.TypeUUID,
										Resolver:    schema.ParentIdResolver,
									},
									{
										Name:        "parameter_name",
										Description: "The name of the parameter.  This member is required.",
										Type:        schema.TypeString,
									},
									{
										Name:        "parameter_value",
										Description: "The parameter value.  This member is required.",
										Type:        schema.TypeString,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchKinesisFirehoses(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
func ResolveKinesisFirehoseTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	panic("not implemented")
}
func fetchKinesisFirehoseOpenSearchDestinations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
func fetchKinesisFirehoseOpenSearchDestinationProcessingConfigurationProcessors(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
func fetchKinesisFirehoseOpenSearchDestinationProcessingConfigurationProcessorParameters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
func fetchKinesisFirehoseElasticsearchDestinations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
func fetchKinesisFirehoseElasticsearchDestinationProcessingConfigurationProcessors(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
func fetchKinesisFirehoseElasticsearchDestinationProcessingConfigurationProcessorParameters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
func fetchKinesisFirehoseExtendedS3Destinations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
func fetchKinesisFirehoseExtendedS3DestinationProcessingConfigurationProcessors(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
func fetchKinesisFirehoseExtendedS3DestinationProcessingConfigurationProcessorParameters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
func fetchKinesisFirehoseHttpDestinations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
func fetchKinesisFirehoseHttpDestinationProcessingConfigurationProcessors(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
func fetchKinesisFirehoseHttpDestinationProcessingConfigurationProcessorParameters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
func fetchKinesisFirehoseHttpDestinationRequestConfigurationCommonAttributes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
func fetchKinesisFirehoseRedshiftDestinations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
func fetchKinesisFirehoseRedshiftDestinationProcessingConfigurationProcessors(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
func fetchKinesisFirehoseRedshiftDestinationProcessingConfigurationProcessorParameters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
func fetchKinesisFirehoseSplunkDestinations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
func fetchKinesisFirehoseSplunkDestinationProcessingConfigurationProcessors(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
func fetchKinesisFirehoseSplunkDestinationProcessingConfigurationProcessorParameters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}

// ====================================================================================================================
//                                                  User Defined Helpers
// ====================================================================================================================

func resolveKinesisFirehoseTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	panic("not implemented")
}
func fetchKinesisFirehoseDestinations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
