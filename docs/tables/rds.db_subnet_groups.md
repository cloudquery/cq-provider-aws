
# Table: aws_rds_subnet_groups

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text||
|region|text||
|db_subnet_group_arn|text||
|db_subnet_group_description|text||
|db_subnet_group_name|text||
|subnet_group_status|text||
|vpc_id|text||
## Relations
## Table: aws_rds_subnet_group_subnets

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|subnet_group_id|uuid||
|subnet_availability_zone_name|text||
|subnet_identifier|text||
|subnet_outpost_arn|text||
|subnet_status|text||
