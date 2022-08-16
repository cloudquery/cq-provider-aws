
# Table: aws_fsx_file_system_openzfs_configs
The configuration for this Amazon FSx for OpenZFS file system.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|fsx_cq_id|uuid|Unique CloudQuery ID of aws_fsx_file_systems table (FK)|
|automatic_backup_retention_days|integer|The number of days to retain automatic backups.|
|copy_tags_to_backups|boolean|A boolean flag indicating whether tags on the file system should be copied to backups.|
|copy_tags_to_volumes|boolean|A Boolean value indicating whether tags for the volume should be copied to snapshots.|
|daily_automatic_backup_start_time|text|A recurring daily time, in the format HH:MM.|
|deployment_type|text|Specifies the file-system deployment type.|
|throughput_capacity|integer|The sustained throughput of an Amazon FSx file system in Megabytes per second (MBps).|
|weekly_maintenance_start_time|text|A recurring weekly time, in the format D:HH:MM .|
|disk_iops_mode|text|Specifies whether the number of IOPS for the file system is using the system default (AUTOMATIC) or was provisioned by the customer (USER_PROVISIONED).|
|disk_iops|float|The total number of SSD IOPS provisioned for the file system.|
