-- Resource: ecs.task_definitions
ALTER TABLE IF EXISTS "aws_ecs_task_definitions" ADD COLUMN IF NOT EXISTS "ephemeral_storage_size" integer;
ALTER TABLE IF EXISTS "aws_ecs_task_definitions" ADD COLUMN IF NOT EXISTS "runtime_platform_cpu_architecture" text;
ALTER TABLE IF EXISTS "aws_ecs_task_definitions" ADD COLUMN IF NOT EXISTS "runtime_platform_os_family" text;

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

