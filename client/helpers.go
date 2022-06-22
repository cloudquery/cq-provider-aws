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
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"golang.org/x/sync/semaphore"
)

type AWSService string

type AwsService struct {
	Regions map[string]*map[string]interface{} `json:"regions"`
}

type AwsPartition struct {
	Id       string                 `json:"partition"`
	Name     string                 `json:"partitionName"`
	Services map[string]*AwsService `json:"services"`
}

type SupportedServiceRegionsData struct {
	Partitions        map[string]AwsPartition `json:"partitions"`
	regionVsPartition map[string]string
}

type ListResolverFunc func(ctx context.Context, meta schema.ClientMeta, detailChan chan<- interface{}) error

type DetailResolverFunc func(ctx context.Context, meta schema.ClientMeta, resultsChan chan<- interface{}, errorChan chan<- error, summary interface{})

const (
	ApigatewayService           AWSService = "apigateway"
	Athena                      AWSService = "athena"
	CloudfrontService           AWSService = "cloudfront"
	CognitoIdentityService      AWSService = "cognito-identity"
	DirectConnectService        AWSService = "directconnect"
	DynamoDBService             AWSService = "dynamodb"
	EC2Service                  AWSService = "ec2"
	EFSService                  AWSService = "elasticfilesystem"
	ElasticLoadBalancingService AWSService = "elasticloadbalancing"
	GuardDutyService            AWSService = "guardduty"
	RedshiftService             AWSService = "redshift"
	Route53Service              AWSService = "route53"
	S3Service                   AWSService = "s3"
	WAFRegional                 AWSService = "waf-regional"
	WorkspacesService           AWSService = "workspaces"
)

const MAX_GOROUTINES = 10

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

var notFoundErrorSubstrings = []string{
	"NoSuch",
	"NotFound",
	"ResourceNotFoundException",
	"WAFNonexistentItemException",
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

	var result SupportedServiceRegionsData
	if err := json.Unmarshal(data, &result); err != nil {
		return nil
	}

	result.regionVsPartition = make(map[string]string)
	for _, p := range result.Partitions {
		for _, svc := range p.Services {
			for reg := range svc.Regions {
				result.regionVsPartition[reg] = p.Id
			}
		}
	}

	return &result
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

	prt, _ := RegionsPartition(region)
	currentPartition := supportedServiceRegion.Partitions[prt]

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

	for _, prt := range supportedServiceRegion.Partitions {
		for _, service := range prt.Services {
			for region := range service.Regions {
				regionsSet[region] = true
			}
		}
	}

	return regionsSet, nil
}

func RegionsPartition(region string) (string, bool) {
	readOnce.Do(func() {
		supportedServiceRegion = readSupportedServiceRegions()
	})

	prt, ok := supportedServiceRegion.regionVsPartition[region]
	if !ok {
		return defaultPartition, false
	}
	return prt, true
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

func IgnoreCommonErrors(err error) bool {
	if IgnoreAccessDeniedServiceDisabled(err) || IgnoreNotAvailableRegion(err) || IgnoreWithInvalidAction(err) || isNotFoundError(err) {
		return true
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

func IgnoreNotAvailableRegion(err error) bool {
	var ae smithy.APIError
	if errors.As(err, &ae) {
		if ae.ErrorCode() == "InvalidRequestException" && strings.Contains(ae.ErrorMessage(), "not available in the current Region") {
			return true
		}
	}
	return false
}

func accountObfusactor(aa []string, msg string) string {
	for _, a := range aa {
		msg = strings.ReplaceAll(msg, a, obfuscateAccountId(a))
	}
	return msg
}

// makeARN creates an ARN using supplied service name, partition, account id, region name and resource id parts.
// Resource id parts are concatenated using forward slash (/).
// See https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html for more information.
func makeARN(service AWSService, partition, accountID, region string, idParts ...string) arn.ARN {
	return arn.ARN{
		Partition: partition,
		Service:   string(service),
		Region:    region,
		AccountID: accountID,
		Resource:  strings.Join(idParts, "/"),
	}
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
		return resource.Set(c.Name, makeARN(service, cl.Partition, accountID, region, idParts...).String())
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

// IsNotFoundError checks if api error should be ignored
func (c *Client) IsNotFoundError(err error) bool {
	if isNotFoundError(err) {
		c.logger.Warn("API returned \"NotFound\" error ignoring it...", "error", err)
		return true
	}
	return false
}

func isNotFoundError(err error) bool {
	var ae smithy.APIError
	if !errors.As(err, &ae) {
		return false
	}
	errorCode := ae.ErrorCode()
	for _, s := range notFoundErrorSubstrings {
		if strings.Contains(errorCode, s) {
			return true
		}
	}
	return false
}

func IsInvalidParameterValueError(err error) bool {
	var apiErr smithy.APIError
	return errors.As(err, &apiErr) && apiErr.ErrorCode() == "InvalidParameterValue"
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

func IsErrorRegex(err error, code string, messageRegex *regexp.Regexp) bool {
	var ae smithy.APIError
	if !errors.As(err, &ae) {
		return false
	}
	if ae.ErrorCode() == code && messageRegex.MatchString(ae.ErrorMessage()) {
		return true
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

func ListAndDetailResolver(ctx context.Context, meta schema.ClientMeta, res chan<- interface{}, list ListResolverFunc, details DetailResolverFunc) error {
	var diags diag.Diagnostics

	errorChan := make(chan error)
	detailChan := make(chan interface{})
	// Channel that will communicate with goroutine that is aggregating the errors
	done := make(chan struct{})
	go func() {
		defer close(done)
		for detailError := range errorChan {
			diags = diags.Add(diag.FromError(detailError, diag.RESOLVING))
		}
	}()
	sem := semaphore.NewWeighted(int64(MAX_GOROUTINES))

	go func() {
		defer close(errorChan)
		for item := range detailChan {
			if err := sem.Acquire(ctx, 1); err != nil {
				continue
			}
			func(summary interface{}) {
				defer sem.Release(1)
				details(ctx, meta, res, errorChan, summary)
			}(item)
		}
	}()

	err := list(ctx, meta, detailChan)
	close(detailChan)
	if err != nil {
		return diag.WrapError(err)
	}

	// All items will be attempted to be fetched, and all errors will be aggregated
	<-done

	// Only return the diags if there is an error
	if diags.HasDiags() {
		return diags
	}
	return nil
}
