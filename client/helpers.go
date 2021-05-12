package client

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/route53/types"

	"github.com/aws/smithy-go"
)

//log-group:([a-zA-Z0-9/]+):
var GroupNameRegex = regexp.MustCompile("arn:aws:logs:[a-z0-9-]+:[0-9]+:log-group:([a-zA-Z0-9-/]+):")

func IgnoreAccessDeniedServiceDisabled(err error) bool {
	var ae smithy.APIError
	if errors.As(err, &ae) {
		switch ae.ErrorCode() {
		case "AccessDenied", "AccessDeniedException", "UnauthorizedOperation":
			return true
		case "OptInRequired", "SubscriptionRequiredException", "InvalidClientTokenId":
			return true
		}
	}
	return false
}

func ParseRoute53HostedZoneId(id string) string {
	return strings.Replace(id, fmt.Sprintf("/%s/", types.TagResourceTypeHostedzone), "", 1)
}
