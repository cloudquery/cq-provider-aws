-- Autogenerated by migration tool on 2022-04-24 08:54:24
-- CHANGEME: Verify or edit this file before proceeding

-- Resource: elasticbeanstalk.application_versions
CREATE TABLE IF NOT EXISTS "aws_elasticbeanstalk_application_versions" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"account_id" text,
	"region" text,
	"application_name" text,
	"arn" text,
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
	CONSTRAINT aws_elasticbeanstalk_application_versions_pk PRIMARY KEY(arn),
	UNIQUE(cq_id)
);
