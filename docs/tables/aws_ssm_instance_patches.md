
# Table: aws_ssm_instance_patches
Information about the state of a patch on a particular instance as it relates to the patch baseline used to patch the instance.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|instance_cq_id|uuid|Unique CloudQuery ID of aws_ssm_instances table (FK)|
|classification|text|The classification of the patch, such as SecurityUpdates, Updates, and CriticalUpdates.|
|installed_time|timestamp without time zone|The date/time the patch was installed on the instance|
|kb_id|text|The operating system-specific ID of the patch.|
|severity|text|The severity of the patchsuch as Critical, Important, and Moderate.|
|state|text|The state of the patch on the instance, such as INSTALLED or FAILED|
|title|text|The title of the patch.|
|cve_ids|text|The IDs of one or more Common Vulnerabilities and Exposure (CVE) issues that are resolved by the patch.|
