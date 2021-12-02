
# Table: aws_dynamodb_table_replica_auto_scaling_replica_auto_scaling_global_secondary_indexes
Represents the auto scaling configuration for a replica global secondary index.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|table_replica_auto_scaling_cq_id|uuid|Unique CloudQuery ID of aws_dynamodb_table_replica_auto_scalings table (FK)|
|index_name|text|The name of the global secondary index.|
|index_status|text|The current state of the replica global secondary index:  * CREATING - The index is being created.  * UPDATING - The index is being updated.  * DELETING - The index is being deleted.  * ACTIVE - The index is ready for use.|
|provisioned_read_capacity_auto_scaling_settings|jsonb|Represents the auto scaling settings for a global table or global secondary index.|
|provisioned_write_capacity_auto_scaling_settings|jsonb|Represents the auto scaling settings for a global table or global secondary index.|
