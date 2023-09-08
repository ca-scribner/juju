// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/worker/servicefactory (interfaces: ControllerServiceFactory,ModelServiceFactory,ServiceFactory,ServiceFactoryGetter,Logger)

// Package servicefactory is a generated GoMock package.
package servicefactory

import (
	reflect "reflect"

	service "github.com/juju/juju/domain/autocert/service"
	service0 "github.com/juju/juju/domain/controllerconfig/service"
	service1 "github.com/juju/juju/domain/controllernode/service"
	service2 "github.com/juju/juju/domain/credential/service"
	service3 "github.com/juju/juju/domain/externalcontroller/service"
	service4 "github.com/juju/juju/domain/modelmanager/service"
	gomock "go.uber.org/mock/gomock"
)

// MockControllerServiceFactory is a mock of ControllerServiceFactory interface.
type MockControllerServiceFactory struct {
	ctrl     *gomock.Controller
	recorder *MockControllerServiceFactoryMockRecorder
}

// MockControllerServiceFactoryMockRecorder is the mock recorder for MockControllerServiceFactory.
type MockControllerServiceFactoryMockRecorder struct {
	mock *MockControllerServiceFactory
}

// NewMockControllerServiceFactory creates a new mock instance.
func NewMockControllerServiceFactory(ctrl *gomock.Controller) *MockControllerServiceFactory {
	mock := &MockControllerServiceFactory{ctrl: ctrl}
	mock.recorder = &MockControllerServiceFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockControllerServiceFactory) EXPECT() *MockControllerServiceFactoryMockRecorder {
	return m.recorder
}

// AutocertCache mocks base method.
func (m *MockControllerServiceFactory) AutocertCache() *service.Service {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AutocertCache")
	ret0, _ := ret[0].(*service.Service)
	return ret0
}

// AutocertCache indicates an expected call of AutocertCache.
func (mr *MockControllerServiceFactoryMockRecorder) AutocertCache() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AutocertCache", reflect.TypeOf((*MockControllerServiceFactory)(nil).AutocertCache))
}

// ControllerConfig mocks base method.
func (m *MockControllerServiceFactory) ControllerConfig() *service0.Service {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerConfig")
	ret0, _ := ret[0].(*service0.Service)
	return ret0
}

// ControllerConfig indicates an expected call of ControllerConfig.
func (mr *MockControllerServiceFactoryMockRecorder) ControllerConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerConfig", reflect.TypeOf((*MockControllerServiceFactory)(nil).ControllerConfig))
}

// ControllerNode mocks base method.
func (m *MockControllerServiceFactory) ControllerNode() *service1.Service {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerNode")
	ret0, _ := ret[0].(*service1.Service)
	return ret0
}

// ControllerNode indicates an expected call of ControllerNode.
func (mr *MockControllerServiceFactoryMockRecorder) ControllerNode() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerNode", reflect.TypeOf((*MockControllerServiceFactory)(nil).ControllerNode))
}

// Credential mocks base method.
func (m *MockControllerServiceFactory) Credential() *service2.Service {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Credential")
	ret0, _ := ret[0].(*service2.Service)
	return ret0
}

// Credential indicates an expected call of Credential.
func (mr *MockControllerServiceFactoryMockRecorder) Credential() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Credential", reflect.TypeOf((*MockControllerServiceFactory)(nil).Credential))
}

// ExternalController mocks base method.
func (m *MockControllerServiceFactory) ExternalController() *service3.Service {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExternalController")
	ret0, _ := ret[0].(*service3.Service)
	return ret0
}

// ExternalController indicates an expected call of ExternalController.
func (mr *MockControllerServiceFactoryMockRecorder) ExternalController() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExternalController", reflect.TypeOf((*MockControllerServiceFactory)(nil).ExternalController))
}

// ModelManager mocks base method.
func (m *MockControllerServiceFactory) ModelManager() *service4.Service {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelManager")
	ret0, _ := ret[0].(*service4.Service)
	return ret0
}

// ModelManager indicates an expected call of ModelManager.
func (mr *MockControllerServiceFactoryMockRecorder) ModelManager() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelManager", reflect.TypeOf((*MockControllerServiceFactory)(nil).ModelManager))
}

// MockModelServiceFactory is a mock of ModelServiceFactory interface.
type MockModelServiceFactory struct {
	ctrl     *gomock.Controller
	recorder *MockModelServiceFactoryMockRecorder
}

// MockModelServiceFactoryMockRecorder is the mock recorder for MockModelServiceFactory.
type MockModelServiceFactoryMockRecorder struct {
	mock *MockModelServiceFactory
}

// NewMockModelServiceFactory creates a new mock instance.
func NewMockModelServiceFactory(ctrl *gomock.Controller) *MockModelServiceFactory {
	mock := &MockModelServiceFactory{ctrl: ctrl}
	mock.recorder = &MockModelServiceFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockModelServiceFactory) EXPECT() *MockModelServiceFactoryMockRecorder {
	return m.recorder
}

// Name mocks base method.
func (m *MockModelServiceFactory) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockModelServiceFactoryMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockModelServiceFactory)(nil).Name))
}

// MockServiceFactory is a mock of ServiceFactory interface.
type MockServiceFactory struct {
	ctrl     *gomock.Controller
	recorder *MockServiceFactoryMockRecorder
}

// MockServiceFactoryMockRecorder is the mock recorder for MockServiceFactory.
type MockServiceFactoryMockRecorder struct {
	mock *MockServiceFactory
}

// NewMockServiceFactory creates a new mock instance.
func NewMockServiceFactory(ctrl *gomock.Controller) *MockServiceFactory {
	mock := &MockServiceFactory{ctrl: ctrl}
	mock.recorder = &MockServiceFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceFactory) EXPECT() *MockServiceFactoryMockRecorder {
	return m.recorder
}

// AutocertCache mocks base method.
func (m *MockServiceFactory) AutocertCache() *service.Service {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AutocertCache")
	ret0, _ := ret[0].(*service.Service)
	return ret0
}

// AutocertCache indicates an expected call of AutocertCache.
func (mr *MockServiceFactoryMockRecorder) AutocertCache() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AutocertCache", reflect.TypeOf((*MockServiceFactory)(nil).AutocertCache))
}

// ControllerConfig mocks base method.
func (m *MockServiceFactory) ControllerConfig() *service0.Service {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerConfig")
	ret0, _ := ret[0].(*service0.Service)
	return ret0
}

// ControllerConfig indicates an expected call of ControllerConfig.
func (mr *MockServiceFactoryMockRecorder) ControllerConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerConfig", reflect.TypeOf((*MockServiceFactory)(nil).ControllerConfig))
}

// ControllerNode mocks base method.
func (m *MockServiceFactory) ControllerNode() *service1.Service {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerNode")
	ret0, _ := ret[0].(*service1.Service)
	return ret0
}

// ControllerNode indicates an expected call of ControllerNode.
func (mr *MockServiceFactoryMockRecorder) ControllerNode() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerNode", reflect.TypeOf((*MockServiceFactory)(nil).ControllerNode))
}

// Credential mocks base method.
func (m *MockServiceFactory) Credential() *service2.Service {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Credential")
	ret0, _ := ret[0].(*service2.Service)
	return ret0
}

// Credential indicates an expected call of Credential.
func (mr *MockServiceFactoryMockRecorder) Credential() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Credential", reflect.TypeOf((*MockServiceFactory)(nil).Credential))
}

// ExternalController mocks base method.
func (m *MockServiceFactory) ExternalController() *service3.Service {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExternalController")
	ret0, _ := ret[0].(*service3.Service)
	return ret0
}

// ExternalController indicates an expected call of ExternalController.
func (mr *MockServiceFactoryMockRecorder) ExternalController() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExternalController", reflect.TypeOf((*MockServiceFactory)(nil).ExternalController))
}

// ModelManager mocks base method.
func (m *MockServiceFactory) ModelManager() *service4.Service {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelManager")
	ret0, _ := ret[0].(*service4.Service)
	return ret0
}

// ModelManager indicates an expected call of ModelManager.
func (mr *MockServiceFactoryMockRecorder) ModelManager() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelManager", reflect.TypeOf((*MockServiceFactory)(nil).ModelManager))
}

// Name mocks base method.
func (m *MockServiceFactory) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockServiceFactoryMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockServiceFactory)(nil).Name))
}

// MockServiceFactoryGetter is a mock of ServiceFactoryGetter interface.
type MockServiceFactoryGetter struct {
	ctrl     *gomock.Controller
	recorder *MockServiceFactoryGetterMockRecorder
}

// MockServiceFactoryGetterMockRecorder is the mock recorder for MockServiceFactoryGetter.
type MockServiceFactoryGetterMockRecorder struct {
	mock *MockServiceFactoryGetter
}

// NewMockServiceFactoryGetter creates a new mock instance.
func NewMockServiceFactoryGetter(ctrl *gomock.Controller) *MockServiceFactoryGetter {
	mock := &MockServiceFactoryGetter{ctrl: ctrl}
	mock.recorder = &MockServiceFactoryGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceFactoryGetter) EXPECT() *MockServiceFactoryGetterMockRecorder {
	return m.recorder
}

// FactoryForModel mocks base method.
func (m *MockServiceFactoryGetter) FactoryForModel(arg0 string) ServiceFactory {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FactoryForModel", arg0)
	ret0, _ := ret[0].(ServiceFactory)
	return ret0
}

// FactoryForModel indicates an expected call of FactoryForModel.
func (mr *MockServiceFactoryGetterMockRecorder) FactoryForModel(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FactoryForModel", reflect.TypeOf((*MockServiceFactoryGetter)(nil).FactoryForModel), arg0)
}

// MockLogger is a mock of Logger interface.
type MockLogger struct {
	ctrl     *gomock.Controller
	recorder *MockLoggerMockRecorder
}

// MockLoggerMockRecorder is the mock recorder for MockLogger.
type MockLoggerMockRecorder struct {
	mock *MockLogger
}

// NewMockLogger creates a new mock instance.
func NewMockLogger(ctrl *gomock.Controller) *MockLogger {
	mock := &MockLogger{ctrl: ctrl}
	mock.recorder = &MockLoggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogger) EXPECT() *MockLoggerMockRecorder {
	return m.recorder
}

// Child mocks base method.
func (m *MockLogger) Child(arg0 string) Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Child", arg0)
	ret0, _ := ret[0].(Logger)
	return ret0
}

// Child indicates an expected call of Child.
func (mr *MockLoggerMockRecorder) Child(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Child", reflect.TypeOf((*MockLogger)(nil).Child), arg0)
}

// Debugf mocks base method.
func (m *MockLogger) Debugf(arg0 string, arg1 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debugf", varargs...)
}

// Debugf indicates an expected call of Debugf.
func (mr *MockLoggerMockRecorder) Debugf(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debugf", reflect.TypeOf((*MockLogger)(nil).Debugf), varargs...)
}

// Tracef mocks base method.
func (m *MockLogger) Tracef(arg0 string, arg1 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Tracef", varargs...)
}

// Tracef indicates an expected call of Tracef.
func (mr *MockLoggerMockRecorder) Tracef(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tracef", reflect.TypeOf((*MockLogger)(nil).Tracef), varargs...)
}

// Warningf mocks base method.
func (m *MockLogger) Warningf(arg0 string, arg1 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Warningf", varargs...)
}

// Warningf indicates an expected call of Warningf.
func (mr *MockLoggerMockRecorder) Warningf(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warningf", reflect.TypeOf((*MockLogger)(nil).Warningf), varargs...)
}
