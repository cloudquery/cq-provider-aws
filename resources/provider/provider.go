package provider

import (

	"github.com/cloudquery/cq-provider-aws/client"
  
  "github.com/cloudquery/cq-provider-aws/resources/services/autoscaling"
  "github.com/cloudquery/cq-provider-aws/resources/services/cloudwatch"
  "github.com/cloudquery/cq-provider-aws/resources/services/ecs"
  "github.com/cloudquery/cq-provider-aws/resources/services/redshift"
  "github.com/cloudquery/cq-provider-aws/resources/services/sqs"
  "github.com/cloudquery/cq-provider-aws/resources/services/accessanalyzer"
  "github.com/cloudquery/cq-provider-aws/resources/services/athena"
  "github.com/cloudquery/cq-provider-aws/resources/services/dynamodb"
  "github.com/cloudquery/cq-provider-aws/resources/services/ec2"
  "github.com/cloudquery/cq-provider-aws/resources/services/efs"
  "github.com/cloudquery/cq-provider-aws/resources/services/eks"
	"github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

var (

	Version = "Development"
)

func Provider() *provider.Provider {
	return &provider.Provider{
		Name:             "aws",
		Version:          Version,
		Configure:        client.Configure,
//		ErrorClassifier:  client.ErrorClassifier,
		ResourceMap: map[string]*schema.Table{
      
      "accessanalyzer.analyzer": accessanalyzer.Aws_accessanalyzer_analyzer(),
      "athena.workgroup": athena.Aws_athena_workgroup(),
      "autoscaling.launchconfiguration": autoscaling.Aws_autoscaling_launchconfiguration(),
      "cloudwatch.compositealarm": cloudwatch.Aws_cloudwatch_compositealarm(),
      "dynamodb.globaltable": dynamodb.Aws_dynamodb_globaltable(),
      "ec2.egressonlyinternetgateway": ec2.Aws_ec2_egressonlyinternetgateway(),
      "ec2.flowlog": ec2.Aws_ec2_flowlog(),
      "ec2.host": ec2.Aws_ec2_host(),
      "ec2.internetgateway": ec2.Aws_ec2_internetgateway(),
      "ec2.networkacl": ec2.Aws_ec2_networkacl(),
      "ec2.routetable": ec2.Aws_ec2_routetable(),
      "ec2.vpc": ec2.Aws_ec2_vpc(),
      "ec2.vpcpeeringconnection": ec2.Aws_ec2_vpcpeeringconnection(),
      "ecs.cluster": ecs.Aws_ecs_cluster(),
      "ecs.taskdefinition": ecs.Aws_ecs_taskdefinition(),
      "efs.filesystem": efs.Aws_efs_filesystem(),
      "eks.cluster": eks.Aws_eks_cluster(),
      "redshift.cluster": redshift.Aws_redshift_cluster(),
      "sqs.queue": sqs.Aws_sqs_queue(),
		},
		Config: func() provider.Config {
			return &client.Config{}
		},
	}
}
