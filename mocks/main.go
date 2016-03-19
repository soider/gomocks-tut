// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/soider/gomocks-tut/ifaces (interfaces: ITaskProcesser)

package mocks

import (
	gomock "github.com/golang/mock/gomock"
)

// Mock of ITaskProcesser interface
type MockITaskProcesser struct {
	ctrl     *gomock.Controller
	recorder *_MockITaskProcesserRecorder
}

// Recorder for MockITaskProcesser (not exported)
type _MockITaskProcesserRecorder struct {
	mock *MockITaskProcesser
}

func NewMockITaskProcesser(ctrl *gomock.Controller) *MockITaskProcesser {
	mock := &MockITaskProcesser{ctrl: ctrl}
	mock.recorder = &_MockITaskProcesserRecorder{mock}
	return mock
}

func (_m *MockITaskProcesser) EXPECT() *_MockITaskProcesserRecorder {
	return _m.recorder
}

func (_m *MockITaskProcesser) ProcessTask(_param0 []int64) error {
	ret := _m.ctrl.Call(_m, "ProcessTask", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockITaskProcesserRecorder) ProcessTask(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ProcessTask", arg0)
}
