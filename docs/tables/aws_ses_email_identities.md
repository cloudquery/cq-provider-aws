
# Table: aws_ses_email_identities
Details about an email identity.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|arn|text|The Amazon Resource Name (ARN) for the resource.|
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|configuration_set_name|text|The configuration set used by default when sending from this identity.|
|dkim_attributes_current_signing_key_length|text|[Easy DKIM] The key length of the DKIM key pair in use.|
|dkim_attributes_last_key_generation_timestamp|timestamp without time zone|[Easy DKIM] The last time a key pair was generated for this identity.|
|dkim_attributes_next_signing_key_length|text|[Easy DKIM] The key length of the future DKIM key pair to be generated|
|dkim_attributes_signing_attributes_origin|text|A string that indicates how DKIM was configured for the identity|
|dkim_attributes_signing_enabled|boolean|If the value is true, then the messages that you send from the identity are signed using DKIM|
|dkim_attributes_status|text|Describes whether or not Amazon SES has successfully located the DKIM records in the DNS records for the domain|
|dkim_attributes_tokens|text[]|If you used Easy DKIM (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/easy-dkim.html) to configure DKIM authentication for the domain, then this object contains a set of unique strings that you use to create a set of CNAME records that you add to the DNS configuration for your domain|
|feedback_forwarding_status|boolean|The feedback forwarding configuration for the identity|
|identity_type|text|The email identity type|
|behavior_on_mx_failure|text|The action to take if the required MX record can't be found when you send an email|
|mail_from_domain|text|The name of a domain that an email identity uses as a custom MAIL FROM domain.  This member is required.|
|mail_from_domain_status|text|The status of the MAIL FROM domain|
|policies|jsonb|A map of policy names to policies.|
|tags|jsonb|An array of objects that define the tags (keys and values) that are associated with the email identity.|
|verified_for_sending_status|boolean|Specifies whether or not the identity is verified|
