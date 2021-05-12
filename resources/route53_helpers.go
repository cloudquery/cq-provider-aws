package resources

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/route53/types"
)

const (
	resourceTypeAssertError = "failed to assert resource type"
)

type Route53HealthCheckWrapper struct {
	types.HealthCheck
	Tags map[string]interface{}
}

type Route53HostedZoneWrapper struct {
	types.HostedZone
	Tags map[string]interface{}
}

func parseRoute53HostedZoneId(id string) string {
	return strings.Replace(id, fmt.Sprintf("/%s/", types.TagResourceTypeHostedzone), "", 1)
}

func getRoute53tagsByResourceID(id string, set []types.ResourceTagSet) []types.Tag {
	for _, s := range set {
		if *s.ResourceId == id {
			return s.Tags
		}
	}
	return nil
}
