-- Resource: backup.global_settings
CREATE TABLE IF NOT EXISTS "aws_backup_global_settings" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"cq_fetch_date" timestamp without time zone NOT NULL,
	"account_id" text,
	"global_settings" jsonb,
	"last_update_time" timestamp without time zone,
	CONSTRAINT aws_backup_global_settings_pk PRIMARY KEY(cq_fetch_date,account_id),
	UNIQUE(cq_fetch_date,cq_id)
);
SELECT setup_tsdb_parent('aws_backup_global_settings');

-- Resource: backup.region_settings
CREATE TABLE IF NOT EXISTS "aws_backup_region_settings" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"cq_fetch_date" timestamp without time zone NOT NULL,
	"account_id" text,
	"region" text,
	"application_name" text,
	"application_version_arn" text,
	"build_arn" text,
	"date_created" timestamp without time zone,
	"date_updated" timestamp without time zone,
	"description" text,
	"source_build_information_source_location" text,
	"source_build_information_source_repository" text,
	"source_build_information_source_type" text,
	"source_bundle_s3_bucket" text,
	"source_bundle_s3_key" text,
	"status" text,
	"version_label" text,
	CONSTRAINT aws_elasticbeanstalk_application_versions_pk PRIMARY KEY(cq_fetch_date,application_version_arn),
	UNIQUE(cq_fetch_date,cq_id)
);
SELECT setup_tsdb_parent('aws_elasticbeanstalk_application_versions');
