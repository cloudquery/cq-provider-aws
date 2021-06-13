
# Table: aws_lambda_layers

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text||
|region|text||
|latest_matching_version_compatible_runtimes|text[]||
|latest_matching_version_created_date|text||
|latest_matching_version_description|text||
|latest_matching_version_layer_version_arn|text||
|latest_matching_version_license_info|text||
|latest_matching_version|bigint||
|layer_arn|text||
|layer_name|text||
## Relations
## Table: aws_lambda_layer_versions

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|layer_id|uuid||
|compatible_runtimes|text[]||
|created_date|text||
|description|text||
|layer_version_arn|text||
|license_info|text||
|version|bigint||
## Relations
## Table: aws_lambda_layer_version_policies

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|layer_version_id|uuid||
|policy|text||
|revision_id|text||
