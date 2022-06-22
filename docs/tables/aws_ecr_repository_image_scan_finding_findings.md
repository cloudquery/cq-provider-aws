
# Table: aws_ecr_repository_image_scan_finding_findings
Contains information about an image scan finding.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|repository_image_scan_finding_cq_id|uuid|Unique CloudQuery ID of aws_ecr_repository_image_scan_findings table (FK)|
|attributes|jsonb|A collection of attributes of the host from which the finding is generated.|
|description|text|The description of the finding.|
|name|text|The name associated with the finding, usually a CVE number.|
|severity|text|The finding severity.|
|uri|text|A link containing additional details about the security vulnerability.|
