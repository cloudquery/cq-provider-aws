// +build integration

package sagemaker

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationSageMakerNotebookInstances(t *testing.T) {
	awsTestIntegrationHelper(t, resources.SagemakerNotebookInstances(),
		"./snapshots")
}
