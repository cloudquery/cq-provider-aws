
# Table: aws_cloudwatch_alarms

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text||
|region|text||
|actions_enabled|boolean||
|alarm_actions|text[]||
|alarm_arn|text||
|alarm_configuration_updated_timestamp|timestamp without time zone||
|alarm_description|text||
|alarm_name|text||
|comparison_operator|text||
|datapoints_to_alarm|integer||
|dimensions|jsonb||
|evaluate_low_sample_count_percentile|text||
|evaluation_periods|integer||
|extended_statistic|text||
|insufficient_data_actions|text[]||
|metric_name|text||
|namespace|text||
|ok_actions|text[]||
|period|integer||
|state_reason|text||
|state_reason_data|text||
|state_updated_timestamp|timestamp without time zone||
|state_value|text||
|statistic|text||
|threshold|float||
|threshold_metric_id|text||
|treat_missing_data|text||
|unit|text||
## Relations
## Table: aws_cloudwatch_alarm_metrics

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|alarm_id|uuid||
|metric_id|text||
|expression|text||
|label|text||
|metric_stat_metric_dimensions|jsonb||
|metric_stat_metric_name|text||
|metric_stat_metric_namespace|text||
|metric_stat_period|integer||
|metric_stat|text||
|metric_stat_unit|text||
|period|integer||
|return_data|boolean||
