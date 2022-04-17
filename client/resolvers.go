package client

import (
	"context"
	"reflect"
	"time"

	"github.com/spf13/cast"
	"github.com/thoas/go-funk"

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

func ResolveTags(ctx context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	return ResolveTagField("Tags")(ctx, meta, r, c)
}

func ResolveTagField(fieldName string) func(context.Context, schema.ClientMeta, *schema.Resource, schema.Column) error {
	return func(_ context.Context, _ schema.ClientMeta, r *schema.Resource, c schema.Column) error {
		var val reflect.Value

		if reflect.TypeOf(r.Item).Kind() == reflect.Ptr {
			val = reflect.ValueOf(r.Item).Elem()
		} else {
			val = reflect.ValueOf(r.Item)
		}

		if val.Kind() != reflect.Struct {
			panic("need struct type")
		}
		f := val.FieldByName(fieldName)
		if f.IsNil() {
			return diag.WrapError(r.Set(c.Name, map[string]string{})) // can't have nil or the integration test will make a fuss
		} else if f.IsZero() {
			panic("no such field " + fieldName)
		}
		data := TagsToMap(f.Interface())
		return diag.WrapError(r.Set(c.Name, data))
	}
}

func parseISODate(d string) (*time.Time, error) {
	if d == "" {
		return nil, nil
	}
	location, err := time.LoadLocation("UTC")
	if err != nil {
		return nil, err
	}
	date, err := time.ParseInLocation(time.RFC3339, d, location)
	if err != nil {
		return nil, err
	}
	return &date, err
}

func ISODateResolver(path string) schema.ColumnResolver {
	return func(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
		data, err := cast.ToStringE(funk.Get(r.Item, path, funk.WithAllowZero()))
		if err != nil {
			return err
		}
		date, err := parseISODate(data)
		if err != nil {
			return err
		}
		return r.Set(c.Name, date)
	}
}
