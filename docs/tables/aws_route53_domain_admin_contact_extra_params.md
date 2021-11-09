
# Table: aws_route53_domain_admin_contact_extra_params
ExtraParam includes the following elements.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|domain_cq_id|uuid|Unique CloudQuery ID of aws_route53_domains table (FK)|
|name|text|The name of an additional parameter that is required by a top-level domain|
|value|text|The value that corresponds with the name of an extra parameter.  This member is required.|
