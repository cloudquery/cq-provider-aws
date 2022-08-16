
# Table: aws_fsx_file_system_aliases
DNS aliases that are currently associated with the Amazon FSx file system.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|fsx_cq_id|uuid|Unique CloudQuery ID of aws_fsx_file_systems table (FK)|
|name|text|The name of the DNS alias.|
|lifecycle|text|The lifecycle status of DNS alias.|
