
# Table: aws_iam_users

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text||
|password_last_used|timestamp without time zone||
|arn|text||
|password_enabled|boolean||
|password_status|text||
|password_last_changed|timestamp without time zone||
|password_next_rotation|timestamp without time zone||
|mfa_active|boolean||
|create_date|timestamp without time zone||
|path|text||
|permissions_boundary_arn|text||
|permissions_boundary_type|text||
|tags|jsonb||
|user_id|text||
|user_name|text||
|access_key_1_active|boolean||
|access_key_1_last_rotated|timestamp without time zone||
|access_key_2_active|boolean||
|access_key_2_last_rotated|timestamp without time zone||
|cert_1_active|boolean||
|cert_1_last_rotated|timestamp without time zone||
|cert_2_active|boolean||
|cert_2_last_rotated|timestamp without time zone||
## Relations
## Table: aws_iam_user_access_keys

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|user_id|uuid||
|access_key_id|text||
|create_date|timestamp without time zone||
|status|text||
|last_used|timestamp without time zone||
|last_rotated|timestamp without time zone||
|last_used_service_name|text||
## Table: aws_iam_user_groups

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|user_id|uuid||
|arn|text||
|create_date|timestamp without time zone||
|group_id|text||
|group_name|text||
|path|text||
## Table: aws_iam_user_attached_policies

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|user_id|uuid||
|policy_arn|text||
|policy_name|text||
## Table: aws_iam_user_policies
Inline policies that are embedded in the specified IAM user
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|user_id|uuid|User ID the policy belongs too.|
|account_id|text|The AWS Account ID of the resource.|
|policy_document|jsonb|The policy document. IAM stores policies in JSON format. However, resources that were created using AWS CloudFormation templates can be formatted in YAML. AWS CloudFormation always converts a YAML policy to JSON format before submitting it to IAM.|
|policy_name|text|The name of the policy.|
|user_name|text|The user the policy is associated with.|
