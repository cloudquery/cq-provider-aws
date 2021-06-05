# public.aws_rds_instances

## Description

## Columns

| Name | Type | Default | Nullable | Children | Parents | Comment |
| ---- | ---- | ------- | -------- | -------- | ------- | ------- |
| id | uuid |  | false | [public.aws_rds_instance_associated_roles](public.aws_rds_instance_associated_roles.md) [public.aws_rds_instance_db_instance_automated_backups_replications](public.aws_rds_instance_db_instance_automated_backups_replications.md) [public.aws_rds_instance_db_parameter_groups](public.aws_rds_instance_db_parameter_groups.md) [public.aws_rds_instance_db_security_groups](public.aws_rds_instance_db_security_groups.md) [public.aws_rds_instance_db_subnet_group_subnets](public.aws_rds_instance_db_subnet_group_subnets.md) [public.aws_rds_instance_domain_memberships](public.aws_rds_instance_domain_memberships.md) [public.aws_rds_instance_option_group_memberships](public.aws_rds_instance_option_group_memberships.md) [public.aws_rds_instance_status_infos](public.aws_rds_instance_status_infos.md) [public.aws_rds_instance_vpc_security_groups](public.aws_rds_instance_vpc_security_groups.md) |  |  |
| account_id | text |  | true |  |  |  |
| region | text |  | true |  |  |  |
| allocated_storage | integer |  | true |  |  |  |
| auto_minor_version_upgrade | boolean |  | true |  |  |  |
| availability_zone | text |  | true |  |  |  |
| aws_backup_recovery_point_arn | text |  | true |  |  |  |
| backup_retention_period | integer |  | true |  |  |  |
| ca_certificate_identifier | text |  | true |  |  |  |
| character_set_name | text |  | true |  |  |  |
| copy_tags_to_snapshot | boolean |  | true |  |  |  |
| customer_owned_ip_enabled | boolean |  | true |  |  |  |
| db_cluster_identifier | text |  | true |  |  |  |
| db_instance_arn | text |  | true |  |  |  |
| db_instance_class | text |  | true |  |  |  |
| db_instance_identifier | text |  | true |  |  |  |
| db_instance_status | text |  | true |  |  |  |
| db_name | text |  | true |  |  |  |
| db_subnet_group_arn | text |  | true |  |  |  |
| db_subnet_group_description | text |  | true |  |  |  |
| db_subnet_group_name | text |  | true |  |  |  |
| db_subnet_group_subnet_group_status | text |  | true |  |  |  |
| db_subnet_group_vpc_id | text |  | true |  |  |  |
| db_instance_port | integer |  | true |  |  |  |
| dbi_resource_id | text |  | true |  |  |  |
| deletion_protection | boolean |  | true |  |  |  |
| enabled_cloudwatch_logs_exports | text[] |  | true |  |  |  |
| endpoint_address | text |  | true |  |  |  |
| endpoint_hosted_zone_id | text |  | true |  |  |  |
| endpoint_port | integer |  | true |  |  |  |
| engine | text |  | true |  |  |  |
| engine_version | text |  | true |  |  |  |
| enhanced_monitoring_resource_arn | text |  | true |  |  |  |
| iam_database_authentication_enabled | boolean |  | true |  |  |  |
| instance_create_time | timestamp without time zone |  | true |  |  |  |
| iops | integer |  | true |  |  |  |
| kms_key_id | text |  | true |  |  |  |
| latest_restorable_time | timestamp without time zone |  | true |  |  |  |
| license_model | text |  | true |  |  |  |
| listener_endpoint_address | text |  | true |  |  |  |
| listener_endpoint_hosted_zone_id | text |  | true |  |  |  |
| listener_endpoint_port | integer |  | true |  |  |  |
| master_username | text |  | true |  |  |  |
| max_allocated_storage | integer |  | true |  |  |  |
| monitoring_interval | integer |  | true |  |  |  |
| monitoring_role_arn | text |  | true |  |  |  |
| multi_az | boolean |  | true |  |  |  |
| nchar_character_set_name | text |  | true |  |  |  |
| pending_modified_values_allocated_storage | integer |  | true |  |  |  |
| pending_modified_values_backup_retention_period | integer |  | true |  |  |  |
| pending_modified_values_ca_certificate_identifier | text |  | true |  |  |  |
| pending_modified_values_db_instance_class | text |  | true |  |  |  |
| pending_modified_values_db_instance_identifier | text |  | true |  |  |  |
| pending_modified_values_db_subnet_group_name | text |  | true |  |  |  |
| pending_modified_values_engine_version | text |  | true |  |  |  |
| pending_modified_values_iam_database_authentication_enabled | boolean |  | true |  |  |  |
| pending_modified_values_iops | integer |  | true |  |  |  |
| pending_modified_values_license_model | text |  | true |  |  |  |
| pending_modified_values_master_user_password | text |  | true |  |  |  |
| pending_modified_values_multi_az | boolean |  | true |  |  |  |
| pending_cloudwatch_logs_types_to_disable | text[] |  | true |  |  |  |
| pending_cloudwatch_logs_types_to_enable | text[] |  | true |  |  |  |
| pending_modified_values_port | integer |  | true |  |  |  |
| pending_modified_values_storage_type | text |  | true |  |  |  |
| performance_insights_enabled | boolean |  | true |  |  |  |
| performance_insights_kms_key_id | text |  | true |  |  |  |
| performance_insights_retention_period | integer |  | true |  |  |  |
| preferred_backup_window | text |  | true |  |  |  |
| preferred_maintenance_window | text |  | true |  |  |  |
| promotion_tier | integer |  | true |  |  |  |
| publicly_accessible | boolean |  | true |  |  |  |
| read_replica_db_cluster_identifiers | text[] |  | true |  |  |  |
| read_replica_db_instance_identifiers | text[] |  | true |  |  |  |
| read_replica_source_db_instance_identifier | text |  | true |  |  |  |
| replica_mode | text |  | true |  |  |  |
| secondary_availability_zone | text |  | true |  |  |  |
| storage_encrypted | boolean |  | true |  |  |  |
| storage_type | text |  | true |  |  |  |
| tags | jsonb |  | true |  |  |  |
| tde_credential_arn | text |  | true |  |  |  |
| timezone | text |  | true |  |  |  |
| aws_rds_instance_pending_modified_values_processor_features | jsonb |  | true |  |  |  |
| aws_rds_instance_processor_features | jsonb |  | true |  |  |  |

## Constraints

| Name | Type | Definition |
| ---- | ---- | ---------- |
| aws_rds_instances_pkey | PRIMARY KEY | PRIMARY KEY (id) |

## Indexes

| Name | Definition |
| ---- | ---------- |
| aws_rds_instances_pkey | CREATE UNIQUE INDEX aws_rds_instances_pkey ON public.aws_rds_instances USING btree (id) |

---

> Generated by [tbls](https://github.com/k1LoW/tbls)
