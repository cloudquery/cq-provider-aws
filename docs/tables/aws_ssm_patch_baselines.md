
# Table: aws_ssm_patch_baselines

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|arn|text|ARN of the resource.|
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|approved_patches|text[]|A list of explicitly approved patches for the baseline.|
|approved_patches_compliance_level|text|Returns the specified compliance severity level for approved patches in the patch baseline.|
|approved_patches_enable_non_security|boolean|Indicates whether the list of approved patches includes non-security updates that should be applied to the instances|
|baseline_id|text|The ARN of the retrieved patch baseline.|
|created_date|timestamp without time zone|The date the patch baseline was created.|
|description|text|A description of the patch baseline.|
|global_filters|jsonb|A set of global filters used to exclude patches from the baseline.|
|modified_date|timestamp without time zone|The date the patch baseline was last modified.|
|name|text|The name of the patch baseline.|
|operating_system|text|Returns the operating system specified for the patch baseline.|
|patch_groups|text[]|Patch groups included in the patch baseline.|
|rejected_patches|text[]|A list of explicitly rejected patches for the baseline.|
|rejected_patches_action|text|The action specified to take on patches included in the RejectedPatches list|
|tags|jsonb|Resource tags.|
