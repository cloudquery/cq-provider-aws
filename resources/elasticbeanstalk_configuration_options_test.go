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

func buildElasticbeanstalkConfigOptions(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockElasticbeanstalkClient(ctrl)

	la := elasticbeanstalkTypes.ApplicationDescription{}
	err := faker.FakeData(&la)
	if err != nil {
		t.Fatal(err)
	}

	// m.EXPECT().DescribeApplications(gomock.Any(), gomock.Any(), gomock.Any()).Return(
	// 	&elasticbeanstalk.DescribeApplicationsOutput{
	// 		Applications: []elasticbeanstalkTypes.ApplicationDescription{la},
	// 	}, nil)

	l := elasticbeanstalkTypes.EnvironmentDescription{
		ApplicationName: la.ApplicationName,
	}
	err = faker.FakeData(&l)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeEnvironments(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&elasticbeanstalk.DescribeEnvironmentsOutput{
			Environments: []elasticbeanstalkTypes.EnvironmentDescription{l},
		}, nil)

	configOutput := elasticbeanstalk.DescribeConfigurationOptionsOutput{}
	err = faker.FakeData(&configOutput)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeConfigurationOptions(gomock.Any(), gomock.Any(), gomock.Any()).Return(&configOutput, nil)

	return client.Services{
		ElasticBeanstalk: m,
	}
}

func TestElasticbeanstalkConfigOptions(t *testing.T) {
	awsTestHelper(t, ElasticbeanstalkConfigurationOptions(), buildElasticbeanstalkConfigOptions, TestOptions{})
}
