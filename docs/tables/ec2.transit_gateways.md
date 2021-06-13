
# Table: aws_ec2_transit_gateways

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text||
|region|text||
|amazon_side_asn|bigint||
|association_default_route_table_id|text||
|auto_accept_shared_attachments|text||
|creation_time|timestamp without time zone||
|default_route_table_association|text||
|default_route_table_propagation|text||
|description|text||
|dns_support|text||
|multicast_support|text||
|owner_id|text||
|propagation_default_route_table_id|text||
|state|text||
|tags|jsonb||
|transit_gateway_arn|text||
|transit_gateway_cidr_blocks|text[]||
|transit_gateway_id|text||
|vpn_ecmp_support|text||
## Relations
## Table: aws_ec2_transit_gateway_attachments

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|transit_gateway_id|uuid||
|association_state|text||
|association_route_table_id|text||
|creation_time|timestamp without time zone||
|resource_id|text||
|resource_owner_id|text||
|resource_type|text||
|state|text||
|tags|jsonb||
|transit_gateway_owner_id|text||
## Table: aws_ec2_transit_gateway_route_tables

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|transit_gateway_id|uuid||
|creation_time|timestamp without time zone||
|default_association_route_table|boolean||
|default_propagation_route_table|boolean||
|state|text||
|tags|jsonb||
|transit_gateway_route_table_id|text||
## Table: aws_ec2_transit_gateway_vpc_attachments

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|transit_gateway_id|uuid||
|creation_time|timestamp without time zone||
|appliance_mode_support|text||
|dns_support|text||
|ipv6_support|text||
|state|text||
|tags|jsonb||
|transit_gateway_attachment_id|text||
|vpc_id|text||
|vpc_owner_id|text||
## Table: aws_ec2_transit_gateway_peering_attachments

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|transit_gateway_id|uuid||
|accepter_owner_id|text||
|accepter_region|text||
|accepter_transit_gateway_id|text||
|creation_time|timestamp without time zone||
|requester_owner_id|text||
|requester_region|text||
|requester_transit_gateway_id|text||
|state|text||
|status_code|text||
|status_message|text||
|tags|jsonb||
|transit_gateway_attachment_id|text||
## Table: aws_ec2_transit_gateway_multicast_domains

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|transit_gateway_id|uuid||
|creation_time|timestamp without time zone||
|auto_accept_shared_associations|text||
|igmpv2_support|text||
|static_sources_support|text||
|owner_id|text||
|state|text||
|tags|jsonb||
|transit_gateway_multicast_domain_arn|text||
|transit_gateway_multicast_domain_id|text||
