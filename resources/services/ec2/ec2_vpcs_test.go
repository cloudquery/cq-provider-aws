// +build integration

package ec2

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationEc2Vpcs(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Ec2Vpcs(),
		"./snapshots")
}

func TestIntegrationEc2VpcPeeringConnections(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Ec2VpcPeeringConnections(),
		"./snapshots")
}

func TestIntegrationEc2VpcEndpoints(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Ec2VpcEndpoints(),
		"./snapshots")
}

func TestIntegrationEc2TransitGateways(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Ec2TransitGateways(),
		"./snapshots")
}

func TestIntegrationEc2Subnets(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Ec2Subnets(),
		"./snapshots")
}

func TestIntegrationEc2SecurityGroups(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Ec2SecurityGroups(),
		"./snapshots")
}

func TestIntegrationEc2RouteTables(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Ec2RouteTables(),
		"./snapshots")
}

func TestIntegrationEc2NetworkAcls(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Ec2NetworkAcls(),
		"./snapshots")
}

func TestIntegrationEc2NatGateways(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Ec2NatGateways(),
		"./snapshots")
}

func TestIntegrationEc2InternetGateways(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Ec2InternetGateways(),
		"./snapshots")
}
