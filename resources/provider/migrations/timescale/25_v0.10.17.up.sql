-- Resource: backup.plans
CREATE TABLE IF NOT EXISTS "aws_backup_plans" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"cq_fetch_date" timestamp without time zone NOT NULL,
	"account_id" text,
	"region" text,
	"arn" text,
	"id" text,
	"name" text,
	"creation_date" timestamp without time zone,
	"creator_request_id" text,
	"last_execution_date" timestamp without time zone,
	"version_id" text,
	"advanced_backup_settings" jsonb,
	"tags" jsonb,
	CONSTRAINT aws_backup_plans_pk PRIMARY KEY(cq_fetch_date,arn),
	UNIQUE(cq_fetch_date,cq_id)
);
SELECT setup_tsdb_parent('aws_backup_plans');
CREATE TABLE IF NOT EXISTS "aws_backup_plan_rules" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"cq_fetch_date" timestamp without time zone NOT NULL,
	"plan_cq_id" uuid,
	"name" text,
	"target_backup_vault_name" text,
	"completion_window_minutes" bigint,
	"copy_actions" jsonb,
	"enable_continuous_backup" boolean,
	"delete_after_days" bigint,
	"move_to_cold_storage_after_days" bigint,
	"recovery_point_tags" jsonb,
	"id" text,
	"schedule_expression" text,
	"start_window_minutes" bigint,
	CONSTRAINT aws_backup_plan_rules_pk PRIMARY KEY(cq_fetch_date,cq_id),
	UNIQUE(cq_fetch_date,cq_id)
);
CREATE INDEX ON aws_backup_plan_rules (cq_fetch_date, plan_cq_id);
SELECT setup_tsdb_child('aws_backup_plan_rules', 'plan_cq_id', 'aws_backup_plans', 'cq_id');
CREATE TABLE IF NOT EXISTS "aws_backup_plan_selections" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"cq_fetch_date" timestamp without time zone NOT NULL,
	"plan_cq_id" uuid,
	"creation_date" timestamp without time zone,
	"creator_request_id" text,
	"iam_role_arn" text,
	"selection_id" text,
	"selection_name" text,
	"conditions" jsonb,
	"list_of_tags" jsonb,
	"not_resources" text[],
	"resources" text[],
	CONSTRAINT aws_backup_plan_selections_pk PRIMARY KEY(cq_fetch_date,cq_id),
	UNIQUE(cq_fetch_date,cq_id)
);
CREATE INDEX ON aws_backup_plan_selections (cq_fetch_date, plan_cq_id);
SELECT setup_tsdb_child('aws_backup_plan_selections', 'plan_cq_id', 'aws_backup_plans', 'cq_id');

-- Resource: backup.vaults
CREATE TABLE IF NOT EXISTS "aws_backup_vaults" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"cq_fetch_date" timestamp without time zone NOT NULL,
	"account_id" text,
	"region" text,
	"arn" text,
	"name" text,
	"creation_date" timestamp without time zone,
	"creator_request_id" text,
	"encryption_key_arn" text,
	"lock_date" timestamp without time zone,
	"locked" boolean,
	"max_retention_days" bigint,
	"min_retention_days" bigint,
	"number_of_recovery_points" bigint,
	"access_policy" jsonb,
	"notification_events" text[],
	"notification_sns_topic_arn" text,
	"tags" jsonb,
	CONSTRAINT aws_backup_vaults_pk PRIMARY KEY(cq_fetch_date,arn),
	UNIQUE(cq_fetch_date,cq_id)
);
SELECT setup_tsdb_parent('aws_backup_vaults');
CREATE TABLE IF NOT EXISTS "aws_backup_vault_recovery_points" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"cq_fetch_date" timestamp without time zone NOT NULL,
	"vault_cq_id" uuid,
	"backup_size" bigint,
	"calculated_delete_at" timestamp without time zone,
	"calculated_move_to_cold_storage_at" timestamp without time zone,
	"completion_date" timestamp without time zone,
	"created_by" jsonb,
	"creation_date" timestamp without time zone,
	"encryption_key_arn" text,
	"iam_role_arn" text,
	"is_encrypted" boolean,
	"last_restore_time" timestamp without time zone,
	"delete_after" bigint,
	"move_to_cold_storage_after" bigint,
	"arn" text,
	"resource_arn" text,
	"resource_type" text,
	"source_backup_vault_arn" text,
	"status" text,
	"status_message" text,
	"tags" jsonb,
	CONSTRAINT aws_backup_vault_recovery_points_pk PRIMARY KEY(cq_fetch_date,cq_id),
	UNIQUE(cq_fetch_date,cq_id)
);
CREATE INDEX ON aws_backup_vault_recovery_points (cq_fetch_date, vault_cq_id);
SELECT setup_tsdb_child('aws_backup_vault_recovery_points', 'vault_cq_id', 'aws_backup_vaults', 'cq_id');

-- Resource: codepipeline.webhooks
CREATE TABLE IF NOT EXISTS "aws_codepipeline_webhooks" (
    "cq_id" uuid NOT NULL,
    "cq_meta" jsonb,
    "cq_fetch_date" timestamp without time zone NOT NULL,
    "account_id" text,
    "region" text,
    "authentication" text,
    "authentication_allowed_ip_range" text,
    "authentication_secret_token" text,
    "name" text,
    "target_action" text,
    "target_pipeline" text,
    "url" text,
    "arn" text,
    "error_code" text,
    "error_message" text,
    "last_triggered" timestamp without time zone,
    "tags" jsonb,
    CONSTRAINT aws_codepipeline_webhooks_pk PRIMARY KEY(cq_fetch_date,arn),
    UNIQUE(cq_fetch_date,cq_id)
);
SELECT setup_tsdb_parent('aws_codepipeline_webhooks');
CREATE TABLE IF NOT EXISTS "aws_codepipeline_webhook_filters" (
    "cq_id" uuid NOT NULL,
    "cq_meta" jsonb,
    "cq_fetch_date" timestamp without time zone NOT NULL,
    "webhook_cq_id" uuid,
    "json_path" text,
    "match_equals" text,
    CONSTRAINT aws_codepipeline_webhook_filters_pk PRIMARY KEY(cq_fetch_date,cq_id),
    UNIQUE(cq_fetch_date,cq_id)
);
CREATE INDEX ON aws_codepipeline_webhook_filters (cq_fetch_date, webhook_cq_id);
SELECT setup_tsdb_child('aws_codepipeline_webhook_filters', 'webhook_cq_id', 'aws_codepipeline_webhooks', 'cq_id');