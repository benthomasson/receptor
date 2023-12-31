// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/workceptor/workceptor.go

// Package mock_workceptor is a generated GoMock package.
package mock_workceptor

import (
	context "context"
	tls "crypto/tls"
	logger "github.com/ansible/receptor/pkg/logger"
	netceptor "github.com/ansible/receptor/pkg/netceptor"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockNetceptorForWorkceptor is a mock of NetceptorForWorkceptor interface
type MockNetceptorForWorkceptor struct {
	ctrl     *gomock.Controller
	recorder *MockNetceptorForWorkceptorMockRecorder
}

// MockNetceptorForWorkceptorMockRecorder is the mock recorder for MockNetceptorForWorkceptor
type MockNetceptorForWorkceptorMockRecorder struct {
	mock *MockNetceptorForWorkceptor
}

// NewMockNetceptorForWorkceptor creates a new mock instance
func NewMockNetceptorForWorkceptor(ctrl *gomock.Controller) *MockNetceptorForWorkceptor {
	mock := &MockNetceptorForWorkceptor{ctrl: ctrl}
	mock.recorder = &MockNetceptorForWorkceptorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNetceptorForWorkceptor) EXPECT() *MockNetceptorForWorkceptorMockRecorder {
	return m.recorder
}

// NodeID mocks base method
func (m *MockNetceptorForWorkceptor) NodeID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NodeID")
	ret0, _ := ret[0].(string)
	return ret0
}

// NodeID indicates an expected call of NodeID
func (mr *MockNetceptorForWorkceptorMockRecorder) NodeID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NodeID", reflect.TypeOf((*MockNetceptorForWorkceptor)(nil).NodeID))
}

// AddWorkCommand mocks base method
func (m *MockNetceptorForWorkceptor) AddWorkCommand(typeName string, verifySignature bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddWorkCommand", typeName, verifySignature)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddWorkCommand indicates an expected call of AddWorkCommand
func (mr *MockNetceptorForWorkceptorMockRecorder) AddWorkCommand(typeName, verifySignature interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddWorkCommand", reflect.TypeOf((*MockNetceptorForWorkceptor)(nil).AddWorkCommand), typeName, verifySignature)
}

// GetClientTLSConfig mocks base method
func (m *MockNetceptorForWorkceptor) GetClientTLSConfig(name, expectedHostName string, expectedHostNameType netceptor.ExpectedHostnameType) (*tls.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClientTLSConfig", name, expectedHostName, expectedHostNameType)
	ret0, _ := ret[0].(*tls.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClientTLSConfig indicates an expected call of GetClientTLSConfig
func (mr *MockNetceptorForWorkceptorMockRecorder) GetClientTLSConfig(name, expectedHostName, expectedHostNameType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClientTLSConfig", reflect.TypeOf((*MockNetceptorForWorkceptor)(nil).GetClientTLSConfig), name, expectedHostName, expectedHostNameType)
}

// GetLogger mocks base method
func (m *MockNetceptorForWorkceptor) GetLogger() *logger.ReceptorLogger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLogger")
	ret0, _ := ret[0].(*logger.ReceptorLogger)
	return ret0
}

// GetLogger indicates an expected call of GetLogger
func (mr *MockNetceptorForWorkceptorMockRecorder) GetLogger() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogger", reflect.TypeOf((*MockNetceptorForWorkceptor)(nil).GetLogger))
}

// DialContext mocks base method
func (m *MockNetceptorForWorkceptor) DialContext(ctx context.Context, node, service string, tlscfg *tls.Config) (*netceptor.Conn, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DialContext", ctx, node, service, tlscfg)
	ret0, _ := ret[0].(*netceptor.Conn)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DialContext indicates an expected call of DialContext
func (mr *MockNetceptorForWorkceptorMockRecorder) DialContext(ctx, node, service, tlscfg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DialContext", reflect.TypeOf((*MockNetceptorForWorkceptor)(nil).DialContext), ctx, node, service, tlscfg)
}
