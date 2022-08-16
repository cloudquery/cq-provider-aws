package fsx

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/fsx"
	"github.com/aws/aws-sdk-go-v2/service/fsx/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildFSXClientForBackups(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockFsxClient(ctrl)

	var b types.Backup
	if err := faker.FakeData(&b); err != nil {
		t.Fatal(err)
	}

	mock.EXPECT().DescribeBackups(
		gomock.Any(),
		&fsx.DescribeBackupsInput{},
		gomock.Any(),
	).Return(
		&fsx.DescribeBackupsOutput{Backups: []types.Backup{b}},
		nil,
	)

	return client.Services{FSX: mock}
}

func TestFSXBackups(t *testing.T) {
	client.AwsMockTestHelper(t, FsxBackups(), buildFSXClientForBackups, client.TestOptions{})
}
