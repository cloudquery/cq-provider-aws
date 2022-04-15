-- Autogenerated by migration tool on 2022-04-13 11:51:48

-- Resource: autoscaling.scheduled_actions
CREATE TABLE IF NOT EXISTS "aws_autoscaling_scheduled_actions" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"account_id" text,
	"region" text,
	"auto_scaling_group_name" text,
	"desired_capacity" integer,
	"end_time" timestamp without time zone,
	"max_size" integer,
	"min_size" integer,
	"recurrence" text,
	"arn" text,
	"name" text,
	"start_time" timestamp without time zone,
	"time" timestamp without time zone,
	"time_zone" text,
	CONSTRAINT aws_autoscaling_scheduled_actions_pk PRIMARY KEY(arn),
	UNIQUE(cq_id)
);


-- Resource: lambda.functions
ALTER TABLE IF EXISTS "aws_lambda_functions" ADD COLUMN IF NOT EXISTS "architectures" text[];
ALTER TABLE IF EXISTS "aws_lambda_functions" ADD COLUMN IF NOT EXISTS "ephemeral_storage_size" integer;
ALTER TABLE IF EXISTS "aws_lambda_functions" ADD COLUMN IF NOT EXISTS "url_config_auth_type" text;
ALTER TABLE IF EXISTS "aws_lambda_functions" ADD COLUMN IF NOT EXISTS "url_config_creation_time" text;
ALTER TABLE IF EXISTS "aws_lambda_functions" ADD COLUMN IF NOT EXISTS "url_config_function_arn" text;
ALTER TABLE IF EXISTS "aws_lambda_functions" ADD COLUMN IF NOT EXISTS "url_config_function_url" text;
ALTER TABLE IF EXISTS "aws_lambda_functions" ADD COLUMN IF NOT EXISTS "url_config_last_modified_time" text;
ALTER TABLE IF EXISTS "aws_lambda_functions" ADD COLUMN IF NOT EXISTS "url_config_cors" jsonb;
ALTER TABLE IF EXISTS "aws_lambda_function_versions" ADD COLUMN IF NOT EXISTS "architectures" text[];
ALTER TABLE IF EXISTS "aws_lambda_function_versions" ADD COLUMN IF NOT EXISTS "ephemeral_storage_size" integer;
ALTER TABLE IF EXISTS "aws_lambda_function_event_source_mappings" ADD COLUMN IF NOT EXISTS "criteria_filters" text[];

