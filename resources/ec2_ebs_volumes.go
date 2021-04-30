package resources

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func Ec2EbsVolumes() *schema.Table {
	return &schema.Table{
		Name:         "aws_ec2_ebs_volumes",
		Resolver:     fetchEc2EbsVolumes,
		Multiplex:    client.AccountRegionMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
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
				Name: "volume_id",
				Type: schema.TypeString,
			},
			{
				Name: "availability_zone",
				Type: schema.TypeString,
			},
			{
				Name: "create_time",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "encrypted",
				Type: schema.TypeBool,
			},
			{
				Name: "fast_restored",
				Type: schema.TypeBool,
			},
			{
				Name: "iops",
				Type: schema.TypeInt,
			},
			{
				Name: "kms_key_id",
				Type: schema.TypeString,
			},
			{
				Name: "multi_attach_enabled",
				Type: schema.TypeBool,
			},
			{
				Name: "outpost_arn",
				Type: schema.TypeString,
			},
			{
				Name: "size",
				Type: schema.TypeInt,
			},
			{
				Name: "snapshot_id",
				Type: schema.TypeString,
			},
			{
				Name: "state",
				Type: schema.TypeString,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveEc2EbsVolumeTags,
			},
			{
				Name: "throughput",
				Type: schema.TypeInt,
			},
			{
				Name: "volume_type",
				Type: schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchEc2EbsVolumes(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().EC2

	response, err := svc.DescribeVolumes(ctx, &ec2.DescribeVolumesInput{}, func(o *ec2.Options) {
		o.Region = c.Region
	})
	if err != nil {
		return err
	}
	for _, volume := range response.Volumes {
		res <- volume
	}
	return nil
}
func resolveEc2EbsVolumeTags(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, _ schema.Column) error {
	r := resource.Item.(types.Volume)
	tags := map[string]*string{}
	for _, t := range r.Tags {
		tags[*t.Key] = t.Value
	}
	resource.Set("tags", tags)
	return nil
}
