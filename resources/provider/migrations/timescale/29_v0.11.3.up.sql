-- Resource: ssm.associations
CREATE TABLE IF NOT EXISTS "aws_ssm_associations" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"cq_fetch_date" timestamp without time zone NOT NULL,
	"arn" text,
	"account_id" text,
	"region" text,
	"apply_only_at_cron_interval" boolean,
	"id" text,
	"association_name" text,
	"association_version" text,
	"automation_target_parameter_name" text,
	"calendar_names" text[],
	"compliance_severity" text,
	"date" timestamp without time zone,
	"document_version" text,
	"instance_id" text,
	"last_execution_date" timestamp without time zone,
	"last_successful_execution_date" timestamp without time zone,
	"last_update_association_date" timestamp without time zone,
	"max_concurrency" text,
	"max_errors" text,
	"name" text,
	"output_location_s3_bucket_name" text,
	"output_location_s3_key_prefix" text,
	"output_location_s3_region" text,
	"overview_association_status_aggregated_count" jsonb,
	"overview_detailed_status" text,
	"overview_status" text,
	"parameters" jsonb,
	"schedule_expression" text,
	"status_date" timestamp without time zone,
	"status_message" text,
	"status_name" text,
	"status_additional_info" text,
	"sync_compliance" text,
	"targets" jsonb,
	CONSTRAINT aws_ssm_associations_pk PRIMARY KEY(cq_fetch_date,arn),
	UNIQUE(cq_fetch_date,cq_id)
);
SELECT setup_tsdb_parent('aws_ssm_associations');
CREATE TABLE IF NOT EXISTS "aws_ssm_association_target_locations" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"cq_fetch_date" timestamp without time zone NOT NULL,
	"association_cq_id" uuid,
	"accounts" text[],
	"execution_role_name" text,
	"regions" text[],
	"target_location_max_concurrency" text,
	"target_location_max_errors" text,
	CONSTRAINT aws_ssm_association_target_locations_pk PRIMARY KEY(cq_fetch_date,cq_id),
	UNIQUE(cq_fetch_date,cq_id)
);
CREATE INDEX ON aws_ssm_association_target_locations (cq_fetch_date, association_cq_id);
SELECT setup_tsdb_child('aws_ssm_association_target_locations', 'association_cq_id', 'aws_ssm_associations', 'cq_id');

-- Resource: ssm.instances
CREATE TABLE IF NOT EXISTS "aws_ssm_instance_patches" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"cq_fetch_date" timestamp without time zone NOT NULL,
	"instance_cq_id" uuid,
	"classification" text,
	"installed_time" timestamp without time zone,
	"kb_id" text,
	"severity" text,
	"state" text,
	"title" text,
	"cve_ids" text,
	CONSTRAINT aws_ssm_instance_patches_pk PRIMARY KEY(cq_fetch_date,instance_cq_id,kb_id),
	UNIQUE(cq_fetch_date,cq_id)
);
CREATE INDEX ON aws_ssm_instance_patches (cq_fetch_date, instance_cq_id);
SELECT setup_tsdb_child('aws_ssm_instance_patches', 'instance_cq_id', 'aws_ssm_instances', 'cq_id');

-- Resource: ssm.patch_baselines
CREATE TABLE IF NOT EXISTS "aws_ssm_patch_baselines" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"cq_fetch_date" timestamp without time zone NOT NULL,
	"arn" text,
	"account_id" text,
	"region" text,
	"approved_patches" text[],
	"approved_patches_compliance_level" text,
	"approved_patches_enable_non_security" boolean,
	"baseline_id" text,
	"created_date" timestamp without time zone,
	"description" text,
	"global_filters" jsonb,
	"modified_date" timestamp without time zone,
	"name" text,
	"operating_system" text,
	"patch_groups" text[],
	"rejected_patches" text[],
	"rejected_patches_action" text,
	"tags" jsonb,
	CONSTRAINT aws_ssm_patch_baselines_pk PRIMARY KEY(cq_fetch_date,arn),
	UNIQUE(cq_fetch_date,cq_id)
);
SELECT setup_tsdb_parent('aws_ssm_patch_baselines');
CREATE TABLE IF NOT EXISTS "aws_ssm_patch_baseline_approval_rules" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"cq_fetch_date" timestamp without time zone NOT NULL,
	"patch_baseline_cq_id" uuid,
	"approve_after_days" integer,
	"approve_until_date" text,
	"compliance_level" text,
	"enable_non_security" boolean,
	"patch_filter_group" jsonb,
	CONSTRAINT aws_ssm_patch_baseline_approval_rules_pk PRIMARY KEY(cq_fetch_date,cq_id),
	UNIQUE(cq_fetch_date,cq_id)
);
CREATE INDEX ON aws_ssm_patch_baseline_approval_rules (cq_fetch_date, patch_baseline_cq_id);
SELECT setup_tsdb_child('aws_ssm_patch_baseline_approval_rules', 'patch_baseline_cq_id', 'aws_ssm_patch_baselines', 'cq_id');
CREATE TABLE IF NOT EXISTS "aws_ssm_patch_baseline_sources" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"cq_fetch_date" timestamp without time zone NOT NULL,
	"patch_baseline_cq_id" uuid,
	"configuration" text,
	"name" text,
	"products" text[],
	CONSTRAINT aws_ssm_patch_baseline_sources_pk PRIMARY KEY(cq_fetch_date,cq_id),
	UNIQUE(cq_fetch_date,cq_id)
);
CREATE INDEX ON aws_ssm_patch_baseline_sources (cq_fetch_date, patch_baseline_cq_id);
SELECT setup_tsdb_child('aws_ssm_patch_baseline_sources', 'patch_baseline_cq_id', 'aws_ssm_patch_baselines', 'cq_id');
