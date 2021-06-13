
# Table: aws_s3_bucket_lifecycles

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|bucket_id|uuid||
|status|text||
|abort_incomplete_multipart_upload_days_after_initiation|integer||
|expiration_date|timestamp without time zone||
|expiration_days|integer||
|expiration_expired_object_delete_marker|boolean||
|filter|jsonb||
|resource_id|text||
|noncurrent_version_expiration_days|integer||
|noncurrent_version_transitions|jsonb||
|prefix|text||
|transitions|jsonb||
