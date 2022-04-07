package client

import (
	"context"

	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func ResolveAWSAccount(_ context.Context, meta schema.ClientMeta, r *schema.Resource, _ schema.Column) error {
	client := meta.(*Client)
	return r.Set("account_id", client.AccountID)
}

func ResolveAWSRegion(_ context.Context, meta schema.ClientMeta, r *schema.Resource, _ schema.Column) error {
	client := meta.(*Client)
	return r.Set("region", client.Region)
}

func ResolveAWSNamespace(_ context.Context, meta schema.ClientMeta, r *schema.Resource, _ schema.Column) error {
	client := meta.(*Client)
	return r.Set("namespace", client.AutoscalingNamespace)
}

func ResolveWAFScope(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	return diag.WrapError(r.Set(c.Name, meta.(*Client).WAFScope))
}
