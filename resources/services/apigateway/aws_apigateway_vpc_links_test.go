// +build integration

package apigateway

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationApigatewayVpcLinks(t *testing.T) {
	awsTestIntegrationHelper(t, resources.ApigatewayVpcLinks(),
		"./snapshots")
}
