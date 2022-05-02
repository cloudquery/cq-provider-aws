DROP TABLE IF EXISTS aws_elasticbeanstalk_application_versions;

ALTER TABLE IF EXISTS "aws_sns_topics" DROP COLUMN IF EXISTS "tags";

-- Resource: waf.web_acls
DROP TABLE IF EXISTS aws_waf_web_acl_logging_configuration;
ALTER TABLE IF EXISTS "aws_waf_web_acls" ADD COLUMN IF NOT EXISTS "logging_configuration" TEXT[];

-- Resource: wafv2.web_acls
DROP TABLE IF EXISTS aws_wafv2_web_acl_logging_configuration;
ALTER TABLE IF EXISTS "aws_wafv2_web_acls" ADD COLUMN IF NOT EXISTS "logging_configuration" TEXT[];