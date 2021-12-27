// +build integration

package apigateway

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationApigatewayRestApis(t *testing.T) {
	awsTestIntegrationHelper(t, resources.ApigatewayRestApis(), "./snapshots")
}
