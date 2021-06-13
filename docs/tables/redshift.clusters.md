
# Table: aws_redshift_clusters

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text||
|region|text||
|allow_version_upgrade|boolean||
|automated_snapshot_retention_period|integer||
|availability_zone|text||
|availability_zone_relocation_status|text||
|cluster_availability_status|text||
|cluster_create_time|timestamp without time zone||
|cluster_identifier|text||
|cluster_namespace_arn|text||
|cluster_public_key|text||
|cluster_revision_number|text||
|cluster_snapshot_copy_status_destination_region|text||
|cluster_snapshot_copy_status_manual_snapshot_retention_period|integer||
|cluster_snapshot_copy_status_retention_period|bigint||
|cluster_snapshot_copy_status_snapshot_copy_grant_name|text||
|cluster_status|text||
|cluster_subnet_group_name|text||
|cluster_version|text||
|db_name|text||
|data_transfer_progress_current_rate_in_mega_bytes_per_second|float||
|data_transfer_progress_data_transferred_in_mega_bytes|bigint||
|data_transfer_progress_elapsed_time_in_seconds|bigint||
|data_transfer_progress_estimated_time_to_completion_in_seconds|bigint||
|data_transfer_progress_status|text||
|data_transfer_progress_total_data_in_mega_bytes|bigint||
|elastic_ip_status_elastic_ip|text||
|elastic_ip_status|text||
|elastic_resize_number_of_node_options|text||
|encrypted|boolean||
|endpoint_address|text||
|endpoint_port|integer||
|enhanced_vpc_routing|boolean||
|expected_next_snapshot_schedule_time|timestamp without time zone||
|expected_next_snapshot_schedule_time_status|text||
|hsm_status_hsm_client_certificate_identifier|text||
|hsm_status_hsm_configuration_identifier|text||
|hsm_status|text||
|kms_key_id|text||
|maintenance_track_name|text||
|manual_snapshot_retention_period|integer||
|master_username|text||
|modify_status|text||
|next_maintenance_window_start_time|timestamp without time zone||
|node_type|text||
|number_of_nodes|integer||
|pending_actions|text[]||
|pending_modified_values_automated_snapshot_retention_period|integer||
|pending_modified_values_cluster_identifier|text||
|pending_modified_values_cluster_type|text||
|pending_modified_values_cluster_version|text||
|pending_modified_values_encryption_type|text||
|pending_modified_values_enhanced_vpc_routing|boolean||
|pending_modified_values_maintenance_track_name|text||
|pending_modified_values_master_user_password|text||
|pending_modified_values_node_type|text||
|pending_modified_values_number_of_nodes|integer||
|pending_modified_values_publicly_accessible|boolean||
|preferred_maintenance_window|text||
|publicly_accessible|boolean||
|resize_info_allow_cancel_resize|boolean||
|resize_info_resize_type|text||
|restore_status_current_restore_rate_in_mega_bytes_per_second|float||
|restore_status_elapsed_time_in_seconds|bigint||
|restore_status_estimated_time_to_completion_in_seconds|bigint||
|restore_status_progress_in_mega_bytes|bigint||
|restore_status_snapshot_size_in_mega_bytes|bigint||
|restore_status|text||
|snapshot_schedule_identifier|text||
|snapshot_schedule_state|text||
|tags|jsonb||
|total_storage_capacity_in_mega_bytes|bigint||
|vpc_id|text||
## Relations
## Table: aws_redshift_cluster_nodes

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|cluster_id|uuid||
|node_role|text||
|private_ip_address|text||
|public_ip_address|text||
## Table: aws_redshift_cluster_parameter_groups

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|cluster_id|uuid||
|parameter_apply_status|text||
|parameter_group_name|text||
## Relations
## Table: aws_redshift_cluster_parameter_group_status_lists

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|cluster_parameter_group_id|uuid||
|parameter_apply_error_description|text||
|parameter_apply_status|text||
|parameter_name|text||
## Table: aws_redshift_cluster_security_groups

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|cluster_id|uuid||
|cluster_security_group_name|text||
|status|text||
## Table: aws_redshift_cluster_deferred_maintenance_windows

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|cluster_id|uuid||
|defer_maintenance_end_time|timestamp without time zone||
|defer_maintenance_identifier|text||
|defer_maintenance_start_time|timestamp without time zone||
## Table: aws_redshift_cluster_endpoint_vpc_endpoints

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|cluster_id|uuid||
|vpc_endpoint_id|text||
|vpc_id|text||
## Relations
## Table: aws_redshift_cluster_endpoint_vpc_endpoint_network_interfaces

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|clusterendpoint_vpc_endpoint_id|uuid||
|availability_zone|text||
|network_interface_id|text||
|private_ip_address|text||
|subnet_id|text||
## Table: aws_redshift_cluster_iam_roles

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|cluster_id|uuid||
|apply_status|text||
|iam_role_arn|text||
## Table: aws_redshift_cluster_vpc_security_groups

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|cluster_id|uuid||
|status|text||
|vpc_security_group_id|text||
