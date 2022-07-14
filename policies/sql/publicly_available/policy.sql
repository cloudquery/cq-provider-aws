\set framework 'publicly_available'
\set execution_time ''''`date '+%Y-%m-%d %H:%M:%S'`''''::timestamp
\i sql/create_aws_policy_results.sql

\set check_id 'API-Gateways'
\echo "Executing check API-Gateways"
\i sql/queries/apigateway/api_gw_publicly_accessible.sql

\set check_id 'API-Gateway-V2'
\echo "Executing check API-Gateway-V2"
\i sql/queries/apigateway/api_gw_v2_publicly_accessible.sql

\set check_id 'CloudFront-Distributions'
\echo "Executing check CloudFront-Distributions"
\i sql/queries/cloudfront/all_distributions.sql

\set check_id 'EC2-Public-Ips'
\echo "Executing check EC2-Public-Ips"
\i sql/queries/ec2/public_ips.sql

\set check_id 'ELB-Classic'
\echo "Executing check ELB-Classic"
\i sql/queries/elb/elbv1_internet_facing.sql

\set check_id 'ELB-V2'
\echo "Executing check ELB-V2"
\i sql/queries/elb/elbv2_internet_facing.sql

\set check_id 'Redshift'
\echo "Executing check Redshift"
\i sql/queries/redshift/cluster_publicly_accessible.sql

\set check_id 'RDS'
\echo "Executing check RDS"
\i sql/queries/rds/rds_db_instances_should_prohibit_public_access.sql
