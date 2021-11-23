package resources

import (
	"embed"

	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

var (
	//go:embed migrations/*.sql
	awsMigrations embed.FS
)

func Provider() *provider.Provider {
	return &provider.Provider{
		Name:            "aws",
		Configure:       client.Configure,
		ErrorClassifier: client.ErrorClassifier,
		Migrations:      awsMigrations,
		ResourceMap: map[string]*schema.Table{
			"accessanalyzer.analyzers":              AccessAnalyzerAnalyzer(),
			"apigateway.api_keys":                   ApigatewayAPIKeys(),
			"apigateway.client_certificates":        ApigatewayClientCertificates(),
			"apigateway.domain_names":               ApigatewayDomainNames(),
			"apigateway.rest_apis":                  ApigatewayRestApis(),
			"apigateway.usage_plans":                ApigatewayUsagePlans(),
			"apigateway.vpc_links":                  ApigatewayVpcLinks(),
			"apigatewayv2.apis":                     Apigatewayv2Apis(),
			"apigatewayv2.domain_names":             Apigatewayv2DomainNames(),
			"apigatewayv2.vpc_links":                Apigatewayv2VpcLinks(),
			"autoscaling.launch_configurations":     AutoscalingLaunchConfigurations(),
			"cloudfront.cache_policies":             CloudfrontCachePolicies(),
			"cloudfront.distributions":              CloudfrontDistributions(),
			"cloudtrail.trails":                     CloudtrailTrails(),
			"cloudwatch.alarms":                     CloudwatchAlarms(),
			"cloudwatchlogs.filters":                CloudwatchlogsFilters(),
			"cognito.identity_pools":                CognitoIdentityPools(),
			"cognito.user_pools":                    CognitoUserPools(),
			"config.configuration_recorders":        ConfigConfigurationRecorders(),
			"config.conformance_packs":              ConfigConformancePack(),
			"directconnect.connections":             DirectconnectConnections(),
			"directconnect.gateways":                DirectconnectGateways(),
			"directconnect.lags":                    DirectconnectLags(),
			"directconnect.virtual_gateways":        DirectconnectVirtualGateways(),
			"directconnect.virtual_interfaces":      DirectconnectVirtualInterfaces(),
			"ec2.byoip_cidrs":                       Ec2ByoipCidrs(),
			"ec2.customer_gateways":                 Ec2CustomerGateways(),
			"ec2.ebs_volumes":                       Ec2EbsVolumes(),
			"ec2.flow_logs":                         Ec2FlowLogs(),
			"ec2.images":                            Ec2Images(),
			"ec2.instances":                         Ec2Instances(),
			"ec2.internet_gateways":                 Ec2InternetGateways(),
			"ec2.nat_gateways":                      Ec2NatGateways(),
			"ec2.network_acls":                      Ec2NetworkAcls(),
			"ec2.regional_config":                   Ec2RegionalConfig(),
			"ec2.route_tables":                      Ec2RouteTables(),
			"ec2.security_groups":                   Ec2SecurityGroups(),
			"ec2.subnets":                           Ec2Subnets(),
			"ec2.transit_gateways":                  Ec2TransitGateways(),
			"ec2.vpc_endpoints":                     Ec2VpcEndpoints(),
			"ec2.vpc_peering_connections":           Ec2VpcPeeringConnections(),
			"ec2.vpcs":                              Ec2Vpcs(),
			"ec2.vpn_gateways":                      Ec2VpnGateways(),
			"ecr.repositories":                      EcrRepositories(),
			"ecs.clusters":                          EcsClusters(),
			"efs.filesystems":                       EfsFilesystems(),
			"eks.clusters":                          EksClusters(),
			"elasticbeanstalk.environments":         ElasticbeanstalkEnvironments(),
			"elasticsearch.domains":                 ElasticsearchDomains(),
			"elbv1.load_balancers":                  Elbv1LoadBalancers(),
			"elbv2.load_balancers":                  Elbv2LoadBalancers(),
			"elbv2.target_groups":                   Elbv2TargetGroups(),
			"emr.clusters":                          EmrClusters(),
			"emr.block_public_access_configs":       EmrBlockPublicAccessConfigs(),
			"fsx.backups":                           FsxBackups(),
			"iam.accounts":                          IamAccounts(),
			"iam.groups":                            IamGroups(),
			"iam.openid_connect_identity_providers": IamOpenidConnectIdentityProviders(),
			"iam.password_policies":                 IamPasswordPolicies(),
			"iam.policies":                          IamPolicies(),
			"iam.roles":                             IamRoles(),
			"iam.saml_identity_providers":           IamSamlIdentityProviders(),
			"iam.server_certificates":               IamServerCertificates(),
			"iam.users":                             IamUsers(),
			"iam.virtual_mfa_devices":               IamVirtualMfaDevices(),
			"kms.keys":                              KmsKeys(),
			"lambda.functions":                      LambdaFunctions(),
			"lambda.layers":                         LambdaLayers(),
			"mq.brokers":                            MqBrokers(),
			"organizations.accounts":                OrganizationsAccounts(),
			"rds.certificates":                      RdsCertificates(),
			"rds.clusters":                          RdsClusters(),
			"rds.db_subnet_groups":                  RdsSubnetGroups(),
			"rds.instances":                         RdsInstances(),
			"redshift.clusters":                     RedshiftClusters(),
			"redshift.subnet_groups":                RedshiftSubnetGroups(),
			"route53.domains":                       Route53Domains(),
			"route53.health_checks":                 Route53HealthChecks(),
			"route53.hosted_zones":                  Route53HostedZones(),
			"route53.reusable_delegation_sets":      Route53ReusableDelegationSets(),
			"route53.traffic_policies":              Route53TrafficPolicies(),
			"s3.buckets":                            S3Buckets(),
			"sagemaker.notebook_instances":          SagemakerNotebookInstances(),
			"sagemaker.models":                      SagemakerModels(),
			"sagemaker.endpoint_configurations":     SagemakerEndpointConfigurations(),
			"sagemaker.training_jobs":               SagemakerTrainingJobs(),
			"sns.subscriptions":                     SnsSubscriptions(),
			"sns.topics":                            SnsTopics(),
			"sqs.queues":                            SQSQueues(),
			"waf.rule_groups":                       WafRuleGroups(),
			"waf.rules":                             WafRules(),
			"waf.subscribed_rule_groups":            WafSubscribedRuleGroups(),
			"waf.web_acls":                          WafWebAcls(),
			"wafv2.managed_rule_groups":             Wafv2ManagedRuleGroups(),
			"wafv2.rule_groups":                     Wafv2RuleGroups(),
			"wafv2.web_acls":                        Wafv2WebAcls(),
		},
		Config: func() provider.Config {
			return &client.Config{}
		},
	}
}
