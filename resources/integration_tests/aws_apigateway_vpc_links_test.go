// +build integration

package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationApigatewayVpcLinks(t *testing.T) {
	awsTestIntegrationHelper(t, resources.ApigatewayVpcLinks()
}
