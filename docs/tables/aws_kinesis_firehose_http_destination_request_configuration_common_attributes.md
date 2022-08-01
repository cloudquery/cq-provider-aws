
# Table: aws_kinesis_firehose_http_destination_request_configuration_common_attributes
Describes the metadata that's delivered to the specified HTTP endpoint destination.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|firehose_http_destination_cq_id|uuid|Unique CloudQuery ID of aws_kinesis_firehose_http_destination table (FK)|
|attribute_name|text|The name of the HTTP endpoint common attribute.  This member is required.|
|attribute_value|text|The value of the HTTP endpoint common attribute.  This member is required.|
