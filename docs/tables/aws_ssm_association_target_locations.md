
# Table: aws_ssm_association_target_locations
The combination of Amazon Web Services Regions and Amazon Web Services accounts targeted by the current Automation execution.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|association_cq_id|uuid|Unique CloudQuery ID of aws_ssm_associations table (FK)|
|accounts|text[]|The Amazon Web Services accounts targeted by the current Automation execution.|
|execution_role_name|text|The Automation execution role used by the currently running Automation|
|regions|text[]|The Amazon Web Services Regions targeted by the current Automation execution.|
|target_location_max_concurrency|text|The maximum number of Amazon Web Services Regions and Amazon Web Services accounts allowed to run the Automation concurrently.|
|target_location_max_errors|text|The maximum number of errors allowed before the system stops queueing additional Automation executions for the currently running Automation.|
