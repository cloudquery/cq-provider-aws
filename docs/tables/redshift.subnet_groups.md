
# Table: aws_redshift_subnet_groups

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text||
|region|text||
|cluster_subnet_group_name|text||
|description|text||
|subnet_group_status|text||
|tags|jsonb||
|vpc_id|text||
## Relations
## Table: aws_redshift_subnet_group_subnets

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|subnet_group_id|uuid||
|subnet_availability_zone_name|text||
|subnet_availability_zone_supported_platforms|text[]||
|subnet_identifier|text||
|subnet_status|text||
