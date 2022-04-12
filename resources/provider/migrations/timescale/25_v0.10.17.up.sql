-- Resource: wafv2.ipsets
CREATE TABLE IF NOT EXISTS "aws_wafv2_ipsets" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"cq_fetch_date" timestamp without time zone NOT NULL,
	"account_id" text,
	"region" text,
	"scope" text,
	"arn" text,
	"addresses" cidr[],
	"ip_address_version" text,
	"id" text,
	"name" text,
	"description" text,
	"tags" jsonb,
	CONSTRAINT aws_wafv2_ipsets_pk PRIMARY KEY(cq_fetch_date,arn),
	UNIQUE(cq_fetch_date,cq_id)
);
SELECT setup_tsdb_parent('aws_wafv2_ipsets');

-- Resource: wafv2.regex_pattern_sets
CREATE TABLE IF NOT EXISTS "aws_wafv2_regex_pattern_sets" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"cq_fetch_date" timestamp without time zone NOT NULL,
	"account_id" text,
	"region" text,
	"scope" text,
	"arn" text,
	"description" text,
	"id" text,
	"name" text,
	"regular_expression_list" text[],
	"tags" jsonb,
	CONSTRAINT aws_wafv2_regex_pattern_sets_pk PRIMARY KEY(cq_fetch_date,arn),
	UNIQUE(cq_fetch_date,cq_id)
);
SELECT setup_tsdb_parent('aws_wafv2_regex_pattern_sets');

-- Resource: ec2.egress_only_internet_gateways
CREATE TABLE IF NOT EXISTS "aws_ec2_egress_only_internet_gateways" (
    "cq_id" uuid NOT NULL,
    "cq_meta" jsonb,
    "cq_fetch_date" timestamp without time zone NOT NULL,
    "account_id" text,
    "region" text,
    "arn" text,
    "attachments" jsonb,
    "id" text,
    "tags" jsonb,
    CONSTRAINT aws_ec2_egress_only_internet_gateways_pk PRIMARY KEY(cq_fetch_date,arn),
    UNIQUE(cq_fetch_date,cq_id)
);
SELECT setup_tsdb_parent('aws_ec2_egress_only_internet_gateways');
