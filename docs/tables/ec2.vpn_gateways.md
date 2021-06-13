
# Table: aws_ec2_vpn_gateways

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text||
|region|text||
|amazon_side_asn|bigint||
|availability_zone|text||
|state|text||
|tags|jsonb||
|type|text||
|vpn_gateway_id|text||
## Relations
## Table: aws_ec2_vpc_attachment

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|vpn_gateway_id|uuid||
|state|text||
|vpc_id|text||
