// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/facades/client/machinemanager (interfaces: ControllerBackend,InstanceConfigBackend)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	machinemanager "github.com/juju/juju/apiserver/facades/client/machinemanager"
	controller "github.com/juju/juju/controller"
	network "github.com/juju/juju/core/network"
	binarystorage "github.com/juju/juju/state/binarystorage"
	names "github.com/juju/names/v5"
	gomock "go.uber.org/mock/gomock"
)

// MockControllerBackend is a mock of ControllerBackend interface.
type MockControllerBackend struct {
	ctrl     *gomock.Controller
	recorder *MockControllerBackendMockRecorder
}

// MockControllerBackendMockRecorder is the mock recorder for MockControllerBackend.
type MockControllerBackendMockRecorder struct {
	mock *MockControllerBackend
}

// NewMockControllerBackend creates a new mock instance.
func NewMockControllerBackend(ctrl *gomock.Controller) *MockControllerBackend {
	mock := &MockControllerBackend{ctrl: ctrl}
	mock.recorder = &MockControllerBackendMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockControllerBackend) EXPECT() *MockControllerBackendMockRecorder {
	return m.recorder
}

// APIHostPortsForAgents mocks base method.
func (m *MockControllerBackend) APIHostPortsForAgents() ([]network.SpaceHostPorts, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "APIHostPortsForAgents")
	ret0, _ := ret[0].([]network.SpaceHostPorts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// APIHostPortsForAgents indicates an expected call of APIHostPortsForAgents.
func (mr *MockControllerBackendMockRecorder) APIHostPortsForAgents() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "APIHostPortsForAgents", reflect.TypeOf((*MockControllerBackend)(nil).APIHostPortsForAgents))
}

// ControllerConfig mocks base method.
func (m *MockControllerBackend) ControllerConfig() (controller.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerConfig")
	ret0, _ := ret[0].(controller.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ControllerConfig indicates an expected call of ControllerConfig.
func (mr *MockControllerBackendMockRecorder) ControllerConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerConfig", reflect.TypeOf((*MockControllerBackend)(nil).ControllerConfig))
}

// ControllerTag mocks base method.
func (m *MockControllerBackend) ControllerTag() names.ControllerTag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerTag")
	ret0, _ := ret[0].(names.ControllerTag)
	return ret0
}

// ControllerTag indicates an expected call of ControllerTag.
func (mr *MockControllerBackendMockRecorder) ControllerTag() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerTag", reflect.TypeOf((*MockControllerBackend)(nil).ControllerTag))
}

// MockInstanceConfigBackend is a mock of InstanceConfigBackend interface.
type MockInstanceConfigBackend struct {
	ctrl     *gomock.Controller
	recorder *MockInstanceConfigBackendMockRecorder
}

// MockInstanceConfigBackendMockRecorder is the mock recorder for MockInstanceConfigBackend.
type MockInstanceConfigBackendMockRecorder struct {
	mock *MockInstanceConfigBackend
}

// NewMockInstanceConfigBackend creates a new mock instance.
func NewMockInstanceConfigBackend(ctrl *gomock.Controller) *MockInstanceConfigBackend {
	mock := &MockInstanceConfigBackend{ctrl: ctrl}
	mock.recorder = &MockInstanceConfigBackendMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInstanceConfigBackend) EXPECT() *MockInstanceConfigBackendMockRecorder {
	return m.recorder
}

// Machine mocks base method.
func (m *MockInstanceConfigBackend) Machine(arg0 string) (machinemanager.Machine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Machine", arg0)
	ret0, _ := ret[0].(machinemanager.Machine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Machine indicates an expected call of Machine.
func (mr *MockInstanceConfigBackendMockRecorder) Machine(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Machine", reflect.TypeOf((*MockInstanceConfigBackend)(nil).Machine), arg0)
}

// Model mocks base method.
func (m *MockInstanceConfigBackend) Model() (machinemanager.Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Model")
	ret0, _ := ret[0].(machinemanager.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Model indicates an expected call of Model.
func (mr *MockInstanceConfigBackendMockRecorder) Model() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Model", reflect.TypeOf((*MockInstanceConfigBackend)(nil).Model))
}

// ToolsStorage mocks base method.
func (m *MockInstanceConfigBackend) ToolsStorage() (binarystorage.StorageCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToolsStorage")
	ret0, _ := ret[0].(binarystorage.StorageCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ToolsStorage indicates an expected call of ToolsStorage.
func (mr *MockInstanceConfigBackendMockRecorder) ToolsStorage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToolsStorage", reflect.TypeOf((*MockInstanceConfigBackend)(nil).ToolsStorage))
}
