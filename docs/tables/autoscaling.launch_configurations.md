
# Table: aws_autoscaling_launch_configurations

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text||
|region|text||
|created_time|timestamp without time zone||
|image_id|text||
|instance_type|text||
|launch_configuration_name|text||
|associate_public_ip_address|boolean||
|classic_link_vpc_id|text||
|classic_link_vpc_security_groups|text[]||
|ebs_optimized|boolean||
|iam_instance_profile|text||
|instance_monitoring_enabled|boolean||
|kernel_id|text||
|key_name|text||
|launch_configuration_arn|text||
|metadata_options_http_endpoint|text||
|metadata_options_http_put_response_hop_limit|integer||
|metadata_options_http_tokens|text||
|placement_tenancy|text||
|ramdisk_id|text||
|security_groups|text[]||
|spot_price|text||
|user_data|text||
## Relations
## Table: aws_autoscaling_launch_configuration_block_device_mappings

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|launch_configuration_id|uuid||
|device_name|text||
|ebs_delete_on_termination|boolean||
|ebs_encrypted|boolean||
|ebs_iops|integer||
|ebs_snapshot_id|text||
|ebs_volume_size|integer||
|ebs_volume_type|text||
|no_device|boolean||
|virtual_name|text||
