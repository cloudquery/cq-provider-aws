
# Table: aws_waf_rules

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text||
|region|text||
|arn|text||
|tags|jsonb||
|rule_id|text||
|metric_name|text||
|name|text||
## Relations
## Table: aws_waf_rule_predicates

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|rule_id|uuid||
|data_id|text||
|negated|boolean||
|type|text||
