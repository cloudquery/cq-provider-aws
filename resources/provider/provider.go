package provider

import (
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/resources/services/accessanalyzer"
	"github.com/cloudquery/cq-provider-aws/resources/services/acm"
	"github.com/cloudquery/cq-provider-aws/resources/services/apigateway"
	"github.com/cloudquery/cq-provider-aws/resources/services/apigatewayv2"
	"github.com/cloudquery/cq-provider-aws/resources/services/applicationautoscaling"
	"github.com/cloudquery/cq-provider-aws/resources/services/athena"
	"github.com/cloudquery/cq-provider-aws/resources/services/autoscaling"
	"github.com/cloudquery/cq-provider-aws/resources/services/backup"
	"github.com/cloudquery/cq-provider-aws/resources/services/cloudformation"
	"github.com/cloudquery/cq-provider-aws/resources/services/cloudfront"
	"github.com/cloudquery/cq-provider-aws/resources/services/cloudtrail"
	"github.com/cloudquery/cq-provider-aws/resources/services/cloudwatch"
	"github.com/cloudquery/cq-provider-aws/resources/services/cloudwatchlogs"
	"github.com/cloudquery/cq-provider-aws/resources/services/codebuild"
	"github.com/cloudquery/cq-provider-aws/resources/services/codepipeline"
	"github.com/cloudquery/cq-provider-aws/resources/services/cognito"
	"github.com/cloudquery/cq-provider-aws/resources/services/config"
	"github.com/cloudquery/cq-provider-aws/resources/services/dax"
	"github.com/cloudquery/cq-provider-aws/resources/services/directconnect"
	"github.com/cloudquery/cq-provider-aws/resources/services/dms"
	"github.com/cloudquery/cq-provider-aws/resources/services/dynamodb"
	"github.com/cloudquery/cq-provider-aws/resources/services/ec2"
	"github.com/cloudquery/cq-provider-aws/resources/services/ecr"
	"github.com/cloudquery/cq-provider-aws/resources/services/ecs"
	"github.com/cloudquery/cq-provider-aws/resources/services/efs"
	"github.com/cloudquery/cq-provider-aws/resources/services/eks"
	"github.com/cloudquery/cq-provider-aws/resources/services/elasticbeanstalk"
	"github.com/cloudquery/cq-provider-aws/resources/services/elasticsearch"
	"github.com/cloudquery/cq-provider-aws/resources/services/elbv1"
	"github.com/cloudquery/cq-provider-aws/resources/services/elbv2"
	"github.com/cloudquery/cq-provider-aws/resources/services/emr"
	"github.com/cloudquery/cq-provider-aws/resources/services/fsx"
	"github.com/cloudquery/cq-provider-aws/resources/services/guardduty"
	"github.com/cloudquery/cq-provider-aws/resources/services/iam"
	"github.com/cloudquery/cq-provider-aws/resources/services/iot"
	"github.com/cloudquery/cq-provider-aws/resources/services/kms"
	"github.com/cloudquery/cq-provider-aws/resources/services/lambda"
	"github.com/cloudquery/cq-provider-aws/resources/services/lightsail"
	"github.com/cloudquery/cq-provider-aws/resources/services/mq"
	"github.com/cloudquery/cq-provider-aws/resources/services/organizations"
	"github.com/cloudquery/cq-provider-aws/resources/services/qldb"
	"github.com/cloudquery/cq-provider-aws/resources/services/rds"
	"github.com/cloudquery/cq-provider-aws/resources/services/redshift"
	"github.com/cloudquery/cq-provider-aws/resources/services/route53"
	"github.com/cloudquery/cq-provider-aws/resources/services/s3"
	"github.com/cloudquery/cq-provider-aws/resources/services/sagemaker"
	"github.com/cloudquery/cq-provider-aws/resources/services/secretsmanager"
	"github.com/cloudquery/cq-provider-aws/resources/services/ses"
	"github.com/cloudquery/cq-provider-aws/resources/services/shield"
	"github.com/cloudquery/cq-provider-aws/resources/services/sns"
	"github.com/cloudquery/cq-provider-aws/resources/services/sqs"
	"github.com/cloudquery/cq-provider-aws/resources/services/ssm"
	"github.com/cloudquery/cq-provider-aws/resources/services/waf"
	"github.com/cloudquery/cq-provider-aws/resources/services/wafregional"
	"github.com/cloudquery/cq-provider-aws/resources/services/wafv2"
	"github.com/cloudquery/cq-provider-aws/resources/services/workspaces"
	"github.com/cloudquery/cq-provider-aws/resources/services/xray"
	"github.com/cloudquery/cq-provider-sdk/plugins"
	"github.com/cloudquery/cq-provider-sdk/schema"
)

var (
	Version = "development"
)

const ExampleConfig = `
# Optional, Repeated. Add an accounts block for every account you want to assume-role into and fetch data from.
# accounts:
# - id: <UNIQUE ACCOUNT IDENTIFIER>
# Optional. Role ARN we want to assume when accessing this account
#  role_arn: < YOUR_ROLE_ARN >
# Optional. Named profile in config or credential file from where CQ should grab credentials
#  local_profile: < PROFILE_NAME >
# Optional. by default assumes all regions
# regions:
# - us-east-1
# - us-west-2
# Optional. Enable AWS SDK debug logging.
# aws_debug: false
# The maximum number of times that a request will be retried for failures. Defaults to 10 retry attempts.
# max_retries: 10
# The maximum back off delay between attempts. The backoff delays exponentially with a jitter based on the number of attempts. Defaults to 30 seconds.
# max_backoff: 30
`

func Provider() *plugins.SourcePlugin {
	return &plugins.SourcePlugin{
		Name:      "aws",
		Version:   Version,
		Configure: client.Configure,
		// ErrorClassifier: client.ErrorClassifier,
		Tables: []*schema.Table{
			accessanalyzer.Analyzers(),
			acm.AcmCertificates(),
			apigateway.ApigatewayAPIKeys(),
			apigateway.ApigatewayClientCertificates(),
			apigateway.ApigatewayDomainNames(),
			apigateway.ApigatewayRestApis(),
			apigateway.ApigatewayUsagePlans(),
			apigateway.ApigatewayVpcLinks(),
			apigatewayv2.Apigatewayv2Apis(),
			apigatewayv2.Apigatewayv2DomainNames(),
			apigatewayv2.Apigatewayv2VpcLinks(),
			applicationautoscaling.ApplicationautoscalingPolicies(),
			athena.DataCatalogs(),
			athena.WorkGroups(),
			autoscaling.AutoscalingGroups(),
			autoscaling.AutoscalingLaunchConfigurations(),
			autoscaling.AutoscalingScheduledActions(),
			ec2.AwsRegions(),
			backup.Plans(),
			backup.Vaults(),
			backup.GlobalSettings(),
			backup.RegionSettings(),
			cloudformation.Stacks(),
			cloudfront.CloudfrontCachePolicies(),
			cloudfront.CloudfrontDistributions(),
			cloudtrail.CloudtrailTrails(),
			cloudwatch.CloudwatchAlarms(),
			cloudwatchlogs.CloudwatchlogsFilters(),
			codebuild.CodebuildProjects(),
			codepipeline.Pipelines(),
			codepipeline.Webhooks(),
			cognito.CognitoIdentityPools(),
			cognito.CognitoUserPools(),
			config.ConfigConfigurationRecorders(),
			config.ConfigConformancePack(),
			dax.DaxClusters(),
			directconnect.DirectconnectConnections(),
			directconnect.DirectconnectGateways(),
			directconnect.DirectconnectLags(),
			directconnect.DirectconnectVirtualGateways(),
			directconnect.DirectconnectVirtualInterfaces(),
			dms.DmsReplicationInstances(),
			dynamodb.DynamodbTables(),
			ec2.Ec2ByoipCidrs(),
			ec2.Ec2CustomerGateways(),
			ec2.Ec2EbsSnapshots(),
			ec2.Ec2EbsVolumes(),
			ec2.EgressOnlyInternetGateways(),
			ec2.Ec2Eips(),
			ec2.Hosts(),
			ec2.Ec2FlowLogs(),
			ec2.Ec2Images(),
			ec2.Ec2InstanceStatuses(),
			ec2.Ec2Instances(),
			ec2.Ec2InternetGateways(),
			ec2.NetworkInterfaces(),
			ec2.Ec2NatGateways(),
			ec2.Ec2NetworkAcls(),
			ec2.Ec2RegionalConfig(),
			ec2.Ec2RouteTables(),
			ec2.Ec2SecurityGroups(),
			ec2.Ec2Subnets(),
			ec2.Ec2TransitGateways(),
			ec2.Ec2VpcEndpointServiceConfigurations(),
			ec2.Ec2VpcEndpointServices(),
			ec2.Ec2VpcEndpoints(),
			ec2.Ec2VpcPeeringConnections(),
			ec2.Ec2Vpcs(),
			ec2.Ec2VpnGateways(),
			ecr.Repositories(),
			ecs.Clusters(),
			ecs.EcsTaskDefinitions(),
			efs.EfsFilesystems(),
			eks.EksClusters(),
			elasticbeanstalk.ElasticbeanstalkApplications(),
			elasticbeanstalk.ApplicationVersions(),
			elasticbeanstalk.ElasticbeanstalkEnvironments(),
			elasticsearch.ElasticsearchDomains(),
			elbv1.Elbv1LoadBalancers(),
			elbv2.Elbv2LoadBalancers(),
			elbv2.Elbv2TargetGroups(),
			emr.EmrBlockPublicAccessConfigs(),
			emr.EmrClusters(),
			fsx.FsxBackups(),
			guardduty.GuarddutyDetectors(),
			iam.IamAccounts(),
			iam.IamGroups(),
			iam.IamOpenidConnectIdentityProviders(),
			iam.IamPasswordPolicies(),
			iam.IamPolicies(),
			iam.IamRoles(),
			iam.IamSamlIdentityProviders(),
			iam.IamServerCertificates(),
			iam.IamUsers(),
			iam.IamVirtualMfaDevices(),
			iot.IotBillingGroups(),
			iot.IotCaCertificates(),
			iot.IotCertificates(),
			iot.IotPolicies(),
			iot.IotStreams(),
			iot.IotThingGroups(),
			iot.IotThingTypes(),
			iot.IotThings(),
			iot.IotTopicRules(),
			kms.Keys(),
			lambda.Functions(),
			lambda.LambdaLayers(),
			lambda.LambdaRuntimes(),
			lightsail.Alarms(),
			lightsail.Buckets(),
			lightsail.Certificates(),
			lightsail.DatabaseSnapshots(),
			lightsail.Databases(),
			lightsail.Disks(),
			lightsail.Instances(),
			lightsail.InstanceSnapshots(),
			lightsail.LoadBalancers(),
			lightsail.StaticIps(),
			mq.Brokers(),
			organizations.Accounts(),
			qldb.Ledgers(),
			rds.RdsCertificates(),
			rds.RdsClusterParameterGroups(),
			rds.RdsClusterSnapshots(),
			rds.RdsClusters(),
			rds.RdsDbParameterGroups(),
			rds.RdsDbSecurityGroups(),
			rds.RdsDbSnapshots(),
			rds.RdsSubnetGroups(),
			rds.RdsEventSubscriptions(),
			rds.RdsInstances(),
			redshift.EventSubscriptions(),
			redshift.RedshiftClusters(),
			redshift.RedshiftSubnetGroups(),
			route53.Route53Domains(),
			route53.Route53HealthChecks(),
			route53.Route53HostedZones(),
			route53.Route53ReusableDelegationSets(),
			route53.Route53TrafficPolicies(),
			s3.Accounts(),
			s3.Buckets(),
			sagemaker.SagemakerEndpointConfigurations(),
			sagemaker.SagemakerModels(),
			sagemaker.SagemakerNotebookInstances(),
			sagemaker.SagemakerTrainingJobs(),
			secretsmanager.SecretsmanagerSecrets(),
			ses.Templates(),
			shield.Attacks(),
			shield.Subscriptions(),
			shield.ProtectionGroups(),
			shield.Protections(),
			sns.SnsSubscriptions(),
			sns.SnsTopics(),
			sqs.SQSQueues(),
			ssm.SsmDocuments(),
			ssm.SsmInstances(),
			waf.WafRuleGroups(),
			waf.WafRules(),
			waf.WafSubscribedRuleGroups(),
			waf.WafWebAcls(),
			wafv2.Ipsets(),
			wafv2.Wafv2ManagedRuleGroups(),
			wafv2.RegexPatternSets(),
			wafv2.Wafv2RuleGroups(),
			wafv2.Wafv2WebAcls(),
			wafregional.RateBasedRules(),
			wafregional.RuleGroups(),
			wafregional.Rules(),
			wafregional.WebAcls(),
			workspaces.Workspaces(),
			workspaces.Directories(),
			xray.EncryptionConfigs(),
			xray.Groups(),
			xray.SamplingRules(),
			//				 iot.IotSecurityProfiles(), //TODO disabled because of api error NotFoundException: No method found matching route security-profiles for http method GET.
		},
		ExampleConfig: ExampleConfig,
		Config: func() interface{} {
			return &client.Config{}
		},
	}
}
