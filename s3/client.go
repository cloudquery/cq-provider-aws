package s3

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/cloudquery/cloudquery/database"
	"github.com/cloudquery/cq-provider-aws/resource"
	"go.uber.org/zap"
)

type Client struct {
	db        *database.Database
	log       *zap.Logger
	accountID string
	region    string
	svc       *s3.Client
	awsConfig aws.Config
}

func NewClient(awsConfig aws.Config, db *database.Database, log *zap.Logger,
	accountID string, _ string) resource.ClientInterface {
	awsConfig.Region = "us-east-1"
	return &Client{
		db:        db,
		log:       log,
		accountID: accountID,
		region:    "us-east-1",
		svc:       s3.NewFromConfig(awsConfig),
		awsConfig: awsConfig,
	}
}

func (c *Client) CollectResource(resource string, config interface{}) error {
	switch resource {
	case "buckets":
		return c.buckets(config)
	default:
		return fmt.Errorf("unsupported resource buckets.%s", resource)
	}
}
