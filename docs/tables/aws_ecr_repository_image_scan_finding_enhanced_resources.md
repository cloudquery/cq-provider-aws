
# Table: aws_ecr_repository_image_scan_finding_enhanced_resources
Details about the resource involved in a finding.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|repository_image_scan_finding_enhanced_cq_id|uuid|Unique CloudQuery ID of aws_ecr_repository_image_scan_finding_enhanced table (FK)|
|aws_ecr_container_image_architecture|text|The architecture of the Amazon ECR container image.|
|aws_ecr_container_image_author|text|The image author of the Amazon ECR container image.|
|aws_ecr_container_image_image_hash|text|The image hash of the Amazon ECR container image.|
|aws_ecr_container_image_image_tags|text[]|The image tags attached to the Amazon ECR container image.|
|aws_ecr_container_image_platform|text|The platform of the Amazon ECR container image.|
|aws_ecr_container_image_pushed_at|timestamp without time zone|The date and time the Amazon ECR container image was pushed.|
|aws_ecr_container_image_registry|text|The registry the Amazon ECR container image belongs to.|
|aws_ecr_container_image_repository_name|text|The name of the repository the Amazon ECR container image resides in.|
|id|text|The ID of the resource.|
|tags|jsonb|The tags attached to the resource.|
|type|text|The type of resource.|
