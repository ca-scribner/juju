// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/internal/resource/charmhub (interfaces: ResourceGetter,CharmHub)
//
// Generated by this command:
//
//	mockgen -typed -package charmhub_test -destination charmhub_mock_test.go github.com/juju/juju/internal/resource/charmhub ResourceGetter,CharmHub
//

// Package charmhub_test is a generated GoMock package.
package charmhub_test

import (
	context "context"
	io "io"
	url "net/url"
	reflect "reflect"

	charmhub "github.com/juju/juju/internal/charmhub"
	transport "github.com/juju/juju/internal/charmhub/transport"
	charmhub0 "github.com/juju/juju/internal/resource/charmhub"
	gomock "go.uber.org/mock/gomock"
)

// MockResourceGetter is a mock of ResourceGetter interface.
type MockResourceGetter struct {
	ctrl     *gomock.Controller
	recorder *MockResourceGetterMockRecorder
}

// MockResourceGetterMockRecorder is the mock recorder for MockResourceGetter.
type MockResourceGetterMockRecorder struct {
	mock *MockResourceGetter
}

// NewMockResourceGetter creates a new mock instance.
func NewMockResourceGetter(ctrl *gomock.Controller) *MockResourceGetter {
	mock := &MockResourceGetter{ctrl: ctrl}
	mock.recorder = &MockResourceGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResourceGetter) EXPECT() *MockResourceGetterMockRecorder {
	return m.recorder
}

// GetResource mocks base method.
func (m *MockResourceGetter) GetResource(arg0 charmhub0.ResourceRequest) (charmhub0.ResourceData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResource", arg0)
	ret0, _ := ret[0].(charmhub0.ResourceData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResource indicates an expected call of GetResource.
func (mr *MockResourceGetterMockRecorder) GetResource(arg0 any) *MockResourceGetterGetResourceCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResource", reflect.TypeOf((*MockResourceGetter)(nil).GetResource), arg0)
	return &MockResourceGetterGetResourceCall{Call: call}
}

// MockResourceGetterGetResourceCall wrap *gomock.Call
type MockResourceGetterGetResourceCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockResourceGetterGetResourceCall) Return(arg0 charmhub0.ResourceData, arg1 error) *MockResourceGetterGetResourceCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockResourceGetterGetResourceCall) Do(f func(charmhub0.ResourceRequest) (charmhub0.ResourceData, error)) *MockResourceGetterGetResourceCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockResourceGetterGetResourceCall) DoAndReturn(f func(charmhub0.ResourceRequest) (charmhub0.ResourceData, error)) *MockResourceGetterGetResourceCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockCharmHub is a mock of CharmHub interface.
type MockCharmHub struct {
	ctrl     *gomock.Controller
	recorder *MockCharmHubMockRecorder
}

// MockCharmHubMockRecorder is the mock recorder for MockCharmHub.
type MockCharmHubMockRecorder struct {
	mock *MockCharmHub
}

// NewMockCharmHub creates a new mock instance.
func NewMockCharmHub(ctrl *gomock.Controller) *MockCharmHub {
	mock := &MockCharmHub{ctrl: ctrl}
	mock.recorder = &MockCharmHubMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCharmHub) EXPECT() *MockCharmHubMockRecorder {
	return m.recorder
}

// DownloadResource mocks base method.
func (m *MockCharmHub) DownloadResource(arg0 context.Context, arg1 *url.URL) (io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownloadResource", arg0, arg1)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DownloadResource indicates an expected call of DownloadResource.
func (mr *MockCharmHubMockRecorder) DownloadResource(arg0, arg1 any) *MockCharmHubDownloadResourceCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadResource", reflect.TypeOf((*MockCharmHub)(nil).DownloadResource), arg0, arg1)
	return &MockCharmHubDownloadResourceCall{Call: call}
}

// MockCharmHubDownloadResourceCall wrap *gomock.Call
type MockCharmHubDownloadResourceCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCharmHubDownloadResourceCall) Return(arg0 io.ReadCloser, arg1 error) *MockCharmHubDownloadResourceCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCharmHubDownloadResourceCall) Do(f func(context.Context, *url.URL) (io.ReadCloser, error)) *MockCharmHubDownloadResourceCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCharmHubDownloadResourceCall) DoAndReturn(f func(context.Context, *url.URL) (io.ReadCloser, error)) *MockCharmHubDownloadResourceCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Refresh mocks base method.
func (m *MockCharmHub) Refresh(arg0 context.Context, arg1 charmhub.RefreshConfig) ([]transport.RefreshResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Refresh", arg0, arg1)
	ret0, _ := ret[0].([]transport.RefreshResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Refresh indicates an expected call of Refresh.
func (mr *MockCharmHubMockRecorder) Refresh(arg0, arg1 any) *MockCharmHubRefreshCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Refresh", reflect.TypeOf((*MockCharmHub)(nil).Refresh), arg0, arg1)
	return &MockCharmHubRefreshCall{Call: call}
}

// MockCharmHubRefreshCall wrap *gomock.Call
type MockCharmHubRefreshCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCharmHubRefreshCall) Return(arg0 []transport.RefreshResponse, arg1 error) *MockCharmHubRefreshCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCharmHubRefreshCall) Do(f func(context.Context, charmhub.RefreshConfig) ([]transport.RefreshResponse, error)) *MockCharmHubRefreshCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCharmHubRefreshCall) DoAndReturn(f func(context.Context, charmhub.RefreshConfig) ([]transport.RefreshResponse, error)) *MockCharmHubRefreshCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
