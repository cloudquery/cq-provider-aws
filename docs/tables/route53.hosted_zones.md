
# Table: aws_route53_hosted_zones

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text||
|tags|jsonb||
|delegation_set_id|text||
|caller_reference|text||
|resource_id|text||
|name|text||
|config_comment|text||
|config_private_zone|boolean||
|linked_service_description|text||
|linked_service_principal|text||
|resource_record_set_count|bigint||
## Relations
## Table: aws_route53_hosted_zone_query_logging_configs

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|hosted_zone_id|uuid||
|cloud_watch_logs_log_group_arn|text||
|query_logging_config_id|text||
## Table: aws_route53_hosted_zone_resource_record_sets

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|hosted_zone_id|uuid||
|resource_records|text[]||
|name|text||
|type|text||
|dns_name|text||
|evaluate_target_health|boolean||
|failover|text||
|geo_location_continent_code|text||
|geo_location_country_code|text||
|geo_location_subdivision_code|text||
|health_check_id|text||
|multi_value_answer|boolean||
|region|text||
|set_identifier|text||
|ttl|bigint||
|traffic_policy_instance_id|text||
|weight|bigint||
## Table: aws_route53_hosted_zone_traffic_policy_instances

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|hosted_zone_id|uuid||
|policy_id|text||
|message|text||
|name|text||
|state|text||
|ttl|bigint||
|traffic_policy_id|text||
|traffic_policy_type|text||
|traffic_policy_version|integer||
## Table: aws_route53_hosted_zone_vpc_association_authorizations

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|hosted_zone_id|uuid||
|vpc_id|text||
|vpc_region|text||
