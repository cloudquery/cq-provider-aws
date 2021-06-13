
# Table: aws_iam_policies

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text||
|region|text||
|arn|text||
|attachment_count|integer||
|create_date|timestamp without time zone||
|default_version_id|text||
|description|text||
|is_attachable|boolean||
|path|text||
|permissions_boundary_usage_count|integer||
|policy_id|text||
|policy_name|text||
|update_date|timestamp without time zone||
## Relations
## Table: aws_iam_policy_versions

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|policy_id|uuid||
|create_date|timestamp without time zone||
|document|jsonb||
|is_default_version|boolean||
|version_id|text||
