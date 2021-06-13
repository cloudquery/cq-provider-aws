
# Table: aws_ec2_vpc_endpoints

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text||
|region|text||
|creation_timestamp|timestamp without time zone||
|last_error_code|text||
|last_error_message|text||
|network_interface_ids|text[]||
|owner_id|text||
|policy_document|text||
|private_dns_enabled|boolean||
|requester_managed|boolean||
|route_table_ids|text[]||
|service_name|text||
|state|text||
|subnet_ids|text[]||
|tags|jsonb||
|vpc_endpoint_id|text||
|vpc_endpoint_type|text||
|vpc_id|text||
## Relations
## Table: aws_ec2_vpc_endpoint_dns_entries

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|vpc_endpoint_id|uuid||
|dns_name|text||
|hosted_zone_id|text||
## Table: aws_ec2_vpc_endpoint_groups

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|vpc_endpoint_id|uuid||
|group_id|text||
|group_name|text||
