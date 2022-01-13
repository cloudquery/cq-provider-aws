module github.com/cloudquery/cq-provider-aws

go 1.16

require (
	github.com/aws/aws-sdk-go-v2 v1.12.0
	github.com/aws/aws-sdk-go-v2/config v1.3.0
	github.com/aws/aws-sdk-go-v2/credentials v1.2.1
	github.com/aws/aws-sdk-go-v2/feature/s3/manager v1.2.1
	github.com/aws/aws-sdk-go-v2/service/accessanalyzer v1.4.1
	github.com/aws/aws-sdk-go-v2/service/acm v1.9.1
	github.com/aws/aws-sdk-go-v2/service/apigateway v1.4.0
	github.com/aws/aws-sdk-go-v2/service/apigatewayv2 v1.3.1
	github.com/aws/aws-sdk-go-v2/service/applicationautoscaling v1.10.2
	github.com/aws/aws-sdk-go-v2/service/autoscaling v1.2.0
	github.com/aws/aws-sdk-go-v2/service/cloudfront v1.3.0
	github.com/aws/aws-sdk-go-v2/service/cloudtrail v1.1.2
	github.com/aws/aws-sdk-go-v2/service/cloudwatch v1.1.2
	github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs v1.1.2
	github.com/aws/aws-sdk-go-v2/service/cognitoidentity v1.3.1
	github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider v1.3.3
	github.com/aws/aws-sdk-go-v2/service/configservice v1.5.1
	github.com/aws/aws-sdk-go-v2/service/databasemigrationservice v1.12.0
	github.com/aws/aws-sdk-go-v2/service/dax v1.7.2
	github.com/aws/aws-sdk-go-v2/service/directconnect v1.4.1
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.10.0
	github.com/aws/aws-sdk-go-v2/service/ec2 v1.16.0
	github.com/aws/aws-sdk-go-v2/service/ecr v1.2.0
	github.com/aws/aws-sdk-go-v2/service/ecs v1.2.0
	github.com/aws/aws-sdk-go-v2/service/efs v1.2.0
	github.com/aws/aws-sdk-go-v2/service/eks v1.2.1
	github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk v1.9.1
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing v1.3.0
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 v1.2.0
	github.com/aws/aws-sdk-go-v2/service/elasticsearchservice v1.4.0
	github.com/aws/aws-sdk-go-v2/service/emr v1.2.0
	github.com/aws/aws-sdk-go-v2/service/fsx v1.2.0
	github.com/aws/aws-sdk-go-v2/service/guardduty v1.7.1
	github.com/aws/aws-sdk-go-v2/service/iam v1.3.0
	github.com/aws/aws-sdk-go-v2/service/iot v1.20.0
	github.com/aws/aws-sdk-go-v2/service/kms v1.2.1
	github.com/aws/aws-sdk-go-v2/service/lambda v1.3.0
	github.com/aws/aws-sdk-go-v2/service/mq v1.2.1
	github.com/aws/aws-sdk-go-v2/service/organizations v1.2.1
	github.com/aws/aws-sdk-go-v2/service/rds v1.2.1
	github.com/aws/aws-sdk-go-v2/service/redshift v1.3.0
	github.com/aws/aws-sdk-go-v2/service/route53 v1.4.0
	github.com/aws/aws-sdk-go-v2/service/s3 v1.19.1
	github.com/aws/aws-sdk-go-v2/service/s3control v1.14.1
	github.com/aws/aws-sdk-go-v2/service/sagemaker v1.19.1
	github.com/aws/aws-sdk-go-v2/service/secretsmanager v1.10.2
	github.com/aws/aws-sdk-go-v2/service/sns v1.1.2
	github.com/aws/aws-sdk-go-v2/service/sqs v1.9.1
	github.com/aws/aws-sdk-go-v2/service/ssm v1.16.0
	github.com/aws/aws-sdk-go-v2/service/sts v1.4.1
	github.com/aws/aws-sdk-go-v2/service/waf v1.2.1
	github.com/aws/aws-sdk-go-v2/service/wafv2 v1.5.1
	github.com/aws/smithy-go v1.9.1
	github.com/bxcodec/faker v2.0.1+incompatible
	github.com/cloudquery/cq-provider-sdk v0.6.1
	github.com/cloudquery/faker/v3 v3.7.5
	github.com/gocarina/gocsv v0.0.0-20210516172204-ca9e8a8ddea8
	github.com/golang/mock v1.6.0
	github.com/hashicorp/go-hclog v1.0.0
	github.com/mitchellh/mapstructure v1.4.2
	github.com/spf13/cast v1.4.1
	github.com/stretchr/testify v1.7.0
)

require (
	github.com/kr/text v0.2.0 // indirect
	github.com/kylelemons/godebug v1.1.0 // indirect
	github.com/sergi/go-diff v1.2.0 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
)

require (
	github.com/aws/aws-sdk-go-v2/service/codebuild v1.14.1
	github.com/aws/aws-sdk-go-v2/service/route53domains v1.6.0
	github.com/google/go-cmp v0.5.6
)
