
# Table: aws_ssm_patch_baseline_sources
Information about the patches to use to update the instances, including target operating systems and source repository
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|patch_baseline_cq_id|uuid|Unique CloudQuery ID of aws_ssm_patch_baselines table (FK)|
|configuration|text|The value of the yum repo configuration|
|name|text|The name specified to identify the patch source.  This member is required.|
|products|text[]|The specific operating system versions a patch repository applies to, such as "Ubuntu16.04", "AmazonLinux2016.09", "RedhatEnterpriseLinux7.2" or "Suse12.7". For lists of supported product values, see PatchFilter.  This member is required.|
