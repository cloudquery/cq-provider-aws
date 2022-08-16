
# Table: aws_fsx_file_system_lustre_configs
The configuration for the Amazon FSx for Lustre file system.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|fsx_cq_id|uuid|Unique CloudQuery ID of aws_fsx_file_systems table (FK)|
|deployment_type|text|The deployment type of the FSx for Lustre file system.|
|per_unit_storage_throughput|integer|Per unit storage throughput represents the megabytes per second of read or write throughput per 1 tebibyte of storage provisioned.|
|mount_name|integer|Mount name of the FSx Lustre file system.|
|daily_automatic_backup_start_time|text|A recurring daily time, in the format HH:MM.|
|automatic_backup_retention_days|integer|The number of days to retain automatic backups.|
|copy_tags_to_backups|boolean|A boolean flag indicating whether tags on the file system are copied to backups.|
|drive_cache_type|text|The type of drive cache used by PERSISTENT_1 file systems that are provisioned with HDD storage devices.|
|data_compression_type|text|The data compression configuration for the file system.|
|log_level|text|The data repository events that are logged by Amazon FSx.|
|log_destination|text|The Amazon Resource Name (ARN) that specifies the destination of the logs.|
|root_squash|text|You enable root squash by setting a user ID (UID) and group ID (GID) for the file system in the format UID:GID (for example, 365534:65534 ).|
|no_squash_nids|integer[]|When root squash is enabled, you can optionally specify an array of NIDs of clients for which root squash does not apply.|
