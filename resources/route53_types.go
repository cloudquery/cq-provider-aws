package resources

import "github.com/aws/aws-sdk-go-v2/service/route53/types"

type Route53HealthCheckWrapper struct {
	types.HealthCheck
	Tags map[string]interface{}
}

type Route53HostedZoneWrapper struct {
	types.HostedZone
	Tags            map[string]interface{}
	DelegationSetId *string
	VPCs            []types.VPC
}
