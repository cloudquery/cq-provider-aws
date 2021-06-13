
# Table: aws_s3_bucket_replication_rules

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|bucket_id|uuid||
|destination_bucket|text||
|destination_access_control_translation_owner|text||
|destination_account|text||
|destination_encryption_configuration_replica_kms_key_id|text||
|destination_metrics_status|text||
|destination_metrics_event_threshold_minutes|integer||
|destination_replication_time_status|text||
|destination_replication_time_minutes|integer||
|destination_storage_class|text||
|status|text||
|delete_marker_replication_status|text||
|existing_object_replication_status|text||
|filter|jsonb||
|resource_id|text||
|prefix|text||
|priority|integer||
|source_replica_modifications_status|text||
|source_sse_kms_encrypted_objects_status|text||
