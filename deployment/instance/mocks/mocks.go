// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/cloudfoundry/bosh-micro-cli/deployment/instance (interfaces: Instance,Manager)

package mocks

import (
	gomock "code.google.com/p/gomock/gomock"
	disk "github.com/cloudfoundry/bosh-micro-cli/deployment/disk"
	instance "github.com/cloudfoundry/bosh-micro-cli/deployment/instance"
	manifest "github.com/cloudfoundry/bosh-micro-cli/deployment/manifest"
	manifest0 "github.com/cloudfoundry/bosh-micro-cli/installation/manifest"
	stemcell "github.com/cloudfoundry/bosh-micro-cli/stemcell"
	ui "github.com/cloudfoundry/bosh-micro-cli/ui"
	time "time"
)

// Mock of Instance interface
type MockInstance struct {
	ctrl     *gomock.Controller
	recorder *_MockInstanceRecorder
}

// Recorder for MockInstance (not exported)
type _MockInstanceRecorder struct {
	mock *MockInstance
}

func NewMockInstance(ctrl *gomock.Controller) *MockInstance {
	mock := &MockInstance{ctrl: ctrl}
	mock.recorder = &_MockInstanceRecorder{mock}
	return mock
}

func (_m *MockInstance) EXPECT() *_MockInstanceRecorder {
	return _m.recorder
}

func (_m *MockInstance) Delete(_param0 time.Duration, _param1 time.Duration, _param2 ui.Stage) error {
	ret := _m.ctrl.Call(_m, "Delete", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockInstanceRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Delete", arg0, arg1, arg2)
}

func (_m *MockInstance) Disks() ([]disk.Disk, error) {
	ret := _m.ctrl.Call(_m, "Disks")
	ret0, _ := ret[0].([]disk.Disk)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInstanceRecorder) Disks() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Disks")
}

func (_m *MockInstance) ID() int {
	ret := _m.ctrl.Call(_m, "ID")
	ret0, _ := ret[0].(int)
	return ret0
}

func (_mr *_MockInstanceRecorder) ID() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ID")
}

func (_m *MockInstance) JobName() string {
	ret := _m.ctrl.Call(_m, "JobName")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockInstanceRecorder) JobName() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "JobName")
}

func (_m *MockInstance) UpdateDisks(_param0 manifest.Manifest, _param1 ui.Stage) ([]disk.Disk, error) {
	ret := _m.ctrl.Call(_m, "UpdateDisks", _param0, _param1)
	ret0, _ := ret[0].([]disk.Disk)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInstanceRecorder) UpdateDisks(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateDisks", arg0, arg1)
}

func (_m *MockInstance) UpdateJobs(_param0 manifest.Manifest, _param1 ui.Stage) error {
	ret := _m.ctrl.Call(_m, "UpdateJobs", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockInstanceRecorder) UpdateJobs(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateJobs", arg0, arg1)
}

func (_m *MockInstance) WaitUntilReady(_param0 manifest0.Registry, _param1 manifest0.SSHTunnel, _param2 ui.Stage) error {
	ret := _m.ctrl.Call(_m, "WaitUntilReady", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockInstanceRecorder) WaitUntilReady(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "WaitUntilReady", arg0, arg1, arg2)
}

// Mock of Manager interface
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *_MockManagerRecorder
}

// Recorder for MockManager (not exported)
type _MockManagerRecorder struct {
	mock *MockManager
}

func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &_MockManagerRecorder{mock}
	return mock
}

func (_m *MockManager) EXPECT() *_MockManagerRecorder {
	return _m.recorder
}

func (_m *MockManager) Create(_param0 string, _param1 int, _param2 manifest.Manifest, _param3 stemcell.CloudStemcell, _param4 manifest0.Registry, _param5 manifest0.SSHTunnel, _param6 ui.Stage) (instance.Instance, []disk.Disk, error) {
	ret := _m.ctrl.Call(_m, "Create", _param0, _param1, _param2, _param3, _param4, _param5, _param6)
	ret0, _ := ret[0].(instance.Instance)
	ret1, _ := ret[1].([]disk.Disk)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockManagerRecorder) Create(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Create", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

func (_m *MockManager) DeleteAll(_param0 time.Duration, _param1 time.Duration, _param2 ui.Stage) error {
	ret := _m.ctrl.Call(_m, "DeleteAll", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockManagerRecorder) DeleteAll(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteAll", arg0, arg1, arg2)
}

func (_m *MockManager) FindCurrent() ([]instance.Instance, error) {
	ret := _m.ctrl.Call(_m, "FindCurrent")
	ret0, _ := ret[0].([]instance.Instance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockManagerRecorder) FindCurrent() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FindCurrent")
}
