
# Table: aws_apigateway_usage_plans

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text||
|region|text||
|description|text||
|resource_id|text||
|name|text||
|product_code|text||
|quota_limit|integer||
|quota_offset|integer||
|quota_period|text||
|tags|jsonb||
|throttle_burst_limit|integer||
|throttle_rate_limit|float||
## Relations
## Table: aws_apigateway_usage_plan_api_stages

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|usage_plan_id|uuid||
|api_id|text||
|stage|text||
|throttle|jsonb||
## Table: aws_apigateway_usage_plan_keys

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|usage_plan_id|uuid||
|resource_id|text||
|name|text||
|type|text||
|value|text||
