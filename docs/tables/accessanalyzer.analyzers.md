
# Table: aws_access_analyzer_analyzers

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text||
|region|text||
|arn|text||
|created_at|timestamp without time zone||
|name|text||
|status|text||
|type|text||
|last_resource_analyzed|text||
|last_resource_analyzed_at|timestamp without time zone||
|status_reason_code|text||
|tags|jsonb||
## Relations
## Table: aws_access_analyzer_analyzer_findings

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|analyzer_id|uuid||
|analyzed_at|timestamp without time zone||
|condition|jsonb||
|created_at|timestamp without time zone||
|finding_id|text||
|resource_owner_account|text||
|resource_type|text||
|status|text||
|updated_at|timestamp without time zone||
|action|text[]||
|error|text||
|is_public|boolean||
|principal|jsonb||
|resource|text||
## Relations
## Table: aws_access_analyzer_analyzer_finding_sources

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|analyzer_finding_id|uuid||
|type|text||
|detail_access_point_arn|text||
