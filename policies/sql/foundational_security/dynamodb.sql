\set check_id 'DynamoDB.1'
\echo "Executing check DynamoDB.1"
\i sql/queries/dynamodb/autoscale_or_ondemand.sql

\set check_id 'DynamoDB.2'
\echo "Executing check DynamoDB.2"
\i sql/queries/dynamodb/point_in_time_recovery.sql

\set check_id 'DynamoDB.3'
\echo "Executing check DynamoDB.3"
\i sql/queries/dynamodb/dax_encrypted_at_rest.sql
