
# Table: aws_route53_traffic_policies

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text||
|resource_id|text||
|latest_version|integer||
|name|text||
|traffic_policy_count|integer||
|type|text||
## Relations
## Table: aws_route53_traffic_policy_versions

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|traffic_policy_id|uuid||
|document|jsonb||
|version_id|text||
|name|text||
|type|text||
|version|integer||
|comment|text||
