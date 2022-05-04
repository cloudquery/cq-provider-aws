-- Autogenerated by migration tool on 2022-05-04 20:31:20
-- CHANGEME: Verify or edit this file before proceeding

-- Resource: wafv2.managed_rule_groups
TRUNCATE TABLE aws_wafv2_managed_rule_groups;
ALTER TABLE IF EXISTS aws_wafv2_managed_rule_groups DROP CONSTRAINT aws_wafv2_managed_rule_groups_pk;
ALTER TABLE IF EXISTS "aws_wafv2_managed_rule_groups" ADD COLUMN IF NOT EXISTS "scope" text;
ALTER TABLE IF EXISTS aws_wafv2_managed_rule_groups ADD CONSTRAINT aws_wafv2_managed_rule_groups_pk PRIMARY KEY (account_id,region,scope,vendor_name,name);

-- Resource: wafv2.rule_groups
TRUNCATE TABLE aws_wafv2_rule_groups;
ALTER TABLE IF EXISTS "aws_wafv2_rule_groups" ADD COLUMN IF NOT EXISTS "scope" text;

-- Resource: wafv2.web_acls
DELETE FROM aws_wafv2_web_acls;
ALTER TABLE IF EXISTS "aws_wafv2_web_acls" ADD COLUMN IF NOT EXISTS "scope" text;
