package ses

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"

	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource email_identities --config gen.hcl --output .
func EmailIdentities() *schema.Table {
	return &schema.Table{
		Name:         "aws_ses_email_identities",
		Description:  "Details about an email identity.",
		Resolver:     fetchSesEmailIdentities,
		Multiplex:    client.ServiceAccountRegionMultiplexer("email"),
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
		Columns: []schema.Column{
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) for the resource.",
				Type:        schema.TypeString,
				Resolver:    resolveSesEmailIdentityArn,
			},
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
				Name:        "configuration_set_name",
				Description: "The configuration set used by default when sending from this identity.",
				Type:        schema.TypeString,
			},
			{
				Name:        "dkim_attributes_current_signing_key_length",
				Description: "[Easy DKIM] The key length of the DKIM key pair in use.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DkimAttributes.CurrentSigningKeyLength"),
			},
			{
				Name:        "dkim_attributes_last_key_generation_timestamp",
				Description: "[Easy DKIM] The last time a key pair was generated for this identity.",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("DkimAttributes.LastKeyGenerationTimestamp"),
			},
			{
				Name:        "dkim_attributes_next_signing_key_length",
				Description: "[Easy DKIM] The key length of the future DKIM key pair to be generated",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DkimAttributes.NextSigningKeyLength"),
			},
			{
				Name:        "dkim_attributes_signing_attributes_origin",
				Description: "A string that indicates how DKIM was configured for the identity",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DkimAttributes.SigningAttributesOrigin"),
			},
			{
				Name:        "dkim_attributes_signing_enabled",
				Description: "If the value is true, then the messages that you send from the identity are signed using DKIM",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("DkimAttributes.SigningEnabled"),
			},
			{
				Name:        "dkim_attributes_status",
				Description: "Describes whether or not Amazon SES has successfully located the DKIM records in the DNS records for the domain",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DkimAttributes.Status"),
			},
			{
				Name:        "dkim_attributes_tokens",
				Description: "If you used Easy DKIM (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/easy-dkim.html) to configure DKIM authentication for the domain, then this object contains a set of unique strings that you use to create a set of CNAME records that you add to the DNS configuration for your domain",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("DkimAttributes.Tokens"),
			},
			{
				Name:        "feedback_forwarding_status",
				Description: "The feedback forwarding configuration for the identity",
				Type:        schema.TypeBool,
			},
			{
				Name:        "identity_type",
				Description: "The email identity type",
				Type:        schema.TypeString,
			},
			{
				Name:        "behavior_on_mx_failure",
				Description: "The action to take if the required MX record can't be found when you send an email",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("MailFromAttributes.BehaviorOnMxFailure"),
			},
			{
				Name:        "mail_from_domain",
				Description: "The name of a domain that an email identity uses as a custom MAIL FROM domain.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("MailFromAttributes.MailFromDomain"),
			},
			{
				Name:        "mail_from_domain_status",
				Description: "The status of the MAIL FROM domain",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("MailFromAttributes.MailFromDomainStatus"),
			},
			{
				Name:        "policies",
				Description: "A map of policy names to policies.",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "tags",
				Description: "An array of objects that define the tags (keys and values) that are associated with the email identity.",
				Type:        schema.TypeJSON,
				Resolver:    client.ResolveTags,
			},
			{
				Name:        "verified_for_sending_status",
				Description: "Specifies whether or not the identity is verified",
				Type:        schema.TypeBool,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchSesEmailIdentities(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	return diag.WrapError(client.ListAndDetailResolver(ctx, meta, res, listSesEmailIdentities, sesEmailIdentityDetail))
}
func resolveSesEmailIdentityArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	return client.ResolveARN(client.SESService, func(resource *schema.Resource) ([]string, error) {
		return []string{"identity", *resource.Item.(*sesv2.GetEmailIdentityOutput).ConfigurationSetName}, nil
	})(ctx, meta, resource, c)
}

func sesEmailIdentityDetail(ctx context.Context, meta schema.ClientMeta, resultsChan chan<- interface{}, errorChan chan<- error, detail interface{}) {
	c := meta.(*client.Client)
	svc := c.Services().SES
	n := detail.(types.IdentityInfo)
	config := sesv2.GetEmailIdentityInput{EmailIdentity: n.IdentityName}
	response, err := svc.GetEmailIdentity(ctx, &config)
	if err != nil {
		errorChan <- diag.WrapError(err)
		return
	}
	resultsChan <- response
}

func listSesEmailIdentities(ctx context.Context, meta schema.ClientMeta, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().SES
	var input sesv2.ListEmailIdentitiesInput
	for {
		response, err := svc.ListEmailIdentities(ctx, &input)
		if err != nil {
			return diag.WrapError(err)
		}
		for _, d := range response.EmailIdentities {
			res <- d
		}
		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}
