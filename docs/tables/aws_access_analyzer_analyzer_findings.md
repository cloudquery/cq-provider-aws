
# Table: aws_access_analyzer_analyzer_findings
Contains information about a finding.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|analyzer_cq_id|uuid|Unique CloudQuery ID of aws_access_analyzer_analyzers table (FK)|
|analyzed_at|timestamp without time zone|The time at which the resource-based policy that generated the finding was analyzed.  This member is required.|
|condition|jsonb|The condition in the analyzed policy statement that resulted in a finding.  This member is required.|
|created_at|timestamp without time zone|The time at which the finding was created.  This member is required.|
|id|text|The ID of the finding.  This member is required.|
|resource_owner_account|text|The AWS account ID that owns the resource.  This member is required.|
|resource_type|text|The type of the resource that the external principal has access to.  This member is required.|
|status|text|The status of the finding.  This member is required.|
|updated_at|timestamp without time zone|The time at which the finding was most recently updated.  This member is required.|
|action|text[]|The action in the analyzed policy statement that an external principal has permission to use.|
|error|text|The error that resulted in an Error finding.|
|is_public|boolean|Indicates whether the finding reports a resource that has a policy that allows public access.|
|principal|jsonb|The external principal that has access to a resource within the zone of trust.|
|resource|text|The resource that the external principal has access to.|
