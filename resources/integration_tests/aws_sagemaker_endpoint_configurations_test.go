package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationSageMakerEndpointConfigurations(t *testing.T) {
	awsTestIntegrationHelper(t, resources.SagemakerEndpointConfigurations(),
		"./snapshots/sagemaker")
}
