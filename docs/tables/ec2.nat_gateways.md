
# Table: aws_ec2_nat_gateways

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text||
|region|text||
|create_time|timestamp without time zone||
|delete_time|timestamp without time zone||
|failure_code|text||
|failure_message|text||
|nat_gateway_id|text||
|provisioned_bandwidth_provision_time|timestamp without time zone||
|provisioned_bandwidth_provisioned|text||
|provisioned_bandwidth_request_time|timestamp without time zone||
|provisioned_bandwidth_requested|text||
|provisioned_bandwidth_status|text||
|state|text||
|subnet_id|text||
|tags|jsonb||
|vpc_id|text||
## Relations
## Table: aws_ec2_nat_gateway_addresses

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|nat_gateway_id|uuid||
|allocation_id|text||
|network_interface_id|text||
|private_ip|text||
|public_ip|text||
