-- Autogenerated by migration tool on 2022-05-04 19:52:06
-- CHANGEME: Verify or edit this file before proceeding

-- Resource: wafv2.managed_rule_groups
ALTER TABLE IF EXISTS aws_wafv2_managed_rule_groups DROP CONSTRAINT aws_wafv2_managed_rule_groups_pk;
ALTER TABLE IF EXISTS aws_wafv2_managed_rule_groups ADD CONSTRAINT aws_wafv2_managed_rule_groups_pk PRIMARY KEY (account_id,region,vendor_name,name);
ALTER TABLE IF EXISTS "aws_wafv2_managed_rule_groups" DROP COLUMN IF EXISTS "scope";
