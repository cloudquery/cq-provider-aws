
# Table: aws_ec2_network_acls

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text||
|region|text||
|is_default|boolean||
|network_acl_id|text||
|owner_id|text||
|tags|jsonb||
|vpc_id|text||
## Relations
## Table: aws_ec2_network_acl_associations

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|network_acl_id|uuid||
|network_acl_association_id|text||
|subnet_id|text||
## Table: aws_ec2_network_acl_entries

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|network_acl_id|uuid||
|cidr_block|text||
|egress|boolean||
|icmp_type_code|integer||
|icmp_type_code_type|integer||
|ipv6_cidr_block|text||
|port_range_from|integer||
|port_range_to|integer||
|protocol|text||
|rule_action|text||
|rule_number|integer||
