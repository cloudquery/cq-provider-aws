
# Table: aws_ssm_associations
Describes the parameters for a document.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|arn|text|ARN of the association|
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|apply_only_at_cron_interval|boolean|By default, when you create a new associations, the system runs it immediately after it is created and then according to the schedule you specified|
|id|text|The association ID.|
|association_name|text|The association name.|
|association_version|text|The association version.|
|automation_target_parameter_name|text|Choose the parameter that will define how your automation will branch out|
|calendar_names|text[]|The names or Amazon Resource Names (ARNs) of the Change Calendar type documents your associations are gated under|
|compliance_severity|text|The severity level that is assigned to the association.|
|date|timestamp without time zone|The date when the association was made.|
|document_version|text|The document version.|
|instance_id|text|The instance ID.|
|last_execution_date|timestamp without time zone|The date on which the association was last run.|
|last_successful_execution_date|timestamp without time zone|The last date on which the association was successfully run.|
|last_update_association_date|timestamp without time zone|The date when the association was last updated.|
|max_concurrency|text|The maximum number of targets allowed to run the association at the same time. You can specify a number, for example 10, or a percentage of the target set, for example 10%|
|max_errors|text|The number of errors that are allowed before the system stops sending requests to run the association on additional targets|
|name|text|The name of the SSM document.|
|output_location_s3_bucket_name|text|The name of the S3 bucket.|
|output_location_s3_key_prefix|text|The S3 bucket subfolder.|
|output_location_s3_region|text|The Amazon Web Services Region of the S3 bucket.|
|overview_association_status_aggregated_count|jsonb|Returns the number of targets for the association status|
|overview_detailed_status|text|A detailed status of the association.|
|overview_status|text|The status of the association|
|parameters|jsonb|A description of the parameters for a document.|
|schedule_expression|text|A cron expression that specifies a schedule when the association runs.|
|status_date|timestamp without time zone|The date when the status changed.|
|status_message|text|The reason for the status.|
|status_name|text|The status.|
|status_additional_info|text|A user-defined string.|
|sync_compliance|text|The mode for generating association compliance|
|targets|jsonb|The managed nodes targeted by the request.|
