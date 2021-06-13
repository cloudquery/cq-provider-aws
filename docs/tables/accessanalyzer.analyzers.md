
# Table: aws_access_analyzer_analyzers
Contains information about the analyzer.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|arn|text|The ARN of the analyzer.|
|created_at|timestamp without time zone|A timestamp for the time at which the analyzer was created.|
|name|text|The name of the analyzer.|
|status|text|The status of the analyzer. An Active analyzer successfully monitors supported resources and generates new findings. The analyzer is Disabled when a user action, such as removing trusted access for AWS IAM Access Analyzer from AWS Organizations, causes the analyzer to stop generating new findings. The status is Creating when the analyzer creation is in progress and Failed when the analyzer creation has failed.|
|type|text|The type of analyzer, which corresponds to the zone of trust chosen for the analyzer.|
|last_resource_analyzed|text|The resource that was most recently analyzed by the analyzer.|
|last_resource_analyzed_at|timestamp without time zone|The time at which the most recently analyzed resource was analyzed.|
|status_reason_code|text|The reason code for the current status of the analyzer.|
|tags|jsonb|The tags added to the analyzer.|
## Relations
## Table: aws_access_analyzer_analyzer_findings
Contains information about a finding.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|analyzer_id|uuid|Analyzer ID that belongs to aws_access_analyzer_analyzers|
|analyzed_at|timestamp without time zone|The time at which the resource-based policy that generated the finding was analyzed.|
|condition|jsonb|The condition in the analyzed policy statement that resulted in a finding.|
|created_at|timestamp without time zone|The time at which the finding was created.|
|finding_id|text|The ID of the finding.|
|resource_owner_account|text|The AWS account ID that owns the resource.|
|resource_type|text|The type of the resource that the external principal has access to.|
|status|text|The status of the finding.|
|updated_at|timestamp without time zone|The time at which the finding was most recently updated.|
|action|text[]|The action in the analyzed policy statement that an external principal has permission to use.|
|error|text|The error that resulted in an Error finding.|
|is_public|boolean|Indicates whether the finding reports a resource that has a policy that allows public access.|
|principal|jsonb|The external principal that has access to a resource within the zone of trust.|
|resource|text|The resource that the external principal has access to.|
## Relations
## Table: aws_access_analyzer_analyzer_finding_sources
The source of the finding.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|analyzer_finding_id|uuid|AnalyzerFinding ID that belongs to aws_access_analyzer_analyzer_findings|
|type|text|Indicates the type of access that generated the finding.|
|detail_access_point_arn|text|The ARN of the access point that generated the finding.|
