--aws_apigatewayv2_vpc_links
ALTER TABLE "aws_apigatewayv2_vpc_links"
DROP
COLUMN "vpc_link_id";

--aws_cloudtrail_trails
ALTER TABLE "aws_cloudtrail_trails"
    ADD COLUMN "tags" json;
ALTER TABLE "aws_cloudtrail_trails"
DROP
COLUMN "home_region";

--aws_elasticbeanstalk_environments
ALTER TABLE "aws_elasticbeanstalk_environments"
    ADD COLUMN "tags" json;
ALTER TABLE "aws_elasticbeanstalk_environments" RENAME COLUMN "environment_name" TO "name";

--aws_elasticsearch_domains
ALTER TABLE "aws_elasticsearch_domains"
    ADD COLUMN "tags" json;

--aws_elbv2_load_balancers
ALTER TABLE "aws_elbv2_load_balancers"
    ADD COLUMN "tags" json;

--aws_elbv2_target_groups
ALTER TABLE "aws_elbv2_target_groups"
    ADD COLUMN "tags" json;

--aws_kms_keys
ALTER TABLE "aws_kms_keys"
    ADD COLUMN "tags" json;
ALTER TABLE "aws_kms_keys" RENAME COLUMN "key_id" TO "id";



