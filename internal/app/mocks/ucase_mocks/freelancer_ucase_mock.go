// Code generated by MockGen. DO NOT EDIT.
// Source: internal/app/freelancer/usecase.go

// Package ucase_mocks is a generated GoMock package.
package ucase_mocks

import (
	model "github.com/go-park-mail-ru/2019_2_Comandus/internal/model"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockFreelancerUsecase is a mock of Usecase interface
type MockFreelancerUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockFreelancerUsecaseMockRecorder
}

// MockFreelancerUsecaseMockRecorder is the mock recorder for MockFreelancerUsecase
type MockFreelancerUsecaseMockRecorder struct {
	mock *MockFreelancerUsecase
}

// NewMockFreelancerUsecase creates a new mock instance
func NewMockFreelancerUsecase(ctrl *gomock.Controller) *MockFreelancerUsecase {
	mock := &MockFreelancerUsecase{ctrl: ctrl}
	mock.recorder = &MockFreelancerUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFreelancerUsecase) EXPECT() *MockFreelancerUsecaseMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockFreelancerUsecase) Create(arg0 int64) (*model.Freelancer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(*model.Freelancer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockFreelancerUsecaseMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockFreelancerUsecase)(nil).Create), arg0)
}

// FindByUser mocks base method
func (m *MockFreelancerUsecase) FindByUser(arg0 int64) (*model.Freelancer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByUser", arg0)
	ret0, _ := ret[0].(*model.Freelancer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByUser indicates an expected call of FindByUser
func (mr *MockFreelancerUsecaseMockRecorder) FindByUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByUser", reflect.TypeOf((*MockFreelancerUsecase)(nil).FindByUser), arg0)
}

// Find mocks base method
func (m *MockFreelancerUsecase) Find(arg0 int64) (*model.Freelancer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0)
	ret0, _ := ret[0].(*model.Freelancer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockFreelancerUsecaseMockRecorder) Find(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockFreelancerUsecase)(nil).Find), arg0)
}

// Edit mocks base method
func (m *MockFreelancerUsecase) Edit(arg0, arg1 *model.Freelancer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Edit", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Edit indicates an expected call of Edit
func (mr *MockFreelancerUsecaseMockRecorder) Edit(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Edit", reflect.TypeOf((*MockFreelancerUsecase)(nil).Edit), arg0, arg1)
}

// PatternSearch mocks base method
func (m *MockFreelancerUsecase) PatternSearch(arg0 string) ([]model.ExtendFreelancer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PatternSearch", arg0)
	ret0, _ := ret[0].([]model.ExtendFreelancer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PatternSearch indicates an expected call of PatternSearch
func (mr *MockFreelancerUsecaseMockRecorder) PatternSearch(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatternSearch", reflect.TypeOf((*MockFreelancerUsecase)(nil).PatternSearch), arg0)
}
