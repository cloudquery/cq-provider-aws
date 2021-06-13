
# Table: aws_cloudtrail_trails

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text||
|region|text||
|cloudwatch_logs_log_group_name|text||
|is_logging|boolean||
|latest_cloud_watch_logs_delivery_error|text||
|latest_cloud_watch_logs_delivery_time|timestamp without time zone||
|latest_delivery_attempt_succeeded|text||
|latest_delivery_attempt_time|text||
|latest_delivery_error|text||
|latest_delivery_time|timestamp without time zone||
|latest_digest_delivery_error|text||
|latest_digest_delivery_time|timestamp without time zone||
|latest_notification_attempt_succeeded|text||
|latest_notification_attempt_time|text||
|latest_notification_error|text||
|latest_notification_time|timestamp without time zone||
|start_logging_time|timestamp without time zone||
|stop_logging_time|timestamp without time zone||
|time_logging_started|text||
|time_logging_stopped|text||
|cloud_watch_logs_log_group_arn|text||
|cloud_watch_logs_role_arn|text||
|has_custom_event_selectors|boolean||
|has_insight_selectors|boolean||
|home_region|text||
|include_global_service_events|boolean||
|is_multi_region_trail|boolean||
|is_organization_trail|boolean||
|kms_key_id|text||
|log_file_validation_enabled|boolean||
|name|text||
|s3_bucket_name|text||
|s3_key_prefix|text||
|sns_topic_arn|text||
|sns_topic_name|text||
|trail_arn|text||
## Relations
## Table: aws_cloudtrail_trail_event_selectors

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|trail_id|uuid||
|exclude_management_event_sources|text[]||
|include_management_events|boolean||
|read_write_type|text||
