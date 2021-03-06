// Code generated by MockGen. DO NOT EDIT.
// Source: internal/app/clients/interfaces/freelancerClient.go

// Package client_mocks is a generated GoMock package.
package client_mocks

import (
	freelancer_grpc "github.com/go-park-mail-ru/2019_2_Comandus/internal/app/freelancer/delivery/grpc/freelancer_grpc"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockClientFreelancer is a mock of ClientFreelancer interface
type MockClientFreelancer struct {
	ctrl     *gomock.Controller
	recorder *MockClientFreelancerMockRecorder
}

// MockClientFreelancerMockRecorder is the mock recorder for MockClientFreelancer
type MockClientFreelancerMockRecorder struct {
	mock *MockClientFreelancer
}

// NewMockClientFreelancer creates a new mock instance
func NewMockClientFreelancer(ctrl *gomock.Controller) *MockClientFreelancer {
	mock := &MockClientFreelancer{ctrl: ctrl}
	mock.recorder = &MockClientFreelancerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClientFreelancer) EXPECT() *MockClientFreelancerMockRecorder {
	return m.recorder
}

// CreateFreelancerOnServer mocks base method
func (m *MockClientFreelancer) CreateFreelancerOnServer(userId int64) (*freelancer_grpc.Freelancer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFreelancerOnServer", userId)
	ret0, _ := ret[0].(*freelancer_grpc.Freelancer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFreelancerOnServer indicates an expected call of CreateFreelancerOnServer
func (mr *MockClientFreelancerMockRecorder) CreateFreelancerOnServer(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFreelancerOnServer", reflect.TypeOf((*MockClientFreelancer)(nil).CreateFreelancerOnServer), userId)
}

// GetFreelancerByUserFromServer mocks base method
func (m *MockClientFreelancer) GetFreelancerByUserFromServer(id int64) (*freelancer_grpc.Freelancer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFreelancerByUserFromServer", id)
	ret0, _ := ret[0].(*freelancer_grpc.Freelancer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFreelancerByUserFromServer indicates an expected call of GetFreelancerByUserFromServer
func (mr *MockClientFreelancerMockRecorder) GetFreelancerByUserFromServer(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFreelancerByUserFromServer", reflect.TypeOf((*MockClientFreelancer)(nil).GetFreelancerByUserFromServer), id)
}
