// Code generated by MockGen. DO NOT EDIT.
// Source: internal/app/clients/interfaces/jobClient.go

// Package client_mocks is a generated GoMock package.
package client_mocks

import (
	job_grpc "github.com/go-park-mail-ru/2019_2_Comandus/internal/app/job/delivery/grpc/job_grpc"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockClientJob is a mock of ClientJob interface
type MockClientJob struct {
	ctrl     *gomock.Controller
	recorder *MockClientJobMockRecorder
}

// MockClientJobMockRecorder is the mock recorder for MockClientJob
type MockClientJobMockRecorder struct {
	mock *MockClientJob
}

// NewMockClientJob creates a new mock instance
func NewMockClientJob(ctrl *gomock.Controller) *MockClientJob {
	mock := &MockClientJob{ctrl: ctrl}
	mock.recorder = &MockClientJobMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClientJob) EXPECT() *MockClientJobMockRecorder {
	return m.recorder
}

// GetJobFromServer mocks base method
func (m *MockClientJob) GetJobFromServer(id int64) (*job_grpc.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetJobFromServer", id)
	ret0, _ := ret[0].(*job_grpc.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJobFromServer indicates an expected call of GetJobFromServer
func (mr *MockClientJobMockRecorder) GetJobFromServer(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJobFromServer", reflect.TypeOf((*MockClientJob)(nil).GetJobFromServer), id)
}
