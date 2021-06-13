
# Table: aws_ec2_images

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text||
|region|text||
|architecture|text||
|creation_date|text||
|description|text||
|ena_support|boolean||
|hypervisor|text||
|image_id|text||
|image_location|text||
|image_owner_alias|text||
|image_type|text||
|kernel_id|text||
|name|text||
|owner_id|text||
|platform|text||
|platform_details|text||
|product_codes|jsonb||
|public|boolean||
|ramdisk_id|text||
|root_device_name|text||
|root_device_type|text||
|sriov_net_support|text||
|state|text||
|state_reason_code|text||
|state_reason_message|text||
|tags|jsonb||
|usage_operation|text||
|virtualization_type|text||
## Relations
## Table: aws_ec2_image_block_device_mappings

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|image_id|uuid||
|device_name|text||
|ebs_delete_on_termination|boolean||
|ebs_encrypted|boolean||
|ebs_iops|integer||
|ebs_kms_key_id|text||
|ebs_outpost_arn|text||
|ebs_snapshot_id|text||
|ebs_throughput|integer||
|ebs_volume_size|integer||
|ebs_volume_type|text||
|no_device|text||
|virtual_name|text||
