
# Table: aws_iam_roles

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text||
|policies|jsonb||
|arn|text||
|create_date|timestamp without time zone||
|path|text||
|role_id|text||
|role_name|text||
|assume_role_policy_document|jsonb||
|description|text||
|max_session_duration|integer||
|permissions_boundary_arn|text||
|permissions_boundary_type|text||
|role_last_used_last_used_date|timestamp without time zone||
|role_last_used_region|text||
|tags|jsonb||
## Relations
## Table: aws_iam_role_policies

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|role_id|uuid||
|account_id|text||
|policy_document|jsonb||
|policy_name|text||
|role_name|text||
