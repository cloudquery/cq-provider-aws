// +build integration

package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationEc2Vpcs(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Ec2Vpcs())
}

func TestIntegrationEc2VpcPeeringConnections(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Ec2VpcPeeringConnections())
}

func TestIntegrationEc2VpcEndpoints(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Ec2VpcEndpoints())
}

func TestIntegrationEc2TransitGateways(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Ec2TransitGateways())
}

func TestIntegrationEc2Subnets(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Ec2Subnets())
}

func TestIntegrationEc2SecurityGroups(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Ec2SecurityGroups())
}

func TestIntegrationEc2RouteTables(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Ec2RouteTables())
}

func TestIntegrationEc2NetworkAcls(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Ec2NetworkAcls())
}

func TestIntegrationEc2NatGateways(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Ec2NatGateways())
}

func TestIntegrationEc2InternetGateways(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Ec2InternetGateways())
}
