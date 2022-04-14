package ssm

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildAssociationsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockSSMClient(ctrl)

	var association types.AssociationDescription
	if err := faker.FakeData(&association); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListAssociations(
		gomock.Any(),
		&ssm.ListAssociationsInput{},
		gomock.Any(),
	).Return(
		&ssm.ListAssociationsOutput{Associations: []types.Association{
			{
				AssociationId:      association.AssociationId,
				AssociationVersion: association.AssociationVersion,
			},
		}},
		nil,
	)

	mock.EXPECT().DescribeAssociation(
		gomock.Any(),
		&ssm.DescribeAssociationInput{
			AssociationId:      association.AssociationId,
			AssociationVersion: association.AssociationVersion,
		},
		gomock.Any(),
	).Return(
		&ssm.DescribeAssociationOutput{AssociationDescription: &association},
		nil,
	)

	return client.Services{SSM: mock}
}

func TestAssociations(t *testing.T) {
	client.AwsMockTestHelper(t, Associations(), buildAssociationsMock, client.TestOptions{})
}
