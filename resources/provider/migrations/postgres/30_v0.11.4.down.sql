-- Autogenerated by migration tool on 2022-05-09 07:43:07
-- CHANGEME: Verify or edit this file before proceeding

-- Resource: waf.web_acls
DROP TABLE IF EXISTS aws_waf_web_acl_logging_configuration;
ALTER TABLE IF EXISTS "aws_waf_web_acls" ADD COLUMN IF NOT EXISTS "logging_configuration" text[];

-- Resource: wafv2.web_acls
DROP TABLE IF EXISTS aws_wafv2_web_acl_logging_configuration;
ALTER TABLE IF EXISTS "aws_wafv2_web_acls" ADD COLUMN IF NOT EXISTS "logging_configuration" text[];
