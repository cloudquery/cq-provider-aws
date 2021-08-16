package resources

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/identitystore"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func IdentityStoreUsers() *schema.Table {
	return &schema.Table{
		Name:         "aws_identity_store_users",
		Description:  "A user object, which contains a specified user’s metadata and attributes.",
		Resolver:     fetchIdentityStoreUsers,
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
				Name:        "id",
				Description: "The identifier for a user in the identity store.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("UserId"),
			},
			{
				Name:        "name",
				Description: "Contains the user’s user name value",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("UserName"),
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchIdentityStoreUsers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	var config identitystore.ListUsersInput
	client := meta.(*client.Client)
	svc := client.Services().IdentityStore
	for {
		response, err := svc.ListUsers(ctx, &config, func(options *identitystore.Options) {
			options.Region = client.Region
		})
		if err != nil {
			return err
		}
		res <- response.Users
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}

	return nil
}
