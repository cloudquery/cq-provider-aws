\set check_id 'Cloudfront.1'
\echo "Executing check Cloudfront.1"
\i sql/queries/cloudfront/default_root_object_configured.sql

\set check_id 'Cloudfront.2'
\echo "Executing check Cloudfront.2"
\i sql/queries/cloudfront/origin_access_identity_enabled.sql

\set check_id 'Cloudfront.3'
\echo "Executing check Cloudfront.3"
\i sql/queries/cloudfront/viewer_policy_https.sql

\set check_id 'Cloudfront.4'
\echo "Executing check Cloudfront.4"
\i sql/queries/cloudfront/origin_failover_enabled.sql

\set check_id 'Cloudfront.5'
\echo "Executing check Cloudfront.5"
\i sql/queries/cloudfront/access_logs_enabled.sql

\set check_id 'Cloudfront.6'
\echo "Executing check Cloudfront.6"
\i sql/queries/cloudfront/associated_with_waf.sql
