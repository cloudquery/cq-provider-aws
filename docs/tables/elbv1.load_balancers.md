
# Table: aws_elbv1_load_balancers
Information about a load balancer.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|attributes_access_log_enabled|boolean||
|attributes_access_log_s3_bucket_name|text||
|attributes_access_log_s3_bucket_prefix|text||
|attributes_access_log_emit_interval|integer||
|attributes_connection_settings_idle_timeout|integer||
|attributes_cross_zone_load_balancing_enabled|boolean||
|attributes_connection_draining_enabled|boolean||
|attributes_connection_draining_timeout|integer||
|attributes_additional_attributes|jsonb||
|tags|jsonb||
|availability_zones|text[]|The Availability Zones for the load balancer.|
|canonical_hosted_zone_name|text|The DNS name of the load balancer.|
|canonical_hosted_zone_name_id|text|The ID of the Amazon Route 53 hosted zone for the load balancer.|
|created_time|timestamp without time zone|The date and time the load balancer was created.|
|dns_name|text|The DNS name of the load balancer.|
|health_check_healthy_threshold|integer|The number of consecutive health checks successes required before moving the instance to the Healthy state.|
|health_check_interval|integer|The approximate interval, in seconds, between health checks of an individual instance.|
|health_check_target|text|The instance being checked.|
|health_check_timeout|integer|The amount of time, in seconds, during which no response means a failed health check.|
|health_check_unhealthy_threshold|integer|The number of consecutive health check failures required before moving the instance to the Unhealthy state.|
|instances|text[]|The IDs of the instances for the load balancer.|
|load_balancer_name|text|The name of the load balancer.|
|other_policies|text[]|The policies other than the stickiness policies.|
|scheme|text|The type of load balancer.|
|security_groups|text[]|The security groups for the load balancer.|
|source_security_group_name|text|The name of the security group.|
|source_security_group_owner_alias|text|The owner of the security group.|
|subnets|text[]|The IDs of the subnets for the load balancer.|
|vpc_id|text|The ID of the VPC for the load balancer.|
## Relations
## Table: aws_elbv1_load_balancer_backend_server_descriptions
Information about the configuration of an EC2 instance.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|load_balancer_id|uuid|Unique ID of aws_elbv1_load_balancers table (FK)|
|instance_port|integer|The port on which the EC2 instance is listening.|
|policy_names|text[]|The names of the policies enabled for the EC2 instance.|
## Table: aws_elbv1_load_balancer_listeners
The policies enabled for a listener.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|load_balancer_id|uuid|Unique ID of aws_elbv1_load_balancers table (FK)|
|listener_instance_port|integer|The port on which the instance is listening.|
|listener_load_balancer_port|integer|The port on which the load balancer is listening.|
|listener_protocol|text|The load balancer transport protocol to use for routing: HTTP, HTTPS, TCP, or SSL.|
|listener_instance_protocol|text|The protocol to use for routing traffic to instances: HTTP, HTTPS, TCP, or SSL.|
|listener_ssl_certificate_id|text|The Amazon Resource Name (ARN) of the server certificate.|
|policy_names|text[]|The policies.|
## Table: aws_elbv1_load_balancer_policies_app_cookie_stickiness_policies
Information about a policy for application-controlled session stickiness.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|load_balancer_id|uuid|Unique ID of aws_elbv1_load_balancers table (FK)|
|cookie_name|text|The name of the application cookie used for stickiness.|
|policy_name|text|The mnemonic name for the policy being created.|
## Table: aws_elbv1_load_balancer_policies_lb_cookie_stickiness_policies
Information about a policy for duration-based session stickiness.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|load_balancer_id|uuid|Unique ID of aws_elbv1_load_balancers table (FK)|
|cookie_expiration_period|bigint|The time period, in seconds, after which the cookie should be considered stale.|
|policy_name|text|The name of the policy.|
## Table: aws_elbv1_load_balancer_policies
Information about a policy.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|load_balancer_id|uuid|Unique ID of aws_elbv1_load_balancers table (FK)|
|policy_attribute_descriptions|jsonb|The policy attributes.|
|policy_name|text|The name of the policy.|
|policy_type_name|text|The name of the policy type.|
