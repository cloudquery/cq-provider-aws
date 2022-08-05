package client

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer"
	"github.com/aws/aws-sdk-go-v2/service/acm"
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	"github.com/aws/aws-sdk-go-v2/service/applicationautoscaling"
	"github.com/aws/aws-sdk-go-v2/service/athena"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	"github.com/aws/aws-sdk-go-v2/service/backup"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/codebuild"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/configservice"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice"
	"github.com/aws/aws-sdk-go-v2/service/dax"
	"github.com/aws/aws-sdk-go-v2/service/directconnect"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/efs"
	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
	elbv1 "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing"
	elbv2 "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice"
	"github.com/aws/aws-sdk-go-v2/service/emr"
	"github.com/aws/aws-sdk-go-v2/service/fsx"
	"github.com/aws/aws-sdk-go-v2/service/guardduty"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/aws/aws-sdk-go-v2/service/mq"
	"github.com/aws/aws-sdk-go-v2/service/organizations"
	"github.com/aws/aws-sdk-go-v2/service/qldb"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/route53domains"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3control"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/shield"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/waf"
	"github.com/aws/aws-sdk-go-v2/service/wafregional"
	"github.com/aws/aws-sdk-go-v2/service/wafv2"
	"github.com/aws/aws-sdk-go-v2/service/workspaces"
	"github.com/aws/aws-sdk-go-v2/service/xray"
)

// S3Manager This is needed because https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/feature/s3/manager
// has different structure then all other services (i.e no service but just a function) and we need
// the ability to mock it.
// Also we need to use s3 manager to be able to query the bucket-region https://github.com/aws/aws-sdk-go-v2/pull/1027#issuecomment-759818990
type S3Manager struct {
	s3Client *s3.Client
}

func newS3ManagerFromConfig(cfg aws.Config) S3Manager {
	return S3Manager{
		s3Client: s3.NewFromConfig(cfg),
	}
}

func (s3Manager S3Manager) GetBucketRegion(ctx context.Context, bucket string, optFns ...func(*s3.Options)) (string, error) {
	return manager.GetBucketRegion(ctx, s3Manager.s3Client, bucket, optFns...)
}

type Services struct {
	ACM                    ACMClient
	Analyzer               AnalyzerClient
	Apigateway             ApigatewayClient
	Apigatewayv2           Apigatewayv2Client
	ApplicationAutoscaling ApplicationAutoscalingClient
	Athena                 AthenaClient
	Autoscaling            AutoscalingClient
	Backup                 BackupClient
	Cloudformation         CloudFormationClient
	Cloudfront             CloudfrontClient
	Cloudtrail             CloudtrailClient
	Cloudwatch             CloudwatchClient
	CloudwatchLogs         CloudwatchLogsClient
	Codebuild              CodebuildClient
	CodePipeline           CodePipelineClient
	CognitoIdentityPools   CognitoIdentityPoolsClient
	CognitoUserPools       CognitoUserPoolsClient
	ConfigService          ConfigServiceClient
	DAX                    DAXClient
	DMS                    DatabasemigrationserviceClient
	Directconnect          DirectconnectClient
	DynamoDB               DynamoDBClient
	EC2                    Ec2Client
	ECR                    EcrClient
	ECS                    EcsClient
	EFS                    EfsClient
	ELBv1                  ElbV1Client
	ELBv2                  ElbV2Client
	EMR                    EmrClient
	Eks                    EksClient
	ElasticBeanstalk       ElasticbeanstalkClient
	ElasticSearch          ElasticSearch
	FSX                    FsxClient
	GuardDuty              GuardDutyClient
	IAM                    IamClient
	IOT                    IOTClient
	KMS                    KmsClient
	Lambda                 LambdaClient
	Lightsail              LightsailClient
	MQ                     MQClient
	Organizations          OrganizationsClient
	QLDB                   QLDBClient
	RDS                    RdsClient
	Redshift               RedshiftClient
	Route53                Route53Client
	Route53Domains         Route53DomainsClient
	S3                     S3Client
	S3Control              S3ControlClient
	S3Manager              S3ManagerClient
	SES                    SESClient
	Shield                 ShieldClient
	SNS                    SnsClient
	SQS                    SQSClient
	SSM                    SSMClient
	SageMaker              SageMakerClient
	SecretsManager         SecretsManagerClient
	Waf                    WafClient
	WafV2                  WafV2Client
	WafRegional            WafRegionalClient
	Workspaces             WorkspacesClient
	Xray                   XrayClient
}

func initServices(region string, c aws.Config) Services {
	awsCfg := c.Copy()
	awsCfg.Region = region
	return Services{
		ACM:                    acm.NewFromConfig(awsCfg),
		Analyzer:               accessanalyzer.NewFromConfig(awsCfg),
		Apigateway:             apigateway.NewFromConfig(awsCfg),
		Apigatewayv2:           apigatewayv2.NewFromConfig(awsCfg),
		ApplicationAutoscaling: applicationautoscaling.NewFromConfig(awsCfg),
		Autoscaling:            autoscaling.NewFromConfig(awsCfg),
		Athena:                 athena.NewFromConfig(awsCfg),
		Backup:                 backup.NewFromConfig(awsCfg),
		Cloudfront:             cloudfront.NewFromConfig(awsCfg),
		Cloudtrail:             cloudtrail.NewFromConfig(awsCfg),
		Cloudwatch:             cloudwatch.NewFromConfig(awsCfg),
		CloudwatchLogs:         cloudwatchlogs.NewFromConfig(awsCfg),
		Cloudformation:         cloudformation.NewFromConfig(awsCfg),
		CognitoIdentityPools:   cognitoidentity.NewFromConfig(awsCfg),
		CognitoUserPools:       cognitoidentityprovider.NewFromConfig(awsCfg),
		Codebuild:              codebuild.NewFromConfig(awsCfg),
		CodePipeline:           codepipeline.NewFromConfig(awsCfg),
		ConfigService:          configservice.NewFromConfig(awsCfg),
		DAX:                    dax.NewFromConfig(awsCfg),
		Directconnect:          directconnect.NewFromConfig(awsCfg),
		DMS:                    databasemigrationservice.NewFromConfig(awsCfg),
		DynamoDB:               dynamodb.NewFromConfig(awsCfg),
		EC2:                    ec2.NewFromConfig(awsCfg),
		ECR:                    ecr.NewFromConfig(awsCfg),
		ECS:                    ecs.NewFromConfig(awsCfg),
		EFS:                    efs.NewFromConfig(awsCfg),
		Eks:                    eks.NewFromConfig(awsCfg),
		ElasticBeanstalk:       elasticbeanstalk.NewFromConfig(awsCfg),
		ElasticSearch:          elasticsearchservice.NewFromConfig(awsCfg),
		ELBv1:                  elbv1.NewFromConfig(awsCfg),
		ELBv2:                  elbv2.NewFromConfig(awsCfg),
		EMR:                    emr.NewFromConfig(awsCfg),
		FSX:                    fsx.NewFromConfig(awsCfg),
		GuardDuty:              guardduty.NewFromConfig(awsCfg),
		IAM:                    iam.NewFromConfig(awsCfg),
		KMS:                    kms.NewFromConfig(awsCfg),
		Lambda:                 lambda.NewFromConfig(awsCfg),
		Lightsail:              lightsail.NewFromConfig(awsCfg),
		MQ:                     mq.NewFromConfig(awsCfg),
		Organizations:          organizations.NewFromConfig(awsCfg),
		QLDB:                   qldb.NewFromConfig(awsCfg),
		RDS:                    rds.NewFromConfig(awsCfg),
		Redshift:               redshift.NewFromConfig(awsCfg),
		Route53:                route53.NewFromConfig(awsCfg),
		Route53Domains:         route53domains.NewFromConfig(awsCfg),
		S3:                     s3.NewFromConfig(awsCfg),
		S3Control:              s3control.NewFromConfig(awsCfg),
		S3Manager:              newS3ManagerFromConfig(awsCfg),
		SageMaker:              sagemaker.NewFromConfig(awsCfg),
		SecretsManager:         secretsmanager.NewFromConfig(awsCfg),
		SES:                    sesv2.NewFromConfig(awsCfg),
		Shield:                 shield.NewFromConfig(awsCfg),
		SNS:                    sns.NewFromConfig(awsCfg),
		SSM:                    ssm.NewFromConfig(awsCfg),
		SQS:                    sqs.NewFromConfig(awsCfg),
		Waf:                    waf.NewFromConfig(awsCfg),
		WafV2:                  wafv2.NewFromConfig(awsCfg),
		WafRegional:            wafregional.NewFromConfig(awsCfg),
		Workspaces:             workspaces.NewFromConfig(awsCfg),
		IOT:                    iot.NewFromConfig(awsCfg),
		Xray:                   xray.NewFromConfig(awsCfg),
	}
}
