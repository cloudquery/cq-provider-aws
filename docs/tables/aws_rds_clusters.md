
# Table: aws_rds_clusters

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text||
|region|text||
|activity_stream_kinesis_stream_name|text||
|activity_stream_kms_key_id|text||
|activity_stream_mode|text||
|activity_stream_status|text||
|allocated_storage|integer||
|availability_zones|text[]||
|backtrack_consumed_change_records|bigint||
|backtrack_window|bigint||
|backup_retention_period|integer||
|capacity|integer||
|character_set_name|text||
|clone_group_id|text||
|cluster_create_time|timestamp without time zone||
|copy_tags_to_snapshot|boolean||
|cross_account_clone|boolean||
|custom_endpoints|text[]||
|db_cluster_arn|text||
|db_cluster_identifier|text||
|db_cluster_parameter_group|text||
|db_subnet_group|text||
|database_name|text||
|db_cluster_resource_id|text||
|deletion_protection|boolean||
|earliest_backtrack_time|timestamp without time zone||
|earliest_restorable_time|timestamp without time zone||
|enabled_cloudwatch_logs_exports|text[]||
|endpoint|text||
|engine|text||
|engine_mode|text||
|engine_version|text||
|global_write_forwarding_requested|boolean||
|global_write_forwarding_status|text||
|hosted_zone_id|text||
|http_endpoint_enabled|boolean||
|iam_database_authentication_enabled|boolean||
|kms_key_id|text||
|latest_restorable_time|timestamp without time zone||
|master_username|text||
|multi_az|boolean||
|pending_modified_values_db_cluster_identifier|text||
|pending_modified_values_engine_version|text||
|pending_modified_values_iam_database_authentication_enabled|boolean||
|pending_modified_values_master_user_password|text||
|pending_cloudwatch_logs_types_to_disable|text[]||
|pending_cloudwatch_logs_types_to_enable|text[]||
|percent_progress|text||
|port|integer||
|preferred_backup_window|text||
|preferred_maintenance_window|text||
|read_replica_identifiers|text[]||
|reader_endpoint|text||
|replication_source_identifier|text||
|scaling_configuration_info_auto_pause|boolean||
|scaling_configuration_info_max_capacity|integer||
|scaling_configuration_info_min_capacity|integer||
|scaling_configuration_info_seconds_until_auto_pause|integer||
|scaling_configuration_info_timeout_action|text||
|status|text||
|storage_encrypted|boolean||
|tags|jsonb||
