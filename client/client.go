package client

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	wafv2types "github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/aws/smithy-go"
	"github.com/cloudquery/cq-provider-sdk/plugins"
	"github.com/cloudquery/cq-provider-sdk/spec"
	"github.com/rs/zerolog"
	"github.com/xeipuuv/gojsonschema"
)

type Client struct {
	plugins.SourcePlugin
	// Those are already normalized values after configure and this is why we don't want to hold
	// config directly.
	ServicesManager ServicesManager
	// this is set by table clientList
	AccountID            string
	GlobalRegion         string
	Region               string
	AutoscalingNamespace string
	WAFScope             wafv2types.Scope
	Partition            string
}

type AssumeRoleAPIClient interface {
	AssumeRole(ctx context.Context, params *sts.AssumeRoleInput, optFns ...func(*sts.Options)) (*sts.AssumeRoleOutput, error)
}

type ServicesPartitionAccountRegionMap map[string]map[string]map[string]*Services

// ServicesManager will hold the entire map of (account X region) services
type ServicesManager struct {
	services         ServicesPartitionAccountRegionMap
	wafScopeServices map[string]map[string]*Services
}

const (
	defaultRegion         = "us-east-1"
	defaultVar            = "default"
	cloudfrontScopeRegion = defaultRegion
)

var errInvalidRegion = fmt.Errorf("region wildcard \"*\" is only supported as first argument")
var errUnknownRegion = func(region string) error {
	return fmt.Errorf("unknown region: %q", region)
}

func (s *ServicesManager) ServicesByPartitionAccountAndRegion(partition, accountId, region string) *Services {
	if region == "" {
		region = defaultRegion
	}
	return s.services[partition][accountId][region]
}

func (s *ServicesManager) ServicesByAccountForWAFScope(partition, accountId string) *Services {
	return s.wafScopeServices[partition][accountId]
}

func (s *ServicesManager) InitServicesForPartitionAccountAndRegion(partition, accountId, region string, services Services) {
	if s.services == nil {
		s.services = make(map[string]map[string]map[string]*Services)
	}
	if s.services[partition] == nil {
		s.services[partition] = make(map[string]map[string]*Services)
	}
	if s.services[partition][accountId] == nil {
		s.services[partition][accountId] = make(map[string]*Services)
	}
	s.services[partition][accountId][region] = &services
}

func (s *ServicesManager) InitServicesForPartitionAccountAndScope(partition, accountId string, services Services) {
	if s.wafScopeServices == nil {
		s.wafScopeServices = make(map[string]map[string]*Services)
	}
	if s.wafScopeServices[partition] == nil {
		s.wafScopeServices[partition] = make(map[string]*Services)
	}
	s.wafScopeServices[partition][accountId] = &services
}

func NewAwsClient(logger zerolog.Logger) Client {
	return Client{
		ServicesManager: ServicesManager{
			services: ServicesPartitionAccountRegionMap{},
		},
		logger: logger,
	}
}

func (s ServicesPartitionAccountRegionMap) Accounts() []string {
	accounts := make([]string, 0)
	for partitions := range s {
		for account := range s[partitions] {
			accounts = append(accounts, account)
		}
	}
	return accounts
}

func (c *Client) Services() *Services {
	s := c.ServicesManager.ServicesByPartitionAccountAndRegion(c.Partition, c.AccountID, c.Region)
	if s == nil && c.WAFScope == wafv2types.ScopeCloudfront {
		return c.ServicesManager.ServicesByAccountForWAFScope(c.Partition, c.AccountID)
	}
	return s
}

// ARN builds an ARN tied to current client's partition, accountID and region
func (c *Client) ARN(service AWSService, idParts ...string) string {
	return makeARN(service, c.Partition, c.AccountID, c.Region, idParts...).String()
}

// AccountGlobalARN builds an ARN tied to current client's partition and accountID
func (c *Client) AccountGlobalARN(service AWSService, idParts ...string) string {
	return makeARN(service, c.Partition, c.AccountID, "", idParts...).String()
}

// PartitionGlobalARN builds an ARN tied to current client's partition
func (c *Client) PartitionGlobalARN(service AWSService, idParts ...string) string {
	return makeARN(service, c.Partition, "", "", idParts...).String()
}

func (c *Client) withPartitionAccountIDAndRegion(partition, accountID, region string) *Client {
	return &Client{
		Partition:            partition,
		ServicesManager:      c.ServicesManager,
		logger:               c.logger.With().Str("account_id", accountID).Str("region", region).Logger(),
		AccountID:            accountID,
		Region:               region,
		AutoscalingNamespace: c.AutoscalingNamespace,
		WAFScope:             c.WAFScope,
	}
}

func (c *Client) withPartitionAccountIDRegionAndNamespace(partition, accountID, region, namespace string) *Client {
	c.Logger().With()
	return &Client{
		Partition:            partition,
		ServicesManager:      c.ServicesManager,
		logger:               c.logger.With().Str("account_id", accountID).Str("region", region).Str("AutoscalingNamespace", namespace).Logger(),
		AccountID:            accountID,
		Region:               region,
		AutoscalingNamespace: namespace,
		WAFScope:             c.WAFScope,
	}
}

func (c *Client) withPartitionAccountIDRegionAndScope(partition, accountID, region string, scope wafv2types.Scope) *Client {
	return &Client{
		Partition:            partition,
		ServicesManager:      c.ServicesManager,
		logger:               c.logger.With().Str("account_id", accountID).Str("region", region).Interface("scope", scope).Logger(),
		AccountID:            accountID,
		Region:               region,
		AutoscalingNamespace: c.AutoscalingNamespace,
		WAFScope:             scope,
	}
}

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

func configureAwsClient(ctx context.Context, logger zerolog.Logger, awsConfig *Config, account Account, stsClient AssumeRoleAPIClient) (aws.Config, error) {
	var err error
	var awsCfg aws.Config
	configFns := []func(*config.LoadOptions) error{
		config.WithDefaultRegion(defaultRegion),
		config.WithRetryer(newRetryer(awsConfig.MaxRetries, awsConfig.MaxBackoff)),
	}

	if account.DefaultRegion != "" {
		// According to the docs: If multiple WithDefaultRegion calls are made, the last call overrides the previous call values
		configFns = append(configFns, config.WithDefaultRegion(account.DefaultRegion))
	}

	if account.LocalProfile != "" {
		configFns = append(configFns, config.WithSharedConfigProfile(account.LocalProfile))
	}

	awsCfg, err = config.LoadDefaultConfig(ctx, configFns...)

	if err != nil {
		return awsCfg, err
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
		awsCfg.Logger = awsLoggerAdapter{logger.With().Str("accountName", account.AccountName).Logger()}
	}

	// Test out retrieving credentials
	if _, err := awsCfg.Credentials.Retrieve(ctx); err != nil {
		logger.Error().Err(err).Msg("error retrieving credentials")

		var ae smithy.APIError
		if errors.As(err, &ae) {
			if strings.Contains(ae.ErrorCode(), "InvalidClientTokenId") {
				logger.Warn().Msg("The credentials being used to assume role are invalid. Please check that your credentials are valid in the partition you are using. If you are using a partition other than the AWS commercial region, be sure set the default_region attribute in the cloudquery.yml file.")
			}
		}
		logger.Warn().Msg("Couldn't find any credentials in environment variables or configuration files.")
		return awsCfg, nil
	}

	return awsCfg, err
}

func (c *Client) Configure(ctx context.Context, spec spec.SourceSpec) (*gojsonschema.Result, error) {
	var awsSpec Config
	if err := spec.Spec.Decode(&awsSpec); err != nil {
		return nil, fmt.Errorf("error decoding spec: %w", err)
	}
	// awsConfig := providerConfig.(*Config)
	var adminAccountSts AssumeRoleAPIClient
	if awsSpec.Organization != nil && len(awsSpec.Accounts) > 0 {
		return nil, fmt.Errorf("specifying accounts via both the Accounts and Org properties is not supported. If you want to do both, you should use multiple provider blocks")
	}

	if len(awsSpec.Accounts) == 0 {
		awsSpec.Accounts = append(awsSpec.Accounts, Account{
			AccountName: defaultVar,
		})
	}

	for _, account := range awsSpec.Accounts {
		c.Logger().Debug().Str("account", account.AccountName).Msg("user defined account")

		localRegions := account.Regions
		if len(localRegions) == 0 {
			localRegions = awsSpec.Regions
		}

		if err := verifyRegions(localRegions); err != nil {
			return nil, err
		}

		if isAllRegions(localRegions) {
			c.Logger().Info().Msg("All regions specified in spec. Assuming all regions")
		}

		awsCfg, err := configureAwsClient(ctx, logger, awsConfig, account, adminAccountSts)
		if err != nil {
			var ae smithy.APIError
			if errors.As(err, &ae) {
				if strings.Contains(ae.ErrorCode(), "AccessDenied") {
					c.Logger().Error().Err(err).Msg("Access denied to account. Ensure that the account has access to be able perform `sts:AssumeRole` on the role")
					continue
				}
			}
			return nil, err
		}

		// This is a work-around to skip disabled regions
		// https://github.com/aws/aws-sdk-go-v2/issues/1068
		res, err := ec2.NewFromConfig(awsCfg).DescribeRegions(ctx,
			&ec2.DescribeRegionsInput{AllRegions: aws.Bool(false)},
			func(o *ec2.Options) {
				o.Region = defaultRegion
				if account.DefaultRegion != "" {
					o.Region = account.DefaultRegion
				}

				if len(localRegions) > 0 && !isAllRegions(localRegions) {
					o.Region = localRegions[0]
				}
			})
		if err != nil {
			c.Logger().Warn().Msg(fmt.Errorf("failed to find disabled regions for account %s. AWS Error: %w", account.AccountName, err).Error())
			continue
		}
		account.Regions = filterDisabledRegions(localRegions, res.Regions)

		if len(account.Regions) == 0 {
			c.Logger().Warn().Str("type", "access").Msg(fmt.Sprintf("no enabled regions provided in config for account %s", account.AccountName))
			continue
		}
		awsCfg.Region = account.Regions[0]
		output, err := getAccountId(ctx, awsCfg)
		if err != nil {
			c.Logger().Warn().Err(err).Msg("failed to get caller identity. AWS Error:")
			continue
		}
		iamArn, err := arn.Parse(*output.Arn)
		if err != nil {
			return nil, fmt.Errorf("failed to parse caller identity: %w", err)
		}

		for _, region := range account.Regions {
			c.ServicesManager.InitServicesForPartitionAccountAndRegion(iamArn.Partition, *output.Account, region, initServices(region, awsCfg))
		}
		c.ServicesManager.InitServicesForPartitionAccountAndScope(iamArn.Partition, *output.Account, initServices(cloudfrontScopeRegion, awsCfg))
	}
	if len(c.ServicesManager.services) == 0 {
		return nil, fmt.Errorf("no accounts instantiated")
	}
	return nil, nil
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
