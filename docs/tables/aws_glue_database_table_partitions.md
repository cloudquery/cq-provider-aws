
# Table: aws_glue_database_table_partitions
Represents a slice of table data
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|database_table_cq_id|uuid|Unique CloudQuery ID of aws_glue_database_tables table (FK)|
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|catalog_id|text|The ID of the Data Catalog in which the partition resides|
|creation_time|timestamp without time zone|The time at which the partition was created|
|database_name|text|The name of the catalog database in which to create the partition|
|last_access_time|timestamp without time zone|The last time at which the partition was accessed|
|last_analyzed_time|timestamp without time zone|The last time at which column statistics were computed for this partition|
|parameters|jsonb|These key-value pairs define partition parameters|
|storage_descriptor_additional_locations|text[]|A list of locations that point to the path where a Delta table is located|
|storage_descriptor_bucket_columns|text[]|A list of reducer grouping columns, clustering columns, and bucketing columns in the table|
|storage_descriptor_compressed|boolean|True if the data in the table is compressed, or False if not|
|storage_descriptor_input_format|text|The input format: SequenceFileInputFormat (binary), or TextInputFormat, or a custom format|
|storage_descriptor_location|text|The physical location of the table|
|storage_descriptor_number_of_buckets|bigint|Must be specified if the table contains any dimension columns|
|storage_descriptor_output_format|text|The output format: SequenceFileOutputFormat (binary), or IgnoreKeyTextOutputFormat, or a custom format|
|storage_descriptor_parameters|jsonb|The user-supplied properties in key-value form|
|storage_descriptor_schema_reference_schema_id_registry_name|text|The name of the schema registry that contains the schema|
|storage_descriptor_schema_reference_schema_id_schema_arn|text|The Amazon Resource Name (ARN) of the schema|
|storage_descriptor_schema_reference_schema_id_schema_name|text|The name of the schema|
|storage_descriptor_schema_reference_schema_version_id|text|The unique ID assigned to a version of the schema|
|storage_descriptor_schema_reference_schema_version_number|bigint|The version number of the schema|
|storage_descriptor_serde_info_name|text|Name of the SerDe|
|storage_descriptor_serde_info_parameters|jsonb|These key-value pairs define initialization parameters for the SerDe|
|storage_descriptor_serde_info_serialization_library|text|Usually the class that implements the SerDe|
|storage_descriptor_skewed_info|jsonb|The information about values that appear frequently in a column (skewed values)|
|storage_descriptor_sort_columns|jsonb|A list specifying the sort order of each bucket in the table|
|storage_descriptor_stored_as_sub_directories|boolean|True if the table data is stored in subdirectories, or False if not|
|table_name|text|The name of the database table in which to create the partition|
|values|text[]|The values of the partition|
