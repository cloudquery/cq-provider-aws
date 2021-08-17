
# Table: aws_sso_admin_instance_users
A user object, which contains a specified user’s metadata and attributes.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|instance_cq_id|uuid|Unique CloudQuery ID of aws_sso_admin_instances table (FK)|
|id|text|The identifier for a user in the identity store.  This member is required.|
|name|text|Contains the user’s user name value|
