
# Table: aws_rds_instances

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text||
|region|text||
|allocated_storage|integer||
|auto_minor_version_upgrade|boolean||
|availability_zone|text||
|aws_backup_recovery_point_arn|text||
|backup_retention_period|integer||
|ca_certificate_identifier|text||
|character_set_name|text||
|copy_tags_to_snapshot|boolean||
|customer_owned_ip_enabled|boolean||
|db_cluster_identifier|text||
|db_instance_arn|text||
|db_instance_class|text||
|db_instance_identifier|text||
|db_instance_status|text||
|db_name|text||
|db_subnet_group_arn|text||
|db_subnet_group_description|text||
|db_subnet_group_name|text||
|db_subnet_group_subnet_group_status|text||
|db_subnet_group_vpc_id|text||
|db_instance_port|integer||
|dbi_resource_id|text||
|deletion_protection|boolean||
|enabled_cloudwatch_logs_exports|text[]||
|endpoint_address|text||
|endpoint_hosted_zone_id|text||
|endpoint_port|integer||
|engine|text||
|engine_version|text||
|enhanced_monitoring_resource_arn|text||
|iam_database_authentication_enabled|boolean||
|instance_create_time|timestamp without time zone||
|iops|integer||
|kms_key_id|text||
|latest_restorable_time|timestamp without time zone||
|license_model|text||
|listener_endpoint_address|text||
|listener_endpoint_hosted_zone_id|text||
|listener_endpoint_port|integer||
|master_username|text||
|max_allocated_storage|integer||
|monitoring_interval|integer||
|monitoring_role_arn|text||
|multi_az|boolean||
|nchar_character_set_name|text||
|pending_modified_values_allocated_storage|integer||
|pending_modified_values_backup_retention_period|integer||
|pending_modified_values_ca_certificate_identifier|text||
|pending_modified_values_db_instance_class|text||
|pending_modified_values_db_instance_identifier|text||
|pending_modified_values_db_subnet_group_name|text||
|pending_modified_values_engine_version|text||
|pending_modified_values_iam_database_authentication_enabled|boolean||
|pending_modified_values_iops|integer||
|pending_modified_values_license_model|text||
|pending_modified_values_master_user_password|text||
|pending_modified_values_multi_az|boolean||
|pending_cloudwatch_logs_types_to_disable|text[]||
|pending_cloudwatch_logs_types_to_enable|text[]||
|pending_modified_values_port|integer||
|pending_modified_values_storage_type|text||
|performance_insights_enabled|boolean||
|performance_insights_kms_key_id|text||
|performance_insights_retention_period|integer||
|preferred_backup_window|text||
|preferred_maintenance_window|text||
|promotion_tier|integer||
|publicly_accessible|boolean||
|read_replica_db_cluster_identifiers|text[]||
|read_replica_db_instance_identifiers|text[]||
|read_replica_source_db_instance_identifier|text||
|replica_mode|text||
|secondary_availability_zone|text||
|storage_encrypted|boolean||
|storage_type|text||
|tags|jsonb||
|tde_credential_arn|text||
|timezone|text||
|aws_rds_instance_pending_modified_values_processor_features|jsonb||
|aws_rds_instance_processor_features|jsonb||
## Relations
## Table: aws_rds_instance_associated_roles

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|instance_id|uuid||
|feature_name|text||
|role_arn|text||
|status|text||
## Table: aws_rds_instance_db_instance_automated_backups_replications

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|instance_id|uuid||
|db_instance_automated_backups_arn|text||
## Table: aws_rds_instance_db_parameter_groups

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|instance_id|uuid||
|db_parameter_group_name|text||
|parameter_apply_status|text||
## Table: aws_rds_instance_db_security_groups

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|instance_id|uuid||
|db_security_group_name|text||
|status|text||
## Table: aws_rds_instance_db_subnet_group_subnets

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|instance_id|uuid||
|subnet_availability_zone_name|text||
|subnet_identifier|text||
|subnet_outpost_arn|text||
|subnet_status|text||
## Table: aws_rds_instance_domain_memberships

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|instance_id|uuid||
|domain|text||
|fqdn|text||
|iam_role_name|text||
|status|text||
## Table: aws_rds_instance_option_group_memberships

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|instance_id|uuid||
|option_group_name|text||
|status|text||
## Table: aws_rds_instance_status_infos

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|instance_id|uuid||
|message|text||
|normal|boolean||
|status|text||
|status_type|text||
## Table: aws_rds_instance_vpc_security_groups

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|instance_id|uuid||
|status|text||
|vpc_security_group_id|text||
