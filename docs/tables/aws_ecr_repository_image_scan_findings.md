
# Table: aws_ecr_repository_image_scan_findings
The details of an image scan.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|repository_image_cq_id|uuid|Unique CloudQuery ID of aws_ecr_repository_images table (FK)|
|finding_severity_counts|jsonb|The image vulnerability counts, sorted by severity.|
|image_scan_completed_at|timestamp without time zone|The time of the last completed image scan.|
|vulnerability_source_updated_at|timestamp without time zone|The time when the vulnerability data was last scanned.|
