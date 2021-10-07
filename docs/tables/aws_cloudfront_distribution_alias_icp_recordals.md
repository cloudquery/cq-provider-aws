
# Table: aws_cloudfront_distribution_alias_icp_recordals
AWS services in China customers must file for an Internet Content Provider (ICP) recordal if they want to serve content publicly on an alternate domain name, also known as a CNAME, that they've added to CloudFront
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|distribution_cq_id|uuid|Unique CloudQuery ID of aws_cloudfront_distributions table (FK)|
|cname|text|A domain name associated with a distribution.|
|icp_recordal_status|text|The Internet Content Provider (ICP) recordal status for a CNAME|
