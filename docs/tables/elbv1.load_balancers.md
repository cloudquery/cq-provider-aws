
# Table: aws_elbv1_load_balancers

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text||
|region|text||
|attributes_access_log_enabled|boolean||
|attributes_access_log_s3_bucket_name|text||
|attributes_access_log_s3_bucket_prefix|text||
|attributes_access_log_emit_interval|integer||
|attributes_connection_settings_idle_timeout|integer||
|attributes_cross_zone_load_balancing_enabled|boolean||
|attributes_connection_draining_enabled|boolean||
|attributes_connection_draining_timeout|integer||
|attributes_additional_attributes|jsonb||
|tags|jsonb||
|availability_zones|text[]||
|canonical_hosted_zone_name|text||
|canonical_hosted_zone_name_id|text||
|created_time|timestamp without time zone||
|dns_name|text||
|health_check_healthy_threshold|integer||
|health_check_interval|integer||
|health_check_target|text||
|health_check_timeout|integer||
|health_check_unhealthy_threshold|integer||
|instances|text[]||
|load_balancer_name|text||
|other_policies|text[]||
|scheme|text||
|security_groups|text[]||
|source_security_group_name|text||
|source_security_group_owner_alias|text||
|subnets|text[]||
|vpc_id|text||
## Relations
## Table: aws_elbv1_load_balancer_backend_server_descriptions

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|load_balancer_id|uuid||
|instance_port|integer||
|policy_names|text[]||
## Table: aws_elbv1_load_balancer_listeners

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|load_balancer_id|uuid||
|listener_instance_port|integer||
|listener_load_balancer_port|integer||
|listener_protocol|text||
|listener_instance_protocol|text||
|listener_ssl_certificate_id|text||
|policy_names|text[]||
## Table: aws_elbv1_load_balancer_policies_app_cookie_stickiness_policies

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|load_balancer_id|uuid||
|cookie_name|text||
|policy_name|text||
## Table: aws_elbv1_load_balancer_policies_lb_cookie_stickiness_policies

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|load_balancer_id|uuid||
|cookie_expiration_period|bigint||
|policy_name|text||
## Table: aws_elbv1_load_balancer_policies

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|load_balancer_id|uuid||
|policy_attribute_descriptions|jsonb||
|policy_name|text||
|policy_type_name|text||
