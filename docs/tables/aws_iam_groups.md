
# Table: aws_iam_groups
Contains information about an IAM group entity
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|policies|jsonb|List of policies attached to group.|
|arn|text|The Amazon Resource Name (ARN) specifying the group|
|create_date|timestamp without time zone|The date and time, in ISO 8601 date-time format (http://wwwisoorg/iso/iso8601), when the group was created|
|id|text|The stable and unique string identifying the group|
|name|text|The friendly name that identifies the group|
|path|text|The path to the group|
