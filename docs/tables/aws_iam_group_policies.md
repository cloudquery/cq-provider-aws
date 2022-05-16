
# Table: aws_iam_group_policies
Contains the response to a successful GetGroupPolicy request
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|group_cq_id|uuid|Unique CloudQuery ID of aws_iam_groups table (FK)|
|group_id|text||
|account_id|text|The AWS Account ID of the resource.|
|group_name|text|The group the policy is associated with|
|policy_document|jsonb|The policy document|
|policy_name|text|The name of the policy|
