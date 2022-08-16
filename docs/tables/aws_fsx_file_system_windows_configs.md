
# Table: aws_fsx_file_system_windows_configs
The configuration for this Amazon FSx for Windows File Server file system.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|fsx_cq_id|uuid|Unique CloudQuery ID of aws_fsx_file_systems table (FK)|
|active_directory_id|text|The ID for an existing Amazon Web Services Managed Microsoft Active Directory instance that the file system is joined to.|
|deployment_type|text|Specifies the file system deployment type.|
|remote_administration_endpoint|text|Specifies the file system remote administration endpoint.|
|preferred_subnet_id|text|Specifies the file system preferred subnet ID.|
|preferred_file_server_ip|text|Specifies the file system preferred file server IP.|
|throughput_capacity|integer|The throughput of the Amazon FSx file system, measured in megabytes per second.|
|maintenance_operations_in_progress|text[]|The list of maintenance operations in progress for this file system.|
|weekly_maintenance_start_time|text|The preferred start time to perform weekly maintenance, formatted d:HH:MM in the UTC time zone. d is the weekday number, from 1 through 7, beginning with Monday and ending with Sunday.|
|daily_automatic_backup_start_time|text|The preferred time to take daily automatic backups, in the UTC time zone.|
|automatic_backup_retention_days|integer|The number of days to retain automatic backups. Setting this to 0 disables automatic backups. You can retain automatic backups for a maximum of 90 days.|
|copy_tags_to_backups|boolean|A boolean flag indicating whether tags on the file system should be copied to backups.|
