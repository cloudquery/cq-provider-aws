
# Table: aws_wafv2_rule_groups

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text||
|region|text||
|tags|jsonb||
|policy|jsonb||
|arn|text||
|capacity|bigint||
|resource_id|text||
|name|text||
|visibility_config_cloud_watch_metrics_enabled|boolean||
|visibility_config_metric_name|text||
|visibility_config_sampled_requests_enabled|boolean||
|custom_response_bodies|jsonb||
|description|text||
|label_namespace|text||
|rules|jsonb||
## Relations
## Table: aws_wafv2_rule_group_available_labels

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|rule_group_id|uuid||
|name|text||
## Table: aws_wafv2_rule_group_consumed_labels

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|rule_group_id|uuid||
|name|text||
