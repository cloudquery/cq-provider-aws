
# Table: aws_iam_groups_access_advisor_entities
An object that contains details about when the IAM entities (users or roles) were last used in an attempt to access the specified AWS service
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|groups_access_advisor_cq_id|uuid|Unique CloudQuery ID of aws_iam_groups_access_advisor table (FK)|
|arn|text|The Amazon Resource Name (ARN)|
|id|text|The identifier of the entity (user or role)|
|name|text|The name of the entity (user or role)|
|type|text|The type of entity (user or role)|
|path|text|The path to the entity (user or role)|
|last_authenticated|timestamp without time zone|The date and time, in ISO 8601 date-time format (http://wwwisoorg/iso/iso8601), when the authenticated entity last attempted to access AWS|
