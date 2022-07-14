\set framework 'public_egress'
\set execution_time ''''`date '+%Y-%m-%d %H:%M:%S'`''''::timestamp
\i sql/create_aws_policy_results.sql

\echo "Creating view_aws_security_group_egress_rules"
\i sql/views/security_group_egress_rules.sql

\set check_id 'ec2-all-instances-with-routes-and-security-groups'
\echo "Executing check ec2-all-instances-with-routes-and-security-groups"
\i sql/queries/ec2/public_egress_sg_and_routing_instances.sql

\set check_id 'ec2-instances'
\echo "Executing check ec2-instances"
\i sql/queries/ec2/public_egress_sg_instances.sql

\set check_id 'lambda-functions'
\echo "Executing check lambda-functions"
\i sql/queries/lambda/functions_with_public_egress.sql
