
# Table: aws_fsx_file_system_ontap_configs
The configuration for this Amazon FSx for NetApp ONTAP file system.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|fsx_cq_id|uuid|Unique CloudQuery ID of aws_fsx_file_systems table (FK)|
|automatic_backup_retention_days|integer|The number of days to retain automatic backups.|
|daily_automatic_backup_start_time|text|A recurring daily time, in the format HH:MM.|
|deployment_type|text|Specifies the FSx for ONTAP file system deployment type in use in the file system.|
|endpoint_ip_address_range|text|(Multi-AZ only) The IP address range in which the endpoints to access your file system are created.|
|intercluster_dns_name|text|The Domain Name Service (DNS) name for the file system intercluster endpoint.|
|intercluster_ip_addresses|text|IP addresses of the file system intercluster endpoint.|
|management_dns_name|text|The Domain Name Service (DNS) name for the file system management endpoint.|
|management_ip_addresses|text|IP addresses of the file system management endpoint.|
|disk_iops_mode|text|Specifies whether the number of IOPS for the file system is using the system default (AUTOMATIC) or was provisioned by the customer (USER_PROVISIONED).|
|disk_iops|float|The total number of SSD IOPS provisioned for the file system.|
|preferred_subnet_id|text|The ID for a subnet.|
|route_table_ids|text[]|(Multi-AZ only) The VPC route tables in which your file system's endpoints are created.|
|throughput_capacity|integer|The sustained throughput of an Amazon FSx file system in Megabytes per second (MBps).|
|weekly_maintenance_start_time|text|A recurring weekly time, in the format D:HH:MM .|
