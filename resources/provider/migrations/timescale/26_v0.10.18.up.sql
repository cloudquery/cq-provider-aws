-- Resource: ecs.task_definitions
ALTER TABLE IF EXISTS "aws_ecs_task_definitions" ADD COLUMN IF NOT EXISTS "ephemeral_storage_size" integer;
ALTER TABLE IF EXISTS "aws_ecs_task_definitions" ADD COLUMN IF NOT EXISTS "runtime_platform_cpu_architecture" text;
ALTER TABLE IF EXISTS "aws_ecs_task_definitions" ADD COLUMN IF NOT EXISTS "runtime_platform_os_family" text;

-- Resource: shield.protections
CREATE TABLE IF NOT EXISTS "aws_shield_protections" (
    "cq_id" uuid NOT NULL,
    "cq_meta" jsonb,
    "cq_fetch_date" timestamp without time zone NOT NULL,
    "account_id" text,
    "region" text,
    "tags" jsonb,
    "application_layer_automatic_response_configuration_status" text,
    "health_check_ids" text[],
    "id" text,
    "name" text,
    "arn" text,
    "resource_arn" text,
    CONSTRAINT aws_shield_protections_pk PRIMARY KEY(cq_fetch_date,arn),
    UNIQUE(cq_fetch_date,cq_id)
);
SELECT setup_tsdb_parent('aws_shield_protections');
