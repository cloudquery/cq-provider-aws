ALTER TABLE IF EXISTS "aws_sns_topics" ADD COLUMN IF NOT EXISTS "tags" jsonb;

CREATE TABLE IF NOT EXISTS "aws_elasticbeanstalk_application_versions" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"cq_fetch_date" timestamp without time zone NOT NULL,
	"account_id" text,
	"region" text,
	"application_name" text,
	"arn" text,
	"build_arn" text,
	"date_created" timestamp without time zone,
	"date_updated" timestamp without time zone,
	"description" text,
	"source_location" text,
	"source_repository" text,
	"source_type" text,
	"source_bundle_s3_bucket" text,
	"source_bundle_s3_key" text,
	"status" text,
	"version_label" text,
	CONSTRAINT aws_elasticbeanstalk_application_versions_pk PRIMARY KEY(cq_fetch_date,arn),
	UNIQUE(cq_fetch_date,cq_id)
);
SELECT setup_tsdb_parent('aws_elasticbeanstalk_application_versions');

-- Resource: waf.web_acls
ALTER TABLE IF EXISTS "aws_waf_web_acls" DROP COLUMN IF EXISTS "logging_configuration";
CREATE TABLE IF NOT EXISTS "aws_waf_web_acl_logging_configuration" (
    "cq_id" uuid NOT NULL,
    "cq_meta" jsonb,
    "cq_fetch_date" timestamp without time zone NOT NULL,
    "web_acl_cq_id" uuid,
    "log_destination_configs" text[],
    "resource_arn" text,
    "redacted_fields" jsonb,
    CONSTRAINT aws_waf_web_acl_logging_configuration_pk PRIMARY KEY(cq_fetch_date,web_acl_cq_id,resource_arn),
    UNIQUE(cq_fetch_date,cq_id)
);
CREATE INDEX ON aws_waf_web_acl_logging_configuration (cq_fetch_date, web_acl_cq_id);
SELECT setup_tsdb_child('aws_waf_web_acl_logging_configuration', 'web_acl_cq_id', 'aws_waf_web_acls', 'cq_id');

-- Resource: wafv2.web_acls
ALTER TABLE IF EXISTS "aws_wafv2_web_acls" DROP COLUMN IF EXISTS "logging_configuration";
CREATE TABLE IF NOT EXISTS "aws_wafv2_web_acl_logging_configuration" (
    "cq_id" uuid NOT NULL,
    "cq_meta" jsonb,
    "cq_fetch_date" timestamp without time zone NOT NULL,
    "web_acl_cq_id" uuid,
    "log_destination_configs" text[],
    "resource_arn" text,
    "logging_filter" jsonb,
    "managed_by_firewall_manager" boolean,
    "redacted_fields" jsonb,
    CONSTRAINT aws_wafv2_web_acl_logging_configuration_pk PRIMARY KEY(cq_fetch_date,web_acl_cq_id,resource_arn),
    UNIQUE(cq_fetch_date,cq_id)
);
CREATE INDEX ON aws_wafv2_web_acl_logging_configuration (cq_fetch_date, web_acl_cq_id);
SELECT setup_tsdb_child('aws_wafv2_web_acl_logging_configuration', 'web_acl_cq_id', 'aws_wafv2_web_acls', 'cq_id');

