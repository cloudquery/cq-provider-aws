
# Table: aws_autoscaling_group_scaling_policy_step_adjustments
Describes information used to create a step adjustment for a step scaling policy
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|group_scaling_policy_cq_id|uuid|Unique CloudQuery ID of aws_autoscaling_group_scaling_policies table (FK)|
|scaling_adjustment|integer|The amount by which to scale, based on the specified adjustment type|
|metric_interval_lower_bound|float|The lower bound for the difference between the alarm threshold and the CloudWatch metric|
|metric_interval_upper_bound|float|The upper bound for the difference between the alarm threshold and the CloudWatch metric|
