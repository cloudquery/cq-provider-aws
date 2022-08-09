
# Table: aws_fsx_filesystem_windows_configuration_aliases
A DNS alias that is associated with the file system
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|filesystem_windows_configuration_cq_id|uuid|Unique CloudQuery ID of aws_fsx_filesystem_windows_configuration table (FK)|
|lifecycle|text|Describes the state of the DNS alias.  * AVAILABLE - The DNS alias is associated with an Amazon FSx file system.  * CREATING - Amazon FSx is creating the DNS alias and associating it with the file system.  * CREATE_FAILED - Amazon FSx was unable to associate the DNS alias with the file system.  * DELETING - Amazon FSx is disassociating the DNS alias from the file system and deleting it.  * DELETE_FAILED - Amazon FSx was unable to disassociate the DNS alias from the file system.|
|name|text|The name of the DNS alias|
