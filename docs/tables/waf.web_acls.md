
# Table: aws_waf_web_acls

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text||
|tags|jsonb||
|region|text||
|default_action_type|text||
|web_acl_id|text||
|metric_name|text||
|name|text||
|web_acl_arn|text||
## Relations
## Table: aws_waf_web_acl_rules

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|web_acl_id|uuid||
|priority|integer||
|rule_id|text||
|action_type|text||
|excluded_rules|text[]||
|override_action_type|text||
|type|text||
