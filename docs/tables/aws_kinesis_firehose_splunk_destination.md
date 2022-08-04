
# Table: aws_kinesis_firehose_splunk_destination
Describes a destination in Splunk
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|firehose_cq_id|uuid|Unique CloudQuery ID of aws_kinesis_firehoses table (FK)|
|processing_configuration_processors|jsonb|The Amazon Resource Name (ARN) of the delivery stream|
|cloud_watch_logging_options_enabled|boolean|Enables or disables CloudWatch logging|
|cloud_watch_logging_options_log_group_name|text|The CloudWatch group name for logging|
|cloud_watch_logging_options_log_stream_name|text|The CloudWatch log stream name for logging|
|h_e_ca_cknowledgment_timeout_in_seconds|bigint|The amount of time that Kinesis Data Firehose waits to receive an acknowledgment from Splunk after it sends it data|
|h_e_c_endpoint|text|The HTTP Event Collector (HEC) endpoint to which Kinesis Data Firehose sends your data|
|h_e_c_endpoint_type|text|This type can be either "Raw" or "Event"|
|h_e_c_token|text|A GUID you obtain from your Splunk cluster when you create a new HEC endpoint|
|processing_configuration_enabled|boolean|Enables or disables data processing|
|retry_options_duration_in_seconds|bigint|The total amount of time that Kinesis Data Firehose spends on retries|
|s3_backup_mode|text|Defines how documents should be delivered to Amazon S3|
|bucket_arn|text|The ARN of the S3 bucket|
|buffering_hints_interval_in_seconds|bigint|Buffer incoming data for the specified period of time, in seconds, before delivering it to the destination|
|buffering_hints_size_in_mb_s|bigint|Buffer incoming data to the specified size, in MiBs, before delivering it to the destination|
|compression_format|text|The compression format|
|encryption_configuration_kms_encryption_config_aws_kms_key_arn|text|The Amazon Resource Name (ARN) of the encryption key|
|encryption_configuration_no_encryption_config|text|Specifically override existing encryption information to ensure that no encryption is used|
|role_arn|text|The Amazon Resource Name (ARN) of the AWS credentials|
|cloud_watch_logging_options_enabled|boolean|Enables or disables CloudWatch logging|
|cloud_watch_logging_options_log_group_name|text|The CloudWatch group name for logging|
|cloud_watch_logging_options_log_stream_name|text|The CloudWatch log stream name for logging|
|error_output_prefix|text|A prefix that Kinesis Data Firehose evaluates and adds to failed records before writing them to S3|
|prefix|text|The "YYYY/MM/DD/HH" time format prefix is automatically used for delivered Amazon S3 files|
