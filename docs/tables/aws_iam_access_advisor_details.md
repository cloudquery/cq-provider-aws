
# Table: aws_iam_access_advisor_details
IAM Access Advisor details of the IAM resources (users, groups, roles, policies)
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|parent_type|text|Type of parent IAM resource|
|service_name|text|The name of the service in which access was attempted|
|service_namespace|text|The namespace of the service in which access was attempted|
|last_authenticated|timestamp without time zone|The date and time, in ISO 8601 date-time format (http://wwwisoorg/iso/iso8601), when an authenticated entity most recently attempted to access the service|
|last_authenticated_entity|text|The ARN of the authenticated entity (user or role) that last attempted to access the service|
|last_authenticated_region|text|The Region from which the authenticated entity (user or role) last attempted to access the service|
|total_authenticated_entities|integer|The total number of authenticated principals (root user, IAM users, or IAM roles) that have attempted to access the service|
