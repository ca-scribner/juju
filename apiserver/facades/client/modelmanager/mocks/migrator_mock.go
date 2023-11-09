// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/facades/client/modelmanager (interfaces: ModelExporter)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	description "github.com/juju/description/v4"
	objectstore "github.com/juju/juju/core/objectstore"
	state "github.com/juju/juju/state"
	gomock "go.uber.org/mock/gomock"
)

// MockModelExporter is a mock of ModelExporter interface.
type MockModelExporter struct {
	ctrl     *gomock.Controller
	recorder *MockModelExporterMockRecorder
}

// MockModelExporterMockRecorder is the mock recorder for MockModelExporter.
type MockModelExporterMockRecorder struct {
	mock *MockModelExporter
}

// NewMockModelExporter creates a new mock instance.
func NewMockModelExporter(ctrl *gomock.Controller) *MockModelExporter {
	mock := &MockModelExporter{ctrl: ctrl}
	mock.recorder = &MockModelExporterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockModelExporter) EXPECT() *MockModelExporterMockRecorder {
	return m.recorder
}

// ExportModelPartial mocks base method.
func (m *MockModelExporter) ExportModelPartial(arg0 context.Context, arg1 state.ExportConfig, arg2 objectstore.ObjectStore) (description.Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExportModelPartial", arg0, arg1, arg2)
	ret0, _ := ret[0].(description.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExportModelPartial indicates an expected call of ExportModelPartial.
func (mr *MockModelExporterMockRecorder) ExportModelPartial(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExportModelPartial", reflect.TypeOf((*MockModelExporter)(nil).ExportModelPartial), arg0, arg1, arg2)
}
