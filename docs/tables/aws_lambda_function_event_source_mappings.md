
# Table: aws_lambda_function_event_source_mappings

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
