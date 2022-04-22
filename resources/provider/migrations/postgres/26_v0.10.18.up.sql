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


-- Resource: shield.attacks
CREATE TABLE IF NOT EXISTS "aws_shield_attacks" (
    "cq_id" uuid NOT NULL,
    "cq_meta" jsonb,
    "account_id" text,
    "attack_counters" jsonb,
    "id" text,
    "end_time" timestamp without time zone,
    "mitigations" text[],
    "resource_arn" text,
    "start_time" timestamp without time zone,
    CONSTRAINT aws_shield_attacks_pk PRIMARY KEY(id),
    UNIQUE(cq_id)
);
CREATE TABLE IF NOT EXISTS "aws_shield_attack_properties" (
    "cq_id" uuid NOT NULL,
    "cq_meta" jsonb,
    "attack_cq_id" uuid,
    "attack_layer" text,
    "attack_property_identifier" text,
    "top_contributors" jsonb,
    "total" bigint,
    "unit" text,
    CONSTRAINT aws_shield_attack_properties_pk PRIMARY KEY(cq_id),
    UNIQUE(cq_id),
    FOREIGN KEY (attack_cq_id) REFERENCES aws_shield_attacks(cq_id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS "aws_shield_attack_sub_resources" (
    "cq_id" uuid NOT NULL,
    "cq_meta" jsonb,
    "attack_cq_id" uuid,
    "attack_vectors" jsonb,
    "counters" jsonb,
    "id" text,
    "type" text,
    CONSTRAINT aws_shield_attack_sub_resources_pk PRIMARY KEY(cq_id),
    UNIQUE(cq_id),
    FOREIGN KEY (attack_cq_id) REFERENCES aws_shield_attacks(cq_id) ON DELETE CASCADE
);

-- Resource: shield.protections
CREATE TABLE IF NOT EXISTS "aws_shield_protections" (
    "cq_id" uuid NOT NULL,
    "cq_meta" jsonb,
    "account_id" text,
    "region" text,
    "tags" jsonb,
    "application_layer_automatic_response_configuration_status" text,
    "health_check_ids" text[],
    "id" text,
    "name" text,
    "arn" text,
    "resource_arn" text,
    CONSTRAINT aws_shield_protections_pk PRIMARY KEY(arn),
    UNIQUE(cq_id)
);

-- Resource: shield.protections_groups
CREATE TABLE IF NOT EXISTS "aws_shield_protection_groups" (
    "cq_id" uuid NOT NULL,
    "cq_meta" jsonb,
    "account_id" text,
    "aggregation" text,
    "members" text[],
    "pattern" text,
    "id" text,
    "arn" text,
    "resource_type" text,
    CONSTRAINT aws_shield_protection_groups_pk PRIMARY KEY(arn),
    UNIQUE(cq_id)
);

-- Resource: shield.subscriptions
CREATE TABLE IF NOT EXISTS "aws_shield_subscriptions" (
    "cq_id" uuid NOT NULL,
    "cq_meta" jsonb,
    "account_id" text,
    "protection_group_limits_max_protection_groups" bigint,
    "protection_group_limits_arbitrary_pattern_limits_max_members" bigint,
    "protected_resource_type_limits" jsonb,
    "auto_renew" text,
    "end_time" timestamp without time zone,
    "limits" jsonb,
    "proactive_engagement_status" text,
    "start_time" timestamp without time zone,
    "arn" text,
    "time_commitment_in_seconds" bigint,
    CONSTRAINT aws_shield_subscriptions_pk PRIMARY KEY(arn),
    UNIQUE(cq_id)
);

