// +build integration

package sagemaker

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationSageMakerModels(t *testing.T) {
	awsTestIntegrationHelper(t, resources.SagemakerModels(),
		"./snapshots")
}
