
# Table: aws_ec2_security_groups

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text||
|region|text||
|description|text||
|group_id|text||
|group_name|text||
|owner_id|text||
|tags|jsonb||
|vpc_id|text||
## Relations
## Table: aws_ec2_security_group_ip_permissions

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|security_group_id|uuid||
|from_port|integer||
|ip_protocol|text||
|to_port|integer||
## Relations
## Table: aws_ec2_security_group_ip_permission_ip_ranges

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|security_group_ip_permission_id|uuid||
|cidr_ip|text||
|description|text||
## Table: aws_ec2_security_group_ip_permission_ipv6_ranges

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|security_group_ip_permission_id|uuid||
|cidr_ipv6|text||
|description|text||
## Table: aws_ec2_security_group_ip_permission_prefix_list_ids

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|security_group_ip_permission_id|uuid||
|description|text||
|prefix_list_id|text||
## Table: aws_ec2_security_group_ip_permission_user_id_group_pairs

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|security_group_ip_permission_id|uuid||
|description|text||
|group_id|text||
|group_name|text||
|peering_status|text||
|user_id|text||
|vpc_id|text||
|vpc_peering_connection_id|text||
## Table: aws_ec2_security_group_ip_permissions_egresses

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|security_group_id|uuid||
|from_port|integer||
|ip_protocol|text||
|to_port|integer||
## Relations
## Table: aws_ec2_security_group_ip_permissions_egress_ip_ranges

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|security_group_ip_permissions_egress_id|uuid||
|cidr_ip|text||
|description|text||
## Table: aws_ec2_security_group_ip_permissions_egress_ipv6_ranges

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|security_group_ip_permissions_egress_id|uuid||
|cidr_ipv6|text||
|description|text||
## Table: aws_ec2_security_group_ip_permissions_egress_prefix_list_ids

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|security_group_ip_permissions_egress_id|uuid||
|description|text||
|prefix_list_id|text||
## Table: aws_ec2_security_group_ip_permissions_egress_user_group_pairs

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|security_group_ip_permissions_egress_id|uuid||
|description|text||
|group_id|text||
|group_name|text||
|peering_status|text||
|user_id|text||
|vpc_id|text||
|vpc_peering_connection_id|text||
