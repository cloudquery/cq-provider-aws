// +build integration

package sqs

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationSQSQueues(t *testing.T) {
	awsTestIntegrationHelper(t, resources.SQSQueues(),
		"./snapshots")
}
