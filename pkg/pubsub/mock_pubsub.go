// Code generated by MockGen. DO NOT EDIT.
// Source: pubsub.go

// Package pubsub is a generated GoMock package.
package pubsub

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPubsub is a mock of Pubsub interface.
type MockPubsub struct {
	ctrl     *gomock.Controller
	recorder *MockPubsubMockRecorder
}

// MockPubsubMockRecorder is the mock recorder for MockPubsub.
type MockPubsubMockRecorder struct {
	mock *MockPubsub
}

// NewMockPubsub creates a new mock instance.
func NewMockPubsub(ctrl *gomock.Controller) *MockPubsub {
	mock := &MockPubsub{ctrl: ctrl}
	mock.recorder = &MockPubsubMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPubsub) EXPECT() *MockPubsubMockRecorder {
	return m.recorder
}

// ChName mocks base method.
func (m *MockPubsub) ChName(ctx context.Context) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChName", ctx)
	ret0, _ := ret[0].(string)
	return ret0
}

// ChName indicates an expected call of ChName.
func (mr *MockPubsubMockRecorder) ChName(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChName", reflect.TypeOf((*MockPubsub)(nil).ChName), ctx)
}

// Publish mocks base method.
func (m *MockPubsub) Publish(ctx context.Context, chName string, payload []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Publish", ctx, chName, payload)
	ret0, _ := ret[0].(error)
	return ret0
}

// Publish indicates an expected call of Publish.
func (mr *MockPubsubMockRecorder) Publish(ctx, chName, payload interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockPubsub)(nil).Publish), ctx, chName, payload)
}

// Receive mocks base method.
func (m *MockPubsub) Receive(ctx context.Context) ([][]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Receive", ctx)
	ret0, _ := ret[0].([][]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Receive indicates an expected call of Receive.
func (mr *MockPubsubMockRecorder) Receive(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Receive", reflect.TypeOf((*MockPubsub)(nil).Receive), ctx)
}

// Subscribe mocks base method.
func (m *MockPubsub) Subscribe(ctx context.Context, chName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscribe", ctx, chName)
	ret0, _ := ret[0].(error)
	return ret0
}

// Subscribe indicates an expected call of Subscribe.
func (mr *MockPubsubMockRecorder) Subscribe(ctx, chName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockPubsub)(nil).Subscribe), ctx, chName)
}
