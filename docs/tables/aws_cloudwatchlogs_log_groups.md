
# Table: aws_cloudwatchlogs_log_groups
CloudWatch Logs enables you to centralize the logs from all of your systems, applications, and AWS services that you use, in a single, highly scalable service. Log groups define groups of log streams that share the same retention, monitoring, and access control settings. Each log stream has to belong to one log group. 
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|arn|text|The ARN (Amazon Resource Name) for the distribution|
|creation_time|bigint|The creation time of the log group, expressed as the number of milliseconds after Jan 1, 1970 00:00:00 UTC.|
|kms_key_id|text|The Amazon Resource Name (ARN) of the CMK to use when encrypting log data.|
|name|text|The name of the log group.|
|metric_filter_count|bigint|The number of metric filters|
|retention_in_days|bigint|The number of days to retain the log events in the specified log group. Possible values are: 1, 3, 5, 7, 14, 30, 60, 90, 120, 150, 180, 365, 400, 545, 731, 1827, 2192, 2557, 2922, 3288, and 3653.|
|stored_bytes|bigint|The number of bytes stored|
