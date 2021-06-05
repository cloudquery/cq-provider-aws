# public.aws_redshift_clusters

## Description

## Columns

| Name | Type | Default | Nullable | Children | Parents | Comment |
| ---- | ---- | ------- | -------- | -------- | ------- | ------- |
| id | uuid |  | false | [public.aws_redshift_cluster_nodes](public.aws_redshift_cluster_nodes.md) [public.aws_redshift_cluster_parameter_groups](public.aws_redshift_cluster_parameter_groups.md) [public.aws_redshift_cluster_security_groups](public.aws_redshift_cluster_security_groups.md) [public.aws_redshift_cluster_deferred_maintenance_windows](public.aws_redshift_cluster_deferred_maintenance_windows.md) [public.aws_redshift_cluster_endpoint_vpc_endpoints](public.aws_redshift_cluster_endpoint_vpc_endpoints.md) [public.aws_redshift_cluster_iam_roles](public.aws_redshift_cluster_iam_roles.md) [public.aws_redshift_cluster_vpc_security_groups](public.aws_redshift_cluster_vpc_security_groups.md) |  |  |
| account_id | text |  | true |  |  |  |
| region | text |  | true |  |  |  |
| allow_version_upgrade | boolean |  | true |  |  |  |
| automated_snapshot_retention_period | integer |  | true |  |  |  |
| availability_zone | text |  | true |  |  |  |
| availability_zone_relocation_status | text |  | true |  |  |  |
| cluster_availability_status | text |  | true |  |  |  |
| cluster_create_time | timestamp without time zone |  | true |  |  |  |
| cluster_identifier | text |  | true |  |  |  |
| cluster_namespace_arn | text |  | true |  |  |  |
| cluster_public_key | text |  | true |  |  |  |
| cluster_revision_number | text |  | true |  |  |  |
| cluster_snapshot_copy_status_destination_region | text |  | true |  |  |  |
| cluster_snapshot_copy_status_manual_snapshot_retention_period | integer |  | true |  |  |  |
| cluster_snapshot_copy_status_retention_period | bigint |  | true |  |  |  |
| cluster_snapshot_copy_status_snapshot_copy_grant_name | text |  | true |  |  |  |
| cluster_status | text |  | true |  |  |  |
| cluster_subnet_group_name | text |  | true |  |  |  |
| cluster_version | text |  | true |  |  |  |
| db_name | text |  | true |  |  |  |
| data_transfer_progress_current_rate_in_mega_bytes_per_second | double precision |  | true |  |  |  |
| data_transfer_progress_data_transferred_in_mega_bytes | bigint |  | true |  |  |  |
| data_transfer_progress_elapsed_time_in_seconds | bigint |  | true |  |  |  |
| data_transfer_progress_estimated_time_to_completion_in_seconds | bigint |  | true |  |  |  |
| data_transfer_progress_status | text |  | true |  |  |  |
| data_transfer_progress_total_data_in_mega_bytes | bigint |  | true |  |  |  |
| elastic_ip_status_elastic_ip | text |  | true |  |  |  |
| elastic_ip_status | text |  | true |  |  |  |
| elastic_resize_number_of_node_options | text |  | true |  |  |  |
| encrypted | boolean |  | true |  |  |  |
| endpoint_address | text |  | true |  |  |  |
| endpoint_port | integer |  | true |  |  |  |
| enhanced_vpc_routing | boolean |  | true |  |  |  |
| expected_next_snapshot_schedule_time | timestamp without time zone |  | true |  |  |  |
| expected_next_snapshot_schedule_time_status | text |  | true |  |  |  |
| hsm_status_hsm_client_certificate_identifier | text |  | true |  |  |  |
| hsm_status_hsm_configuration_identifier | text |  | true |  |  |  |
| hsm_status | text |  | true |  |  |  |
| kms_key_id | text |  | true |  |  |  |
| maintenance_track_name | text |  | true |  |  |  |
| manual_snapshot_retention_period | integer |  | true |  |  |  |
| master_username | text |  | true |  |  |  |
| modify_status | text |  | true |  |  |  |
| next_maintenance_window_start_time | timestamp without time zone |  | true |  |  |  |
| node_type | text |  | true |  |  |  |
| number_of_nodes | integer |  | true |  |  |  |
| pending_actions | text[] |  | true |  |  |  |
| pending_modified_values_automated_snapshot_retention_period | integer |  | true |  |  |  |
| pending_modified_values_cluster_identifier | text |  | true |  |  |  |
| pending_modified_values_cluster_type | text |  | true |  |  |  |
| pending_modified_values_cluster_version | text |  | true |  |  |  |
| pending_modified_values_encryption_type | text |  | true |  |  |  |
| pending_modified_values_enhanced_vpc_routing | boolean |  | true |  |  |  |
| pending_modified_values_maintenance_track_name | text |  | true |  |  |  |
| pending_modified_values_master_user_password | text |  | true |  |  |  |
| pending_modified_values_node_type | text |  | true |  |  |  |
| pending_modified_values_number_of_nodes | integer |  | true |  |  |  |
| pending_modified_values_publicly_accessible | boolean |  | true |  |  |  |
| preferred_maintenance_window | text |  | true |  |  |  |
| publicly_accessible | boolean |  | true |  |  |  |
| resize_info_allow_cancel_resize | boolean |  | true |  |  |  |
| resize_info_resize_type | text |  | true |  |  |  |
| restore_status_current_restore_rate_in_mega_bytes_per_second | double precision |  | true |  |  |  |
| restore_status_elapsed_time_in_seconds | bigint |  | true |  |  |  |
| restore_status_estimated_time_to_completion_in_seconds | bigint |  | true |  |  |  |
| restore_status_progress_in_mega_bytes | bigint |  | true |  |  |  |
| restore_status_snapshot_size_in_mega_bytes | bigint |  | true |  |  |  |
| restore_status | text |  | true |  |  |  |
| snapshot_schedule_identifier | text |  | true |  |  |  |
| snapshot_schedule_state | text |  | true |  |  |  |
| tags | jsonb |  | true |  |  |  |
| total_storage_capacity_in_mega_bytes | bigint |  | true |  |  |  |
| vpc_id | text |  | true |  |  |  |

## Constraints

| Name | Type | Definition |
| ---- | ---- | ---------- |
| aws_redshift_clusters_pkey | PRIMARY KEY | PRIMARY KEY (id) |

## Indexes

| Name | Definition |
| ---- | ---------- |
| aws_redshift_clusters_pkey | CREATE UNIQUE INDEX aws_redshift_clusters_pkey ON public.aws_redshift_clusters USING btree (id) |

---

> Generated by [tbls](https://github.com/k1LoW/tbls)
