
# Table: aws_ecr_repositories

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text||
|region|text||
|created_at|timestamp without time zone||
|encryption_configuration_encryption_type|text||
|encryption_configuration_kms_key|text||
|image_scanning_configuration_scan_on_push|boolean||
|image_tag_mutability|text||
|registry_id|text||
|arn|text||
|name|text||
|uri|text||
## Relations
## Table: aws_ecr_repository_images

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|repository_id|uuid||
|account_id|text||
|region|text||
|artifact_media_type|text||
|image_digest|text||
|image_manifest_media_type|text||
|image_pushed_at|timestamp without time zone||
|image_scan_findings_summary_finding_severity_counts|jsonb||
|image_scan_findings_summary_image_scan_completed_at|timestamp without time zone||
|image_scan_findings_summary_vulnerability_source_updated_at|timestamp without time zone||
|image_scan_status_description|text||
|image_scan_status|text||
|image_size_in_bytes|bigint||
|image_tags|text[]||
|registry_id|text||
|repository_name|text||
