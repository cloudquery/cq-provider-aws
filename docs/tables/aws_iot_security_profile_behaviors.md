
# Table: aws_iot_security_profile_behaviors
A Device Defender security profile behavior.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|security_profile_cq_id|uuid|Unique CloudQuery ID of aws_iot_security_profiles table (FK)|
|name|text|The name you've given to the behavior.  This member is required.|
|criteria_comparison_operator|text|The operator that relates the thing measured (metric) to the criteria (containing a value or statisticalThreshold)|
|criteria_consecutive_datapoints_to_alarm|integer|If a device is in violation of the behavior for the specified number of consecutive datapoints, an alarm occurs|
|criteria_consecutive_datapoints_to_clear|integer|If an alarm has occurred and the offending device is no longer in violation of the behavior for the specified number of consecutive datapoints, the alarm is cleared|
|criteria_duration_seconds|integer|Use this to specify the time duration over which the behavior is evaluated, for those criteria that have a time dimension (for example, NUM_MESSAGES_SENT)|
|criteria__ml_detection_config_confidence_level|text|The sensitivity of anomalous behavior evaluation|
|criteria__statistical_threshold_statistic|text|The percentile that resolves to a threshold value by which compliance with a behavior is determined|
|criteria_value|jsonb|The value to be compared with the metric.|
|metric|text|What is measured by the behavior.|
|metric_dimension_dimension_name|text|A unique identifier for the dimension.  This member is required.|
|metric_dimension_operator|text|Defines how the dimensionValues of a dimension are interpreted|
|suppress_alerts|boolean|Suppresses alerts.|
