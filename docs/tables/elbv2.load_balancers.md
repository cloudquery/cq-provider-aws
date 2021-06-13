
# Table: aws_elbv2_load_balancers

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text||
|region|text||
|canonical_hosted_zone_id|text||
|created_time|timestamp without time zone||
|customer_owned_ipv4_pool|text||
|dns_name|text||
|ip_address_type|text||
|load_balancer_arn|text||
|load_balancer_name|text||
|scheme|text||
|security_groups|text[]||
|state_code|text||
|state_reason|text||
|type|text||
|vpc_id|text||
## Relations
## Table: aws_elbv2_load_balancer_availability_zones

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|load_balancer_id|uuid||
|outpost_id|text||
|subnet_id|text||
|zone_name|text||
## Relations
## Table: aws_elbv2_load_balancer_availability_zone_addresses

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|load_balancer_availability_zone_id|uuid||
|allocation_id|text||
|ip_v6_address|text||
|ip_address|text||
|private_ip_v4_address|text||
