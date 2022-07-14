\set check_id 'SSM.1'
\echo "Executing check SSM.1"
\i sql/queries/ssm/ec2_instances_should_be_managed_by_ssm.sql

\set check_id 'SSM.2'
\echo "Executing check SSM.2"
\i sql/queries/ssm/instances_should_have_patch_compliance_status_of_compliant.sql

\set check_id 'SSM.3'
\echo "Executing check SSM.3"
\i sql/queries/ssm/instances_should_have_association_compliance_status_of_compliant.sql

\set check_id 'SSM.4'
\echo "Executing check SSM.4"
\i sql/queries/ssm/documents_should_not_be_public.sql
