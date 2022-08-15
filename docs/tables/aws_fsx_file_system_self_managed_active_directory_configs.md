
# Table: aws_fsx_file_system_self_managed_active_directory_configs
The configuration of the self-managed Microsoft Active Directory (AD) directory to which the Windows File Server or ONTAP storage virtual machine (SVM) instance is joined.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|windows_config_cq_id|uuid|Unique CloudQuery ID of aws_fsx_file_system_windows_configs table (FK)|
|active_directory_id|text|The Active Directory ID.|
|domain_name|text|The fully qualified domain name of the self-managed AD directory.|
|organizational_unit_distinguished_id|text|The fully qualified distinguished name of the organizational unit within the self-managed AD directory to which the Windows File Server or ONTAP storage virtual machine (SVM) instance is joined.|
|file_system_administrators_group|text|The name of the domain group whose members have administrative privileges for the FSx file system.|
|user_name|text|The user name for the service account on your self-managed AD domain that FSx uses to join to your AD domain.|
|dns_ips|text[]|A list of up to three IP addresses of DNS servers or domain controllers in the self-managed AD directory.|
