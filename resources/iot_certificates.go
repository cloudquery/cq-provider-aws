package resources

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func IotCertificates() *schema.Table {
	return &schema.Table{
		Name:         "aws_iot_certificates",
		Description:  "Describes a certificate.",
		Resolver:     fetchIotCertificates,
		Multiplex:    client.AccountRegionMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
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
				Name:     "policies",
				Type:     schema.TypeStringArray,
				Resolver: resolveIotCertificatePolicies,
			},
			{
				Name:        "ca_certificate_id",
				Description: "The certificate ID of the CA certificate used to sign this certificate.",
				Type:        schema.TypeString,
			},
			{
				Name:        "certificate_arn",
				Description: "The ARN of the certificate.",
				Type:        schema.TypeString,
			},
			{
				Name:        "certificate_id",
				Description: "The ID of the certificate.",
				Type:        schema.TypeString,
			},
			{
				Name:        "certificate_mode",
				Description: "The mode of the certificate.",
				Type:        schema.TypeString,
			},
			{
				Name:        "certificate_pem",
				Description: "The certificate data, in PEM format.",
				Type:        schema.TypeString,
			},
			{
				Name:        "creation_date",
				Description: "The date and time the certificate was created.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "customer_version",
				Description: "The customer version of the certificate.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "generation_id",
				Description: "The generation ID of the certificate.",
				Type:        schema.TypeString,
			},
			{
				Name:        "last_modified_date",
				Description: "The date and time the certificate was last modified.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "owned_by",
				Description: "The ID of the AWS account that owns the certificate.",
				Type:        schema.TypeString,
			},
			{
				Name:        "previous_owned_by",
				Description: "The ID of the AWS account of the previous owner of the certificate.",
				Type:        schema.TypeString,
			},
			{
				Name:        "status",
				Description: "The status of the certificate.",
				Type:        schema.TypeString,
			},
			{
				Name:        "transfer_data_accept_date",
				Description: "The date the transfer was accepted.",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("TransferData.AcceptDate"),
			},
			{
				Name:        "transfer_data_reject_date",
				Description: "The date the transfer was rejected.",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("TransferData.RejectDate"),
			},
			{
				Name:        "transfer_data_reject_reason",
				Description: "The reason why the transfer was rejected.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("TransferData.RejectReason"),
			},
			{
				Name:        "transfer_data_transfer_date",
				Description: "The date the transfer took place.",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("TransferData.TransferDate"),
			},
			{
				Name:        "transfer_data_transfer_message",
				Description: "The transfer message.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("TransferData.TransferMessage"),
			},
			{
				Name:        "validity_not_after",
				Description: "The certificate is not valid after this date.",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("Validity.NotAfter"),
			},
			{
				Name:        "validity_not_before",
				Description: "The certificate is not valid before this date.",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("Validity.NotBefore"),
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchIotCertificates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	client := meta.(*client.Client)
	svc := client.Services().IOT
	input := iot.ListCertificatesInput{}

	for {
		response, err := svc.ListCertificates(ctx, &input, func(options *iot.Options) {
			options.Region = client.Region
		})
		if err != nil {
			return err
		}

		for _, ct := range response.Certificates {
			cert, err := svc.DescribeCertificate(ctx, &iot.DescribeCertificateInput{
				CertificateId: ct.CertificateId,
			}, func(options *iot.Options) {
				options.Region = client.Region
			})
			if err != nil {
				return err
			}
			res <- cert.CertificateDescription
		}

		if aws.ToString(response.NextMarker) == "" {
			break
		}
		input.Marker = response.NextMarker
	}
	return nil
}
func resolveIotCertificatePolicies(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	i, ok := resource.Item.(*types.CertificateDescription)
	if !ok {
		return fmt.Errorf("expected *types.CertificateDescription but got %T", resource.Item)
	}
	client := meta.(*client.Client)
	svc := client.Services().IOT
	input := iot.ListAttachedPoliciesInput{
		Target: i.CertificateArn,
	}

	var policies []string
	for {
		response, err := svc.ListAttachedPolicies(ctx, &input, func(options *iot.Options) {
			options.Region = client.Region
		})
		if err != nil {
			return err
		}

		for _, p := range response.Policies {
			policies = append(policies, *p.PolicyArn)
		}

		if aws.ToString(response.NextMarker) == "" {
			break
		}
		input.Marker = response.NextMarker
	}
	return resource.Set(c.Name, policies)
}
