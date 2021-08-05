
# Table: aws_iot_thing_groups

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|things_in_group|text[]|Lists the things in the specified group|
|policies|text[]||
|tags|jsonb||
|index_name|text|The dynamic thing group index name.|
|query_string|text|The dynamic thing group search query string.|
|query_version|text|The dynamic thing group query version.|
|status|text|The dynamic thing group status.|
|arn|text|The thing group ARN.|
|id|text|The thing group ID.|
|creation_date|timestamp without time zone||
|parent_group_name|text||
|root_to_parent_thing_groups|jsonb||
|name|text|The name of the thing group.|
|attribute_payload_attributes|jsonb||
|attribute_payload_merge|boolean||
|thing_group_description|text||
|version|bigint|The version of the thing group.|
