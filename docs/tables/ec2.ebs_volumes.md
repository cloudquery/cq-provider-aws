
# Table: aws_ec2_ebs_volumes

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text||
|region|text||
|volume_id|text||
|availability_zone|text||
|create_time|timestamp without time zone||
|encrypted|boolean||
|fast_restored|boolean||
|iops|integer||
|kms_key_id|text||
|multi_attach_enabled|boolean||
|outpost_arn|text||
|size|integer||
|snapshot_id|text||
|state|text||
|tags|jsonb||
|throughput|integer||
|volume_type|text||
## Relations
## Table: aws_ec2_ebs_volume_attachments

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|ebs_volume_id|uuid||
|attach_time|timestamp without time zone||
|delete_on_termination|boolean||
|device|text||
|instance_id|text||
|state|text||
|volume_id|text||
