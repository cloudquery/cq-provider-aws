
# Table: aws_glue_database_table_partitions
Represents a slice of table data
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|database_table_cq_id|uuid|Unique CloudQuery ID of aws_glue_database_tables table (FK)|
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|parameters|jsonb||
|storage_parameters|jsonb||
|catalog_id|text|The ID of the Data Catalog in which the partition resides|
|creation_time|timestamp without time zone|The time at which the partition was created|
|database_name|text|The name of the catalog database in which to create the partition|
|last_access_time|timestamp without time zone|The last time at which the partition was accessed|
|last_analyzed_time|timestamp without time zone|The last time at which column statistics were computed for this partition|
|additional_locations|text[]|A list of locations that point to the path where a Delta table is located|
|bucket_columns|text[]|A list of reducer grouping columns, clustering columns, and bucketing columns in the table|
|compressed|boolean|True if the data in the table is compressed, or False if not|
|input_format|text|The input format: SequenceFileInputFormat (binary), or TextInputFormat, or a custom format|
|location|text|The physical location of the table|
|number_of_buckets|bigint|Must be specified if the table contains any dimension columns|
|output_format|text|The output format: SequenceFileOutputFormat (binary), or IgnoreKeyTextOutputFormat, or a custom format|
|schema_reference_schema_id|jsonb|A structure that contains schema identity fields|
|schema_reference_schema_version_id|text|The unique ID assigned to a version of the schema|
|schema_reference_schema_version_number|bigint|The version number of the schema|
|serde_info|jsonb|The serialization/deserialization (SerDe) information|
|skewed_info|jsonb|The information about values that appear frequently in a column (skewed values)|
|sort_columns|jsonb|A list specifying the sort order of each bucket in the table|
|stored_as_sub_directories|boolean|True if the table data is stored in subdirectories, or False if not|
|table_name|text|The name of the database table in which to create the partition|
|values|text[]|The values of the partition|
