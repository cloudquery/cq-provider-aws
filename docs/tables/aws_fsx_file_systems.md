
# Table: aws_fsx_file_systems
An Amazon FSx file system.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|id|text|The system-generated, unique 17-digit ID of the file system.|
|creation_time|timestamp without time zone|The time that the file system was created.|
|owner_id|text|The AWS account that created the file system.|
|file_system_type|text|The type of Amazon FSx file system.|
|lifecycle|text|The lifecycle status of the file system.|
|failure_details_message|text|A message describing the data repository association failure.|
|storage_capacity|integer|The storage capacity of the file system in gibibytes (GiB).|
|storage_type|text|The type of storage the file system is using.|
|vpc_id|text|The ID of the primary virtual private cloud (VPC) for the file system.|
|subnet_ids|text[]|Specifies the IDs of the subnets that the file system is accessible from.|
|network_interface_ids|text[]|The IDs of the elastic network interfaces from which a specific file system is accessible.|
|dns_name|text|The Domain Name System (DNS) name for the file system.|
|kms_key_id|text|The ID of the Key Management Service (KMS) key used to encrypt Amazon FSx file system data.|
|arn|text|The Amazon Resource Name (ARN) of the file system resource.|
|tags|jsonb|Tags associated with a particular file system.|
|file_system_type_version|text|The Lustre version of the Amazon FSx for Lustre file system, either 2.10 or 2.12 .|
