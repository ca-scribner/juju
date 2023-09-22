// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/facades/client/client (interfaces: Backend,Model)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	time "time"

	charm "github.com/juju/charm/v11"
	client "github.com/juju/juju/apiserver/facades/client/client"
	controller "github.com/juju/juju/controller"
	constraints "github.com/juju/juju/core/constraints"
	crossmodel "github.com/juju/juju/core/crossmodel"
	instance "github.com/juju/juju/core/instance"
	network "github.com/juju/juju/core/network"
	permission "github.com/juju/juju/core/permission"
	status "github.com/juju/juju/core/status"
	config "github.com/juju/juju/environs/config"
	state "github.com/juju/juju/state"
	names "github.com/juju/names/v4"
	version "github.com/juju/version/v2"
	gomock "go.uber.org/mock/gomock"
)

// MockBackend is a mock of Backend interface.
type MockBackend struct {
	ctrl     *gomock.Controller
	recorder *MockBackendMockRecorder
}

// MockBackendMockRecorder is the mock recorder for MockBackend.
type MockBackendMockRecorder struct {
	mock *MockBackend
}

// NewMockBackend creates a new mock instance.
func NewMockBackend(ctrl *gomock.Controller) *MockBackend {
	mock := &MockBackend{ctrl: ctrl}
	mock.recorder = &MockBackendMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBackend) EXPECT() *MockBackendMockRecorder {
	return m.recorder
}

// APIHostPortsForClients mocks base method.
func (m *MockBackend) APIHostPortsForClients() ([]network.SpaceHostPorts, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "APIHostPortsForClients")
	ret0, _ := ret[0].([]network.SpaceHostPorts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// APIHostPortsForClients indicates an expected call of APIHostPortsForClients.
func (mr *MockBackendMockRecorder) APIHostPortsForClients() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "APIHostPortsForClients", reflect.TypeOf((*MockBackend)(nil).APIHostPortsForClients))
}

// AbortCurrentUpgrade mocks base method.
func (m *MockBackend) AbortCurrentUpgrade() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AbortCurrentUpgrade")
	ret0, _ := ret[0].(error)
	return ret0
}

// AbortCurrentUpgrade indicates an expected call of AbortCurrentUpgrade.
func (mr *MockBackendMockRecorder) AbortCurrentUpgrade() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AbortCurrentUpgrade", reflect.TypeOf((*MockBackend)(nil).AbortCurrentUpgrade))
}

// AddControllerUser mocks base method.
func (m *MockBackend) AddControllerUser(arg0 state.UserAccessSpec) (permission.UserAccess, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddControllerUser", arg0)
	ret0, _ := ret[0].(permission.UserAccess)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddControllerUser indicates an expected call of AddControllerUser.
func (mr *MockBackendMockRecorder) AddControllerUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddControllerUser", reflect.TypeOf((*MockBackend)(nil).AddControllerUser), arg0)
}

// AddMachineInsideMachine mocks base method.
func (m *MockBackend) AddMachineInsideMachine(arg0 state.MachineTemplate, arg1 string, arg2 instance.ContainerType) (*state.Machine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddMachineInsideMachine", arg0, arg1, arg2)
	ret0, _ := ret[0].(*state.Machine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddMachineInsideMachine indicates an expected call of AddMachineInsideMachine.
func (mr *MockBackendMockRecorder) AddMachineInsideMachine(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMachineInsideMachine", reflect.TypeOf((*MockBackend)(nil).AddMachineInsideMachine), arg0, arg1, arg2)
}

// AddMachineInsideNewMachine mocks base method.
func (m *MockBackend) AddMachineInsideNewMachine(arg0, arg1 state.MachineTemplate, arg2 instance.ContainerType) (*state.Machine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddMachineInsideNewMachine", arg0, arg1, arg2)
	ret0, _ := ret[0].(*state.Machine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddMachineInsideNewMachine indicates an expected call of AddMachineInsideNewMachine.
func (mr *MockBackendMockRecorder) AddMachineInsideNewMachine(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMachineInsideNewMachine", reflect.TypeOf((*MockBackend)(nil).AddMachineInsideNewMachine), arg0, arg1, arg2)
}

// AddOneMachine mocks base method.
func (m *MockBackend) AddOneMachine(arg0 state.MachineTemplate) (*state.Machine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddOneMachine", arg0)
	ret0, _ := ret[0].(*state.Machine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddOneMachine indicates an expected call of AddOneMachine.
func (mr *MockBackendMockRecorder) AddOneMachine(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddOneMachine", reflect.TypeOf((*MockBackend)(nil).AddOneMachine), arg0)
}

// AddRelation mocks base method.
func (m *MockBackend) AddRelation(arg0 ...state.Endpoint) (*state.Relation, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddRelation", varargs...)
	ret0, _ := ret[0].(*state.Relation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddRelation indicates an expected call of AddRelation.
func (mr *MockBackendMockRecorder) AddRelation(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRelation", reflect.TypeOf((*MockBackend)(nil).AddRelation), arg0...)
}

// AllApplicationOffers mocks base method.
func (m *MockBackend) AllApplicationOffers() ([]*crossmodel.ApplicationOffer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllApplicationOffers")
	ret0, _ := ret[0].([]*crossmodel.ApplicationOffer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllApplicationOffers indicates an expected call of AllApplicationOffers.
func (mr *MockBackendMockRecorder) AllApplicationOffers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllApplicationOffers", reflect.TypeOf((*MockBackend)(nil).AllApplicationOffers))
}

// AllApplications mocks base method.
func (m *MockBackend) AllApplications() ([]*state.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllApplications")
	ret0, _ := ret[0].([]*state.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllApplications indicates an expected call of AllApplications.
func (mr *MockBackendMockRecorder) AllApplications() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllApplications", reflect.TypeOf((*MockBackend)(nil).AllApplications))
}

// AllIPAddresses mocks base method.
func (m *MockBackend) AllIPAddresses() ([]*state.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllIPAddresses")
	ret0, _ := ret[0].([]*state.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllIPAddresses indicates an expected call of AllIPAddresses.
func (mr *MockBackendMockRecorder) AllIPAddresses() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllIPAddresses", reflect.TypeOf((*MockBackend)(nil).AllIPAddresses))
}

// AllLinkLayerDevices mocks base method.
func (m *MockBackend) AllLinkLayerDevices() ([]*state.LinkLayerDevice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllLinkLayerDevices")
	ret0, _ := ret[0].([]*state.LinkLayerDevice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllLinkLayerDevices indicates an expected call of AllLinkLayerDevices.
func (mr *MockBackendMockRecorder) AllLinkLayerDevices() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllLinkLayerDevices", reflect.TypeOf((*MockBackend)(nil).AllLinkLayerDevices))
}

// AllMachines mocks base method.
func (m *MockBackend) AllMachines() ([]*state.Machine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllMachines")
	ret0, _ := ret[0].([]*state.Machine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllMachines indicates an expected call of AllMachines.
func (mr *MockBackendMockRecorder) AllMachines() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllMachines", reflect.TypeOf((*MockBackend)(nil).AllMachines))
}

// AllModelUUIDs mocks base method.
func (m *MockBackend) AllModelUUIDs() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllModelUUIDs")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllModelUUIDs indicates an expected call of AllModelUUIDs.
func (mr *MockBackendMockRecorder) AllModelUUIDs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllModelUUIDs", reflect.TypeOf((*MockBackend)(nil).AllModelUUIDs))
}

// AllRelations mocks base method.
func (m *MockBackend) AllRelations() ([]*state.Relation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllRelations")
	ret0, _ := ret[0].([]*state.Relation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllRelations indicates an expected call of AllRelations.
func (mr *MockBackendMockRecorder) AllRelations() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllRelations", reflect.TypeOf((*MockBackend)(nil).AllRelations))
}

// AllRemoteApplications mocks base method.
func (m *MockBackend) AllRemoteApplications() ([]*state.RemoteApplication, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllRemoteApplications")
	ret0, _ := ret[0].([]*state.RemoteApplication)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllRemoteApplications indicates an expected call of AllRemoteApplications.
func (mr *MockBackendMockRecorder) AllRemoteApplications() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllRemoteApplications", reflect.TypeOf((*MockBackend)(nil).AllRemoteApplications))
}

// AllSpaceInfos mocks base method.
func (m *MockBackend) AllSpaceInfos() (network.SpaceInfos, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllSpaceInfos")
	ret0, _ := ret[0].(network.SpaceInfos)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllSpaceInfos indicates an expected call of AllSpaceInfos.
func (mr *MockBackendMockRecorder) AllSpaceInfos() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllSpaceInfos", reflect.TypeOf((*MockBackend)(nil).AllSpaceInfos))
}

// AllSubnets mocks base method.
func (m *MockBackend) AllSubnets() ([]*state.Subnet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllSubnets")
	ret0, _ := ret[0].([]*state.Subnet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllSubnets indicates an expected call of AllSubnets.
func (mr *MockBackendMockRecorder) AllSubnets() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllSubnets", reflect.TypeOf((*MockBackend)(nil).AllSubnets))
}

// Annotations mocks base method.
func (m *MockBackend) Annotations(arg0 state.GlobalEntity) (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Annotations", arg0)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Annotations indicates an expected call of Annotations.
func (mr *MockBackendMockRecorder) Annotations(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Annotations", reflect.TypeOf((*MockBackend)(nil).Annotations), arg0)
}

// Application mocks base method.
func (m *MockBackend) Application(arg0 string) (client.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Application", arg0)
	ret0, _ := ret[0].(client.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Application indicates an expected call of Application.
func (mr *MockBackendMockRecorder) Application(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Application", reflect.TypeOf((*MockBackend)(nil).Application), arg0)
}

// Charm mocks base method.
func (m *MockBackend) Charm(arg0 string) (*state.Charm, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Charm", arg0)
	ret0, _ := ret[0].(*state.Charm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Charm indicates an expected call of Charm.
func (mr *MockBackendMockRecorder) Charm(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Charm", reflect.TypeOf((*MockBackend)(nil).Charm), arg0)
}

// ControllerConfig mocks base method.
func (m *MockBackend) ControllerConfig() (controller.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerConfig")
	ret0, _ := ret[0].(controller.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ControllerConfig indicates an expected call of ControllerConfig.
func (mr *MockBackendMockRecorder) ControllerConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerConfig", reflect.TypeOf((*MockBackend)(nil).ControllerConfig))
}

// ControllerNodes mocks base method.
func (m *MockBackend) ControllerNodes() ([]state.ControllerNode, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerNodes")
	ret0, _ := ret[0].([]state.ControllerNode)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ControllerNodes indicates an expected call of ControllerNodes.
func (mr *MockBackendMockRecorder) ControllerNodes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerNodes", reflect.TypeOf((*MockBackend)(nil).ControllerNodes))
}

// ControllerTag mocks base method.
func (m *MockBackend) ControllerTag() names.ControllerTag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerTag")
	ret0, _ := ret[0].(names.ControllerTag)
	return ret0
}

// ControllerTag indicates an expected call of ControllerTag.
func (mr *MockBackendMockRecorder) ControllerTag() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerTag", reflect.TypeOf((*MockBackend)(nil).ControllerTag))
}

// ControllerTimestamp mocks base method.
func (m *MockBackend) ControllerTimestamp() (*time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerTimestamp")
	ret0, _ := ret[0].(*time.Time)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ControllerTimestamp indicates an expected call of ControllerTimestamp.
func (mr *MockBackendMockRecorder) ControllerTimestamp() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerTimestamp", reflect.TypeOf((*MockBackend)(nil).ControllerTimestamp))
}

// EndpointsRelation mocks base method.
func (m *MockBackend) EndpointsRelation(arg0 ...state.Endpoint) (*state.Relation, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EndpointsRelation", varargs...)
	ret0, _ := ret[0].(*state.Relation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EndpointsRelation indicates an expected call of EndpointsRelation.
func (mr *MockBackendMockRecorder) EndpointsRelation(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EndpointsRelation", reflect.TypeOf((*MockBackend)(nil).EndpointsRelation), arg0...)
}

// FindEntity mocks base method.
func (m *MockBackend) FindEntity(arg0 names.Tag) (state.Entity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindEntity", arg0)
	ret0, _ := ret[0].(state.Entity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindEntity indicates an expected call of FindEntity.
func (mr *MockBackendMockRecorder) FindEntity(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindEntity", reflect.TypeOf((*MockBackend)(nil).FindEntity), arg0)
}

// HAPrimaryMachine mocks base method.
func (m *MockBackend) HAPrimaryMachine() (names.MachineTag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HAPrimaryMachine")
	ret0, _ := ret[0].(names.MachineTag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HAPrimaryMachine indicates an expected call of HAPrimaryMachine.
func (mr *MockBackendMockRecorder) HAPrimaryMachine() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HAPrimaryMachine", reflect.TypeOf((*MockBackend)(nil).HAPrimaryMachine))
}

// InferEndpoints mocks base method.
func (m *MockBackend) InferEndpoints(arg0 ...string) ([]state.Endpoint, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "InferEndpoints", varargs...)
	ret0, _ := ret[0].([]state.Endpoint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InferEndpoints indicates an expected call of InferEndpoints.
func (mr *MockBackendMockRecorder) InferEndpoints(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InferEndpoints", reflect.TypeOf((*MockBackend)(nil).InferEndpoints), arg0...)
}

// IsController mocks base method.
func (m *MockBackend) IsController() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsController")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsController indicates an expected call of IsController.
func (mr *MockBackendMockRecorder) IsController() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsController", reflect.TypeOf((*MockBackend)(nil).IsController))
}

// LatestMigration mocks base method.
func (m *MockBackend) LatestMigration() (state.ModelMigration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LatestMigration")
	ret0, _ := ret[0].(state.ModelMigration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LatestMigration indicates an expected call of LatestMigration.
func (mr *MockBackendMockRecorder) LatestMigration() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LatestMigration", reflect.TypeOf((*MockBackend)(nil).LatestMigration))
}

// LatestPlaceholderCharm mocks base method.
func (m *MockBackend) LatestPlaceholderCharm(arg0 *charm.URL) (*state.Charm, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LatestPlaceholderCharm", arg0)
	ret0, _ := ret[0].(*state.Charm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LatestPlaceholderCharm indicates an expected call of LatestPlaceholderCharm.
func (mr *MockBackendMockRecorder) LatestPlaceholderCharm(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LatestPlaceholderCharm", reflect.TypeOf((*MockBackend)(nil).LatestPlaceholderCharm), arg0)
}

// Machine mocks base method.
func (m *MockBackend) Machine(arg0 string) (*state.Machine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Machine", arg0)
	ret0, _ := ret[0].(*state.Machine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Machine indicates an expected call of Machine.
func (mr *MockBackendMockRecorder) Machine(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Machine", reflect.TypeOf((*MockBackend)(nil).Machine), arg0)
}

// Model mocks base method.
func (m *MockBackend) Model() (client.Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Model")
	ret0, _ := ret[0].(client.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Model indicates an expected call of Model.
func (mr *MockBackendMockRecorder) Model() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Model", reflect.TypeOf((*MockBackend)(nil).Model))
}

// ModelConfig mocks base method.
func (m *MockBackend) ModelConfig() (*config.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelConfig")
	ret0, _ := ret[0].(*config.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModelConfig indicates an expected call of ModelConfig.
func (mr *MockBackendMockRecorder) ModelConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelConfig", reflect.TypeOf((*MockBackend)(nil).ModelConfig))
}

// ModelConstraints mocks base method.
func (m *MockBackend) ModelConstraints() (constraints.Value, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelConstraints")
	ret0, _ := ret[0].(constraints.Value)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModelConstraints indicates an expected call of ModelConstraints.
func (mr *MockBackendMockRecorder) ModelConstraints() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelConstraints", reflect.TypeOf((*MockBackend)(nil).ModelConstraints))
}

// ModelTag mocks base method.
func (m *MockBackend) ModelTag() names.ModelTag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelTag")
	ret0, _ := ret[0].(names.ModelTag)
	return ret0
}

// ModelTag indicates an expected call of ModelTag.
func (mr *MockBackendMockRecorder) ModelTag() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelTag", reflect.TypeOf((*MockBackend)(nil).ModelTag))
}

// ModelUUID mocks base method.
func (m *MockBackend) ModelUUID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelUUID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ModelUUID indicates an expected call of ModelUUID.
func (mr *MockBackendMockRecorder) ModelUUID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelUUID", reflect.TypeOf((*MockBackend)(nil).ModelUUID))
}

// MongoSession mocks base method.
func (m *MockBackend) MongoSession() client.MongoSession {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MongoSession")
	ret0, _ := ret[0].(client.MongoSession)
	return ret0
}

// MongoSession indicates an expected call of MongoSession.
func (mr *MockBackendMockRecorder) MongoSession() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MongoSession", reflect.TypeOf((*MockBackend)(nil).MongoSession))
}

// RemoteApplication mocks base method.
func (m *MockBackend) RemoteApplication(arg0 string) (*state.RemoteApplication, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoteApplication", arg0)
	ret0, _ := ret[0].(*state.RemoteApplication)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoteApplication indicates an expected call of RemoteApplication.
func (mr *MockBackendMockRecorder) RemoteApplication(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoteApplication", reflect.TypeOf((*MockBackend)(nil).RemoteApplication), arg0)
}

// RemoteConnectionStatus mocks base method.
func (m *MockBackend) RemoteConnectionStatus(arg0 string) (*state.RemoteConnectionStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoteConnectionStatus", arg0)
	ret0, _ := ret[0].(*state.RemoteConnectionStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoteConnectionStatus indicates an expected call of RemoteConnectionStatus.
func (mr *MockBackendMockRecorder) RemoteConnectionStatus(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoteConnectionStatus", reflect.TypeOf((*MockBackend)(nil).RemoteConnectionStatus), arg0)
}

// RemoveUserAccess mocks base method.
func (m *MockBackend) RemoveUserAccess(arg0 names.UserTag, arg1 names.Tag) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveUserAccess", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveUserAccess indicates an expected call of RemoveUserAccess.
func (mr *MockBackendMockRecorder) RemoveUserAccess(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveUserAccess", reflect.TypeOf((*MockBackend)(nil).RemoveUserAccess), arg0, arg1)
}

// SetAnnotations mocks base method.
func (m *MockBackend) SetAnnotations(arg0 state.GlobalEntity, arg1 map[string]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetAnnotations", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetAnnotations indicates an expected call of SetAnnotations.
func (mr *MockBackendMockRecorder) SetAnnotations(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAnnotations", reflect.TypeOf((*MockBackend)(nil).SetAnnotations), arg0, arg1)
}

// SetModelAgentVersion mocks base method.
func (m *MockBackend) SetModelAgentVersion(arg0 version.Number, arg1 *string, arg2 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetModelAgentVersion", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetModelAgentVersion indicates an expected call of SetModelAgentVersion.
func (mr *MockBackendMockRecorder) SetModelAgentVersion(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetModelAgentVersion", reflect.TypeOf((*MockBackend)(nil).SetModelAgentVersion), arg0, arg1, arg2)
}

// SetModelConstraints mocks base method.
func (m *MockBackend) SetModelConstraints(arg0 constraints.Value) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetModelConstraints", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetModelConstraints indicates an expected call of SetModelConstraints.
func (mr *MockBackendMockRecorder) SetModelConstraints(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetModelConstraints", reflect.TypeOf((*MockBackend)(nil).SetModelConstraints), arg0)
}

// Unit mocks base method.
func (m *MockBackend) Unit(arg0 string) (client.Unit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unit", arg0)
	ret0, _ := ret[0].(client.Unit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Unit indicates an expected call of Unit.
func (mr *MockBackendMockRecorder) Unit(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unit", reflect.TypeOf((*MockBackend)(nil).Unit), arg0)
}

// UpdateModelConfig mocks base method.
func (m *MockBackend) UpdateModelConfig(arg0 map[string]interface{}, arg1 []string, arg2 ...state.ValidateConfigFunc) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateModelConfig", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateModelConfig indicates an expected call of UpdateModelConfig.
func (mr *MockBackendMockRecorder) UpdateModelConfig(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateModelConfig", reflect.TypeOf((*MockBackend)(nil).UpdateModelConfig), varargs...)
}

// MockModel is a mock of Model interface.
type MockModel struct {
	ctrl     *gomock.Controller
	recorder *MockModelMockRecorder
}

// MockModelMockRecorder is the mock recorder for MockModel.
type MockModelMockRecorder struct {
	mock *MockModel
}

// NewMockModel creates a new mock instance.
func NewMockModel(ctrl *gomock.Controller) *MockModel {
	mock := &MockModel{ctrl: ctrl}
	mock.recorder = &MockModelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockModel) EXPECT() *MockModelMockRecorder {
	return m.recorder
}

// AddUser mocks base method.
func (m *MockModel) AddUser(arg0 state.UserAccessSpec) (permission.UserAccess, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUser", arg0)
	ret0, _ := ret[0].(permission.UserAccess)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddUser indicates an expected call of AddUser.
func (mr *MockModelMockRecorder) AddUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUser", reflect.TypeOf((*MockModel)(nil).AddUser), arg0)
}

// CloudCredentialTag mocks base method.
func (m *MockModel) CloudCredentialTag() (names.CloudCredentialTag, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloudCredentialTag")
	ret0, _ := ret[0].(names.CloudCredentialTag)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// CloudCredentialTag indicates an expected call of CloudCredentialTag.
func (mr *MockModelMockRecorder) CloudCredentialTag() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudCredentialTag", reflect.TypeOf((*MockModel)(nil).CloudCredentialTag))
}

// CloudName mocks base method.
func (m *MockModel) CloudName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloudName")
	ret0, _ := ret[0].(string)
	return ret0
}

// CloudName indicates an expected call of CloudName.
func (mr *MockModelMockRecorder) CloudName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudName", reflect.TypeOf((*MockModel)(nil).CloudName))
}

// CloudRegion mocks base method.
func (m *MockModel) CloudRegion() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloudRegion")
	ret0, _ := ret[0].(string)
	return ret0
}

// CloudRegion indicates an expected call of CloudRegion.
func (mr *MockModelMockRecorder) CloudRegion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudRegion", reflect.TypeOf((*MockModel)(nil).CloudRegion))
}

// Config mocks base method.
func (m *MockModel) Config() (*config.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Config")
	ret0, _ := ret[0].(*config.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Config indicates an expected call of Config.
func (mr *MockModelMockRecorder) Config() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Config", reflect.TypeOf((*MockModel)(nil).Config))
}

// LatestToolsVersion mocks base method.
func (m *MockModel) LatestToolsVersion() version.Number {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LatestToolsVersion")
	ret0, _ := ret[0].(version.Number)
	return ret0
}

// LatestToolsVersion indicates an expected call of LatestToolsVersion.
func (mr *MockModelMockRecorder) LatestToolsVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LatestToolsVersion", reflect.TypeOf((*MockModel)(nil).LatestToolsVersion))
}

// Life mocks base method.
func (m *MockModel) Life() state.Life {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Life")
	ret0, _ := ret[0].(state.Life)
	return ret0
}

// Life indicates an expected call of Life.
func (mr *MockModelMockRecorder) Life() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Life", reflect.TypeOf((*MockModel)(nil).Life))
}

// MeterStatus mocks base method.
func (m *MockModel) MeterStatus() state.MeterStatus {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MeterStatus")
	ret0, _ := ret[0].(state.MeterStatus)
	return ret0
}

// MeterStatus indicates an expected call of MeterStatus.
func (mr *MockModelMockRecorder) MeterStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MeterStatus", reflect.TypeOf((*MockModel)(nil).MeterStatus))
}

// Name mocks base method.
func (m *MockModel) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockModelMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockModel)(nil).Name))
}

// Owner mocks base method.
func (m *MockModel) Owner() names.UserTag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Owner")
	ret0, _ := ret[0].(names.UserTag)
	return ret0
}

// Owner indicates an expected call of Owner.
func (mr *MockModelMockRecorder) Owner() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Owner", reflect.TypeOf((*MockModel)(nil).Owner))
}

// SLALevel mocks base method.
func (m *MockModel) SLALevel() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SLALevel")
	ret0, _ := ret[0].(string)
	return ret0
}

// SLALevel indicates an expected call of SLALevel.
func (mr *MockModelMockRecorder) SLALevel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SLALevel", reflect.TypeOf((*MockModel)(nil).SLALevel))
}

// SLAOwner mocks base method.
func (m *MockModel) SLAOwner() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SLAOwner")
	ret0, _ := ret[0].(string)
	return ret0
}

// SLAOwner indicates an expected call of SLAOwner.
func (mr *MockModelMockRecorder) SLAOwner() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SLAOwner", reflect.TypeOf((*MockModel)(nil).SLAOwner))
}

// Status mocks base method.
func (m *MockModel) Status() (status.StatusInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Status")
	ret0, _ := ret[0].(status.StatusInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Status indicates an expected call of Status.
func (mr *MockModelMockRecorder) Status() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockModel)(nil).Status))
}

// StatusHistory mocks base method.
func (m *MockModel) StatusHistory(arg0 status.StatusHistoryFilter) ([]status.StatusInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StatusHistory", arg0)
	ret0, _ := ret[0].([]status.StatusInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StatusHistory indicates an expected call of StatusHistory.
func (mr *MockModelMockRecorder) StatusHistory(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatusHistory", reflect.TypeOf((*MockModel)(nil).StatusHistory), arg0)
}

// Type mocks base method.
func (m *MockModel) Type() state.ModelType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Type")
	ret0, _ := ret[0].(state.ModelType)
	return ret0
}

// Type indicates an expected call of Type.
func (mr *MockModelMockRecorder) Type() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Type", reflect.TypeOf((*MockModel)(nil).Type))
}

// UUID mocks base method.
func (m *MockModel) UUID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UUID")
	ret0, _ := ret[0].(string)
	return ret0
}

// UUID indicates an expected call of UUID.
func (mr *MockModelMockRecorder) UUID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UUID", reflect.TypeOf((*MockModel)(nil).UUID))
}

// Users mocks base method.
func (m *MockModel) Users() ([]permission.UserAccess, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Users")
	ret0, _ := ret[0].([]permission.UserAccess)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Users indicates an expected call of Users.
func (mr *MockModelMockRecorder) Users() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Users", reflect.TypeOf((*MockModel)(nil).Users))
}
