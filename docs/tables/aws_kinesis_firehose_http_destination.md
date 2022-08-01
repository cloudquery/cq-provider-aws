
# Table: aws_kinesis_firehose_http_destination
Describes the HTTP endpoint destination.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|firehose_cq_id|uuid|Unique CloudQuery ID of aws_kinesis_firehoses table (FK)|
|buffering_hints_interval_in_seconds|bigint|Buffer incoming data for the specified period of time, in seconds, before delivering it to the destination|
|buffering_hints_size_in_mb_s|bigint|Buffer incoming data to the specified size, in MBs, before delivering it to the destination|
|cloud_watch_logging_options_enabled|boolean|Enables or disables CloudWatch logging.|
|cloud_watch_logging_options_log_group_name|text|The CloudWatch group name for logging|
|cloud_watch_logging_options_log_stream_name|text|The CloudWatch log stream name for logging|
|endpoint_configuration_name|text|The name of the HTTP endpoint selected as the destination.|
|endpoint_configuration_url|text|The URL of the HTTP endpoint selected as the destination.|
|processing_configuration_enabled|boolean|Enables or disables data processing.|
|request_configuration_content_encoding|text|Kinesis Data Firehose uses the content encoding to compress the body of a request before sending the request to the destination|
|retry_options_duration_in_seconds|bigint|The total amount of time that Kinesis Data Firehose spends on retries|
|role_arn|text|Kinesis Data Firehose uses this IAM role for all the permissions that the delivery stream needs.|
|s3_backup_mode|text|Describes the S3 bucket backup options for the data that Kinesis Firehose delivers to the HTTP endpoint destination|
|bucket_arn|text|The ARN of the S3 bucket|
|buffering_hints_interval_in_seconds|bigint|Buffer incoming data for the specified period of time, in seconds, before delivering it to the destination|
|buffering_hints_size_in_mb_s|bigint|Buffer incoming data to the specified size, in MiBs, before delivering it to the destination|
|compression_format|text|The compression format|
|encryption_configuration_kms_encryption_config_aws_kms_key_arn|text|The Amazon Resource Name (ARN) of the encryption key|
|encryption_configuration_no_encryption_config|text|Specifically override existing encryption information to ensure that no encryption is used.|
|role_arn|text|The Amazon Resource Name (ARN) of the AWS credentials|
|cloud_watch_logging_options_enabled|boolean|Enables or disables CloudWatch logging.|
|cloud_watch_logging_options_log_group_name|text|The CloudWatch group name for logging|
|cloud_watch_logging_options_log_stream_name|text|The CloudWatch log stream name for logging|
|error_output_prefix|text|A prefix that Kinesis Data Firehose evaluates and adds to failed records before writing them to S3|
|prefix|text|The "YYYY/MM/DD/HH" time format prefix is automatically used for delivered Amazon S3 files|
