-- Resource: ecs.task_definitions
ALTER TABLE IF EXISTS "aws_ecs_task_definitions" ADD COLUMN IF NOT EXISTS "ephemeral_storage_size" integer;
ALTER TABLE IF EXISTS "aws_ecs_task_definitions" ADD COLUMN IF NOT EXISTS "runtime_platform_cpu_architecture" text;
ALTER TABLE IF EXISTS "aws_ecs_task_definitions" ADD COLUMN IF NOT EXISTS "runtime_platform_os_family" text;


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

