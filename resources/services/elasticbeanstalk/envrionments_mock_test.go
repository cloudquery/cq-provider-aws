package elasticbeanstalk

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
	elasticbeanstalkTypes "github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildElasticbeanstalkEnvironments(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockElasticbeanstalkClient(ctrl)

	la := elasticbeanstalkTypes.ApplicationDescription{}
	err := faker.FakeData(&la)
	if err != nil {
		t.Fatal(err)
	}

	// l := elasticbeanstalkTypes.EnvironmentDescription{
	// 	ApplicationName: la.ApplicationName,
	// }
	environmentDescriptionsList, err := faker.FakeDataNullablePermutations(elasticbeanstalkTypes.EnvironmentDescription{})
	if err != nil {
		t.Fatal(err)
	}
	describeEnvOutput := elasticbeanstalk.DescribeEnvironmentsOutput{
		Environments: environmentDescriptionsList.([]elasticbeanstalkTypes.EnvironmentDescription),
	}
	for i := range describeEnvOutput.Environments {
		describeEnvOutput.Environments[i].ApplicationName = la.ApplicationName
		describeEnvOutput.Environments[i].PlatformArn = aws.String(fmt.Sprintf("arn-%d", i))
		describeEnvOutput.Environments[i].EnvironmentArn = aws.String(fmt.Sprintf("arn-%d", i))
		describeEnvOutput.Environments[i].EnvironmentId = aws.String(fmt.Sprintf("id-%d", i))
	}

	m.EXPECT().DescribeEnvironments(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&describeEnvOutput, nil)

	tags := elasticbeanstalk.ListTagsForResourceOutput{}
	err = faker.FakeData(&tags)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).MinTimes(1).Return(
		&tags, nil)

	configSettingsOutput := elasticbeanstalk.DescribeConfigurationSettingsOutput{}
	err = faker.FakeData(&configSettingsOutput)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeConfigurationSettings(gomock.Any(), gomock.Any(), gomock.Any()).MinTimes(1).Return(&configSettingsOutput, nil)

	configOptsOutput := elasticbeanstalk.DescribeConfigurationOptionsOutput{}
	err = faker.FakeData(&configOptsOutput)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeConfigurationOptions(gomock.Any(), gomock.Any(), gomock.Any()).MinTimes(1).Return(&configOptsOutput, nil)

	return client.Services{
		ElasticBeanstalk: m,
	}
}

func TestElasticbeanstalkEnvironments(t *testing.T) {
	client.AwsMockTestHelper(t, ElasticbeanstalkEnvironments(), buildElasticbeanstalkEnvironments, client.TestOptions{})
}
