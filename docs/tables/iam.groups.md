
# Table: aws_iam_groups
Contains information about an IAM group entity.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|policies|jsonb|List of policies attached to group.|
|arn|text|The Amazon Resource Name (ARN) specifying the group. For more information about ARNs and how to use them in policies, see IAM identifiers (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) in the IAM User Guide.|
|create_date|timestamp without time zone|The date and time, in ISO 8601 date-time format (http://www.iso.org/iso/iso8601), when the group was created.|
|group_id|text|The stable and unique string identifying the group. For more information about IDs, see IAM identifiers (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) in the IAM User Guide.|
|group_name|text|The friendly name that identifies the group.|
|path|text|The path to the group. For more information about paths, see IAM identifiers (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) in the IAM User Guide.|
## Relations
## Table: aws_iam_group_policies
Inline policies that are embedded in the specified IAM group
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|group_name|text|The group the policy is associated with.|
|policy_document|jsonb|The policy document. IAM stores policies in JSON format. However, resources that were created using AWS CloudFormation templates can be formatted in YAML. AWS CloudFormation always converts a YAML policy to JSON format before submitting it to IAM.|
|policy_name|text|The name of the policy.|
