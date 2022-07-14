\echo "Creating view_aws_apigateway_method_settings"
\i sql/views/api_gateway_method_settings.sql

\set check_id 'ApiGateway.1'
\echo "Executing check ApiGateway.1"
\i sql/queries/apigateway/api_gw_execution_logging_enabled.sql

\set check_id 'ApiGateway.2'
\echo "Executing check ApiGateway.2"
\i sql/queries/apigateway/api_gw_ssl_enabled.sql

\set check_id 'ApiGateway.3'
\echo "Executing check ApiGateway.3"
\i sql/queries/apigateway/api_gw_xray_enabled.sql

\set check_id 'ApiGateway.4'
\echo "Executing check ApiGateway.4"
\i sql/queries/apigateway/api_gw_associated_with_waf.sql

\set check_id 'ApiGateway.5'
\echo "Executing check ApiGateway.5"
\i sql/queries/apigateway/api_gw_cache_encrypted.sql
