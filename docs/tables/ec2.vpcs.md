
# Table: aws_ec2_vpcs

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text||
|region|text||
|cidr_block|text||
|dhcp_options_id|text||
|instance_tenancy|text||
|is_default|boolean||
|owner_id|text||
|state|text||
|tags|jsonb||
|vpc_id|text||
## Relations
## Table: aws_ec2_vpc_cidr_block_association_sets

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|vpc_id|uuid||
|association_id|text||
|cidr_block|text||
|cidr_block_state|text||
|cidr_block_state_status_message|text||
## Table: aws_ec2_vpc_ipv6_cidr_block_association_sets

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|vpc_id|uuid||
|association_id|text||
|ipv6_cidr_block|text||
|ipv6_cidr_block_state|text||
|ipv6_cidr_block_state_status_message|text||
|ipv6_pool|text||
|network_border_group|text||
