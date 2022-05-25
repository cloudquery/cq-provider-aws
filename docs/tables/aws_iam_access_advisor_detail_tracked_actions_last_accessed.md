
# Table: aws_iam_access_advisor_detail_tracked_actions_last_accessed
Contains details about the most recent attempt to access an action within the service
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|access_advisor_detail_cq_id|uuid|Unique CloudQuery ID of aws_iam_access_advisor_details table (FK)|
|action_name|text|The name of the tracked action to which access was attempted|
|last_accessed_entity|text|The Amazon Resource Name (ARN)|
|last_accessed_region|text|The Region from which the authenticated entity (user or role) last attempted to access the tracked action|
|last_accessed_time|timestamp without time zone|The date and time, in ISO 8601 date-time format (http://wwwisoorg/iso/iso8601), when an authenticated entity most recently attempted to access the tracked service|
