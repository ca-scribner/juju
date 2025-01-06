// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/internal/worker/charmrevision (interfaces: CharmhubClient,ModelConfigService,ApplicationService,ModelService)
//
// Generated by this command:
//
//	mockgen -typed -package charmrevision -destination package_mocks_test.go github.com/juju/juju/internal/worker/charmrevision CharmhubClient,ModelConfigService,ApplicationService,ModelService
//

// Package charmrevision is a generated GoMock package.
package charmrevision

import (
	context "context"
	reflect "reflect"

	model "github.com/juju/juju/core/model"
	watcher "github.com/juju/juju/core/watcher"
	application "github.com/juju/juju/domain/application"
	config "github.com/juju/juju/environs/config"
	charmhub "github.com/juju/juju/internal/charmhub"
	transport "github.com/juju/juju/internal/charmhub/transport"
	gomock "go.uber.org/mock/gomock"
)

// MockCharmhubClient is a mock of CharmhubClient interface.
type MockCharmhubClient struct {
	ctrl     *gomock.Controller
	recorder *MockCharmhubClientMockRecorder
}

// MockCharmhubClientMockRecorder is the mock recorder for MockCharmhubClient.
type MockCharmhubClientMockRecorder struct {
	mock *MockCharmhubClient
}

// NewMockCharmhubClient creates a new mock instance.
func NewMockCharmhubClient(ctrl *gomock.Controller) *MockCharmhubClient {
	mock := &MockCharmhubClient{ctrl: ctrl}
	mock.recorder = &MockCharmhubClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCharmhubClient) EXPECT() *MockCharmhubClientMockRecorder {
	return m.recorder
}

// RefreshWithMetricsOnly mocks base method.
func (m *MockCharmhubClient) RefreshWithMetricsOnly(arg0 context.Context, arg1 charmhub.Metrics) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshWithMetricsOnly", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RefreshWithMetricsOnly indicates an expected call of RefreshWithMetricsOnly.
func (mr *MockCharmhubClientMockRecorder) RefreshWithMetricsOnly(arg0, arg1 any) *MockCharmhubClientRefreshWithMetricsOnlyCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshWithMetricsOnly", reflect.TypeOf((*MockCharmhubClient)(nil).RefreshWithMetricsOnly), arg0, arg1)
	return &MockCharmhubClientRefreshWithMetricsOnlyCall{Call: call}
}

// MockCharmhubClientRefreshWithMetricsOnlyCall wrap *gomock.Call
type MockCharmhubClientRefreshWithMetricsOnlyCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCharmhubClientRefreshWithMetricsOnlyCall) Return(arg0 error) *MockCharmhubClientRefreshWithMetricsOnlyCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCharmhubClientRefreshWithMetricsOnlyCall) Do(f func(context.Context, charmhub.Metrics) error) *MockCharmhubClientRefreshWithMetricsOnlyCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCharmhubClientRefreshWithMetricsOnlyCall) DoAndReturn(f func(context.Context, charmhub.Metrics) error) *MockCharmhubClientRefreshWithMetricsOnlyCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// RefreshWithRequestMetrics mocks base method.
func (m *MockCharmhubClient) RefreshWithRequestMetrics(arg0 context.Context, arg1 charmhub.RefreshConfig, arg2 charmhub.Metrics) ([]transport.RefreshResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshWithRequestMetrics", arg0, arg1, arg2)
	ret0, _ := ret[0].([]transport.RefreshResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RefreshWithRequestMetrics indicates an expected call of RefreshWithRequestMetrics.
func (mr *MockCharmhubClientMockRecorder) RefreshWithRequestMetrics(arg0, arg1, arg2 any) *MockCharmhubClientRefreshWithRequestMetricsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshWithRequestMetrics", reflect.TypeOf((*MockCharmhubClient)(nil).RefreshWithRequestMetrics), arg0, arg1, arg2)
	return &MockCharmhubClientRefreshWithRequestMetricsCall{Call: call}
}

// MockCharmhubClientRefreshWithRequestMetricsCall wrap *gomock.Call
type MockCharmhubClientRefreshWithRequestMetricsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCharmhubClientRefreshWithRequestMetricsCall) Return(arg0 []transport.RefreshResponse, arg1 error) *MockCharmhubClientRefreshWithRequestMetricsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCharmhubClientRefreshWithRequestMetricsCall) Do(f func(context.Context, charmhub.RefreshConfig, charmhub.Metrics) ([]transport.RefreshResponse, error)) *MockCharmhubClientRefreshWithRequestMetricsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCharmhubClientRefreshWithRequestMetricsCall) DoAndReturn(f func(context.Context, charmhub.RefreshConfig, charmhub.Metrics) ([]transport.RefreshResponse, error)) *MockCharmhubClientRefreshWithRequestMetricsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockModelConfigService is a mock of ModelConfigService interface.
type MockModelConfigService struct {
	ctrl     *gomock.Controller
	recorder *MockModelConfigServiceMockRecorder
}

// MockModelConfigServiceMockRecorder is the mock recorder for MockModelConfigService.
type MockModelConfigServiceMockRecorder struct {
	mock *MockModelConfigService
}

// NewMockModelConfigService creates a new mock instance.
func NewMockModelConfigService(ctrl *gomock.Controller) *MockModelConfigService {
	mock := &MockModelConfigService{ctrl: ctrl}
	mock.recorder = &MockModelConfigServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockModelConfigService) EXPECT() *MockModelConfigServiceMockRecorder {
	return m.recorder
}

// ModelConfig mocks base method.
func (m *MockModelConfigService) ModelConfig(arg0 context.Context) (*config.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelConfig", arg0)
	ret0, _ := ret[0].(*config.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModelConfig indicates an expected call of ModelConfig.
func (mr *MockModelConfigServiceMockRecorder) ModelConfig(arg0 any) *MockModelConfigServiceModelConfigCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelConfig", reflect.TypeOf((*MockModelConfigService)(nil).ModelConfig), arg0)
	return &MockModelConfigServiceModelConfigCall{Call: call}
}

// MockModelConfigServiceModelConfigCall wrap *gomock.Call
type MockModelConfigServiceModelConfigCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockModelConfigServiceModelConfigCall) Return(arg0 *config.Config, arg1 error) *MockModelConfigServiceModelConfigCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockModelConfigServiceModelConfigCall) Do(f func(context.Context) (*config.Config, error)) *MockModelConfigServiceModelConfigCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockModelConfigServiceModelConfigCall) DoAndReturn(f func(context.Context) (*config.Config, error)) *MockModelConfigServiceModelConfigCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Watch mocks base method.
func (m *MockModelConfigService) Watch() (watcher.Watcher[[]string], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch")
	ret0, _ := ret[0].(watcher.Watcher[[]string])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch.
func (mr *MockModelConfigServiceMockRecorder) Watch() *MockModelConfigServiceWatchCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockModelConfigService)(nil).Watch))
	return &MockModelConfigServiceWatchCall{Call: call}
}

// MockModelConfigServiceWatchCall wrap *gomock.Call
type MockModelConfigServiceWatchCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockModelConfigServiceWatchCall) Return(arg0 watcher.Watcher[[]string], arg1 error) *MockModelConfigServiceWatchCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockModelConfigServiceWatchCall) Do(f func() (watcher.Watcher[[]string], error)) *MockModelConfigServiceWatchCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockModelConfigServiceWatchCall) DoAndReturn(f func() (watcher.Watcher[[]string], error)) *MockModelConfigServiceWatchCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockApplicationService is a mock of ApplicationService interface.
type MockApplicationService struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationServiceMockRecorder
}

// MockApplicationServiceMockRecorder is the mock recorder for MockApplicationService.
type MockApplicationServiceMockRecorder struct {
	mock *MockApplicationService
}

// NewMockApplicationService creates a new mock instance.
func NewMockApplicationService(ctrl *gomock.Controller) *MockApplicationService {
	mock := &MockApplicationService{ctrl: ctrl}
	mock.recorder = &MockApplicationServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApplicationService) EXPECT() *MockApplicationServiceMockRecorder {
	return m.recorder
}

// GetApplicationsForRevisionUpdater mocks base method.
func (m *MockApplicationService) GetApplicationsForRevisionUpdater(arg0 context.Context) ([]application.RevisionUpdaterApplication, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetApplicationsForRevisionUpdater", arg0)
	ret0, _ := ret[0].([]application.RevisionUpdaterApplication)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetApplicationsForRevisionUpdater indicates an expected call of GetApplicationsForRevisionUpdater.
func (mr *MockApplicationServiceMockRecorder) GetApplicationsForRevisionUpdater(arg0 any) *MockApplicationServiceGetApplicationsForRevisionUpdaterCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApplicationsForRevisionUpdater", reflect.TypeOf((*MockApplicationService)(nil).GetApplicationsForRevisionUpdater), arg0)
	return &MockApplicationServiceGetApplicationsForRevisionUpdaterCall{Call: call}
}

// MockApplicationServiceGetApplicationsForRevisionUpdaterCall wrap *gomock.Call
type MockApplicationServiceGetApplicationsForRevisionUpdaterCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockApplicationServiceGetApplicationsForRevisionUpdaterCall) Return(arg0 []application.RevisionUpdaterApplication, arg1 error) *MockApplicationServiceGetApplicationsForRevisionUpdaterCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockApplicationServiceGetApplicationsForRevisionUpdaterCall) Do(f func(context.Context) ([]application.RevisionUpdaterApplication, error)) *MockApplicationServiceGetApplicationsForRevisionUpdaterCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockApplicationServiceGetApplicationsForRevisionUpdaterCall) DoAndReturn(f func(context.Context) ([]application.RevisionUpdaterApplication, error)) *MockApplicationServiceGetApplicationsForRevisionUpdaterCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockModelService is a mock of ModelService interface.
type MockModelService struct {
	ctrl     *gomock.Controller
	recorder *MockModelServiceMockRecorder
}

// MockModelServiceMockRecorder is the mock recorder for MockModelService.
type MockModelServiceMockRecorder struct {
	mock *MockModelService
}

// NewMockModelService creates a new mock instance.
func NewMockModelService(ctrl *gomock.Controller) *MockModelService {
	mock := &MockModelService{ctrl: ctrl}
	mock.recorder = &MockModelServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockModelService) EXPECT() *MockModelServiceMockRecorder {
	return m.recorder
}

// GetModelMetrics mocks base method.
func (m *MockModelService) GetModelMetrics(arg0 context.Context) (model.ModelMetrics, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModelMetrics", arg0)
	ret0, _ := ret[0].(model.ModelMetrics)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetModelMetrics indicates an expected call of GetModelMetrics.
func (mr *MockModelServiceMockRecorder) GetModelMetrics(arg0 any) *MockModelServiceGetModelMetricsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModelMetrics", reflect.TypeOf((*MockModelService)(nil).GetModelMetrics), arg0)
	return &MockModelServiceGetModelMetricsCall{Call: call}
}

// MockModelServiceGetModelMetricsCall wrap *gomock.Call
type MockModelServiceGetModelMetricsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockModelServiceGetModelMetricsCall) Return(arg0 model.ModelMetrics, arg1 error) *MockModelServiceGetModelMetricsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockModelServiceGetModelMetricsCall) Do(f func(context.Context) (model.ModelMetrics, error)) *MockModelServiceGetModelMetricsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockModelServiceGetModelMetricsCall) DoAndReturn(f func(context.Context) (model.ModelMetrics, error)) *MockModelServiceGetModelMetricsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
