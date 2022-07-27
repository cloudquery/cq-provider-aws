
# Table: aws_cloudwatchlogs_log_group_streams
The Log Streams connected to a particular Log Group, each of which is a sequence of log events from a single emitter of logs
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|log_group_cq_id|uuid|Unique CloudQuery ID of aws_cloudwatchlogs_log_groups table (FK)|
|arn|text|The ARN (Amazon Resource Name) for the stream.|
|creation_time|bigint|The creation time of the log group, expressed as the number of milliseconds after Jan 1, 1970 00:00:00 UTC.|
|first_event_timestamp|bigint|The time of the first event, expressed as the number of milliseconds after Jan 1, 1970 00:00:00 UTC.|
|last_event_timestamp|bigint|The time of the most recent log event in the log stream in CloudWatch Logs. This number is expressed as the number of milliseconds after Jan 1, 1970 00:00:00 UTC. |
|last_ingestion_time|bigint|The ingestion time, expressed as the number of milliseconds after Jan 1, 1970 00:00:00 UTC.|
|log_stream_name|text|The name of the CloudWatch metric.|
|stored_bytes|bigint|The number of bytes stored.|
