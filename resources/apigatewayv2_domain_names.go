package resources

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
	"github.com/aws/smithy-go"
	smithyio "github.com/aws/smithy-go/io"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"

	"github.com/cloudquery/cq-provider-aws/client"
)

func Apigatewayv2DomainNames() *schema.Table {
	return &schema.Table{
		Name:         "aws_apigatewayv2_domain_names",
		Description:  "Represents a domain name.",
		Resolver:     fetchApigatewayv2DomainNames,
		Multiplex:    client.AccountRegionMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "region", "domain_name"}},
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:        "domain_name",
				Description: "The name of the DomainName resource.",
				Type:        schema.TypeString,
			},
			{
				Name:        "api_mapping_selection_expression",
				Description: "The API mapping selection expression.",
				Type:        schema.TypeString,
			},
			{
				Name:        "mutual_tls_authentication_truststore_uri",
				Description: "An Amazon S3 URL that specifies the truststore for mutual TLS authentication, for example, s3://bucket-name/key-name. The truststore can contain certificates from public or private certificate authorities. To update the truststore, upload a new version to S3, and then update your custom domain name to use the new version. To update the truststore, you must have permissions to access the S3 object.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("MutualTlsAuthentication.TruststoreUri"),
			},
			{
				Name:        "mutual_tls_authentication_truststore_version",
				Description: "The version of the S3 object that contains your truststore. To specify a version, you must have versioning enabled for the S3 bucket.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("MutualTlsAuthentication.TruststoreVersion"),
			},
			{
				Name:        "mutual_tls_authentication_truststore_warnings",
				Description: "A list of warnings that API Gateway returns while processing your truststore. Invalid certificates produce warnings. Mutual TLS is still enabled, but some clients might not be able to access your API. To resolve warnings, upload a new truststore to S3, and then update you domain name to use the new version.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("MutualTlsAuthentication.TruststoreWarnings"),
			},
			{
				Name:        "tags",
				Description: "The collection of tags associated with a domain name.",
				Type:        schema.TypeJSON,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_apigatewayv2_domain_name_configurations",
				Description: "The domain name configuration.",
				Resolver:    fetchApigatewayv2DomainNameConfigurations,
				Columns: []schema.Column{
					{
						Name:        "domain_name_cq_id",
						Description: "Unique CloudQuery ID of aws_apigatewayv2_domain_names table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "api_gateway_domain_name",
						Description: "A domain name for the API.",
						Type:        schema.TypeString,
					},
					{
						Name:        "certificate_arn",
						Description: "An AWS-managed certificate that will be used by the edge-optimized endpoint for this domain name. AWS Certificate Manager is the only supported source.",
						Type:        schema.TypeString,
					},
					{
						Name:        "certificate_name",
						Description: "The user-friendly name of the certificate that will be used by the edge-optimized endpoint for this domain name.",
						Type:        schema.TypeString,
					},
					{
						Name:        "certificate_upload_date",
						Description: "The timestamp when the certificate that was used by edge-optimized endpoint for this domain name was uploaded.",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "domain_name_status",
						Description: "The status of the domain name migration. The valid values are AVAILABLE and UPDATING. If the status is UPDATING, the domain cannot be modified further until the existing operation is complete. If it is AVAILABLE, the domain can be updated.",
						Type:        schema.TypeString,
					},
					{
						Name:        "domain_name_status_message",
						Description: "An optional text message containing detailed information about status of the domain name migration.",
						Type:        schema.TypeString,
					},
					{
						Name:        "endpoint_type",
						Description: "The endpoint type.",
						Type:        schema.TypeString,
					},
					{
						Name:        "hosted_zone_id",
						Description: "The Amazon Route 53 Hosted Zone ID of the endpoint.",
						Type:        schema.TypeString,
					},
					{
						Name:        "security_policy",
						Description: "The Transport Layer Security (TLS) version of the security policy for this domain name. The valid values are TLS_1_0 and TLS_1_2.",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "aws_apigatewayv2_domain_name_rest_api_mappings",
				Description: "Represents an API mapping.",
				Resolver:    fetchApigatewayv2DomainNameRestApiMappings,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"domain_name_cq_id", "api_mapping_id"}},
				Columns: []schema.Column{
					{
						Name:        "domain_name_cq_id",
						Description: "Unique CloudQuery ID of aws_apigatewayv2_domain_names table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "api_id",
						Description: "The API identifier.",
						Type:        schema.TypeString,
					},
					{
						Name:        "stage",
						Description: "The API stage.",
						Type:        schema.TypeString,
					},
					{
						Name:        "api_mapping_id",
						Description: "The API mapping identifier.",
						Type:        schema.TypeString,
					},
					{
						Name:        "api_mapping_key",
						Description: "The API mapping key.",
						Type:        schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchApigatewayv2DomainNames(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	var config apigatewayv2.GetDomainNamesInput
	c := meta.(*client.Client)
	svc := c.Services().Apigatewayv2
	for {
		response, err := svc.GetDomainNames(ctx, &config, func(options *apigatewayv2.Options) {
			options.Region = c.Region
			// NOTE: Using custom OperationDeserializer until this is fixed: https://github.com/aws/aws-sdk-go-v2/issues/1282
			options.APIOptions = append(options.APIOptions, func(stack *middleware.Stack) error {
				_, err := stack.Deserialize.Swap("OperationDeserializer", &awsRestjson1_deserializeOpGetDomainNames{})
				if err != nil {
					return err
				}
				return nil
			})
		})

		if err != nil {
			return err
		}
		res <- response.Items
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}

func fetchApigatewayv2DomainNameConfigurations(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	r, ok := parent.Item.(types.DomainName)
	if !ok {
		return fmt.Errorf("expected DomainName but got %T", r)
	}
	res <- r.DomainNameConfigurations
	return nil
}

func fetchApigatewayv2DomainNameRestApiMappings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	r, ok := parent.Item.(types.DomainName)
	if !ok {
		return fmt.Errorf("expected DomainName but got %T", r)
	}
	config := apigatewayv2.GetApiMappingsInput{
		DomainName: r.DomainName,
	}
	c := meta.(*client.Client)
	svc := c.Services().Apigatewayv2
	for {
		response, err := svc.GetApiMappings(ctx, &config, func(o *apigatewayv2.Options) {
			o.Region = c.Region
		})

		if err != nil {
			return err
		}
		res <- response.Items
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}

// ====================================================================================================================
//                                               Forked GetDomainNames OperationDeserializer
// ====================================================================================================================

type awsRestjson1_deserializeOpGetDomainNames struct {
}

func (*awsRestjson1_deserializeOpGetDomainNames) ID() string {
	return "OperationDeserializer"
}

func (m *awsRestjson1_deserializeOpGetDomainNames) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	out, metadata, err = next.HandleDeserialize(ctx, in)
	if err != nil {
		return out, metadata, err
	}

	response, ok := out.RawResponse.(*smithyhttp.Response)
	if !ok {
		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return out, metadata, awsRestjson1_deserializeOpErrorGetDomainNames(response, &metadata)
	}
	output := &apigatewayv2.GetDomainNamesOutput{}
	out.Result = output

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(response.Body, ringBuffer)

	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return out, metadata, err
	}

	err = awsRestjson1_deserializeOpDocumentGetDomainNamesOutput(&output, shape)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		return out, metadata, &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body with invalid JSON, %w", err),
			Snapshot: snapshot.Bytes(),
		}
	}

	return out, metadata, err
}

func awsRestjson1_deserializeOpErrorGetDomainNames(response *smithyhttp.Response, metadata *middleware.Metadata) error {
	var errorBuffer bytes.Buffer
	if _, err := io.Copy(&errorBuffer, response.Body); err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("failed to copy error response body, %w", err)}
	}
	errorBody := bytes.NewReader(errorBuffer.Bytes())

	errorCode := "UnknownError"
	errorMessage := errorCode

	code := response.Header.Get("X-Amzn-ErrorType")
	if len(code) != 0 {
		errorCode = restjson.SanitizeErrorCode(code)
	}

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	code, message, err := restjson.GetErrorInfo(decoder)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)
	if len(code) != 0 {
		errorCode = restjson.SanitizeErrorCode(code)
	}
	if len(message) != 0 {
		errorMessage = message
	}

	switch {
	case strings.EqualFold("BadRequestException", errorCode):
		return awsRestjson1_deserializeErrorBadRequestException(response, errorBody)

	case strings.EqualFold("NotFoundException", errorCode):
		return awsRestjson1_deserializeErrorNotFoundException(response, errorBody)

	case strings.EqualFold("TooManyRequestsException", errorCode):
		return awsRestjson1_deserializeErrorTooManyRequestsException(response, errorBody)

	default:
		genericError := &smithy.GenericAPIError{
			Code:    errorCode,
			Message: errorMessage,
		}
		return genericError

	}
}

func awsRestjson1_deserializeOpDocumentGetDomainNamesOutput(v **apigatewayv2.GetDomainNamesOutput, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *apigatewayv2.GetDomainNamesOutput
	if *v == nil {
		sv = &apigatewayv2.GetDomainNamesOutput{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "items":
			if err := awsRestjson1_deserializeDocument__listOfDomainName(&sv.Items, value); err != nil {
				return err
			}

		case "nextToken":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected NextToken to be of type string, got %T instead", value)
				}
				sv.NextToken = ptr.String(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeErrorBadRequestException(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	output := &types.BadRequestException{}
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	err := awsRestjson1_deserializeDocumentBadRequestException(&output, shape)

	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)

	return output
}

func awsRestjson1_deserializeErrorNotFoundException(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	output := &types.NotFoundException{}
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	err := awsRestjson1_deserializeDocumentNotFoundException(&output, shape)

	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)

	return output
}

func awsRestjson1_deserializeErrorTooManyRequestsException(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	output := &types.TooManyRequestsException{}
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	err := awsRestjson1_deserializeDocumentTooManyRequestsException(&output, shape)

	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)

	return output
}

func awsRestjson1_deserializeDocument__listOfDomainName(v *[]types.DomainName, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.([]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var cv []types.DomainName
	if *v == nil {
		cv = []types.DomainName{}
	} else {
		cv = *v
	}

	for _, value := range shape {
		var col types.DomainName
		destAddr := &col
		if err := awsRestjson1_deserializeDocumentDomainName(&destAddr, value); err != nil {
			return err
		}
		col = *destAddr
		cv = append(cv, col)

	}
	*v = cv
	return nil
}

func awsRestjson1_deserializeDocumentBadRequestException(v **types.BadRequestException, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.BadRequestException
	if *v == nil {
		sv = &types.BadRequestException{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "message":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected __string to be of type string, got %T instead", value)
				}
				sv.Message = ptr.String(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentTooManyRequestsException(v **types.TooManyRequestsException, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.TooManyRequestsException
	if *v == nil {
		sv = &types.TooManyRequestsException{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "limitType":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected __string to be of type string, got %T instead", value)
				}
				sv.LimitType = ptr.String(jtv)
			}

		case "message":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected __string to be of type string, got %T instead", value)
				}
				sv.Message = ptr.String(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentDomainName(v **types.DomainName, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.DomainName
	if *v == nil {
		sv = &types.DomainName{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "apiMappingSelectionExpression":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected SelectionExpression to be of type string, got %T instead", value)
				}
				sv.ApiMappingSelectionExpression = ptr.String(jtv)
			}

		case "domainName":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected StringWithLengthBetween1And512 to be of type string, got %T instead", value)
				}
				sv.DomainName = ptr.String(jtv)
			}

		case "domainNameConfigurations":
			if err := awsRestjson1_deserializeDocumentDomainNameConfigurations(&sv.DomainNameConfigurations, value); err != nil {
				return err
			}

		case "mutualTlsAuthentication":
			if err := awsRestjson1_deserializeDocumentMutualTlsAuthentication(&sv.MutualTlsAuthentication, value); err != nil {
				return err
			}

		case "tags":
			if err := awsRestjson1_deserializeDocumentTags(&sv.Tags, value); err != nil {
				return err
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentDomainNameConfigurations(v *[]types.DomainNameConfiguration, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.([]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var cv []types.DomainNameConfiguration
	if *v == nil {
		cv = []types.DomainNameConfiguration{}
	} else {
		cv = *v
	}

	for _, value := range shape {
		var col types.DomainNameConfiguration
		destAddr := &col
		if err := awsRestjson1_deserializeDocumentDomainNameConfiguration(&destAddr, value); err != nil {
			return err
		}
		col = *destAddr
		cv = append(cv, col)

	}
	*v = cv
	return nil
}

func awsRestjson1_deserializeDocumentMutualTlsAuthentication(v **types.MutualTlsAuthentication, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.MutualTlsAuthentication
	if *v == nil {
		sv = &types.MutualTlsAuthentication{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "truststoreUri":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected UriWithLengthBetween1And2048 to be of type string, got %T instead", value)
				}
				sv.TruststoreUri = ptr.String(jtv)
			}

		case "truststoreVersion":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected StringWithLengthBetween1And64 to be of type string, got %T instead", value)
				}
				sv.TruststoreVersion = ptr.String(jtv)
			}

		case "truststoreWarnings":
			if err := awsRestjson1_deserializeDocument__listOf__string(&sv.TruststoreWarnings, value); err != nil {
				return err
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentTags(v *map[string]string, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var mv map[string]string
	if *v == nil {
		mv = map[string]string{}
	} else {
		mv = *v
	}

	for key, value := range shape {
		var parsedVal string
		if value != nil {
			jtv, ok := value.(string)
			if !ok {
				return fmt.Errorf("expected StringWithLengthBetween1And1600 to be of type string, got %T instead", value)
			}
			parsedVal = jtv
		}
		mv[key] = parsedVal

	}
	*v = mv
	return nil
}

func awsRestjson1_deserializeDocumentNotFoundException(v **types.NotFoundException, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.NotFoundException
	if *v == nil {
		sv = &types.NotFoundException{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "message":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected __string to be of type string, got %T instead", value)
				}
				sv.Message = ptr.String(jtv)
			}

		case "resourceType":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected __string to be of type string, got %T instead", value)
				}
				sv.ResourceType = ptr.String(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentDomainNameConfiguration(v **types.DomainNameConfiguration, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.DomainNameConfiguration
	if *v == nil {
		sv = &types.DomainNameConfiguration{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "apiGatewayDomainName":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected __string to be of type string, got %T instead", value)
				}
				sv.ApiGatewayDomainName = ptr.String(jtv)
			}

		case "certificateArn":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected Arn to be of type string, got %T instead", value)
				}
				sv.CertificateArn = ptr.String(jtv)
			}

		case "certificateName":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected StringWithLengthBetween1And128 to be of type string, got %T instead", value)
				}
				sv.CertificateName = ptr.String(jtv)
			}

		case "certificateUploadDate":
			if value != nil {
				// NOTE: Fixes this issue: https://github.com/aws/aws-sdk-go-v2/issues/1282
				// Once it's resolved the whole fork can be deleted.
				jtv, ok := value.(json.Number)
				if !ok {
					return fmt.Errorf("expected json.Number to be of type string, got %T instead", value)
				}
				f, err := jtv.Float64()
				if err != nil {
					return err
				}
				t := smithytime.ParseEpochSeconds(f)
				sv.CertificateUploadDate = ptr.Time(t)
			}

		case "domainNameStatus":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected DomainNameStatus to be of type string, got %T instead", value)
				}
				sv.DomainNameStatus = types.DomainNameStatus(jtv)
			}

		case "domainNameStatusMessage":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected __string to be of type string, got %T instead", value)
				}
				sv.DomainNameStatusMessage = ptr.String(jtv)
			}

		case "endpointType":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected EndpointType to be of type string, got %T instead", value)
				}
				sv.EndpointType = types.EndpointType(jtv)
			}

		case "hostedZoneId":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected __string to be of type string, got %T instead", value)
				}
				sv.HostedZoneId = ptr.String(jtv)
			}

		case "securityPolicy":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected SecurityPolicy to be of type string, got %T instead", value)
				}
				sv.SecurityPolicy = types.SecurityPolicy(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocument__listOf__string(v *[]string, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.([]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var cv []string
	if *v == nil {
		cv = []string{}
	} else {
		cv = *v
	}

	for _, value := range shape {
		var col string
		if value != nil {
			jtv, ok := value.(string)
			if !ok {
				return fmt.Errorf("expected __string to be of type string, got %T instead", value)
			}
			col = jtv
		}
		cv = append(cv, col)

	}
	*v = cv
	return nil
}
