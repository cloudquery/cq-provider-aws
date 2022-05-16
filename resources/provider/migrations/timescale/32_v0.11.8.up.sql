
-- Resource: iam.groups
CREATE TABLE IF NOT EXISTS "aws_iam_group_accessed_details" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"cq_fetch_date" timestamp without time zone NOT NULL,
	"group_cq_id" uuid,
	"service_name" text,
	"service_namespace" text,
	"last_authenticated" timestamp without time zone,
	"last_authenticated_entity" text,
	"last_authenticated_region" text,
	"total_authenticated_entities" integer,
	CONSTRAINT aws_iam_group_accessed_details_pk PRIMARY KEY(cq_fetch_date,cq_id),
	UNIQUE(cq_fetch_date,cq_id)
);
CREATE INDEX ON aws_iam_group_accessed_details (cq_fetch_date, group_cq_id);
SELECT setup_tsdb_child('aws_iam_group_accessed_details', 'group_cq_id', 'aws_iam_groups', 'cq_id');
CREATE TABLE IF NOT EXISTS "aws_iam_group_accessed_detail_tracked_actions_last_accessed" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"cq_fetch_date" timestamp without time zone NOT NULL,
	"group_accessed_detail_cq_id" uuid,
	"action_name" text,
	"last_accessed_entity" text,
	"last_accessed_region" text,
	"last_accessed_time" timestamp without time zone,
	CONSTRAINT aws_iam_group_accessed_detail_tracked_actions_last_accessed_pk PRIMARY KEY(cq_fetch_date,cq_id),
	UNIQUE(cq_fetch_date,cq_id)
);
CREATE INDEX ON aws_iam_group_accessed_detail_tracked_actions_last_accessed (cq_fetch_date, group_accessed_detail_cq_id);
SELECT setup_tsdb_child('aws_iam_group_accessed_detail_tracked_actions_last_accessed', 'group_accessed_detail_cq_id', 'aws_iam_group_accessed_details', 'cq_id');
CREATE TABLE IF NOT EXISTS "aws_iam_group_accessed_detail_entities" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"cq_fetch_date" timestamp without time zone NOT NULL,
	"group_accessed_detail_cq_id" uuid,
	"arn" text,
	"id" text,
	"name" text,
	"type" text,
	"path" text,
	"last_authenticated" timestamp without time zone,
	CONSTRAINT aws_iam_group_accessed_detail_entities_pk PRIMARY KEY(cq_fetch_date,cq_id),
	UNIQUE(cq_fetch_date,cq_id)
);
CREATE INDEX ON aws_iam_group_accessed_detail_entities (cq_fetch_date, group_accessed_detail_cq_id);
SELECT setup_tsdb_child('aws_iam_group_accessed_detail_entities', 'group_accessed_detail_cq_id', 'aws_iam_group_accessed_details', 'cq_id');
