// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/facades/controller/migrationtarget (interfaces: ControllerConfigService,ExternalControllerService,UpgradeService,ModelImporter,ModelManagerService)
//
// Generated by this command:
//
//	mockgen -package migrationtarget_test -destination domain_mock_test.go github.com/juju/juju/apiserver/facades/controller/migrationtarget ControllerConfigService,ExternalControllerService,UpgradeService,ModelImporter,ModelManagerService
//

// Package migrationtarget_test is a generated GoMock package.
package migrationtarget_test

import (
	context "context"
	reflect "reflect"

	controller "github.com/juju/juju/controller"
	crossmodel "github.com/juju/juju/core/crossmodel"
	model "github.com/juju/juju/domain/model"
	state "github.com/juju/juju/state"
	gomock "go.uber.org/mock/gomock"
)

// MockControllerConfigService is a mock of ControllerConfigService interface.
type MockControllerConfigService struct {
	ctrl     *gomock.Controller
	recorder *MockControllerConfigServiceMockRecorder
}

// MockControllerConfigServiceMockRecorder is the mock recorder for MockControllerConfigService.
type MockControllerConfigServiceMockRecorder struct {
	mock *MockControllerConfigService
}

// NewMockControllerConfigService creates a new mock instance.
func NewMockControllerConfigService(ctrl *gomock.Controller) *MockControllerConfigService {
	mock := &MockControllerConfigService{ctrl: ctrl}
	mock.recorder = &MockControllerConfigServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockControllerConfigService) EXPECT() *MockControllerConfigServiceMockRecorder {
	return m.recorder
}

// ControllerConfig mocks base method.
func (m *MockControllerConfigService) ControllerConfig(arg0 context.Context) (controller.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerConfig", arg0)
	ret0, _ := ret[0].(controller.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ControllerConfig indicates an expected call of ControllerConfig.
func (mr *MockControllerConfigServiceMockRecorder) ControllerConfig(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerConfig", reflect.TypeOf((*MockControllerConfigService)(nil).ControllerConfig), arg0)
}

// MockExternalControllerService is a mock of ExternalControllerService interface.
type MockExternalControllerService struct {
	ctrl     *gomock.Controller
	recorder *MockExternalControllerServiceMockRecorder
}

// MockExternalControllerServiceMockRecorder is the mock recorder for MockExternalControllerService.
type MockExternalControllerServiceMockRecorder struct {
	mock *MockExternalControllerService
}

// NewMockExternalControllerService creates a new mock instance.
func NewMockExternalControllerService(ctrl *gomock.Controller) *MockExternalControllerService {
	mock := &MockExternalControllerService{ctrl: ctrl}
	mock.recorder = &MockExternalControllerServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExternalControllerService) EXPECT() *MockExternalControllerServiceMockRecorder {
	return m.recorder
}

// ControllerForModel mocks base method.
func (m *MockExternalControllerService) ControllerForModel(arg0 context.Context, arg1 string) (*crossmodel.ControllerInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerForModel", arg0, arg1)
	ret0, _ := ret[0].(*crossmodel.ControllerInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ControllerForModel indicates an expected call of ControllerForModel.
func (mr *MockExternalControllerServiceMockRecorder) ControllerForModel(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerForModel", reflect.TypeOf((*MockExternalControllerService)(nil).ControllerForModel), arg0, arg1)
}

// UpdateExternalController mocks base method.
func (m *MockExternalControllerService) UpdateExternalController(arg0 context.Context, arg1 crossmodel.ControllerInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateExternalController", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateExternalController indicates an expected call of UpdateExternalController.
func (mr *MockExternalControllerServiceMockRecorder) UpdateExternalController(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateExternalController", reflect.TypeOf((*MockExternalControllerService)(nil).UpdateExternalController), arg0, arg1)
}

// MockUpgradeService is a mock of UpgradeService interface.
type MockUpgradeService struct {
	ctrl     *gomock.Controller
	recorder *MockUpgradeServiceMockRecorder
}

// MockUpgradeServiceMockRecorder is the mock recorder for MockUpgradeService.
type MockUpgradeServiceMockRecorder struct {
	mock *MockUpgradeService
}

// NewMockUpgradeService creates a new mock instance.
func NewMockUpgradeService(ctrl *gomock.Controller) *MockUpgradeService {
	mock := &MockUpgradeService{ctrl: ctrl}
	mock.recorder = &MockUpgradeServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUpgradeService) EXPECT() *MockUpgradeServiceMockRecorder {
	return m.recorder
}

// IsUpgrading mocks base method.
func (m *MockUpgradeService) IsUpgrading(arg0 context.Context) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsUpgrading", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsUpgrading indicates an expected call of IsUpgrading.
func (mr *MockUpgradeServiceMockRecorder) IsUpgrading(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsUpgrading", reflect.TypeOf((*MockUpgradeService)(nil).IsUpgrading), arg0)
}

// MockModelImporter is a mock of ModelImporter interface.
type MockModelImporter struct {
	ctrl     *gomock.Controller
	recorder *MockModelImporterMockRecorder
}

// MockModelImporterMockRecorder is the mock recorder for MockModelImporter.
type MockModelImporterMockRecorder struct {
	mock *MockModelImporter
}

// NewMockModelImporter creates a new mock instance.
func NewMockModelImporter(ctrl *gomock.Controller) *MockModelImporter {
	mock := &MockModelImporter{ctrl: ctrl}
	mock.recorder = &MockModelImporterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockModelImporter) EXPECT() *MockModelImporterMockRecorder {
	return m.recorder
}

// ImportModel mocks base method.
func (m *MockModelImporter) ImportModel(arg0 context.Context, arg1 []byte) (*state.Model, *state.State, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImportModel", arg0, arg1)
	ret0, _ := ret[0].(*state.Model)
	ret1, _ := ret[1].(*state.State)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ImportModel indicates an expected call of ImportModel.
func (mr *MockModelImporterMockRecorder) ImportModel(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImportModel", reflect.TypeOf((*MockModelImporter)(nil).ImportModel), arg0, arg1)
}

// MockModelManagerService is a mock of ModelManagerService interface.
type MockModelManagerService struct {
	ctrl     *gomock.Controller
	recorder *MockModelManagerServiceMockRecorder
}

// MockModelManagerServiceMockRecorder is the mock recorder for MockModelManagerService.
type MockModelManagerServiceMockRecorder struct {
	mock *MockModelManagerService
}

// NewMockModelManagerService creates a new mock instance.
func NewMockModelManagerService(ctrl *gomock.Controller) *MockModelManagerService {
	mock := &MockModelManagerService{ctrl: ctrl}
	mock.recorder = &MockModelManagerServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockModelManagerService) EXPECT() *MockModelManagerServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockModelManagerService) Create(arg0 context.Context, arg1 model.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockModelManagerServiceMockRecorder) Create(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockModelManagerService)(nil).Create), arg0, arg1)
}
