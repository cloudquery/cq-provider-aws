
# Table: aws_sso_admin_instance_permission_sets
An entity that contains IAM policies.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|instance_cq_id|uuid|Unique CloudQuery ID of aws_sso_admin_instances table (FK)|
|inline_policy|jsonb||
|created_date|timestamp without time zone|The date that the permission set was created.|
|description|text|The description of the PermissionSet.|
|name|text|The name of the permission set.|
|arn|text|The ARN of the permission set|
|relay_state|text|Used to redirect users within the application during the federation authentication process.|
|session_duration|text|The length of time that the application user sessions are valid for in the ISO-8601 standard.|
