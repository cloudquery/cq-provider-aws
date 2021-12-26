// +build integration

package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationEc2Vpcs(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Ec2Vpcs(),
		"./snapshots/ec2")
}

func TestIntegrationEc2VpcPeeringConnections(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Ec2VpcPeeringConnections(),
		"./snapshots/ec2")
}

func TestIntegrationEc2VpcEndpoints(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Ec2VpcEndpoints(),
		"./snapshots/ec2")
}

func TestIntegrationEc2TransitGateways(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Ec2TransitGateways(),
		"./snapshots/ec2")
}

func TestIntegrationEc2Subnets(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Ec2Subnets(),
		"./snapshots/ec2")
}

func TestIntegrationEc2SecurityGroups(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Ec2SecurityGroups(),
		"./snapshots/ec2")
}

func TestIntegrationEc2RouteTables(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Ec2RouteTables(),
		"./snapshots/ec2")
}

func TestIntegrationEc2NetworkAcls(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Ec2NetworkAcls(),
		"./snapshots/ec2")
}

func TestIntegrationEc2NatGateways(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Ec2NatGateways(),
		"./snapshots/ec2")
}

func TestIntegrationEc2InternetGateways(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Ec2InternetGateways(),
		"./snapshots/ec2")
}
