
# Table: aws_ecr_repository_image_scan_finding_enhanced
The details of an enhanced image scan
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|repository_image_scan_finding_cq_id|uuid|Unique CloudQuery ID of aws_ecr_repository_image_scan_findings table (FK)|
|aws_account_id|text|The Amazon Web Services account ID associated with the image.|
|description|text|The description of the finding.|
|finding_arn|text|The Amazon Resource Number (ARN) of the finding.|
|first_observed_at|timestamp without time zone|The date and time that the finding was first observed.|
|last_observed_at|timestamp without time zone|The date and time that the finding was last observed.|
|package_vulnerability_details_cvss|jsonb|An object that contains details about the CVSS score of a finding.|
|package_vulnerability_details_reference_urls|text[]|One or more URLs that contain details about this vulnerability type.|
|package_vulnerability_details_related_vulnerabilities|text[]|One or more vulnerabilities related to the one identified in this finding.|
|package_vulnerability_details_source|text|The source of the vulnerability information.|
|package_vulnerability_details_source_url|text|A URL to the source of the vulnerability information.|
|package_vulnerability_details_vendor_created_at|timestamp without time zone|The date and time that this vulnerability was first added to the vendor's database.|
|package_vulnerability_details_vendor_severity|text|The severity the vendor has given to this vulnerability type.|
|package_vulnerability_details_vendor_updated_at|timestamp without time zone|The date and time the vendor last updated this vulnerability in their database.|
|package_vulnerability_details_vulnerability_id|text|The ID given to this vulnerability.|
|package_vulnerability_details_vulnerable_packages|jsonb|The packages impacted by this vulnerability.|
|remediation_recommendation_text|text|The recommended course of action to remediate the finding.|
|remediation_recommendation_url|text|The URL address to the CVE remediation recommendations.|
|score|float|The Amazon Inspector score given to the finding.|
|score_details_cvss_adjustments|jsonb|An object that contains details about adjustment Amazon Inspector made to the CVSS score.|
|score_details_cvss_score|float|The CVSS score.|
|score_details_cvss_score_source|text|The source for the CVSS score.|
|score_details_cvss_scoring_vector|text|The vector for the CVSS score.|
|score_details_cvss_version|text|The CVSS version used in scoring.|
|severity|text|The severity of the finding.|
|status|text|The status of the finding.|
|title|text|The title of the finding.|
|type|text|The type of the finding.|
|updated_at|timestamp without time zone|The date and time the finding was last updated at.|
