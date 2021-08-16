package resources

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/identitystore"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func IdentityStoreGroups() *schema.Table {
	return &schema.Table{
		Name:         "aws_identity_store_groups",
		Description:  "A group object, which contains a specified group’s metadata and attributes.",
		Resolver:     fetchIdentityStoreGroups,
		Multiplex:    client.AccountRegionMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "id"}},
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
				Name:        "display_name",
				Description: "Contains the group’s display name value",
				Type:        schema.TypeString,
			},
			{
				Name:        "id",
				Description: "The identifier for a group in the identity store.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("GroupId"),
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchIdentityStoreGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	var config identitystore.ListGroupsInput
	client := meta.(*client.Client)
	svc := client.Services().IdentityStore
	for {
		response, err := svc.ListGroups(ctx, &config, func(options *identitystore.Options) {
			options.Region = client.Region
		})
		if err != nil {
			return err
		}
		res <- response.Groups
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}

	return nil
}
