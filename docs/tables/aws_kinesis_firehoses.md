
# Table: aws_kinesis_firehoses
Contains information about a delivery stream.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|tags|jsonb||
|delivery_stream_arn|text|The Amazon Resource Name (ARN) of the delivery stream|
|delivery_stream_name|text|The name of the delivery stream.  This member is required.|
|delivery_stream_status|text|The status of the delivery stream|
|delivery_stream_type|text|The delivery stream type|
|has_more_destinations|boolean|Indicates whether there are more destinations available to list.  This member is required.|
|version_id|text|Each time the destination is updated for a delivery stream, the version ID is changed, and the current version ID is required when updating the destination. This is so that the service knows it is applying the changes to the correct version of the delivery stream.  This member is required.|
|create_timestamp|timestamp without time zone|The date and time that the delivery stream was created.|
|failure_description_details|text|A message providing details about the error that caused the failure.  This member is required.|
|failure_description_type|text|The type of error that caused the failure.  This member is required.|
|key_arn|text|If KeyType is CUSTOMER_MANAGED_CMK, this field contains the ARN of the customer managed CMK|
|key_type|text|Indicates the type of customer master key (CMK) that is used for encryption|
|status|text|This is the server-side encryption (SSE) status for the delivery stream|
|failure_description_details|text|A message providing details about the error that caused the failure.  This member is required.|
|failure_description_type|text|The type of error that caused the failure.  This member is required.|
|last_update_timestamp|timestamp without time zone|The date and time that the delivery stream was last updated.|
|source_kinesis_stream_source_description_kinesis_stream_arn|text|The Amazon Resource Name (ARN) of the source Kinesis data stream|
|source_kinesis_stream_source_description_role_arn|text|The ARN of the role used by the source Kinesis data stream|
