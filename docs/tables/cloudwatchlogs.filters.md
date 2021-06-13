
# Table: aws_cloudwatchlogs_filters

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text||
|region|text||
|creation_time|bigint||
|name|text||
|pattern|text||
|log_group_name|text||
## Relations
## Table: aws_cloudwatchlogs_filter_metric_transformations

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|filter_id|uuid||
|metric_name|text||
|metric_namespace|text||
|metric_value|text||
|default_value|float||
