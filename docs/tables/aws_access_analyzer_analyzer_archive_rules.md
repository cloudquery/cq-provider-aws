
# Table: aws_access_analyzer_analyzer_archive_rules
Contains information about an archive rule.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|analyzer_cq_id|uuid|Unique CloudQuery ID of aws_access_analyzer_analyzers table (FK)|
|created_at|timestamp without time zone|The time at which the archive rule was created.  This member is required.|
|filter|jsonb|A filter used to define the archive rule.  This member is required.|
|rule_name|text|The name of the archive rule.  This member is required.|
|updated_at|timestamp without time zone|The time at which the archive rule was last updated.  This member is required.|
