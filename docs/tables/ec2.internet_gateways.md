
# Table: aws_ec2_internet_gateways

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text||
|region|text||
|internet_gateway_id|text||
|owner_id|text||
|tags|jsonb||
## Relations
## Table: aws_ec2_internet_gateway_attachments

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|internet_gateway_id|uuid||
|state|text||
|vpc_id|text||
