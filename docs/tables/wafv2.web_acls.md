
# Table: aws_wafv2_web_acls

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text||
|region|text||
|resources_for_web_acl|text[]||
|tags|jsonb||
|arn|text||
|default_action|jsonb||
|resource_id|text||
|name|text||
|visibility_config_cloud_watch_metrics_enabled|boolean||
|visibility_config_metric_name|text||
|visibility_config_sampled_requests_enabled|boolean||
|capacity|bigint||
|custom_response_bodies|jsonb||
|description|text||
|label_namespace|text||
|managed_by_firewall_manager|boolean||
## Relations
## Table: aws_wafv2_web_acl_rules

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|web_acl_id|uuid||
|name|text||
|priority|integer||
|statement|jsonb||
|visibility_config_cloud_watch_metrics_enabled|boolean||
|visibility_config_metric_name|text||
|visibility_config_sampled_requests_enabled|boolean||
|action|jsonb||
|override_action|jsonb||
|labels|text[]||
## Table: aws_wafv2_web_acl_post_process_firewall_manager_rule_groups

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|web_acl_id|uuid||
|statement|jsonb||
|name|text||
|override_action|jsonb||
|priority|integer||
|visibility_config_cloud_watch_metrics_enabled|boolean||
|visibility_config_metric_name|text||
|visibility_config_sampled_requests_enabled|boolean||
## Table: aws_wafv2_web_acl_pre_process_firewall_manager_rule_groups

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|web_acl_id|uuid||
|statement|jsonb||
|name|text||
|override_action|jsonb||
|priority|integer||
|visibility_config_cloud_watch_metrics_enabled|boolean||
|visibility_config_metric_name|text||
|visibility_config_sampled_requests_enabled|boolean||
