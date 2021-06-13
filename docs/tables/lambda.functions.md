
# Table: aws_lambda_functions

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text||
|region|text||
|policy_document|jsonb||
|policy_revision_id|text||
|code_signing_allowed_publishers_version_arns|text[]||
|code_signing_config_arn|text||
|code_signing_config_id|text||
|code_signing_policies_untrusted_artifact_on_deployment|text||
|code_signing_description|text||
|code_signing_last_modified|timestamp without time zone||
|code_image_uri|text||
|code_location|text||
|code_repository_type|text||
|code_resolved_image_uri|text||
|concurrency_reserved_concurrent_executions|integer||
|code_sha256|text||
|code_size|bigint||
|dead_letter_config_target_arn|text||
|description|text||
|environment_error_code|text||
|environment_error_message|text||
|environment_variables|jsonb||
|function_arn|text||
|function_name|text||
|handler|text||
|error_code|text||
|error_message|text||
|image_config_command|text[]||
|image_config_entry_point|text[]||
|image_config_working_directory|text||
|kms_key_arn|text||
|last_modified|text||
|last_update_status|text||
|last_update_status_reason|text||
|last_update_status_reason_code|text||
|master_arn|text||
|memory_size|integer||
|package_type|text||
|revision_id|text||
|role|text||
|runtime|text||
|signing_job_arn|text||
|signing_profile_version_arn|text||
|state|text||
|state_reason|text||
|state_reason_code|text||
|timeout|integer||
|tracing_config_mode|text||
|version|text||
|vpc_config_security_group_ids|text[]||
|vpc_config_subnet_ids|text[]||
|vpc_config_vpc_id|text||
|tags|jsonb||
## Relations
## Table: aws_lambda_function_file_system_configs

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|function_id|uuid||
|arn|text||
|local_mount_path|text||
## Table: aws_lambda_function_layers

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|function_id|uuid||
|arn|text||
|code_size|bigint||
|signing_job_arn|text||
|signing_profile_version_arn|text||
## Table: aws_lambda_function_aliases

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|function_id|uuid||
|alias_arn|text||
|description|text||
|function_version|text||
|name|text||
|revision_id|text||
|routing_config_additional_version_weights|jsonb||
## Table: aws_lambda_function_event_invoke_configs

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|function_id|uuid||
|on_failure_destination|text||
|on_success_destination|text||
|function_arn|text||
|last_modified|timestamp without time zone||
|maximum_event_age_in_seconds|integer||
|maximum_retry_attempts|integer||
## Table: aws_lambda_function_versions

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|function_id|uuid||
|code_sha256|text||
|code_size|bigint||
|dead_letter_config_target_arn|text||
|description|text||
|environment_error_error_code|text||
|environment_error_message|text||
|environment_variables|jsonb||
|function_arn|text||
|function_name|text||
|handler|text||
|error_code|text||
|error_message|text||
|image_config_command|text[]||
|image_config_entry_point|text[]||
|image_config_working_directory|text||
|kms_key_arn|text||
|last_modified|text||
|last_update_status|text||
|last_update_status_reason|text||
|last_update_status_reason_code|text||
|master_arn|text||
|memory_size|integer||
|package_type|text||
|revision_id|text||
|role|text||
|runtime|text||
|signing_job_arn|text||
|signing_profile_version_arn|text||
|state|text||
|state_reason|text||
|state_reason_code|text||
|timeout|integer||
|tracing_config_mode|text||
|version|text||
|vpc_config_security_group_ids|text[]||
|vpc_config_subnet_ids|text[]||
|vpc_config_vpc_id|text||
## Relations
## Table: aws_lambda_function_version_file_system_configs

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|function_version_id|uuid||
|arn|text||
|local_mount_path|text||
## Table: aws_lambda_function_version_layers

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|function_version_id|uuid||
|arn|text||
|code_size|bigint||
|signing_job_arn|text||
|signing_profile_version_arn|text||
## Table: aws_lambda_function_concurrency_configs

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|function_id|uuid||
|allocated_provisioned_concurrent_executions|integer||
|available_provisioned_concurrent_executions|integer||
|function_arn|text||
|last_modified|text||
|requested_provisioned_concurrent_executions|integer||
|status|text||
|status_reason|text||
## Table: aws_lambda_function_event_source_mappings

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|function_id|uuid||
|batch_size|integer||
|bisect_batch_on_function_error|boolean||
|on_failure_destination|text||
|on_success_destination|text||
|event_source_arn|text||
|function_arn|text||
|function_response_types|text[]||
|last_modified|timestamp without time zone||
|last_processing_result|text||
|maximum_batching_window_in_seconds|integer||
|maximum_record_age_in_seconds|integer||
|maximum_retry_attempts|integer||
|parallelization_factor|integer||
|queues|text[]||
|self_managed_event_source_endpoints|jsonb||
|starting_position|text||
|starting_position_timestamp|timestamp without time zone||
|state|text||
|state_transition_reason|text||
|topics|text[]||
|tumbling_window_in_seconds|integer||
|uuid|text||
## Relations
## Table: aws_lambda_function_event_source_mapping_access_configurations

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|function_event_source_mapping_id|uuid||
|type|text||
|uri|text||
