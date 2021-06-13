
# Table: aws_lambda_function_versions

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
