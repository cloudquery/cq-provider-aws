
# Table: aws_sso_admin_instance_permission_set_account_assignments
The assignment that indicates a principal's limited access to a specified Amazon Web Services account with a specified permission set
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|instance_permission_set_cq_id|uuid|Unique CloudQuery ID of aws_sso_admin_instance_permission_sets table (FK)|
|account_id|text|The identifier of the Amazon Web Services account.|
|permission_set_arn|text|The ARN of the permission set|
|principal_id|text|An identifier for an object in Amazon Web Services SSO, such as a user or group. PrincipalIds are GUIDs (For example, f81d4fae-7dec-11d0-a765-00a0c91e6bf6)|
|principal_type|text|The entity type for which the assignment will be created.|
