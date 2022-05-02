-- Resource: backup.global_settings
CREATE TABLE IF NOT EXISTS "aws_backup_global_settings"
(
    "cq_id"            uuid NOT NULL,
    "cq_meta"          jsonb,
    "account_id"       text,
    "global_settings"  jsonb,
    "last_update_time" timestamp without time zone,
    CONSTRAINT aws_backup_global_settings_pk PRIMARY KEY (account_id),
    UNIQUE (cq_id)
);

-- Resource: backup.region_settings
CREATE TABLE IF NOT EXISTS "aws_backup_region_settings"
(
    "cq_id"                               uuid NOT NULL,
    "cq_meta"                             jsonb,
    "account_id"                          text,
    "region"                              text,
    "resource_type_management_preference" jsonb,
    "resource_type_opt_in_preference"     jsonb,
    CONSTRAINT aws_backup_region_settings_pk PRIMARY KEY (account_id, region),
    UNIQUE (cq_id)
);

-- Resource: athena.data_catalogs
CREATE TABLE IF NOT EXISTS "aws_athena_data_catalogs"
(
    "cq_id"       uuid NOT NULL,
    "cq_meta"     jsonb,
    "account_id"  text,
    "region"      text,
    "tags"        jsonb,
    "name"        text,
    "type"        text,
    "description" text,
    "parameters"  jsonb,
    CONSTRAINT aws_athena_data_catalogs_pk PRIMARY KEY (account_id, region, name),
    UNIQUE (cq_id)
);
CREATE TABLE IF NOT EXISTS "aws_athena_data_catalog_databases"
(
    "cq_id"              uuid NOT NULL,
    "cq_meta"            jsonb,
    "data_catalog_cq_id" uuid,
    "name"               text,
    "description"        text,
    "parameters"         jsonb,
    CONSTRAINT aws_athena_data_catalog_databases_pk PRIMARY KEY (cq_id),
    UNIQUE (cq_id),
    FOREIGN KEY (data_catalog_cq_id) REFERENCES aws_athena_data_catalogs (cq_id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS "aws_athena_data_catalog_database_tables"
(
    "cq_id"                       uuid NOT NULL,
    "cq_meta"                     jsonb,
    "data_catalog_database_cq_id" uuid,
    "name"                        text,
    "create_time"                 timestamp without time zone,
    "last_access_time"            timestamp without time zone,
    "parameters"                  jsonb,
    "table_type"                  text,
    CONSTRAINT aws_athena_data_catalog_database_tables_pk PRIMARY KEY (cq_id),
    UNIQUE (cq_id),
    FOREIGN KEY (data_catalog_database_cq_id) REFERENCES aws_athena_data_catalog_databases (cq_id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS "aws_athena_data_catalog_database_table_columns"
(
    "cq_id"                             uuid NOT NULL,
    "cq_meta"                           jsonb,
    "data_catalog_database_table_cq_id" uuid,
    "name"                              text,
    "comment"                           text,
    "type"                              text,
    CONSTRAINT aws_athena_data_catalog_database_table_columns_pk PRIMARY KEY (cq_id),
    UNIQUE (cq_id),
    FOREIGN KEY (data_catalog_database_table_cq_id) REFERENCES aws_athena_data_catalog_database_tables (cq_id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS "aws_athena_data_catalog_database_table_partition_keys"
(
    "cq_id"                             uuid NOT NULL,
    "cq_meta"                           jsonb,
    "data_catalog_database_table_cq_id" uuid,
    "name"                              text,
    "comment"                           text,
    "type"                              text,
    CONSTRAINT aws_athena_data_catalog_database_table_partition_keys_pk PRIMARY KEY (cq_id),
    UNIQUE (cq_id),
    FOREIGN KEY (data_catalog_database_table_cq_id) REFERENCES aws_athena_data_catalog_database_tables (cq_id) ON DELETE CASCADE
);

-- Resource: athena.work_groups
CREATE TABLE IF NOT EXISTS "aws_athena_work_groups"
(
    "cq_id"                                      uuid NOT NULL,
    "cq_meta"                                    jsonb,
    "account_id"                                 text,
    "region"                                     text,
    "tags"                                       jsonb,
    "name"                                       text,
    "bytes_scanned_cutoff_per_query"             bigint,
    "enforce_work_group_configuration"           boolean,
    "effective_engine_version"                   text,
    "selected_engine_version"                    text,
    "publish_cloud_watch_metrics_enabled"        boolean,
    "requester_pays_enabled"                     boolean,
    "acl_configuration_s3_acl_option"            text,
    "encryption_configuration_encryption_option" text,
    "encryption_configuration_kms_key"           text,
    "expected_bucket_owner"                      text,
    "output_location"                            text,
    "creation_time"                              timestamp without time zone,
    "description"                                text,
    "state"                                      text,
    CONSTRAINT aws_athena_work_groups_pk PRIMARY KEY (account_id, region, name),
    UNIQUE (cq_id)
);
CREATE TABLE IF NOT EXISTS "aws_athena_work_group_prepared_statements"
(
    "cq_id"              uuid NOT NULL,
    "cq_meta"            jsonb,
    "work_group_cq_id"   uuid,
    "description"        text,
    "last_modified_time" timestamp without time zone,
    "query_statement"    text,
    "statement_name"     text,
    "work_group_name"    text,
    CONSTRAINT aws_athena_work_group_prepared_statements_pk PRIMARY KEY (cq_id),
    UNIQUE (cq_id),
    FOREIGN KEY (work_group_cq_id) REFERENCES aws_athena_work_groups (cq_id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS "aws_athena_work_group_query_executions"
(
    "cq_id"                                      uuid NOT NULL,
    "cq_meta"                                    jsonb,
    "work_group_cq_id"                           uuid,
    "effective_engine_version"                   text,
    "selected_engine_version"                    text,
    "query"                                      text,
    "catalog"                                    text,
    "database"                                   text,
    "id"                                         text,
    "acl_configuration_s3_acl_option"            text,
    "encryption_configuration_encryption_option" text,
    "encryption_configuration_kms_key"           text,
    "expected_bucket_owner"                      text,
    "output_location"                            text,
    "statement_type"                             text,
    "data_manifest_location"                     text,
    "data_scanned_in_bytes"                      bigint,
    "engine_execution_time_in_millis"            bigint,
    "query_planning_time_in_millis"              bigint,
    "query_queue_time_in_millis"                 bigint,
    "service_processing_time_in_millis"          bigint,
    "total_execution_time_in_millis"             bigint,
    "athena_error_error_category"                integer,
    "athena_error_error_message"                 text,
    "athena_error_error_type"                    integer,
    "athena_error_retryable"                     boolean,
    "completion_date_time"                       timestamp without time zone,
    "state"                                      text,
    "state_change_reason"                        text,
    "submission_date_time"                       timestamp without time zone,
    "work_group"                                 text,
    CONSTRAINT aws_athena_work_group_query_executions_pk PRIMARY KEY (cq_id),
    UNIQUE (cq_id),
    FOREIGN KEY (work_group_cq_id) REFERENCES aws_athena_work_groups (cq_id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS "aws_athena_work_group_named_queries"
(
    "cq_id"            uuid NOT NULL,
    "cq_meta"          jsonb,
    "work_group_cq_id" uuid,
    "database"         text,
    "name"             text,
    "query_string"     text,
    "description"      text,
    "named_query_id"   text,
    "work_group"       text,
    CONSTRAINT aws_athena_work_group_named_queries_pk PRIMARY KEY (cq_id),
    UNIQUE (cq_id),
    FOREIGN KEY (work_group_cq_id) REFERENCES aws_athena_work_groups (cq_id) ON DELETE CASCADE
);