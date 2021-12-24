package resources

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/identitystore"

	"github.com/aws/aws-sdk-go-v2/service/ssoadmin"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildSsoAdminInstances(t *testing.T, ctrl *gomock.Controller) client.Services {
	sso := mocks.NewMockSSOAdminClient(ctrl)
	identityStore := mocks.NewMockIdentityStoreClient(ctrl)

	i := ssoadmin.ListInstancesOutput{}
	if err := faker.FakeData(&i); err != nil {
		t.Fatal(err)
	}
	i.NextToken = nil
	sso.EXPECT().ListInstances(gomock.Any(), gomock.Any(), gomock.Any()).Return(&i, nil)

	p := ssoadmin.ListPermissionSetsOutput{}
	if err := faker.FakeData(&p); err != nil {
		t.Fatal(err)
	}
	p.NextToken = nil
	sso.EXPECT().ListPermissionSets(gomock.Any(), gomock.Any(), gomock.Any()).Return(&p, nil)

	ps := ssoadmin.DescribePermissionSetOutput{}
	if err := faker.FakeData(&ps); err != nil {
		t.Fatal(err)
	}
	sso.EXPECT().DescribePermissionSet(gomock.Any(), gomock.Any(), gomock.Any()).Return(&ps, nil)

	tags := ssoadmin.ListTagsForResourceOutput{}
	if err := faker.FakeData(&tags); err != nil {
		t.Fatal(err)
	}
	tags.NextToken = nil
	sso.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(&tags, nil)

	aa := ssoadmin.ListAccountAssignmentsOutput{}
	if err := faker.FakeData(&aa); err != nil {
		t.Fatal(err)
	}
	aa.NextToken = nil
	sso.EXPECT().ListAccountAssignments(gomock.Any(), gomock.Any(), gomock.Any()).Return(&aa, nil)

	mp := ssoadmin.ListManagedPoliciesInPermissionSetOutput{}
	if err := faker.FakeData(&mp); err != nil {
		t.Fatal(err)
	}
	mp.NextToken = nil
	sso.EXPECT().ListManagedPoliciesInPermissionSet(gomock.Any(), gomock.Any(), gomock.Any()).Return(&mp, nil)

	policy := "{\"hello\":1}"
	sso.EXPECT().GetInlinePolicyForPermissionSet(gomock.Any(), gomock.Any(), gomock.Any()).Return(&ssoadmin.GetInlinePolicyForPermissionSetOutput{
		InlinePolicy: &policy,
	}, nil)

	g := identitystore.ListGroupsOutput{}
	if err := faker.FakeData(&g); err != nil {
		t.Fatal(err)
	}
	g.NextToken = nil
	identityStore.EXPECT().ListGroups(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&g, nil)

	u := identitystore.ListUsersOutput{}
	if err := faker.FakeData(&u); err != nil {
		t.Fatal(err)
	}
	u.NextToken = nil
	identityStore.EXPECT().ListUsers(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&u, nil)

	return client.Services{
		SSOAdmin:      sso,
		IdentityStore: identityStore,
	}
}

func TestSsoAdminInstances(t *testing.T) {
	awsTestHelper(t, SsoAdminInstances(), buildSsoAdminInstances, TestOptions{})
}
