--aws_apigatewayv2_vpc_links
ALTER TABLE "aws_apigatewayv2_vpc_links"
    ADD COLUMN "vpc_link_id" TEXT;

UPDATE "aws_apigatewayv2_vpc_links"
SET "vpc_link_id" = "id";

--aws_cloudtrail_trails
ALTER TABLE "aws_cloudtrail_trails"
DROP
COLUMN "tags";
ALTER TABLE "aws_cloudtrail_trails"
    ADD COLUMN "home_region" TEXT;
UPDATE "aws_cloudtrail_trails"
SET "home_region" = "region";

--aws_elasticbeanstalk_environments
ALTER TABLE "aws_elasticbeanstalk_environments"
DROP
COLUMN "tags";
ALTER TABLE "aws_elasticbeanstalk_environments" RENAME COLUMN "name" TO "environment_name";

--aws_elasticsearch_domains
ALTER TABLE "aws_elasticsearch_domains"
DROP
COLUMN "tags";

--aws_elbv2_load_balancers
ALTER TABLE "aws_elbv2_load_balancers"
DROP
COLUMN "tags";

--aws_elbv2_target_groups
ALTER TABLE "aws_elbv2_target_groups"
DROP
COLUMN "tags";

--aws_kms_keys
ALTER TABLE "aws_kms_keys"
DROP
COLUMN "tags";
ALTER TABLE "aws_kms_keys" RENAME COLUMN "id" TO "key_id";
