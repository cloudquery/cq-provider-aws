--aws_apigatewayv2_vpc_links
ALTER TABLE "aws_apigatewayv2_vpc_links"
    ADD COLUMN "vpc_link_id" TEXT;

UPDATE "aws_apigatewayv2_vpc_links"
SET "vpc_link_id" = "id";


--aws_cloudfront_distributions
ALTER TABLE "aws_cloudfront_distributions"
DROP
COLUMN "tags";
ALTER TABLE "aws_cloudfront_distributions" RENAME COLUMN "aliases" TO "aliases_items";
-- ALTER TABLE "aws_cloudfront_distributions" RENAME COLUMN "cache_behavior_target_origin_id" TO "cache_behaviour_target_origin_id";
-- ALTER TABLE "aws_cloudfront_distributions" RENAME COLUMN "cache_behavior_viewer_protocol_policy" TO "cache_behaviour_viewer_protocol_policy";
-- ALTER TABLE "aws_cloudfront_distributions" RENAME COLUMN "cache_behavior_allowed_methods" TO "cache_behaviour_allowed_methods";
-- ALTER TABLE "aws_cloudfront_distributions" RENAME COLUMN "cache_behavior_allowed_methods_cached_methods" TO "cache_behaviour_allowed_methods_cached_methods";
-- ALTER TABLE "aws_cloudfront_distributions" RENAME COLUMN "cache_behavior_cache_policy_id" TO "cache_behaviour_cache_policy_id";
-- ALTER TABLE "aws_cloudfront_distributions" RENAME COLUMN "cache_behavior_compress" TO "cache_behaviour_compress";
-- ALTER TABLE "aws_cloudfront_distributions" RENAME COLUMN "cache_behavior_default_ttl" TO "cache_behaviour_default_ttl";
-- ALTER TABLE "aws_cloudfront_distributions" RENAME COLUMN "cache_behavior_field_level_encryption_id" TO "cache_behaviour_field_level_encryption_id";
-- ALTER TABLE "aws_cloudfront_distributions" RENAME COLUMN "cache_behavior_forwarded_values_cookies_forward" TO "cache_behaviour_forwarded_values_cookies_forward";
-- ALTER TABLE "aws_cloudfront_distributions" RENAME COLUMN "cache_behavior_forwarded_values_cookies_whitelisted_names" TO "cache_behaviour_forwarded_values_cookies_white_listed_names";
-- ALTER TABLE "aws_cloudfront_distributions" RENAME COLUMN "cache_behavior_forwarded_values_query_string" TO "cache_behaviour_forwarded_values_query_string";
-- ALTER TABLE "aws_cloudfront_distributions" RENAME COLUMN "cache_behavior_forwarded_values_headers" TO "cache_behaviour_forwarded_values_headers";
-- ALTER TABLE "aws_cloudfront_distributions" RENAME COLUMN "cache_behavior_forwarded_values_query_string_cache_keys" TO "cache_behaviour_forwarded_values_query_string_cache_keys";
-- ALTER TABLE "aws_cloudfront_distributions" RENAME COLUMN "cache_behavior_max_ttl" TO "cache_behaviour_max_ttl";
-- ALTER TABLE "aws_cloudfront_distributions" RENAME COLUMN "cache_behavior_min_ttl" TO "cache_behaviour_min_ttl";
-- ALTER TABLE "aws_cloudfront_distributions" RENAME COLUMN "cache_behavior_origin_request_policy_id" TO "cache_behaviour_origin_request_policy_id";
-- ALTER TABLE "aws_cloudfront_distributions" RENAME COLUMN "cache_behavior_realtime_log_config_arn" TO "cache_behaviour_realtime_log_config_arn";
-- ALTER TABLE "aws_cloudfront_distributions" RENAME COLUMN "cache_behavior_smooth_streaming" TO "cache_behaviour_smooth_streaming";
-- ALTER TABLE "aws_cloudfront_distributions" RENAME COLUMN "cache_behavior_trusted_key_groups_enabled" TO "cache_behaviour_trusted_key_groups_enabled";
-- ALTER TABLE "aws_cloudfront_distributions" RENAME COLUMN "cache_behavior_trusted_key_groups" TO "cache_behaviour_trusted_key_groups";
-- ALTER TABLE "aws_cloudfront_distributions" RENAME COLUMN "cache_behavior_trusted_signers_enabled" TO "cache_behaviour_trusted_signers_enabled";
-- ALTER TABLE "aws_cloudfront_distributions" RENAME COLUMN "cache_behavior_trusted_signers" TO "cache_behaviour_trusted_signers";
ALTER TABLE "aws_cloudfront_distributions" RENAME COLUMN "restrictions_geo_restriction_restrictions" TO "restrictions_geo_restriction_restriction_items";

--aws_cloudfront_distribution_cache_behaviours
ALTER TABLE "aws_cloudfront_distributions"
DROP
COLUMN "cache_policy_id";
ALTER TABLE "aws_cloudfront_distributions"
DROP
COLUMN "compress";
ALTER TABLE "aws_cloudfront_distributions"
DROP
COLUMN "default_ttl";
ALTER TABLE "aws_cloudfront_distributions"
DROP
COLUMN "field_level_encryption_id";
ALTER TABLE "aws_cloudfront_distributions"
DROP
COLUMN "forwarded_values_cookies_forward";
ALTER TABLE "aws_cloudfront_distributions"
DROP
COLUMN "forwarded_values_cookies_whitelisted_names";
ALTER TABLE "aws_cloudfront_distributions"
DROP
COLUMN "forwarded_values_query_string";
ALTER TABLE "aws_cloudfront_distributions"
DROP
COLUMN "forwarded_values_headers";
ALTER TABLE "aws_cloudfront_distributions"
DROP
COLUMN "forwarded_values_query_string_cache_keys";
ALTER TABLE "aws_cloudfront_distributions"
DROP
COLUMN "max_ttl";
ALTER TABLE "aws_cloudfront_distributions"
DROP
COLUMN "min_ttl";
ALTER TABLE "aws_cloudfront_distributions"
DROP
COLUMN "origin_request_policy_id";
ALTER TABLE "aws_cloudfront_distributions"
DROP
COLUMN "realtime_log_config_arn";
ALTER TABLE "aws_cloudfront_distributions"
DROP
COLUMN "smooth_streaming";
ALTER TABLE "aws_cloudfront_distributions"
DROP
COLUMN "trusted_key_groups_enabled";
ALTER TABLE "aws_cloudfront_distributions"
DROP
COLUMN "trusted_key_groups";
ALTER TABLE "aws_cloudfront_distributions"
DROP
COLUMN "trusted_signers_enabled";
ALTER TABLE "aws_cloudfront_distributions"
DROP
COLUMN "trusted_signers";

--aws_cloudfront_distribution_cache_behavior_lambda_functions
DROP TABLE aws_cloudfront_distribution_cache_behavior_lambda_functions;

--aws_cloudfront_distribution_default_cache_behavior_lambda_functions
ALTER TABLE "aws_cloudfront_distribution_default_behaviour_lambda_functions" RENAME TO "aws_cache_behaviour_lambda_function_associations";

--aws_cloudfront_distribution_origin_groups
ALTER TABLE "aws_cloudfront_distribution_origin_groups" RENAME COLUMN "failover_criteria_status_codes" TO "failover_criteria_status_codes_items";

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

--aws_elbv2_load_balancer_availability_zones
ALTER TABLE "aws_elbv2_load_balancer_availability_zones"
    ADD CONSTRAINT "aws_elbv2_load_balancer_availability_zones_pk"
        PRIMARY KEY(load_balancer_cq_id, zone_name);

--aws_elbv2_load_balancer_availability_zone_addresses
ALTER TABLE "aws_elbv2_load_balancer_availability_zone_addresses"
    ADD CONSTRAINT "aws_elbv2_load_balancer_availability_zone_addresses_pk"
        PRIMARY KEY(load_balancer_availability_zone_cq_id, ip_address);

