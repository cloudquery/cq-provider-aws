package client

type Account struct {
	ID              string `yaml:"id,label"`
	AccountID       string
	AccountName     string   `yaml:"account_name,omitempty"`
	LocalProfile    string   `yaml:"local_profile,omitempty"`
	RoleARN         string   `yaml:"role_arn,omitempty"`
	RoleSessionName string   `yaml:"role_session_name,omitempty"`
	ExternalID      string   `yaml:"external_id,omitempty"`
	Regions         []string `yaml:"regions,omitempty"`
	source          string
}

type AwsOrg struct {
	OrganizationUnits           []string `yaml:"organization_units,optional"`
	AdminAccount                *Account `yaml:"admin_account,block"`
	MemberCredentials           *Account `yaml:"member_trusted_principal,block"`
	ChildAccountRoleName        string   `yaml:"member_role_name,optional"`
	ChildAccountRoleSessionName string   `yaml:"member_role_session_name,optional"`
	ChildAccountExternalID      string   `yaml:"member_external_id,optional"`
	ChildAccountRegions         []string `yaml:"member_regions,optional"`
}

type Config struct {
	Regions      []string  `yaml:"regions,omitempty"`
	Accounts     []Account `yaml:"accounts,omitempty"`
	Organization *AwsOrg   `yaml:"org,omitempty"`
	AWSDebug     bool      `yaml:"aws_debug,omitempty"`
	MaxRetries   int       `yaml:"max_retries,omitempty" default:"10"`
	MaxBackoff   int       `yaml:"max_backoff,omitempty" default:"30"`
}

func (c Config) Example() string {
	return `
		configuration:
//			accounts: // Optional, Repeated. Add an 'accounts' block for every account you want to assume-role into and fetch data from.
//				- name: "unique_account_identifier" // Optional, Repeated. Add an 'accounts' block for every account you want to assume-role into and fetch data from.  
//				  role_arn: < YOUR_ROLE_ARN > // Optional. Role ARN we want to assume when accessing this account
//					local_profile = < PROFILE_NAME > // Optional. Named profile in config or credential file from where CQ should grab credentials
//		  regions: ["us-east-1", "us-west-2"] // Optional. by default assumes all regions
//      aws_debug: false
//      max_retries = 10 // The maximum number of times that a request will be retried for failures. Defaults to 10 retry attempts.      
//      max_backoff: 30 //The maximum back off delay between attempts. The backoff delays exponentially with a jitter based on the number of attempts. Defaults to 30 seconds.
`
}
