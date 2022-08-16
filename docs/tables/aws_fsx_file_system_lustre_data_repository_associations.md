
# Table: aws_fsx_file_system_lustre_data_repository_associations
The data repository configuration object for Lustre file systems returned in the response of the CreateFileSystem operation.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|lustre_config_cq_id|uuid|Unique CloudQuery ID of aws_fsx_file_system_lustre_configs table (FK)|
|import_path|text|The import path to the Amazon S3 bucket (and optional prefix) that you're using as the data repository for your FSx for Lustre file system.|
|export_path|text|The export path to the Amazon S3 bucket (and prefix) that you are using to store new and changed Lustre file system files in S3.|
|imported_file_chunk_size|integer|For files imported from a data repository, this value determines the stripe count and maximum amount of data per file (in MiB) stored on a single physical disk.|
|auto_import_policy|text|Describes the file system's linked S3 data repository's AutoImportPolicy. The AutoImportPolicy configures how Amazon FSx keeps your file and directory listings up to date as you add or modify objects in your linked S3 bucket.|
|failure_details_message|text|A message describing the data repository association failure.|
