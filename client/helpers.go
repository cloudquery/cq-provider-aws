package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/smithy-go"
)

//log-group:([a-zA-Z0-9/]+):
var GroupNameRegex = regexp.MustCompile("arn:aws:logs:[a-z0-9-]+:[0-9]+:log-group:([a-zA-Z0-9-/]+):")

type SupportedServicesData struct {
	Prices []struct {
		Attributes struct {
			Region  string `json:"aws:region"`
			Service string `json:"aws:serviceName"`
		} `json:"attributes"`
		Id string `json:"id"`
	} `json:"prices"`
}

var supportedServices *SupportedServicesData

func downloadSupportedResourcesForRegions() (*SupportedServicesData, error) {
	// Get the data
	req, err := http.NewRequest(http.MethodGet, "https://api.regional-table.region-services.aws.a2z.com/index.json", nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get aws supported resources for region, status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data SupportedServicesData
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	for i, item := range data.Prices {
		// replacing the name because name from api does not always fit
		parts := strings.Split(item.Id, ":")
		data.Prices[i].Attributes.Service = parts[0]
	}

	return &data, nil
}

var once sync.Once

func checkUnsupportedResourceForRegionError(err error) bool {
	once.Do(func() {
		supportedServices, _ = downloadSupportedResourcesForRegions()
	})
	errText := err.Error()

	if supportedServices != nil && strings.Contains(errText, "no such host") {
		var ae *smithy.OperationError
		if errors.As(err, &ae) {
			for _, p := range supportedServices.Prices {
				pattern := fmt.Sprintf("lookup %s.+%s\\.amazonaws\\.com", p.Attributes.Service, p.Attributes.Region)
				// match means that error is related to a supported service
				if match, _ := regexp.MatchString(pattern, errText); match {
					return false
				}
			}
		}
	}
	return true
}

func IgnoreAccessDeniedServiceDisabled(err error) bool {
	if checkUnsupportedResourceForRegionError(err) {
		return true
	}
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

func IgnoreWithInvalidAction(err error) bool {
	if IgnoreAccessDeniedServiceDisabled(err) {
		return true
	}
	var ae smithy.APIError
	if errors.As(err, &ae) {
		if ae.ErrorCode() == "InvalidAction" {
			return true
		}
	}
	return false
}

// GenerateResourceARN generates the arn for a resource.
// Service: The service name e.g. waf or elb or s3
// ResourceType: The sub resource type e.g. rule or instance (for an ec2 instance)
// ResourceID: The resource id e.g. i-1234567890abcdefg
// Region: The resource region e.g. us-east-1
// AccountID: The account id e.g. 123456789012
// See https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html for
// more information.
func GenerateResourceARN(service, resourceType, resourceID, region, accountID string) string {

	// if resource type is empty
	// for example s3 bucket
	resource := ""
	if resourceType == "" {
		resource = resourceID
	} else {
		resource = fmt.Sprintf("%s/%s", resourceType, resourceID)
	}

	return arn.ARN{
		// TODO: Make this configurable in the future
		Partition: "aws",
		Service:   service,
		Region:    region,
		AccountID: accountID,
		Resource:  resource,
	}.String()
}
