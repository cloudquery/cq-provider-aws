
# Table: aws_iam_user_policies
Inline policies that are embedded in the specified IAM user
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|user_id|uuid|User ID the policy belongs too.|
|account_id|text|The AWS Account ID of the resource.|
|policy_document|jsonb|The policy document. IAM stores policies in JSON format. However, resources that were created using AWS CloudFormation templates can be formatted in YAML. AWS CloudFormation always converts a YAML policy to JSON format before submitting it to IAM.|
|policy_name|text|The name of the policy.|
|user_name|text|The user the policy is associated with.|
