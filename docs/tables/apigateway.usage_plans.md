
# Table: aws_apigateway_usage_plans
Represents a usage plan than can specify who can assess associated API stages with specified request limits and quotas.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|description|text|The description of a usage plan.|
|resource_id|text|The identifier of a UsagePlan resource.|
|name|text|The name of a usage plan.|
|product_code|text|The AWS Markeplace product identifier to associate with the usage plan as a SaaS product on AWS Marketplace.|
|quota_limit|integer|The maximum number of requests that can be made in a given time period.|
|quota_offset|integer|The day that a time period starts. For example, with a time period of WEEK, an offset of 0 starts on Sunday, and an offset of 1 starts on Monday.|
|quota_period|text|The time period in which the limit applies. Valid values are "DAY", "WEEK" or "MONTH".|
|tags|jsonb|The collection of tags. Each tag element is associated with a given resource.|
|throttle_burst_limit|integer|The API request burst limit, the maximum rate limit over a time ranging from one to a few seconds, depending upon whether the underlying token bucket is at its full capacity.|
|throttle_rate_limit|float|The API request steady-state rate limit.|
## Relations
## Table: aws_apigateway_usage_plan_api_stages
API stage name of the associated API stage in a usage plan.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|usage_plan_id|uuid|Unique ID of aws_apigateway_usage_plans table (FK)|
|api_id|text|API Id of the associated API stage in a usage plan.|
|stage|text|API stage name of the associated API stage in a usage plan.|
|throttle|jsonb|Map containing method level throttling information for API stage in a usage plan.|
## Table: aws_apigateway_usage_plan_keys
Represents a usage plan key to identify a plan customer.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|usage_plan_id|uuid|Unique ID of aws_apigateway_usage_plans table (FK)|
|resource_id|text|The Id of a usage plan key.|
|name|text|The name of a usage plan key.|
|type|text|The type of a usage plan key. Currently, the valid key type is API_KEY.|
|value|text|The value of a usage plan key.|
