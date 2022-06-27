// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/cmd/juju/commands (interfaces: ModelManagerAPI)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	names "github.com/juju/names/v4"
	version "github.com/juju/version/v2"
)

// MockModelManagerAPI is a mock of ModelManagerAPI interface.
type MockModelManagerAPI struct {
	ctrl     *gomock.Controller
	recorder *MockModelManagerAPIMockRecorder
}

// MockModelManagerAPIMockRecorder is the mock recorder for MockModelManagerAPI.
type MockModelManagerAPIMockRecorder struct {
	mock *MockModelManagerAPI
}

// NewMockModelManagerAPI creates a new mock instance.
func NewMockModelManagerAPI(ctrl *gomock.Controller) *MockModelManagerAPI {
	mock := &MockModelManagerAPI{ctrl: ctrl}
	mock.recorder = &MockModelManagerAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockModelManagerAPI) EXPECT() *MockModelManagerAPIMockRecorder {
	return m.recorder
}

// AbortCurrentUpgrade mocks base method.
func (m *MockModelManagerAPI) AbortCurrentUpgrade() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AbortCurrentUpgrade")
	ret0, _ := ret[0].(error)
	return ret0
}

// AbortCurrentUpgrade indicates an expected call of AbortCurrentUpgrade.
func (mr *MockModelManagerAPIMockRecorder) AbortCurrentUpgrade() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AbortCurrentUpgrade", reflect.TypeOf((*MockModelManagerAPI)(nil).AbortCurrentUpgrade))
}

// BestAPIVersion mocks base method.
func (m *MockModelManagerAPI) BestAPIVersion() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BestAPIVersion")
	ret0, _ := ret[0].(int)
	return ret0
}

// BestAPIVersion indicates an expected call of BestAPIVersion.
func (mr *MockModelManagerAPIMockRecorder) BestAPIVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BestAPIVersion", reflect.TypeOf((*MockModelManagerAPI)(nil).BestAPIVersion))
}

// Close mocks base method.
func (m *MockModelManagerAPI) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockModelManagerAPIMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockModelManagerAPI)(nil).Close))
}

// UpgradeModel mocks base method.
func (m *MockModelManagerAPI) UpgradeModel(arg0 names.ModelTag, arg1 version.Number, arg2 string, arg3, arg4 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpgradeModel", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpgradeModel indicates an expected call of UpgradeModel.
func (mr *MockModelManagerAPIMockRecorder) UpgradeModel(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpgradeModel", reflect.TypeOf((*MockModelManagerAPI)(nil).UpgradeModel), arg0, arg1, arg2, arg3, arg4)
}

// ValidateModelUpgrade mocks base method.
func (m *MockModelManagerAPI) ValidateModelUpgrade(arg0 names.ModelTag, arg1 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateModelUpgrade", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateModelUpgrade indicates an expected call of ValidateModelUpgrade.
func (mr *MockModelManagerAPIMockRecorder) ValidateModelUpgrade(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateModelUpgrade", reflect.TypeOf((*MockModelManagerAPI)(nil).ValidateModelUpgrade), arg0, arg1)
}
