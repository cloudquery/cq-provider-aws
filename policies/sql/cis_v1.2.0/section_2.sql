\echo  "Executing CIS V1.2.0 Section 2"
\set check_id "2.1"
\echo "Executing check 2.1"
\i sql/queries/cloudtrail/enabled_in_all_regions.sql
\set check_id "2.2"
\echo "Executing check 2.2"
\i sql/queries/cloudtrail/log_file_validation_enabled.sql
\set check_id "2.4"
\echo "Executing check 2.4"
\i sql/queries/cloudtrail/integrated_with_cloudwatch_logs.sql
\set check_id "2.6"
\echo "Executing check 2.6"
\i sql/queries/cloudtrail/bucket_access_logging.sql
\set check_id "2.7"
\echo "Executing check 2.7"
\i sql/queries/cloudtrail/logs_encrypted.sql
\set check_id "2.8"
\echo "Executing check 2.8"
\i sql/queries/kms/rotation_enabled_for_customer_key.sql
\set check_id "2.9"
\echo "Executing check 2.9"
\i sql/queries/ec2/flow_logs_enabled_in_all_vpcs.sql