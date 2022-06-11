package client

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/aws-sdk-go-v2/service/sts"

	"github.com/aws/smithy-go/logging"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/hashicorp/go-hclog"
)

type Client struct {
	// Those are already normalized values after configure and this is why we don't want to hold
	// config directly.
	Accounts        []Account
	Cfgs            []aws.Config
	logLevel        *string
	maxRetries      int
	maxBackoff      int
	ServicesManager ServicesManager
	logger          hclog.Logger
	// this is set by table clientList
	AccountID            string
	AWSCfg               aws.Config
	Region               string
	AutoscalingNamespace string
	// WAFScope             wafv2types.Scope
	// Partition string
}

// S3Manager This is needed because https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/feature/s3/manager
// has different structure then all other services (i.e no service but just a function) and we need
// the ability to mock it.
// Also we need to use s3 manager to be able to query the bucket-region https://github.com/aws/aws-sdk-go-v2/pull/1027#issuecomment-759818990
// type S3Manager struct {
// 	s3Client *s3.Client
// }

type AwsLogger struct {
	l hclog.Logger
}

type AssumeRoleAPIClient interface {
	AssumeRole(ctx context.Context, params *sts.AssumeRoleInput, optFns ...func(*sts.Options)) (*sts.AssumeRoleOutput, error)
}

type Services struct {
	CloudControl CloudControlClient
}

type ServicesAccountRegionMap map[string]map[string]*Services

// ServicesManager will hold the entire map of (account X region) services
type ServicesManager struct {
	services         ServicesAccountRegionMap
	wafScopeServices map[string]*Services
}

const (
	defaultRegion              = "us-east-1"
	awsFailedToConfigureErrMsg = "failed to retrieve credentials for account %s. AWS Error: %w, detected aws env variables: %s"
	defaultVar                 = "default"
	cloudfrontScopeRegion      = defaultRegion
)

var envVarsToCheck = []string{
	"AWS_PROFILE",
	"AWS_ACCESS_KEY_ID",
	"AWS_SECRET_ACCESS_KEY",
	"AWS_CONFIG_FILE",
	"AWS_ROLE_ARN",
	"AWS_SESSION_TOKEN",
	"AWS_SHARED_CREDENTIALS_FILE",
}

var errInvalidRegion = fmt.Errorf("region wildcard \"*\" is only supported as first argument")
var errUnknownRegion = func(region string) error {
	return fmt.Errorf("unknown region: %q", region)
}

var (
	_ schema.ClientMeta       = (*Client)(nil)
	_ schema.ClientIdentifier = (*Client)(nil)
)

func (s *ServicesManager) ServicesByAccountAndRegion(accountId string, region string) *Services {
	if region == "" {
		region = defaultRegion
	}
	return s.services[accountId][region]
}

func (s *ServicesManager) ServicesByAccountForWAFScope(accountId string) *Services {
	return s.wafScopeServices[accountId]
}

func (s *ServicesManager) InitServicesForAccountAndRegion(accountId string, region string, services Services) {
	if s.services[accountId] == nil {
		s.services[accountId] = make(map[string]*Services)
	}
	s.services[accountId][region] = &services
}

// func (s *ServicesManager) InitServicesForAccountAndScope(accountId string, services Services) {
// 	if s.wafScopeServices == nil {
// 		s.wafScopeServices = make(map[string]*Services)
// 	}
// 	s.wafScopeServices[accountId] = &services
// }

// func newS3ManagerFromConfig(cfg aws.Config) S3Manager {
// 	return S3Manager{
// 		s3Client: s3.NewFromConfig(cfg),
// 	}
// }

// func (s3Manager S3Manager) GetBucketRegion(ctx context.Context, bucket string, optFns ...func(*s3.Options)) (string, error) {
// 	return manager.GetBucketRegion(ctx, s3Manager.s3Client, bucket, optFns...)
// }

func NewAwsClient(logger hclog.Logger, accounts []Account) Client {
	return Client{
		ServicesManager: ServicesManager{
			services: ServicesAccountRegionMap{},
		},
		logger:   logger,
		Accounts: accounts,
	}
}
func (c *Client) Logger() hclog.Logger {
	return &awsLogger{c.logger, c.Accounts}
}

// Identify the given client
func (c *Client) Identify() string {
	return strings.TrimRight(strings.Join([]string{
		obfuscateAccountId(c.AccountID),
		c.Region,
		c.AutoscalingNamespace,
		// string(c.WAFScope),
	}, ":"), ":")
}

func (c *Client) Services() *Services {
	s := c.ServicesManager.ServicesByAccountAndRegion(c.AccountID, c.Region)
	// if s == nil && c.WAFScope == wafv2types.ScopeCloudfront {
	// 	return c.ServicesManager.ServicesByAccountForWAFScope(c.AccountID)
	// }
	return s
}

// ARN builds an ARN tied to current client's partition, accountID and region
// func (c *Client) ARN(service AWSService, idParts ...string) string {
// 	return makeARN(service, c.Partition, c.AccountID, c.Region, idParts...).String()
// }

// AccountGlobalARN builds an ARN tied to current client's partition and accountID
// func (c *Client) AccountGlobalARN(service AWSService, idParts ...string) string {
// 	return makeARN(service, c.Partition, c.AccountID, "", idParts...).String()
// }

// PartitionGlobalARN builds an ARN tied to current client's partition
// func (c *Client) PartitionGlobalARN(service AWSService, idParts ...string) string {
// 	return makeARN(service, c.Partition, "", "", idParts...).String()
// }

func (c *Client) withAccountID(accountID string) *Client {
	return &Client{
		// Partition:            c.Partition,
		Accounts:             c.Accounts,
		logLevel:             c.logLevel,
		maxRetries:           c.maxRetries,
		maxBackoff:           c.maxBackoff,
		ServicesManager:      c.ServicesManager,
		logger:               c.logger.With("account_id", obfuscateAccountId(accountID)),
		AccountID:            accountID,
		Region:               c.Region,
		AutoscalingNamespace: c.AutoscalingNamespace,
	}
}

func (c *Client) withAccountIDAndRegion(accountID, region string, cfg aws.Config) *Client {
	return &Client{
		// Partition:            c.Partition,
		AWSCfg:               cfg,
		Accounts:             c.Accounts,
		logLevel:             c.logLevel,
		maxRetries:           c.maxRetries,
		maxBackoff:           c.maxBackoff,
		ServicesManager:      c.ServicesManager,
		logger:               c.logger.With("account_id", obfuscateAccountId(accountID), "Region", region),
		AccountID:            accountID,
		Region:               region,
		AutoscalingNamespace: c.AutoscalingNamespace,
		// WAFScope:             c.WAFScope,
	}
}

func (c *Client) withAccountIDRegionAndNamespace(accountID, region, namespace string) *Client {
	return &Client{
		// Partition:            c.Partition,
		Accounts:             c.Accounts,
		logLevel:             c.logLevel,
		maxRetries:           c.maxRetries,
		maxBackoff:           c.maxBackoff,
		ServicesManager:      c.ServicesManager,
		logger:               c.logger.With("account_id", obfuscateAccountId(accountID), "Region", region, "AutoscalingNamespace", namespace),
		AccountID:            accountID,
		Region:               region,
		AutoscalingNamespace: namespace,
		// WAFScope:             c.WAFScope,
	}
}

// func (c *Client) withAccountIDRegionAndScope(accountID, region string, scope wafv2types.Scope) *Client {
// 	return &Client{
// 		Partition:            c.Partition,
// 		Accounts:             c.Accounts,
// 		logLevel:             c.logLevel,
// 		maxRetries:           c.maxRetries,
// 		maxBackoff:           c.maxBackoff,
// 		ServicesManager:      c.ServicesManager,
// 		logger:               c.logger.With("account_id", obfuscateAccountId(accountID), "Region", region, "Scope", scope),
// 		AccountID:            accountID,
// 		Region:               region,
// 		AutoscalingNamespace: c.AutoscalingNamespace,
// 		// WAFScope:             scope,
// 	}
// }

func verifyRegions(regions []string) error {
	availableRegions, err := getAvailableRegions()
	if err != nil {
		return err
	}

	// validate regions values
	var hasWildcard bool
	for i, region := range regions {
		if region == "*" {
			hasWildcard = true
		}
		if i != 0 && region == "*" {
			return errInvalidRegion
		}
		if i > 0 && hasWildcard {
			return errInvalidRegion
		}
		regionExist := availableRegions[region]
		if !hasWildcard && !regionExist {
			return errUnknownRegion(region)
		}
	}
	return nil
}
func isAllRegions(regions []string) bool {
	// if regions array is not valid return false
	err := verifyRegions(regions)
	if err != nil {
		return false
	}

	wildcardAllRegions := false
	if (len(regions) == 1 && regions[0] == "*") || (len(regions) == 0) {
		wildcardAllRegions = true
	}
	return wildcardAllRegions
}

func getAccountId(ctx context.Context, awsCfg aws.Config) (*sts.GetCallerIdentityOutput, error) {
	svc := sts.NewFromConfig(awsCfg)
	return svc.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
}

func configureAwsClient(ctx context.Context, logger hclog.Logger, awsConfig *Config, account Account, stsClient AssumeRoleAPIClient) (aws.Config, error) {
	var err error
	var awsCfg aws.Config
	configFns := []func(*config.LoadOptions) error{
		config.WithDefaultRegion(defaultRegion),
		config.WithRetryer(newRetryer(logger, awsConfig.MaxRetries, awsConfig.MaxBackoff)),
	}

	if account.LocalProfile != "" {
		configFns = append(configFns, config.WithSharedConfigProfile(account.LocalProfile))
	}

	awsCfg, err = config.LoadDefaultConfig(ctx, configFns...)

	if err != nil {
		logger.Error("error loading default config", "err", err)
		return awsCfg, fmt.Errorf(awsFailedToConfigureErrMsg, account.AccountName, err, checkEnvVariables())
	}

	if account.RoleARN != "" {
		opts := make([]func(*stscreds.AssumeRoleOptions), 0, 1)
		if account.ExternalID != "" {
			opts = append(opts, func(opts *stscreds.AssumeRoleOptions) {
				opts.ExternalID = &account.ExternalID
			})
		}
		if account.RoleSessionName != "" {
			opts = append(opts, func(opts *stscreds.AssumeRoleOptions) {
				opts.RoleSessionName = account.RoleSessionName
			})
		}
		if stsClient == nil {
			stsClient = sts.NewFromConfig(awsCfg)
		}
		provider := stscreds.NewAssumeRoleProvider(stsClient, account.RoleARN, opts...)

		awsCfg.Credentials = aws.NewCredentialsCache(provider)
	}

	if awsConfig.AWSDebug {
		awsCfg.ClientLogMode = aws.LogRequest | aws.LogResponse | aws.LogRetries
		awsCfg.Logger = AwsLogger{logger.With("accountName", account.AccountName)}
	}

	// Test out retrieving credentials
	if _, err := awsCfg.Credentials.Retrieve(ctx); err != nil {
		logger.Error("error retrieving credentials", "err", err)
		return awsCfg, classifyError(fmt.Errorf(awsFailedToConfigureErrMsg, account.AccountName, err, checkEnvVariables()), diag.INTERNAL, nil, diag.WithSeverity(diag.ERROR), diag.WithNoOverwrite())
	}

	return awsCfg, err
}

func Configure(logger hclog.Logger, providerConfig interface{}) (schema.ClientMeta, diag.Diagnostics) {
	var diags diag.Diagnostics

	ctx := context.Background()
	awsConfig := providerConfig.(*Config)
	client := NewAwsClient(logger, awsConfig.Accounts)
	var adminAccountSts AssumeRoleAPIClient

	if len(awsConfig.Accounts) == 0 {
		awsConfig.Accounts = append(awsConfig.Accounts, Account{
			ID: defaultVar,
		})
	}

	for _, account := range awsConfig.Accounts {
		logger.Debug("user defined account", "account", account)
		if account.AccountID != "" {
			return nil, diags.Add(diag.FromError(errors.New("account_id is no longer supported. To specify a profile use `local_profile`. To specify an account alias use `account_name`"), diag.USER))
		}

		if account.AccountName == "" {
			account.AccountName = account.ID
		}

		localRegions := account.Regions
		if len(localRegions) == 0 {
			localRegions = awsConfig.Regions
		}

		if err := verifyRegions(localRegions); err != nil {
			return nil, diags.Add(classifyError(err, diag.USER, nil))
		}

		if isAllRegions(localRegions) {
			logger.Info("All regions specified in config.yml. Assuming all regions")
		}

		awsCfg, err := configureAwsClient(ctx, logger, awsConfig, account, adminAccountSts)
		if err != nil {
			return nil, diags.Add(diag.FromError(err, diag.ACCESS))
		}

		// This is a work-around to skip disabled regions
		// https://github.com/aws/aws-sdk-go-v2/issues/1068
		res, err := ec2.NewFromConfig(awsCfg).DescribeRegions(ctx,
			&ec2.DescribeRegionsInput{AllRegions: aws.Bool(false)},
			func(o *ec2.Options) {
				o.Region = defaultRegion
				if len(localRegions) > 0 && !isAllRegions(localRegions) {
					o.Region = localRegions[0]
				}
			})
		if err != nil {
			return nil, diags.Add(classifyError(fmt.Errorf("failed to find disabled regions for account %s. AWS Error: %w", account.AccountName, err), diag.INTERNAL, nil, diag.WithSeverity(diag.ERROR), diag.WithNoOverwrite()))
		}
		account.Regions = filterDisabledRegions(localRegions, res.Regions)

		if len(account.Regions) == 0 {
			return nil, diags.Add(diag.FromError(fmt.Errorf("no enabled regions provided in config for account %s", account.AccountName), diag.USER))
		}
		awsCfg.Region = account.Regions[0]
		output, err := getAccountId(ctx, awsCfg)
		if err != nil {
			return nil, diags.Add(classifyError(err, diag.INTERNAL, nil))
		}

		client.Accounts = append(client.Accounts,
			Account{
				ID:      *output.Account,
				RoleARN: *output.Arn,
				Regions: account.Regions,
			})
		client.Cfgs = append(client.Cfgs, awsCfg)
		// for _, region := range account.Regions {
		// 	client.ServicesManager.InitServicesForAccountAndRegion(*output.Account, region, initServices(region, awsCfg))
		// }
		// client.ServicesManager.InitServicesForAccountAndScope(*output.Account, initServices(cloudfrontScopeRegion, awsCfg))
	}
	if len(client.Accounts) == 0 {
		return nil, diags.Add(diag.FromError(errors.New("no accounts instantiated"), diag.USER))
	}
	return &client, diags
}

func initServices(region string, c aws.Config) Services {
	awsCfg := c.Copy()
	awsCfg.Region = region
	// cc := cloudcontrol.NewFromConfig(awsCfg)

	return Services{
		CloudControl: cloudcontrol.NewFromConfig(awsCfg),
	}
}

func filterDisabledRegions(regions []string, enabledRegions []types.Region) []string {
	regionsMap := map[string]bool{}
	for _, r := range enabledRegions {
		if r.RegionName != nil && r.OptInStatus != nil && *r.OptInStatus != "not-opted-in" {
			regionsMap[*r.RegionName] = true
		}
	}

	var filteredRegions []string
	// Our list of regions might not always be the latest and most up to date list
	// if a user specifies all regions via a "*" then they should get the most broad list possible
	if isAllRegions(regions) {
		for region := range regionsMap {
			filteredRegions = append(filteredRegions, region)
		}
	} else {
		for _, r := range regions {
			if regionsMap[r] {
				filteredRegions = append(filteredRegions, r)
			}
		}
	}
	return filteredRegions
}

func (a AwsLogger) Logf(classification logging.Classification, format string, v ...interface{}) {
	if classification == logging.Warn {
		a.l.Warn(fmt.Sprintf(format, v...))
	} else {
		a.l.Debug(fmt.Sprintf(format, v...))
	}
}

func obfuscateAccountId(accountId string) string {
	if len(accountId) <= 4 {
		return accountId
	}
	return accountId[:4] + "xxxxxxxx"
}

// checkEnvVariables checks which aws environment variables are set
func checkEnvVariables() string {
	var result []string
	for _, v := range envVarsToCheck {
		if _, present := os.LookupEnv(v); present {
			result = append(result, v)
		}
	}
	return strings.Join(result, ",")
}
