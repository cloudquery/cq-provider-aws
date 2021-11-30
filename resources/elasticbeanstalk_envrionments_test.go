package resources

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
	elasticbeanstalkTypes "github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildElasticbeanstalkEnvironments(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockElasticbeanstalkClient(ctrl)
	l := elasticbeanstalkTypes.EnvironmentDescription{}
	err := faker.FakeData(&l)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeEnvironments(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&elasticbeanstalk.DescribeEnvironmentsOutput{
			Environments: []elasticbeanstalkTypes.EnvironmentDescription{l},
		}, nil)

	tags := elasticbeanstalk.ListTagsForResourceOutput{}
	err = faker.FakeData(&tags)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&tags, nil)
	return client.Services{
		ElasticBeanstalk: m,
	}
}

func TestElasticbeanstalkApplications(t *testing.T) {
	awsTestHelper(t, ElasticbeanstalkApplications(), buildElasticbeanstalkEnvironments, TestOptions{})
}
