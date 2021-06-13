
# Table: aws_ecr_repositories
An object representing a repository.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|created_at|timestamp without time zone|The date and time, in JavaScript date format, when the repository was created.|
|encryption_configuration_encryption_type|text|The encryption type to use.|
|encryption_configuration_kms_key|text|If you use the KMS encryption type, specify the CMK to use for encryption.|
|image_scanning_configuration_scan_on_push|boolean|The setting that determines whether images are scanned after being pushed to a repository.|
|image_tag_mutability|text|The tag mutability setting for the repository.|
|registry_id|text|The AWS account ID associated with the registry that contains the repository.|
|arn|text|The Amazon Resource Name (ARN) that identifies the repository.|
|name|text|The name of the repository.|
|uri|text|The URI for the repository.|
## Relations
## Table: aws_ecr_repository_images
An object that describes an image returned by a DescribeImages operation.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|repository_id|uuid|Unique ID of aws_ecr_repositories table (FK)|
|account_id|text|The AWS Account ID of the resource.|
|region|text||
|artifact_media_type|text|The artifact media type of the image.|
|image_digest|text|The sha256 digest of the image manifest.|
|image_manifest_media_type|text|The media type of the image manifest.|
|image_pushed_at|timestamp without time zone|The date and time, expressed in standard JavaScript date format, at which the current image was pushed to the repository.|
|image_scan_findings_summary_finding_severity_counts|jsonb|The image vulnerability counts, sorted by severity.|
|image_scan_findings_summary_image_scan_completed_at|timestamp without time zone|The time of the last completed image scan.|
|image_scan_findings_summary_vulnerability_source_updated_at|timestamp without time zone|The time when the vulnerability data was last scanned.|
|image_scan_status_description|text|The description of the image scan status.|
|image_scan_status|text|The current state of an image scan.|
|image_size_in_bytes|bigint|The size, in bytes, of the image in the repository.|
|image_tags|text[]|The list of tags associated with this image.|
|registry_id|text|The AWS account ID associated with the registry to which this image belongs.|
|repository_name|text|The name of the repository to which this image belongs.|
