
# Table: aws_elasticbeanstalk_environments

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text||
|region|text||
|abortable_operation_in_progress|boolean||
|application_name|text||
|cname|text||
|date_created|timestamp without time zone||
|date_updated|timestamp without time zone||
|description|text||
|endpoint_url|text||
|environment_arn|text||
|environment_id|text||
|environment_name|text||
|health|text||
|health_status|text||
|operations_role|text||
|platform_arn|text||
|load_balancer_domain|text||
|load_balancer_name|text||
|solution_stack_name|text||
|status|text||
|template_name|text||
|tier_name|text||
|tier_type|text||
|tier_version|text||
|version_label|text||
## Relations
## Table: aws_elasticbeanstalk_environment_links

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|environment_id|uuid||
|environment_name|text||
|link_name|text||
## Table: aws_elasticbeanstalk_env_resources_load_balancer_listeners

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|environment_id|uuid||
|port|integer||
|protocol|text||
