package resources

import (
	"context"
	"errors"

	aws "github.com/aws/aws-sdk-go-v2/aws"
	s3control "github.com/aws/aws-sdk-go-v2/service/s3control"
	s3controlTypes "github.com/aws/aws-sdk-go-v2/service/s3control/types"
	"github.com/aws/smithy-go"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func S3Accounts() *schema.Table {
	return &schema.Table{
		Name:         "aws_s3_accounts",
		Description:  "TODO-FILL-THIS-IN",
		Resolver:     fetchS3AccountConfig,
		Multiplex:    client.AccountMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"account_id"}},
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:        "config_exists",
				Type:        schema.TypeBool,
				Description: "Specifies whether Amazon S3 access control exists",
			},
			{
				Name:        "block_public_acls",
				Type:        schema.TypeBool,
				Description: "Specifies whether Amazon S3 should block public access control lists (ACLs) for buckets in this account",
			},
			{
				Name:        "block_public_policy",
				Type:        schema.TypeBool,
				Description: "Specifies whether Amazon S3 should block public bucket policies for buckets in this account.",
			},

			{
				Name:        "ignore_public_acls",
				Type:        schema.TypeBool,
				Description: "Specifies whether Amazon S3 should ignore public ACLs for buckets in this account",
			},
			{
				Name:        "restrict_public_buckets",
				Type:        schema.TypeBool,
				Description: "Specifies whether Amazon S3 should restrict public bucket policies for buckets in this account.",
			},
		},
	}
}

func fetchS3AccountConfig(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	var s3AccountConfig S3AccountConfig
	c := meta.(*client.Client)

	svc := c.Services().S3Control
	var accountConfig s3control.GetPublicAccessBlockInput
	accountConfig.AccountId = aws.String(c.AccountID)
	s3AccountConfig.ConfigExists = false
	resp, err := svc.GetPublicAccessBlock(ctx, &accountConfig, func(options *s3control.Options) {
		options.Region = c.Region
	})
	if err != nil {
		// If we received any error other than NoSuchPublicAccessBlockConfiguration, we return and error
		var ae smithy.APIError
		if errors.As(err, &ae) && ae.ErrorCode() != "NoSuchPublicAccessBlockConfiguration" {
			return err
		}
	}
	if resp != nil {
		res <- S3AccountConfig{*resp.PublicAccessBlockConfiguration, true}
	} else {
		res <- s3AccountConfig
	}

	return nil
}

type S3AccountConfig struct {
	s3controlTypes.PublicAccessBlockConfiguration
	ConfigExists bool
}
