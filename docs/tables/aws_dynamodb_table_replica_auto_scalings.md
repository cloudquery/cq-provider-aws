
# Table: aws_dynamodb_table_replica_auto_scalings
Represents the auto scaling settings of the replica.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|table_cq_id|uuid|Unique CloudQuery ID of aws_dynamodb_tables table (FK)|
|region_name|text|The Region where the replica exists.|
|replica_provisioned_read_capacity_auto_scaling_settings|jsonb|Represents the auto scaling settings for a global table or global secondary index.|
|replica_provisioned_write_capacity_auto_scaling_settings|jsonb|Represents the auto scaling settings for a global table or global secondary index.|
|replica_status|text|The current state of the replica:  * CREATING - The replica is being created.  * UPDATING - The replica is being updated.  * DELETING - The replica is being deleted.  * ACTIVE - The replica is ready for use.|
