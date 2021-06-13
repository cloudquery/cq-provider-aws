
# Table: aws_ec2_instances

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text||
|region|text||
|ami_launch_index|integer||
|architecture|text||
|capacity_reservation_id|text||
|cap_reservation_preference|text||
|cap_reservation_target_capacity_reservation_id|text||
|cap_reservation_target_capacity_reservation_rg_arn|text||
|client_token|text||
|cpu_options_core_count|integer||
|cpu_options_threads_per_core|integer||
|ebs_optimized|boolean||
|ena_support|boolean||
|enclave_options_enabled|boolean||
|hibernation_options_configured|boolean||
|hypervisor|text||
|iam_instance_profile_arn|text||
|iam_instance_profile_id|text||
|image_id|text||
|instance_id|text||
|instance_lifecycle|text||
|instance_type|text||
|kernel_id|text||
|key_name|text||
|launch_time|timestamp without time zone||
|metadata_options_http_endpoint|text||
|metadata_options_http_put_response_hop_limit|integer||
|metadata_options_http_tokens|text||
|metadata_options_state|text||
|monitoring_state|text||
|outpost_arn|text||
|placement_affinity|text||
|placement_availability_zone|text||
|placement_group_name|text||
|placement_host_id|text||
|placement_host_resource_group_arn|text||
|placement_partition_number|integer||
|placement_spread_domain|text||
|placement_tenancy|text||
|platform|text||
|private_dns_name|text||
|private_ip_address|text||
|public_dns_name|text||
|public_ip_address|text||
|ramdisk_id|text||
|root_device_name|text||
|root_device_type|text||
|source_dest_check|boolean||
|spot_instance_request_id|text||
|sriov_net_support|text||
|state_code|integer||
|state_name|text||
|state_reason_code|text||
|state_reason_message|text||
|state_transition_reason|text||
|subnet_id|text||
|tags|jsonb||
|virtualization_type|text||
|vpc_id|text||
## Relations
## Table: aws_ec2_instance_block_device_mappings

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|instance_id|uuid||
|device_name|text||
|ebs_attach_time|timestamp without time zone||
|ebs_delete_on_termination|boolean||
|ebs_status|text||
|ebs_volume_id|text||
## Table: aws_ec2_instance_elastic_gpu_associations

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|instance_id|uuid||
|elastic_gpu_association_id|text||
|elastic_gpu_association_state|text||
|elastic_gpu_association_time|text||
|elastic_gpu_id|text||
## Table: aws_ec2_instance_elastic_inference_accelerator_associations

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|instance_id|uuid||
|elastic_inference_accelerator_arn|text||
|elastic_inference_accelerator_association_id|text||
|elastic_inference_accelerator_association_state|text||
|elastic_inference_accelerator_association_time|timestamp without time zone||
## Table: aws_ec2_instance_licenses

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|instance_id|uuid||
|license_configuration_arn|text||
## Table: aws_ec2_instance_network_interfaces

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|instance_id|uuid||
|association_carrier_ip|text||
|association_ip_owner_id|text||
|association_public_dns_name|text||
|association_public_ip|text||
|attachment_attach_time|timestamp without time zone||
|attachment_id|text||
|attachment_delete_on_termination|boolean||
|attachment_device_index|integer||
|attachment_network_card_index|integer||
|attachment_status|text||
|description|text||
|interface_type|text||
|mac_address|text||
|network_interface_id|text||
|owner_id|text||
|private_dns_name|text||
|private_ip_address|text||
|source_dest_check|boolean||
|status|text||
|subnet_id|text||
|vpc_id|text||
## Relations
## Table: aws_ec2_instance_network_interface_groups

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|instance_network_interface_id|uuid||
|group_id|text||
|group_name|text||
## Table: aws_ec2_instance_network_interface_ipv6_addresses

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|instance_network_interface_id|uuid||
|ipv6_address|text||
## Table: aws_ec2_instance_network_interface_private_ip_addresses

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|instance_network_interface_id|uuid||
|association_carrier_ip|text||
|association_ip_owner_id|text||
|association_public_dns_name|text||
|association_public_ip|text||
|is_primary|boolean||
|private_dns_name|text||
|private_ip_address|text||
## Table: aws_ec2_instance_product_codes

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|instance_id|uuid||
|product_code_id|text||
|product_code_type|text||
## Table: aws_ec2_instance_security_groups

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|instance_id|uuid||
|group_id|text||
|group_name|text||
