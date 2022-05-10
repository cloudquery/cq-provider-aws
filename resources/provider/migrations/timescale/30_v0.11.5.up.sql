-- Resource: shield.attacks
CREATE TABLE IF NOT EXISTS "aws_shield_attacks" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"cq_fetch_date" timestamp WITHOUT TIME ZONE NOT NULL,
	"account_id" text,
	"attack_counters" jsonb,
	"id" text,
	"end_time" timestamp WITHOUT TIME ZONE,
	"mitigations" text[],
	"resource_arn" text,
	"start_time" timestamp WITHOUT TIME ZONE,
	CONSTRAINT aws_shield_attacks_pk PRIMARY KEY (cq_fetch_date, id),
	UNIQUE (cq_fetch_date, cq_id)
);
SELECT setup_tsdb_parent('aws_shield_attacks');
CREATE TABLE IF NOT EXISTS "aws_shield_attack_properties" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"cq_fetch_date" timestamp WITHOUT TIME ZONE NOT NULL,
	"attack_cq_id" uuid,
	"attack_layer" text,
	"attack_property_identifier" text,
	"top_contributors" jsonb,
	"total" bigint,
	"unit" text,
	CONSTRAINT aws_shield_attack_properties_pk PRIMARY KEY (cq_fetch_date, cq_id),
	UNIQUE (cq_fetch_date, cq_id)
);
CREATE INDEX ON aws_shield_attack_properties(cq_fetch_date, attack_cq_id);
SELECT setup_tsdb_child('aws_shield_attack_properties', 'attack_cq_id', 'aws_shield_attacks', 'cq_id');
CREATE TABLE IF NOT EXISTS "aws_shield_attack_sub_resources" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"cq_fetch_date" timestamp WITHOUT TIME ZONE NOT NULL,
	"attack_cq_id" uuid,
	"attack_vectors" jsonb,
	"counters" jsonb,
	"id" text,
	"type" text,
	CONSTRAINT aws_shield_attack_sub_resources_pk PRIMARY KEY (cq_fetch_date, cq_id),
	UNIQUE (cq_fetch_date, cq_id)
);
CREATE INDEX ON aws_shield_attack_sub_resources(cq_fetch_date, attack_cq_id);
SELECT setup_tsdb_child('aws_shield_attack_sub_resources', 'attack_cq_id', 'aws_shield_attacks', 'cq_id');

-- Resource: shield.protections
CREATE TABLE IF NOT EXISTS "aws_shield_protections" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"cq_fetch_date" timestamp WITHOUT TIME ZONE NOT NULL,
	"account_id" text,
	"region" text,
	"tags" jsonb,
	"application_automatic_response_configuration_status" text,
	"health_check_ids" text[],
	"id" text,
	"name" text,
	"arn" text,
	"resource_arn" text,
	CONSTRAINT aws_shield_protections_pk PRIMARY KEY (cq_fetch_date, arn),
	UNIQUE (cq_fetch_date, cq_id)
);
SELECT setup_tsdb_parent('aws_shield_protections');

-- Resource: shield.protections_groups
CREATE TABLE IF NOT EXISTS "aws_shield_protection_groups" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"cq_fetch_date" timestamp WITHOUT TIME ZONE NOT NULL,
	"account_id" text,
	"aggregation" text,
	"members" text[],
	"pattern" text,
	"id" text,
	"arn" text,
	"resource_type" text,
	"tags" jsonb,
	CONSTRAINT aws_shield_protection_groups_pk PRIMARY KEY (cq_fetch_date, arn),
	UNIQUE (cq_fetch_date, cq_id)
);
SELECT setup_tsdb_parent('aws_shield_protection_groups');

-- Resource: shield.subscriptions
CREATE TABLE IF NOT EXISTS "aws_shield_subscriptions" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"cq_fetch_date" timestamp WITHOUT TIME ZONE NOT NULL,
	"account_id" text,
	"protection_group_limits_max_protection_groups" int,
	"protection_group_limits_arbitrary_pattern_limits_max_members" int,
	"protected_resource_type_limits" jsonb,
	"auto_renew" text,
	"end_time" timestamp WITHOUT TIME ZONE,
	"limits" jsonb,
	"proactive_engagement_status" text,
	"start_time" timestamp WITHOUT TIME ZONE,
	"arn" text,
	"time_commitment_in_seconds" int,
	CONSTRAINT aws_shield_subscriptions_pk PRIMARY KEY (cq_fetch_date, arn),
	UNIQUE (cq_fetch_date, cq_id)
);
SELECT setup_tsdb_parent('aws_shield_subscriptions');

-- Resource: athena.data_catalogs
CREATE TABLE IF NOT EXISTS "aws_athena_data_catalogs"(
    "cq_id" uuid NOT NULL,
    "cq_meta" jsonb,
    "cq_fetch_date" timestamp WITHOUT TIME ZONE NOT NULL,
    "account_id" text,
    "region" text,
    "tags" jsonb,
    "name" text,
    "type" text,
    "description" text,
    "parameters" jsonb,
    CONSTRAINT aws_athena_data_catalogs_pk PRIMARY KEY(cq_fetch_date,account_id,region,name),
    UNIQUE(cq_fetch_date,cq_id)
);
SELECT setup_tsdb_parent('aws_athena_data_catalogs');
CREATE TABLE IF NOT EXISTS "aws_athena_data_catalog_databases"(
    "cq_id" uuid NOT NULL,
    "cq_meta" jsonb,
    "cq_fetch_date" timestamp WITHOUT TIME ZONE NOT NULL,
    "data_catalog_cq_id" uuid,
    "name" text,
    "description" text,
    "parameters" jsonb,
    CONSTRAINT aws_athena_data_catalog_databases_pk PRIMARY KEY(cq_fetch_date,cq_id),
    UNIQUE(cq_fetch_date,cq_id)
);
CREATE INDEX ON aws_athena_data_catalog_databases(cq_fetch_date,data_catalog_cq_id);
SELECT setup_tsdb_child('aws_athena_data_catalog_databases','data_catalog_cq_id','aws_athena_data_catalogs','cq_id');
CREATE TABLE IF NOT EXISTS "aws_athena_data_catalog_database_tables"(
    "cq_id" uuid NOT NULL,
    "cq_meta" jsonb,
    "cq_fetch_date" timestamp WITHOUT TIME ZONE NOT NULL,
    "data_catalog_database_cq_id" uuid,
    "name" text,
    "create_time" timestamp WITHOUT TIME ZONE,
    "last_access_time" timestamp WITHOUT TIME ZONE,
    "parameters" jsonb,
    "table_type" text,
    CONSTRAINT aws_athena_data_catalog_database_tables_pk PRIMARY KEY(cq_fetch_date,cq_id),
    UNIQUE(cq_fetch_date,cq_id)
);
CREATE INDEX ON aws_athena_data_catalog_database_tables(cq_fetch_date,data_catalog_database_cq_id);
SELECT setup_tsdb_child('aws_athena_data_catalog_database_tables','data_catalog_database_cq_id',
                        'aws_athena_data_catalog_databases','cq_id');
CREATE TABLE IF NOT EXISTS "aws_athena_data_catalog_database_table_columns"(
    "cq_id" uuid NOT NULL,
    "cq_meta" jsonb,
    "cq_fetch_date" timestamp WITHOUT TIME ZONE NOT NULL,
    "data_catalog_database_table_cq_id" uuid,
    "name" text,
    "comment" text,
    "type" text,
    CONSTRAINT aws_athena_data_catalog_database_table_columns_pk PRIMARY KEY(cq_fetch_date,cq_id),
    UNIQUE(cq_fetch_date,cq_id)
);
CREATE INDEX ON aws_athena_data_catalog_database_table_columns(cq_fetch_date,data_catalog_database_table_cq_id);
SELECT setup_tsdb_child('aws_athena_data_catalog_database_table_columns','data_catalog_database_table_cq_id',
                        'aws_athena_data_catalog_database_tables','cq_id');
CREATE TABLE IF NOT EXISTS "aws_athena_data_catalog_database_table_partition_keys"(
    "cq_id" uuid NOT NULL,
    "cq_meta" jsonb,
    "cq_fetch_date" timestamp WITHOUT TIME ZONE NOT NULL,
    "data_catalog_database_table_cq_id" uuid,
    "name" text,
    "comment" text,
    "type" text,
    CONSTRAINT aws_athena_data_catalog_database_table_partition_keys_pk PRIMARY KEY(cq_fetch_date,cq_id),
    UNIQUE(cq_fetch_date,cq_id)
);
CREATE INDEX ON aws_athena_data_catalog_database_table_partition_keys(cq_fetch_date,data_catalog_database_table_cq_id);
SELECT setup_tsdb_child('aws_athena_data_catalog_database_table_partition_keys','data_catalog_database_table_cq_id',
                        'aws_athena_data_catalog_database_tables','cq_id');

-- Resource: athena.work_groups
CREATE TABLE IF NOT EXISTS "aws_athena_work_groups"(
    "cq_id" uuid NOT NULL,
    "cq_meta" jsonb,
    "cq_fetch_date" timestamp WITHOUT TIME ZONE NOT NULL,
    "account_id" text,
    "region" text,
    "tags" jsonb,
    "name" text,
    "bytes_scanned_cutoff_per_query" bigint,
    "enforce_work_group_configuration" boolean,
    "effective_engine_version" text,
    "selected_engine_version" text,
    "publish_cloud_watch_metrics_enabled" boolean,
    "requester_pays_enabled" boolean,
    "acl_configuration_s3_acl_option" text,
    "encryption_configuration_encryption_option" text,
    "encryption_configuration_kms_key" text,
    "expected_bucket_owner" text,
    "output_location" text,
    "creation_time" timestamp WITHOUT TIME ZONE,
    "description" text,
    "state" text,
    CONSTRAINT aws_athena_work_groups_pk PRIMARY KEY(cq_fetch_date,account_id,region,name),
    UNIQUE(cq_fetch_date,cq_id)
);
SELECT setup_tsdb_parent('aws_athena_work_groups');
CREATE TABLE IF NOT EXISTS "aws_athena_work_group_prepared_statements"(
    "cq_id" uuid NOT NULL,
    "cq_meta" jsonb,
    "cq_fetch_date" timestamp WITHOUT TIME ZONE NOT NULL,
    "work_group_cq_id" uuid,
    "description" text,
    "last_modified_time" timestamp WITHOUT TIME ZONE,
    "query_statement" text,
    "statement_name" text,
    "work_group_name" text,
    CONSTRAINT aws_athena_work_group_prepared_statements_pk PRIMARY KEY(cq_fetch_date,cq_id),
    UNIQUE(cq_fetch_date,cq_id)
);
CREATE INDEX ON aws_athena_work_group_prepared_statements(cq_fetch_date,work_group_cq_id);
SELECT setup_tsdb_child('aws_athena_work_group_prepared_statements','work_group_cq_id','aws_athena_work_groups',
                        'cq_id');
CREATE TABLE IF NOT EXISTS "aws_athena_work_group_query_executions"(
    "cq_id" uuid NOT NULL,
    "cq_meta" jsonb,
    "cq_fetch_date" timestamp WITHOUT TIME ZONE NOT NULL,
    "work_group_cq_id" uuid,
    "effective_engine_version" text,
    "selected_engine_version" text,
    "query" text,
    "catalog" text,
    "database" text,
    "id" text,
    "acl_configuration_s3_acl_option" text,
    "encryption_configuration_encryption_option" text,
    "encryption_configuration_kms_key" text,
    "expected_bucket_owner" text,
    "output_location" text,
    "statement_type" text,
    "data_manifest_location" text,
    "data_scanned_in_bytes" bigint,
    "engine_execution_time_in_millis" bigint,
    "query_planning_time_in_millis" bigint,
    "query_queue_time_in_millis" bigint,
    "service_processing_time_in_millis" bigint,
    "total_execution_time_in_millis" bigint,
    "athena_error_error_category" integer,
    "athena_error_error_message" text,
    "athena_error_error_type" integer,
    "athena_error_retryable" boolean,
    "completion_date_time" timestamp WITHOUT TIME ZONE,
    "state" text,
    "state_change_reason" text,
    "submission_date_time" timestamp WITHOUT TIME ZONE,
    "work_group" text,
    CONSTRAINT aws_athena_work_group_query_executions_pk PRIMARY KEY(cq_fetch_date,cq_id),
    UNIQUE(cq_fetch_date,cq_id)
);
CREATE INDEX ON aws_athena_work_group_query_executions(cq_fetch_date,work_group_cq_id);
SELECT setup_tsdb_child('aws_athena_work_group_query_executions','work_group_cq_id','aws_athena_work_groups',
                        'cq_id');
CREATE TABLE IF NOT EXISTS "aws_athena_work_group_named_queries"(
    "cq_id" uuid NOT NULL,
    "cq_meta" jsonb,
    "cq_fetch_date" timestamp WITHOUT TIME ZONE NOT NULL,
    "work_group_cq_id" uuid,
    "database" text,
    "name" text,
    "query_string" text,
    "description" text,
    "named_query_id" text,
    "work_group" text,
    CONSTRAINT aws_athena_work_group_named_queries_pk PRIMARY KEY(cq_fetch_date,cq_id),
    UNIQUE(cq_fetch_date,cq_id)
);
CREATE INDEX ON aws_athena_work_group_named_queries(cq_fetch_date,work_group_cq_id);
SELECT setup_tsdb_child('aws_athena_work_group_named_queries','work_group_cq_id','aws_athena_work_groups','cq_id');
