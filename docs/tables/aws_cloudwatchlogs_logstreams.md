
# Table: aws_cloudwatchlogs_logstreams
Represents a log stream, which is a sequence of log events from a single emitter of logs.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|loggroup_cq_id|uuid|Unique CloudQuery ID of aws_cloudwatchlogs_loggroups table (FK)|
|arn|text|The Amazon Resource Name (ARN) of the log stream.|
|creation_time|bigint|The creation time of the stream, expressed as the number of milliseconds after Jan 1, 1970 00:00:00 UTC.|
|first_event_timestamp|bigint|The time of the first event, expressed as the number of milliseconds after Jan 1, 1970 00:00:00 UTC.|
|last_event_timestamp|bigint|The time of the most recent log event in the log stream in CloudWatch Logs|
|last_ingestion_time|bigint|The ingestion time, expressed as the number of milliseconds after Jan 1, 1970 00:00:00 UTC.|
|log_stream_name|text|The name of the log stream.|
|stored_bytes|bigint|The number of bytes stored|
|upload_sequence_token|text|The sequence token.|
