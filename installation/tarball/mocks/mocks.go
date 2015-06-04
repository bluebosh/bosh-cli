// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/cloudfoundry/bosh-init/installation/tarball (interfaces: Provider)

package mocks

import (
	gomock "code.google.com/p/gomock/gomock"
	tarball "github.com/cloudfoundry/bosh-init/installation/tarball"
	ui "github.com/cloudfoundry/bosh-init/ui"
)

// Mock of Provider interface
type MockProvider struct {
	ctrl     *gomock.Controller
	recorder *_MockProviderRecorder
}

// Recorder for MockProvider (not exported)
type _MockProviderRecorder struct {
	mock *MockProvider
}

func NewMockProvider(ctrl *gomock.Controller) *MockProvider {
	mock := &MockProvider{ctrl: ctrl}
	mock.recorder = &_MockProviderRecorder{mock}
	return mock
}

func (_m *MockProvider) EXPECT() *_MockProviderRecorder {
	return _m.recorder
}

func (_m *MockProvider) Get(_param0 tarball.Source, _param1 ui.Stage) (string, error) {
	ret := _m.ctrl.Call(_m, "Get", _param0, _param1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockProviderRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Get", arg0, arg1)
}
