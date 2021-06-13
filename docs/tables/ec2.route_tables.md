
# Table: aws_ec2_route_tables

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text||
|region|text||
|owner_id|text||
|resource_id|text||
|tags|jsonb||
|vpc_id|text||
## Relations
## Table: aws_ec2_route_table_associations

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|route_table_id|uuid||
|association_state|text||
|association_state_status_message|text||
|gateway_id|text||
|main|boolean||
|route_table_association_id|text||
|subnet_id|text||
## Table: aws_ec2_route_table_propagating_vgws

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|route_table_id|uuid||
|gateway_id|text||
## Table: aws_ec2_route_table_routes

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|route_table_id|uuid||
|carrier_gateway_id|text||
|destination_cidr_block|text||
|destination_ipv6_cidr_block|text||
|destination_prefix_list_id|text||
|egress_only_internet_gateway_id|text||
|gateway_id|text||
|instance_id|text||
|instance_owner_id|text||
|local_gateway_id|text||
|nat_gateway_id|text||
|network_interface_id|text||
|origin|text||
|state|text||
|transit_gateway_id|text||
|vpc_peering_connection_id|text||
