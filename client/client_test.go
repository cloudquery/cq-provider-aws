package client

import (
	"context"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/hashicorp/go-hclog"
	"github.com/stretchr/testify/assert"

	"github.com/aws/aws-sdk-go-v2/aws"
)

// emptyInterfaceFieldNames looks at value s, which should be a struct (or a pointer to a struct),
// and returns the list of its field names which represent interface values but have nil value.
func emptyInterfaceFieldNames(s interface{}) []string {
	if s == nil {
		return nil
	}
	v := reflect.ValueOf(s)
	if v.Type().Kind() == reflect.Ptr {
		if v.IsNil() {
			return nil
		}
		v = reflect.Indirect(v)
	}
	if v.Type().Kind() != reflect.Struct {
		return nil
	}
	var empty []string
	for i := 0; i < v.Type().NumField(); i++ {
		field := v.Field(i)
		if t := field.Type(); t == nil || t.Kind() != reflect.Interface {
			continue
		}
		if field.IsNil() {
			empty = append(empty, v.Type().Field(i).Name)
		}
	}
	return empty
}

func Test_emptyInterfaceFieldNames(t *testing.T) {
	// emptyInterfaceFieldNames is a test helper but it is not trivial and uses reflection. So let's test it too.
	tests := []struct {
		s    interface{}
		want []string
	}{
		{nil, nil},
		{
			struct {
				x int
				y *string
			}{}, nil,
		},
		{
			struct {
				x interface{}
				y interface{}
			}{0, "test"}, nil,
		},
		{
			struct {
				x interface{}
				y interface{}
			}{},
			[]string{"x", "y"},
		},
		{
			struct {
				x interface{}
				y interface{}
			}{nil, 1},
			[]string{"x"},
		},
		{
			struct {
				x interface{}
				y interface{}
			}{1, nil},
			[]string{"y"},
		},
		{
			&struct { // test that pointer to a struct works too
				x interface{}
				y interface{}
			}{1, nil},
			[]string{"y"},
		},
	}
	for _, tt := range tests {
		got := emptyInterfaceFieldNames(tt.s)
		results := cmp.Diff(got, tt.want)
		if results != "" {
			t.Errorf(results)
		}
	}
}

func Test_initServices_NoNilValues(t *testing.T) {
	// the purpose of this test is to call initServices and check that returned Services struct
	// has no nil values in its fields.
	empty := emptyInterfaceFieldNames(initServices("us-east-1", aws.Config{}))
	for _, name := range empty {
		t.Errorf("initServices().%s == nil", name)
	}
}

func Test_obfuscateAccountId(t *testing.T) {
	assert.Equal(t, "1111xxxxxxxx", obfuscateAccountId("1111111111"))
	assert.Equal(t, "11", obfuscateAccountId("11"))
}

func Test_isValidRegions(t *testing.T) {
	tests := []struct {
		regions []string
		want    error
	}{
		{
			regions: []string{"us-east-1"},
			want:    nil,
		}, {
			regions: []string{"us-east-1", "*"},
			want:    errInvalidRegion,
		}, {
			regions: []string{"*"},
			want:    nil,
		}, {
			regions: []string{"*", "us-east-1"},
			want:    errInvalidRegion,
		},
	}
	for i, tt := range tests {
		err := isValidRegions(tt.regions)
		results := cmp.Diff(err, tt.want, cmpopts.EquateErrors())
		if results != "" {
			t.Errorf("Case-%d failed: %s", i, results)
		}
	}
}
func Test_isAllRegions(t *testing.T) {
	tests := []struct {
		regions []string
		want    bool
	}{
		{
			regions: []string{"us-east-1"},
			want:    false,
		}, {
			regions: []string{"us-east-1", "*"},
			want:    false,
		}, {
			regions: []string{"*"},
			want:    true,
		}, {
			regions: []string{"*", "us-east-1"},
			want:    false,
		},
	}
	for i, tt := range tests {
		err := isAllRegions(tt.regions)
		results := cmp.Diff(err, tt.want)
		if results != "" {
			t.Errorf("Case-%d failed: %s", i, results)
		}
	}
}

func Test_Configure(t *testing.T) {
	ctx := context.Background()
	logger := hclog.New(&hclog.LoggerOptions{})
	tests := []struct {
		account   Account
		awsConfig *Config
		keyId     string
	}{}
	for i, tt := range tests {
		awsClient, err := configureAwsClient(ctx, logger, tt.awsConfig, tt.account)
		if err != nil {
			t.Errorf("Case-%d failed: %+v", i, err)
		}
		a, err := awsClient.Credentials.Retrieve(ctx)
		if err != nil {
			t.Errorf("Case-%d failed: %+v", i, err)
		}
		if a.AccessKeyID != tt.keyId {
			t.Errorf("Case-%d failed: %+v", i, err)
		}

	}
}
