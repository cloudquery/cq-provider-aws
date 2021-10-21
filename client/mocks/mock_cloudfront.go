// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cq-provider-aws/client (interfaces: CloudfrontClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	cloudfront "github.com/aws/aws-sdk-go-v2/service/cloudfront"
	gomock "github.com/golang/mock/gomock"
)

// MockCloudfrontClient is a mock of CloudfrontClient interface.
type MockCloudfrontClient struct {
	ctrl     *gomock.Controller
	recorder *MockCloudfrontClientMockRecorder
}

// MockCloudfrontClientMockRecorder is the mock recorder for MockCloudfrontClient.
type MockCloudfrontClientMockRecorder struct {
	mock *MockCloudfrontClient
}

// NewMockCloudfrontClient creates a new mock instance.
func NewMockCloudfrontClient(ctrl *gomock.Controller) *MockCloudfrontClient {
	mock := &MockCloudfrontClient{ctrl: ctrl}
	mock.recorder = &MockCloudfrontClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCloudfrontClient) EXPECT() *MockCloudfrontClientMockRecorder {
	return m.recorder
}

// ListCachePolicies mocks base method.
func (m *MockCloudfrontClient) ListCachePolicies(arg0 context.Context, arg1 *cloudfront.ListCachePoliciesInput, arg2 ...func(*cloudfront.Options)) (*cloudfront.ListCachePoliciesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListCachePolicies", varargs...)
	ret0, _ := ret[0].(*cloudfront.ListCachePoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCachePolicies indicates an expected call of ListCachePolicies.
func (mr *MockCloudfrontClientMockRecorder) ListCachePolicies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCachePolicies", reflect.TypeOf((*MockCloudfrontClient)(nil).ListCachePolicies), varargs...)
}

// ListDistributions mocks base method.
func (m *MockCloudfrontClient) ListDistributions(arg0 context.Context, arg1 *cloudfront.ListDistributionsInput, arg2 ...func(*cloudfront.Options)) (*cloudfront.ListDistributionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListDistributions", varargs...)
	ret0, _ := ret[0].(*cloudfront.ListDistributionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDistributions indicates an expected call of ListDistributions.
func (mr *MockCloudfrontClientMockRecorder) ListDistributions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDistributions", reflect.TypeOf((*MockCloudfrontClient)(nil).ListDistributions), varargs...)
}

// ListTagsForResource mocks base method.
func (m *MockCloudfrontClient) ListTagsForResource(arg0 context.Context, arg1 *cloudfront.ListTagsForResourceInput, arg2 ...func(*cloudfront.Options)) (*cloudfront.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForResource", varargs...)
	ret0, _ := ret[0].(*cloudfront.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResource indicates an expected call of ListTagsForResource.
func (mr *MockCloudfrontClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockCloudfrontClient)(nil).ListTagsForResource), varargs...)
}
