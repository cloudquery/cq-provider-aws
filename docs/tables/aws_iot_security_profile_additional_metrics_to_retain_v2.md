
# Table: aws_iot_security_profile_additional_metrics_to_retain_v2
The metric you want to retain
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|security_profile_cq_id|uuid|Unique CloudQuery ID of aws_iot_security_profiles table (FK)|
|metric|text|What is measured by the behavior.  This member is required.|
|metric_dimension_name|text|A unique identifier for the dimension.  This member is required.|
|metric_dimension_operator|text|Defines how the dimensionValues of a dimension are interpreted|
