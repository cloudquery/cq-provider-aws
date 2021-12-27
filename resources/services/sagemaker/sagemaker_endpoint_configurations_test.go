// +build integration

package sagemaker

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationSageMakerEndpointConfigurations(t *testing.T) {
	awsTestIntegrationHelper(t, resources.SagemakerEndpointConfigurations(),
		"./snapshots")
}
