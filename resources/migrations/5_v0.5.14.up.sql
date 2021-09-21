--aws_apigatewayv2_vpc_links
ALTER TABLE "aws_apigatewayv2_vpc_links"
DROP
COLUMN "vpc_link_id";

--aws_cloudfront_distributions
ALTER TABLE "aws_cloudfront_distributions"
    ADD COLUMN "tags" json;
ALTER TABLE "aws_cloudfront_distributions" RENAME COLUMN "aliases_items" TO "aliases";
-- ALTER TABLE "aws_cloudfront_distributions" RENAME COLUMN "cache_behaviour_target_origin_id" TO "cache_behavior_target_origin_id";
-- ALTER TABLE "aws_cloudfront_distributions" RENAME COLUMN "cache_behaviour_viewer_protocol_policy" TO "cache_behavior_viewer_protocol_policy";
-- ALTER TABLE "aws_cloudfront_distributions" RENAME COLUMN "cache_behaviour_allowed_methods" TO "cache_behavior_allowed_methods";
-- ALTER TABLE "aws_cloudfront_distributions" RENAME COLUMN "cache_behaviour_allowed_methods_cached_methods" TO "cache_behavior_allowed_methods_cached_methods";
-- ALTER TABLE "aws_cloudfront_distributions" RENAME COLUMN "cache_behaviour_cache_policy_id" TO "cache_behavior_cache_policy_id";
-- ALTER TABLE "aws_cloudfront_distributions" RENAME COLUMN "cache_behaviour_compress" TO "cache_behavior_compress";
-- ALTER TABLE "aws_cloudfront_distributions" RENAME COLUMN "cache_behaviour_default_ttl" TO "cache_behavior_default_ttl";
-- ALTER TABLE "aws_cloudfront_distributions" RENAME COLUMN "cache_behaviour_field_level_encryption_id" TO "cache_behavior_field_level_encryption_id";
-- ALTER TABLE "aws_cloudfront_distributions" RENAME COLUMN "cache_behaviour_forwarded_values_cookies_forward" TO "cache_behavior_forwarded_values_cookies_forward";
-- ALTER TABLE "aws_cloudfront_distributions" RENAME COLUMN "cache_behaviour_forwarded_values_cookies_white_listed_names" TO "cache_behavior_forwarded_values_cookies_whitelisted_names";
-- ALTER TABLE "aws_cloudfront_distributions" RENAME COLUMN "cache_behaviour_forwarded_values_query_string" TO "cache_behavior_forwarded_values_query_string";
-- ALTER TABLE "aws_cloudfront_distributions" RENAME COLUMN "cache_behaviour_forwarded_values_headers" TO "cache_behavior_forwarded_values_headers";
-- ALTER TABLE "aws_cloudfront_distributions" RENAME COLUMN "cache_behaviour_forwarded_values_query_string_cache_keys" TO "cache_behavior_forwarded_values_query_string_cache_keys";
-- ALTER TABLE "aws_cloudfront_distributions" RENAME COLUMN "cache_behaviour_max_ttl" TO "cache_behavior_max_ttl";
-- ALTER TABLE "aws_cloudfront_distributions" RENAME COLUMN "cache_behaviour_min_ttl" TO "cache_behavior_min_ttl";
-- ALTER TABLE "aws_cloudfront_distributions" RENAME COLUMN "cache_behaviour_origin_request_policy_id" TO "cache_behavior_origin_request_policy_id";
-- ALTER TABLE "aws_cloudfront_distributions" RENAME COLUMN "cache_behaviour_realtime_log_config_arn" TO "cache_behavior_realtime_log_config_arn";
-- ALTER TABLE "aws_cloudfront_distributions" RENAME COLUMN "cache_behaviour_smooth_streaming" TO "cache_behavior_smooth_streaming";
-- ALTER TABLE "aws_cloudfront_distributions" RENAME COLUMN "cache_behaviour_trusted_key_groups_enabled" TO "cache_behavior_trusted_key_groups_enabled";
-- ALTER TABLE "aws_cloudfront_distributions" RENAME COLUMN "cache_behaviour_trusted_key_groups" TO "cache_behavior_trusted_key_groups";
-- ALTER TABLE "aws_cloudfront_distributions" RENAME COLUMN "cache_behaviour_trusted_signers_enabled" TO "cache_behavior_trusted_signers_enabled";
-- ALTER TABLE "aws_cloudfront_distributions" RENAME COLUMN "cache_behaviour_trusted_signers" TO "cache_behavior_trusted_signers";
ALTER TABLE "aws_cloudfront_distributions" RENAME COLUMN "restrictions_geo_restriction_restriction_items" TO "restrictions_geo_restriction_restrictions";

--aws_cloudfront_distribution_cache_behaviours
ALTER TABLE "aws_cloudfront_distributions"
    ADD COLUMN "cache_policy_id" TEXT;
ALTER TABLE "aws_cloudfront_distributions"
    ADD COLUMN "compress" BOOL;
ALTER TABLE "aws_cloudfront_distributions"
    ADD COLUMN "default_ttl" INT8;
ALTER TABLE "aws_cloudfront_distributions"
    ADD COLUMN "field_level_encryption_id" TEXT;
ALTER TABLE "aws_cloudfront_distributions"
    ADD COLUMN "forwarded_values_cookies_forward" TEXT;
ALTER TABLE "aws_cloudfront_distributions"
    ADD COLUMN "forwarded_values_cookies_whitelisted_names" _TEXT;
ALTER TABLE "aws_cloudfront_distributions"
    ADD COLUMN "forwarded_values_query_string" BOOL;
ALTER TABLE "aws_cloudfront_distributions"
    ADD COLUMN "forwarded_values_headers" _TEXT;
ALTER TABLE "aws_cloudfront_distributions"
    ADD COLUMN "forwarded_values_query_string_cache_keys" _TEXT;
ALTER TABLE "aws_cloudfront_distributions"
    ADD COLUMN "max_ttl" INT8;
ALTER TABLE "aws_cloudfront_distributions"
    ADD COLUMN "min_ttl" INT8;
ALTER TABLE "aws_cloudfront_distributions"
    ADD COLUMN "origin_request_policy_id" TEXT;
ALTER TABLE "aws_cloudfront_distributions"
    ADD COLUMN "realtime_log_config_arn" TEXT;
ALTER TABLE "aws_cloudfront_distributions"
    ADD COLUMN "smooth_streaming" BOOL;
ALTER TABLE "aws_cloudfront_distributions"
    ADD COLUMN "trusted_key_groups_enabled" BOOL;
ALTER TABLE "aws_cloudfront_distributions"
    ADD COLUMN "trusted_key_groups" _TEXT;
ALTER TABLE "aws_cloudfront_distributions"
    ADD COLUMN "trusted_signers_enabled" BOOL;
ALTER TABLE "aws_cloudfront_distributions"
    ADD COLUMN "trusted_signers" _TEXT;

--aws_cloudfront_distribution_cache_behavior_lambda_functions
CREATE TABLE aws_cloudfront_distribution_cache_behavior_lambda_functions
(
    cq_id                             uuid NOT NULL,
    meta                              jsonb NULL,
    distribution_cache_behavior_cq_id uuid NULL,
    event_type                        text NULL,
    lambda_function_arn               text NULL,
    include_body                      bool NULL,
    CONSTRAINT aws_cloudfront_distribution_cache_behavior_lambda_functions_pk PRIMARY KEY (cq_id)
);

--aws_cloudfront_distribution_default_cache_behavior_lambda_functions
ALTER TABLE "aws_cache_behaviour_lambda_function_associations" RENAME TO "aws_cloudfront_distribution_default_behaviour_lambda_functions";

--aws_cloudfront_distribution_origin_groups
ALTER TABLE "aws_cloudfront_distribution_origin_groups" RENAME COLUMN "failover_criteria_status_codes_items" TO "failover_criteria_status_codes";


--aws_cloudtrail_trails
ALTER TABLE "aws_cloudtrail_trails"
    ADD COLUMN "tags" json;
ALTER TABLE "aws_cloudtrail_trails"
    DROP COLUMN "home_region";

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

--aws_elbv2_load_balancer_availability_zones  drop pk

--aws_elbv2_load_balancer_availability_zone_addresses  drop pk

