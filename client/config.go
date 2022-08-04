package client

type Account struct {
	ID              string   `yaml:"id"`
	AccountName     string   `yaml:"account_name,omitempty"`
	LocalProfile    string   `yaml:"local_profile,omitempty"`
	RoleARN         string   `yaml:"role_arn,omitempty"`
	RoleSessionName string   `yaml:"role_session_name,omitempty"`
	ExternalID      string   `yaml:"external_id,omitempty"`
	DefaultRegion   string   `yaml:"default_region,omitempty"`
	Regions         []string `yaml:"regions,omitempty"`
	source          string
}

type AwsOrg struct {
	OrganizationUnits           []string `yaml:"organization_units,omitempty"`
	AdminAccount                *Account `yaml:"admin_account"`
	MemberCredentials           *Account `yaml:"member_trusted_principal"`
	ChildAccountRoleName        string   `yaml:"member_role_name,omitempty"`
	ChildAccountRoleSessionName string   `yaml:"member_role_session_name,omitempty"`
	ChildAccountExternalID      string   `yaml:"member_external_id,omitempty"`
	ChildAccountRegions         []string `yaml:"member_regions,omitempty"`
}

type Config struct {
	Regions      []string  `yaml:"regions,omitempty"`
	Accounts     []Account `yaml:"accounts"`
	Organization *AwsOrg   `yaml:"org"`
	AWSDebug     bool      `yaml:"aws_debug,omitempty"`
	MaxRetries   int       `yaml:"max_retries,omitempty" default:"10"`
	MaxBackoff   int       `yaml:"max_backoff,omitempty" default:"30"`
	GlobalRegion string    `yaml:"global_region,omitempty" default:"us-east-1"`
}
