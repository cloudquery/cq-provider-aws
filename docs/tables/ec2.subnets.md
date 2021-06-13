
# Table: aws_ec2_subnets

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text||
|region|text||
|assign_ipv6_address_on_creation|boolean||
|availability_zone|text||
|availability_zone_id|text||
|available_ip_address_count|integer||
|cidr_block|text||
|customer_owned_ipv4_pool|text||
|default_for_az|boolean||
|map_customer_owned_ip_on_launch|boolean||
|map_public_ip_on_launch|boolean||
|outpost_arn|text||
|owner_id|text||
|state|text||
|subnet_arn|text||
|subnet_id|text||
|tags|jsonb||
|vpc_id|text||
## Relations
## Table: aws_ec2_subnet_ipv6_cidr_block_association_sets

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|subnet_id|uuid||
|association_id|text||
|ipv6_cidr_block|text||
|ipv6_cidr_block_state|text||
|ipv6_cidr_block_state_status_message|text||
