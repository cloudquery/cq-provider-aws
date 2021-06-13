
# Table: aws_elbv2_load_balancers
Information about a load balancer.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|canonical_hosted_zone_id|text|The ID of the Amazon Route 53 hosted zone associated with the load balancer.|
|created_time|timestamp without time zone|The date and time the load balancer was created.|
|customer_owned_ipv4_pool|text|[Application Load Balancers on Outposts] The ID of the customer-owned address pool.|
|dns_name|text|The public DNS name of the load balancer.|
|ip_address_type|text|The type of IP addresses used by the subnets for your load balancer.|
|load_balancer_arn|text|The Amazon Resource Name (ARN) of the load balancer.|
|load_balancer_name|text|The name of the load balancer.|
|scheme|text|The nodes of an Internet-facing load balancer have public IP addresses.|
|security_groups|text[]|The IDs of the security groups for the load balancer.|
|state_code|text|The state code.|
|state_reason|text|A description of the state.|
|type|text|The type of load balancer.|
|vpc_id|text|The ID of the VPC for the load balancer.|
## Relations
## Table: aws_elbv2_load_balancer_availability_zones
Information about an Availability Zone.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|load_balancer_id|uuid|Unique ID of aws_elbv2_load_balancers table (FK)|
|outpost_id|text|[Application Load Balancers on Outposts] The ID of the Outpost.|
|subnet_id|text|The ID of the subnet.|
|zone_name|text|The name of the Availability Zone.|
## Relations
## Table: aws_elbv2_load_balancer_availability_zone_addresses
Information about a static IP address for a load balancer.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|load_balancer_availability_zone_id|uuid|Unique ID of aws_elbv2_load_balancer_availability_zones table (FK)|
|allocation_id|text|[Network Load Balancers] The allocation ID of the Elastic IP address for an internal-facing load balancer.|
|ipv6_address|text|[Network Load Balancers] The IPv6 address.|
|ip_address|text|The static IP address.|
|private_ipv4_address|text|[Network Load Balancers] The private IPv4 address for an internal load balancer.|
