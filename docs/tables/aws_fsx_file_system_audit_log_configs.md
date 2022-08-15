
# Table: aws_fsx_file_system_audit_log_configs
The configuration that Amazon FSx for Windows File Server uses to audit and log user accesses of files, folders, and file shares on the Amazon FSx for Windows File Server file system.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|fsx_cq_id|uuid|Unique CloudQuery ID of aws_fsx_file_systems table (FK)|
|file_access_audit_log_level|text|Sets which attempt type is logged by Amazon FSx for file and folder accesses.|
|file_share_access_audit_log_level|text|Sets which attempt type is logged by Amazon FSx for file share accesses.|
|audit_log_destination|text|The Amazon Resource Name (ARN) for the destination of the audit logs.|
