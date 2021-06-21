// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cq-provider-aws/client (interfaces: EcsClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	ecs "github.com/aws/aws-sdk-go-v2/service/ecs"
	gomock "github.com/golang/mock/gomock"
)

// MockEcsClient is a mock of EcsClient interface.
type MockEcsClient struct {
	ctrl     *gomock.Controller
	recorder *MockEcsClientMockRecorder
}

// MockEcsClientMockRecorder is the mock recorder for MockEcsClient.
type MockEcsClientMockRecorder struct {
	mock *MockEcsClient
}

// NewMockEcsClient creates a new mock instance.
func NewMockEcsClient(ctrl *gomock.Controller) *MockEcsClient {
	mock := &MockEcsClient{ctrl: ctrl}
	mock.recorder = &MockEcsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEcsClient) EXPECT() *MockEcsClientMockRecorder {
	return m.recorder
}

// DescribeClusters mocks base method.
func (m *MockEcsClient) DescribeClusters(arg0 context.Context, arg1 *ecs.DescribeClustersInput, arg2 ...func(*ecs.Options)) (*ecs.DescribeClustersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeClusters", varargs...)
	ret0, _ := ret[0].(*ecs.DescribeClustersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeClusters indicates an expected call of DescribeClusters.
func (mr *MockEcsClientMockRecorder) DescribeClusters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeClusters", reflect.TypeOf((*MockEcsClient)(nil).DescribeClusters), varargs...)
}

// ListClusters mocks base method.
func (m *MockEcsClient) ListClusters(arg0 context.Context, arg1 *ecs.ListClustersInput, arg2 ...func(*ecs.Options)) (*ecs.ListClustersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListClusters", varargs...)
	ret0, _ := ret[0].(*ecs.ListClustersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListClusters indicates an expected call of ListClusters.
func (mr *MockEcsClientMockRecorder) ListClusters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListClusters", reflect.TypeOf((*MockEcsClient)(nil).ListClusters), varargs...)
}
