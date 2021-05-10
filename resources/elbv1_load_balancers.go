package resources

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	elbv1 "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func Elbv1LoadBalancers() *schema.Table {
	return &schema.Table{
		Name:         "aws_elbv1_load_balancers",
		Resolver:     fetchElbv1LoadBalancers,
		Multiplex:    client.AccountRegionMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name: "availability_zones",
				Type: schema.TypeStringArray,
			},
			{
				Name: "canonical_hosted_zone_name",
				Type: schema.TypeString,
			},
			{
				Name:     "canonical_hosted_zone_name_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CanonicalHostedZoneNameID"),
			},
			{
				Name: "created_time",
				Type: schema.TypeTimestamp,
			},
			{
				Name:     "dns_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DNSName"),
			},
			{
				Name:     "health_check_healthy_threshold",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("HealthCheck.HealthyThreshold"),
			},
			{
				Name:     "health_check_interval",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("HealthCheck.Interval"),
			},
			{
				Name:     "health_check_target",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("HealthCheck.Target"),
			},
			{
				Name:     "health_check_timeout",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("HealthCheck.Timeout"),
			},
			{
				Name:     "health_check_unhealthy_threshold",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("HealthCheck.UnhealthyThreshold"),
			},
			{
				Name: "load_balancer_name",
				Type: schema.TypeString,
			},
			{
				Name:     "policies_other_policies",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Policies.OtherPolicies"),
			},
			{
				Name: "scheme",
				Type: schema.TypeString,
			},
			{
				Name: "security_groups",
				Type: schema.TypeStringArray,
			},
			{
				Name:     "source_security_group_group_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SourceSecurityGroup.GroupName"),
			},
			{
				Name:     "source_security_group_owner_alias",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SourceSecurityGroup.OwnerAlias"),
			},
			{
				Name: "subnets",
				Type: schema.TypeStringArray,
			},
			{
				Name:     "vpc_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VPCId"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "aws_elbv1_load_balancer_backend_server_descriptions",
				Resolver: fetchElbv1LoadBalancerBackendServerDescriptions,
				Columns: []schema.Column{
					{
						Name:     "load_balancer_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "instance_port",
						Type: schema.TypeInt,
					},
					{
						Name: "policy_names",
						Type: schema.TypeStringArray,
					},
				},
			},
			{
				Name:     "aws_elbv1_load_balancer_instances",
				Resolver: fetchElbv1LoadBalancerInstances,
				Columns: []schema.Column{
					{
						Name:     "load_balancer_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "instance_id",
						Type: schema.TypeString,
					},
				},
			},
			{
				Name:     "aws_elbv1_load_balancer_listener_descriptions",
				Resolver: fetchElbv1LoadBalancerListenerDescriptions,
				Columns: []schema.Column{
					{
						Name:     "load_balancer_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name:     "listener_instance_port",
						Type:     schema.TypeInt,
						Resolver: schema.PathResolver("Listener.InstancePort"),
					},
					{
						Name:     "listener_load_balancer_port",
						Type:     schema.TypeInt,
						Resolver: schema.PathResolver("Listener.LoadBalancerPort"),
					},
					{
						Name:     "listener_protocol",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Listener.Protocol"),
					},
					{
						Name:     "listener_instance_protocol",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Listener.InstanceProtocol"),
					},
					{
						Name:     "listener_s_s_l_certificate_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Listener.SSLCertificateId"),
					},
					{
						Name: "policy_names",
						Type: schema.TypeStringArray,
					},
				},
			},
			{
				Name:     "aws_elbv1_load_balancer_policies_app_cookie_stickiness_policies",
				Resolver: fetchElbv1LoadBalancerPoliciesAppCookieStickinessPolicies,
				Columns: []schema.Column{
					{
						Name:     "load_balancer_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "cookie_name",
						Type: schema.TypeString,
					},
					{
						Name: "policy_name",
						Type: schema.TypeString,
					},
				},
			},
			{
				Name:     "aws_elbv1_load_balancer_policies_l_b_cookie_stickiness_policies",
				Resolver: fetchElbv1LoadBalancerPoliciesLBCookieStickinessPolicies,
				Columns: []schema.Column{
					{
						Name:     "load_balancer_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "cookie_expiration_period",
						Type: schema.TypeBigInt,
					},
					{
						Name: "policy_name",
						Type: schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchElbv1LoadBalancers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	var config elbv1.DescribeLoadBalancersInput
	c := meta.(*client.Client)
	svc := c.Services().ELBv1
	for {
		response, err := svc.DescribeLoadBalancers(ctx, &config, func(options *elbv1.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		res <- response.LoadBalancerDescriptions
		if aws.ToString(response.NextMarker) == "" {
			break
		}
		config.Marker = response.NextMarker
	}
	return nil
}

func fetchElbv1LoadBalancerBackendServerDescriptions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func fetchElbv1LoadBalancerInstances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func fetchElbv1LoadBalancerListenerDescriptions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func fetchElbv1LoadBalancerPoliciesAppCookieStickinessPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
func fetchElbv1LoadBalancerPoliciesLBCookieStickinessPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	panic("not implemented")
}
