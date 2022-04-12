-- Autogenerated by migration tool on 2022-04-12 11:25:59
-- CHANGEME: Verify or edit this file before proceeding

-- Resource: ec2.egress_only_internet_gateways
CREATE TABLE IF NOT EXISTS "aws_ec2_egress_only_internet_gateways" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"account_id" text,
	"region" text,
	"arn" text,
	"attachments" jsonb,
	"id" text,
	"tags" jsonb,
	CONSTRAINT aws_ec2_egress_only_internet_gateways_pk PRIMARY KEY(arn),
	UNIQUE(cq_id)
);
