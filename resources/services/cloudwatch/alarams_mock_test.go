package cloudwatch

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildCloudWatchAlarmsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockCloudwatchClient(ctrl)
	services := client.Services{
		Cloudwatch: m,
	}

	alarmList, err := faker.FakeDataNullablePermutations(types.MetricAlarm{})
	if err != nil {
		t.Fatal(err)
	}
	describeAlarmsOutput := cloudwatch.DescribeAlarmsOutput{
		MetricAlarms: alarmList.([]types.MetricAlarm),
	}
	for i := range describeAlarmsOutput.MetricAlarms {
		describeAlarmsOutput.MetricAlarms[i].AlarmArn = aws.String(fmt.Sprintf("arn:%d", i))
	}
	m.EXPECT().DescribeAlarms(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&describeAlarmsOutput, nil)
	return services
}

func TestCloudwatchAlarms(t *testing.T) {
	client.AwsMockTestHelper(t, CloudwatchAlarms(), buildCloudWatchAlarmsMock, client.TestOptions{})
}
