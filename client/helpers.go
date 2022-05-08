package client

import (
	"context"
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/smithy-go"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

const (
	PartitionServiceRegionFile = "data/partition_service_region.json"
	defaultPartition           = "aws"
)

var (
	//go:embed data/partition_service_region.json
	supportedServiceRegionFile embed.FS
	readOnce                   sync.Once
	supportedServiceRegion     *SupportedServiceRegionsData
)

// GroupNameRegex log-group:([a-zA-Z0-9/]+):
var GroupNameRegex = regexp.MustCompile("arn:aws:logs:[a-z0-9-]+:[0-9]+:log-group:([a-zA-Z0-9-/]+):")

type AwsService struct {
	Regions map[string]*map[string]interface{} `json:"regions"`
}

type AwsPartition struct {
	Id       string                 `json:"partition"`
	Name     string                 `json:"partitionName"`
	Services map[string]*AwsService `json:"services"`
}

type SupportedServiceRegionsData struct {
	Partitions map[string]AwsPartition `json:"partitions"`
}

func readSupportedServiceRegions() *SupportedServiceRegionsData {
	f, err := supportedServiceRegionFile.Open(PartitionServiceRegionFile)
	if err != nil {
		return nil
	}
	stat, err := f.Stat()
	if err != nil {
		return nil
	}
	data := make([]byte, stat.Size())
	if _, err := f.Read(data); err != nil {
		return nil
	}
	var result *SupportedServiceRegionsData
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil
	}
	return result
}

func isSupportedServiceForRegion(service string, region string) bool {
	readOnce.Do(func() {
		supportedServiceRegion = readSupportedServiceRegions()
	})

	if supportedServiceRegion == nil {
		return false
	}

	if supportedServiceRegion.Partitions == nil {
		return false
	}

	currentPartition := supportedServiceRegion.Partitions[defaultPartition]

	if currentPartition.Services[service] == nil {
		return false
	}

	if currentPartition.Services[service].Regions[region] == nil {
		return false
	}

	return true
}

func getAvailableRegions() (map[string]bool, error) {
	readOnce.Do(func() {
		supportedServiceRegion = readSupportedServiceRegions()
	})

	regionsSet := make(map[string]bool)

	if supportedServiceRegion == nil {
		return nil, fmt.Errorf("could not get AWS regions/services data")
	}

	if supportedServiceRegion.Partitions == nil {
		return nil, fmt.Errorf("could not found any AWS partitions")
	}

	currentPartition := supportedServiceRegion.Partitions[defaultPartition]

	for _, service := range currentPartition.Services {
		for region := range service.Regions {
			regionsSet[region] = true
		}
	}

	return regionsSet, nil
}

func IgnoreAccessDeniedServiceDisabled(err error) bool {
	var ae smithy.APIError
	if errors.As(err, &ae) {
		switch ae.ErrorCode() {
		case "UnrecognizedClientException":
			return strings.Contains(ae.Error(), "The security token included in the request is invalid")
		case "AWSOrganizationsNotInUseException":
			return true
		case "AuthorizationError", "AccessDenied", "AccessDeniedException", "InsufficientPrivilegesException", "UnauthorizedOperation":
			return true
		case "OptInRequired", "SubscriptionRequiredException", "InvalidClientTokenId":
			return true
		}
	}
	return false
}

func IgnoreWithInvalidAction(err error) bool {
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

func accountObfusactor(aa []Account, msg string) string {
	for _, a := range aa {
		msg = strings.ReplaceAll(msg, a.ID, obfuscateAccountId(a.ID))
	}
	return msg
}

type AWSService string

const (
	ApigatewayService           AWSService = "apigateway"
	CloudfrontService           AWSService = "cloudfront"
	CognitoIdentityService      AWSService = "cognito-identity"
	DirectConnectService        AWSService = "directconnect"
	EC2Service                  AWSService = "ec2"
	ElasticLoadBalancingService AWSService = "elasticloadbalancing"
	GuardDutyService            AWSService = "guardduty"
	RedshiftService             AWSService = "redshift"
	Route53Service              AWSService = "route53"
	S3Service                   AWSService = "s3"
	WAFRegional                 AWSService = "waf-regional"
	WorkspacesService           AWSService = "workspaces"
)

// MakeARN creates an ARN using supplied service name, account id, region name and resource id parts.
// Resource id parts are concatenated using forward slash (/).
func MakeARN(service AWSService, accountID, region string, idParts ...string) string {
	return arn.ARN{
		Partition: "aws",
		Service:   string(service),
		Region:    region,
		AccountID: accountID,
		Resource:  strings.Join(idParts, "/"),
	}.String()
}

func resolveARN(service AWSService, resourceID func(resource *schema.Resource) ([]string, error), useRegion, useAccountID bool) schema.ColumnResolver {
	return func(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
		cl := meta.(*Client)
		idParts, err := resourceID(resource)
		if err != nil {
			return fmt.Errorf("error resolving resource id: %w", err)
		}
		var accountID, region string
		if useAccountID {
			accountID = cl.AccountID
		}
		if useRegion {
			region = cl.Region
		}
		return resource.Set(c.Name, MakeARN(service, accountID, region, idParts...))
	}
}

// ResolveARNWithAccount returns a column resolver that will set a field value to a proper ARN
// based on provided AWS service and resource id value returned by resourceID function.
// Region is left empty and account id is set to the value of the client.
func ResolveARNWithAccount(service AWSService, resourceID func(resource *schema.Resource) ([]string, error)) schema.ColumnResolver {
	return resolveARN(service, resourceID, false, true)
}

// ResolveARNWithRegion returns a column resolver that will set a field value to a proper ARN
// based on provided AWS service and resource id value returned by resourceID function.
// Region is set to the value of the client and account id is left empty.
func ResolveARNWithRegion(service AWSService, resourceID func(resource *schema.Resource) ([]string, error)) schema.ColumnResolver {
	return resolveARN(service, resourceID, true, false)
}

// ResolveARN returns a column resolver that will set a field value to a proper ARN
// based on provided AWS service and resource id value returned by resourceID function.
// Region and account id are set to the values of the client.
func ResolveARN(service AWSService, resourceID func(resource *schema.Resource) ([]string, error)) schema.ColumnResolver {
	return resolveARN(service, resourceID, true, true)
}

// ResolveARNGlobal returns a column resolver that will set a field value to a proper ARN
// based on provided AWS service and resource id value returned by resourceID function.
// Region  and account id are left empty.
func ResolveARNGlobal(service AWSService, resourceID func(resource *schema.Resource) ([]string, error)) schema.ColumnResolver {
	return resolveARN(service, resourceID, false, false)
}

var notFoundErrorPrefixes = []string{
	"ResourceNotFoundException",
	"WAFNonexistentItemException",
	"NoSuch",
	"NotFound",
	"NotFoundError",
}

// IsNotFoundError checks if api error should be ignored
func (c *Client) IsNotFoundError(err error) bool {
	var ae smithy.APIError
	if !errors.As(err, &ae) {
		return false
	}
	errorCode := ae.ErrorCode()
	for _, s := range notFoundErrorPrefixes {
		if strings.Contains(errorCode, s) {
			c.logger.Warn("API returned \"NotFound\" error ignoring it...", "error", err)
			return true
		}
	}
	return false
}

func IsAWSError(err error, code ...string) bool {
	var ae smithy.APIError
	if !errors.As(err, &ae) {
		return false
	}
	for _, c := range code {
		if strings.Contains(ae.ErrorCode(), c) {
			return true
		}
	}
	return false
}

// TagsIntoMap expects []T (usually "[]Tag") where T has "Key" and "Value" fields (of type string or *string) and writes them into the given map
func TagsIntoMap(tagSlice interface{}, dst map[string]string) {
	stringify := func(v reflect.Value) string {
		vt := v.Type()
		if vt.Kind() == reflect.String {
			return v.String()
		}
		if vt.Kind() != reflect.Ptr || vt.Elem().Kind() != reflect.String {
			panic("field is not string or *string")
		}

		return v.Elem().String()
	}

	if k := reflect.TypeOf(tagSlice).Kind(); k != reflect.Slice {
		panic("invalid usage: Only slices are supported as input: " + k.String())
	}
	slc := reflect.ValueOf(tagSlice)

	for i := 0; i < slc.Len(); i++ {
		val := slc.Index(i)
		if k := val.Kind(); k != reflect.Struct {
			panic("slice member is not struct: " + k.String())
		}

		keyField, valField := val.FieldByName("Key"), val.FieldByName("Value")
		if (keyField.Type().Kind() == reflect.Ptr && keyField.IsNil()) || (valField.Type().Kind() == reflect.Ptr && valField.IsNil()) {
			continue
		}

		if keyField.IsZero() || valField.IsZero() {
			panic("slice member is missing Key or Value fields")
		}

		dst[stringify(keyField)] = stringify(valField)
	}
}

// TagsToMap expects []T (usually "[]Tag") where T has "Key" and "Value" fields (of type string or *string) and returns a map
func TagsToMap(tagSlice interface{}) map[string]string {
	if k := reflect.TypeOf(tagSlice).Kind(); k != reflect.Slice {
		panic("invalid usage: Only slices are supported as input: " + k.String())
	}
	slc := reflect.ValueOf(tagSlice)

	ret := make(map[string]string, slc.Len())
	TagsIntoMap(tagSlice, ret)
	return ret
}
