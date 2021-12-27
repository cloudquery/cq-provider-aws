
# Table: aws_iot_security_profiles
A security profile defines a set of expected behaviors for devices in your account and specifies the actions to take when an anomaly is detected
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|targets|text[]||
|additional_metrics_to_retain|text[]|Please use DescribeSecurityProfileResponse$additionalMetricsToRetainV2 instead. A list of metrics whose data is retained (stored)|
|alert_targets|jsonb|Where the alerts are sent|
|creation_date|timestamp without time zone|The time the security profile was created.|
|last_modified_date|timestamp without time zone|The time the security profile was last modified.|
|arn|text|The ARN of the security profile.|
|description|text|A description of the security profile (associated with the security profile when it was created or updated).|
|name|text|The name of the security profile.|
|version|bigint|The version of the security profile|
