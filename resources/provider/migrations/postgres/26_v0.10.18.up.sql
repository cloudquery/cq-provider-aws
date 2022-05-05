-- Resource: config.conformance_packs
CREATE TABLE IF NOT EXISTS "aws_config_conformance_pack_rule_compliances" (
    "cq_id" uuid NOT NULL,
    "cq_meta" jsonb,
    "conformance_pack_cq_id" uuid,
    "compliance_type" text,
    "config_rule_name" text,
    "controls" text[],
    "config_rule_invoked_time" timestamp without time zone,
    "resource_id" text,
    "resource_type" text,
    "ordering_timestamp" timestamp without time zone,
    "result_recorded_time" timestamp without time zone,
    "annotation" text,
    CONSTRAINT aws_config_conformance_pack_rule_compliances_pk PRIMARY KEY(conformance_pack_cq_id),
    UNIQUE(cq_id),
    FOREIGN KEY (conformance_pack_cq_id) REFERENCES aws_config_conformance_packs(cq_id) ON DELETE CASCADE
);

-- Resource: ecs.task_definitions
ALTER TABLE IF EXISTS "aws_ecs_task_definitions" ADD COLUMN IF NOT EXISTS "ephemeral_storage_size" integer;
ALTER TABLE IF EXISTS "aws_ecs_task_definitions" ADD COLUMN IF NOT EXISTS "runtime_platform_cpu_architecture" text;
ALTER TABLE IF EXISTS "aws_ecs_task_definitions" ADD COLUMN IF NOT EXISTS "runtime_platform_os_family" text;

-- Resource: kms.keys
ALTER TABLE IF EXISTS "aws_kms_keys" ADD COLUMN IF NOT EXISTS "aws_account_id" text;
ALTER TABLE "aws_kms_keys" RENAME COLUMN "customer_master_key_spec" TO "key_spec";
ALTER TABLE IF EXISTS "aws_kms_keys" ADD COLUMN IF NOT EXISTS "mac_algorithms" text[];
ALTER TABLE IF EXISTS "aws_kms_keys" ADD COLUMN IF NOT EXISTS "multi_region" boolean;
ALTER TABLE IF EXISTS "aws_kms_keys" ADD COLUMN IF NOT EXISTS "multi_region_key_type" text;
ALTER TABLE IF EXISTS "aws_kms_keys" ADD COLUMN IF NOT EXISTS "primary_key_arn" text;
ALTER TABLE IF EXISTS "aws_kms_keys" ADD COLUMN IF NOT EXISTS "primary_key_region" text;
ALTER TABLE IF EXISTS "aws_kms_keys" ADD COLUMN IF NOT EXISTS "replica_keys" jsonb;
ALTER TABLE IF EXISTS "aws_kms_keys" ADD COLUMN IF NOT EXISTS "pending_deletion_window_in_days" integer;

-- Resource: lambda.functions
ALTER TABLE IF EXISTS "aws_lambda_functions" ADD COLUMN IF NOT EXISTS "architectures" text[];
ALTER TABLE IF EXISTS "aws_lambda_functions" ADD COLUMN IF NOT EXISTS "ephemeral_storage_size" integer;
ALTER TABLE IF EXISTS "aws_lambda_function_aliases" ADD COLUMN IF NOT EXISTS "url_config_auth_type" text;
ALTER TABLE IF EXISTS "aws_lambda_function_aliases" ADD COLUMN IF NOT EXISTS "url_config_creation_time" timestamp;
ALTER TABLE IF EXISTS "aws_lambda_function_aliases" ADD COLUMN IF NOT EXISTS "url_config_function_arn" text;
ALTER TABLE IF EXISTS "aws_lambda_function_aliases" ADD COLUMN IF NOT EXISTS "url_config_function_url" text;
ALTER TABLE IF EXISTS "aws_lambda_function_aliases" ADD COLUMN IF NOT EXISTS "url_config_last_modified_time" timestamp;
ALTER TABLE IF EXISTS "aws_lambda_function_aliases" ADD COLUMN IF NOT EXISTS "url_config_cors" jsonb;
ALTER TABLE IF EXISTS "aws_lambda_function_versions" ADD COLUMN IF NOT EXISTS "architectures" text[];
ALTER TABLE IF EXISTS "aws_lambda_function_versions" ADD COLUMN IF NOT EXISTS "ephemeral_storage_size" integer;
ALTER TABLE IF EXISTS "aws_lambda_function_event_source_mappings" ADD COLUMN IF NOT EXISTS "criteria_filters" text[];