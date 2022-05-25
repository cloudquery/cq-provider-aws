package iam

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	iamTypes "github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildIamGroupsAccessAdvisors(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIamClient(ctrl)
	g := iamTypes.Group{}
	err := faker.FakeData(&g)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListGroups(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&iam.ListGroupsOutput{
			Groups: []iamTypes.Group{g},
		}, nil)
	gad := iam.GenerateServiceLastAccessedDetailsOutput{}
	err = faker.FakeData(&gad)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GenerateServiceLastAccessedDetails(gomock.Any(), gomock.Any(), gomock.Any()).Return(&gad, nil)

	ad := iam.GetServiceLastAccessedDetailsOutput{}
	err = faker.FakeData(&ad)
	if err != nil {
		t.Fatal(err)
	}
	ad.ServicesLastAccessed[0].TotalAuthenticatedEntities = aws.Int32(2)
	ad.JobStatus = iamTypes.JobStatusTypeCompleted
	ad.Marker = nil
	m.EXPECT().GetServiceLastAccessedDetails(gomock.Any(), gomock.Any(), gomock.Any()).Return(&ad, nil)

	ade := iam.GetServiceLastAccessedDetailsWithEntitiesOutput{}
	err = faker.FakeData(&ade)
	ade.Marker = nil
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetServiceLastAccessedDetailsWithEntities(gomock.Any(), gomock.Any(), gomock.Any()).Return(&ade, nil)

	return client.Services{
		IAM: m,
	}
}

func TestIamGroupsAccessAdvisors(t *testing.T) {
	client.AwsMockTestHelper(t, GroupsAccessAdvisorDetails(), buildIamGroupsAccessAdvisors, client.TestOptions{NotParallel: true})
}
