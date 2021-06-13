package resources

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func IamServerCertificates() *schema.Table {
	return &schema.Table{
		Name:         "aws_iam_server_certificates",
		Description:  "Contains information about a server certificate without its certificate body, certificate chain, and private key.",
		Resolver:     fetchIamServerCertificates,
		Multiplex:    client.AccountMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountFilter,
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) specifying the server certificate. For more information about ARNs and how to use them in policies, see IAM identifiers (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) in the IAM User Guide.",
				Type:        schema.TypeString,
			},
			{
				Name:        "path",
				Description: "The path to the server certificate. For more information about paths, see IAM identifiers (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) in the IAM User Guide.",
				Type:        schema.TypeString,
			},
			{
				Name:        "server_certificate_id",
				Description: "The stable and unique string identifying the server certificate. For more information about IDs, see IAM identifiers (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) in the IAM User Guide.",
				Type:        schema.TypeString,
			},
			{
				Name:        "server_certificate_name",
				Description: "The name that identifies the server certificate.",
				Type:        schema.TypeString,
			},
			{
				Name:        "expiration",
				Description: "The date on which the certificate is set to expire. ",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "upload_date",
				Description: "The date when the server certificate was uploaded. ",
				Type:        schema.TypeTimestamp,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchIamServerCertificates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	var config iam.ListServerCertificatesInput
	svc := meta.(*client.Client).Services().IAM
	for {
		response, err := svc.ListServerCertificates(ctx, &config)
		if err != nil {
			return err
		}
		res <- response.ServerCertificateMetadataList
		if aws.ToString(response.Marker) == "" {
			break
		}
		config.Marker = response.Marker
	}
	return nil
}
