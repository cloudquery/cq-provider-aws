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
