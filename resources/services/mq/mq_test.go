// +build integration

package mq

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationMQBrokers(t *testing.T) {
	awsTestIntegrationHelper(t, resources.MqBrokers(),
		"./snapshots")
}
