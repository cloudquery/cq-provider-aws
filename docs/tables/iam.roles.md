
# Table: aws_iam_roles
An IAM role is an IAM identity that you can create in your account that has specific permissions.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|policies|jsonb|List of policies attached to group.|
|arn|text|The Amazon Resource Name (ARN) specifying the role. For more information about ARNs and how to use them in policies, see IAM identifiers (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) in the IAM User Guide guide.|
|create_date|timestamp without time zone|The date and time, in ISO 8601 date-time format (http://www.iso.org/iso/iso8601), when the role was created.|
|path|text|The path to the role. For more information about paths, see IAM identifiers (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) in the IAM User Guide.|
|role_id|text|The stable and unique string identifying the role. For more information about IDs, see IAM identifiers (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) in the IAM User Guide.|
|role_name|text|The friendly name that identifies the role.|
|assume_role_policy_document|jsonb|The policy that grants an entity permission to assume the role. |
|description|text|A description of the role that you provide. |
|max_session_duration|integer|The maximum session duration (in seconds) for the specified role. Anyone who uses the AWS CLI, or API to assume the role can specify the duration using the optional DurationSeconds API parameter or duration-seconds CLI parameter. |
|permissions_boundary_arn|text|The ARN of the policy used to set the permissions boundary for the user or role. |
|permissions_boundary_type|text|The permissions boundary usage type that indicates what type of IAM resource is used as the permissions boundary for an entity. This data type can only have a value of Policy. |
|role_last_used_last_used_date|timestamp without time zone|The date and time, in ISO 8601 date-time format (http://www.iso.org/iso/iso8601) that the role was last used. This field is null if the role has not been used within the IAM tracking period. For more information about the tracking period, see Regions where data is tracked (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_access-advisor.html#access-advisor_tracking-period) in the IAM User Guide. |
|role_last_used_region|text|The name of the AWS Region in which the role was last used. |
|tags|jsonb|A list of tags that are attached to the role. For more information about tagging, see Tagging IAM resources (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html) in the IAM User Guide. |
## Relations
## Table: aws_iam_role_policies
Inline policies that are embedded in the specified IAM role
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|role_id|uuid|Role ID the policy belongs too.|
|account_id|text|The AWS Account ID of the resource.|
|policy_document|jsonb|The policy document. IAM stores policies in JSON format. However, resources that were created using AWS CloudFormation templates can be formatted in YAML. AWS CloudFormation always converts a YAML policy to JSON format before submitting it to IAM.|
|policy_name|text|The name of the policy.|
|role_name|text|The role the policy is associated with.|
