
# Table: aws_ssm_patch_baseline_approval_rules
Defines an approval rule for a patch baseline.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|patch_baseline_cq_id|uuid|Unique CloudQuery ID of aws_ssm_patch_baselines table (FK)|
|approve_after_days|integer|The number of days after the release date of each patch matched by the rule that the patch is marked as approved in the patch baseline|
|approve_until_date|text|The cutoff date for auto approval of released patches|
|compliance_level|text|A compliance severity level for all approved patches in a patch baseline.|
|enable_non_security|boolean|For instances identified by the approval rule filters, enables a patch baseline to apply non-security updates available in the specified repository|
|patch_filter_group|jsonb|The patch filter group that defines the criteria for the rule.|
