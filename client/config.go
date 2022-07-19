package client

type Account struct {
	ID              string `yaml:"id" hcl:",label"`
	AccountID       string
	AccountName     string   `yaml:"account_name,omitempty" hcl:"account_name,optional"`
	LocalProfile    string   `yaml:"local_profile,omitempty" hcl:"local_profile,optional"`
	RoleARN         string   `yaml:"role_arn,omitempty" hcl:"role_arn,optional"`
	RoleSessionName string   `yaml:"role_session_name,omitempty" hcl:"role_session_name,optional"`
	ExternalID      string   `yaml:"external_id,omitempty" hcl:"external_id,optional"`
	DefaultRegion   string   `yaml:"default_region,omitempty" hcl:"default_region,optional"`
	Regions         []string `yaml:"regions,omitempty" hcl:"regions,optional"`
	source          string
}

type AwsOrg struct {
	OrganizationUnits           []string `yaml:"organization_units,omitempty" hcl:"organization_units,optional"`
	AdminAccount                *Account `yaml:"admin_account" hcl:"admin_account,block"`
	MemberCredentials           *Account `yaml:"member_trusted_principal" hcl:"member_trusted_principal,block"`
	ChildAccountRoleName        string   `yaml:"member_role_name,omitempty" hcl:"member_role_name,optional"`
	ChildAccountRoleSessionName string   `yaml:"member_role_session_name,omitempty" hcl:"member_role_session_name,optional"`
	ChildAccountExternalID      string   `yaml:"member_external_id,omitempty" hcl:"member_external_id,optional"`
	ChildAccountRegions         []string `yaml:"member_regions,omitempty" hcl:"member_regions,optional"`
}

type Config struct {
	Regions      []string  `yaml:"regions,omitempty" hcl:"regions,optional"`
	Accounts     []Account `yaml:"accounts" hcl:"accounts,block"`
	Organization *AwsOrg   `yaml:"org" hcl:"org,block"`
	AWSDebug     bool      `yaml:"aws_debug,omitempty" hcl:"aws_debug,optional"`
	MaxRetries   int       `yaml:"max_retries,omitempty" hcl:"max_retries,optional" default:"10"`
	MaxBackoff   int       `yaml:"max_backoff,omitempty" hcl:"max_backoff,optional" default:"30"`
	GlobalRegion string    `yaml:"global_region,omitempty" hcl:"global_region,optional" default:"us-east-1"`
}

func (c Config) Example() string {
	return `
Optional, Repeated. Add an accounts block for every account you want to assume-role into and fetch data from.
accounts:
  - id: <UNIQUE ACCOUNT IDENTIFIER>
Optional. Role ARN we want to assume when accessing this account
    role_arn: < YOUR_ROLE_ARN >
Optional. Named profile in config or credential file from where CQ should grab credentials
    local_profile: < PROFILE_NAME >
Optional. by default assumes all regions
regions:
  - us-east-1
  - us-west-2
Optional. Enable AWS SDK debug logging.
  aws_debug: false
The maximum number of times that a request will be retried for failures. Defaults to 10 retry attempts.
max_retries: 10
The maximum back off delay between attempts. The backoff delays exponentially with a jitter based on the number of attempts. Defaults to 30 seconds.
max_backoff: 30
`
}
