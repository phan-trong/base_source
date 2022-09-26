// Code generated by MockGen. DO NOT EDIT.
// Source: car_repository.go

// Package domain is a generated GoMock package.
package domain

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCarRepository is a mock of CarRepository interface.
type MockCarRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCarRepositoryMockRecorder
}

// MockCarRepositoryMockRecorder is the mock recorder for MockCarRepository.
type MockCarRepositoryMockRecorder struct {
	mock *MockCarRepository
}

// NewMockCarRepository creates a new mock instance.
func NewMockCarRepository(ctrl *gomock.Controller) *MockCarRepository {
	mock := &MockCarRepository{ctrl: ctrl}
	mock.recorder = &MockCarRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCarRepository) EXPECT() *MockCarRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockCarRepository) Create(ctx context.Context, c *Car) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, c)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockCarRepositoryMockRecorder) Create(ctx, c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCarRepository)(nil).Create), ctx, c)
}

// FindByID mocks base method.
func (m *MockCarRepository) FindByID(ctx context.Context, id int64) (*Car, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", ctx, id)
	ret0, _ := ret[0].(*Car)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockCarRepositoryMockRecorder) FindByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockCarRepository)(nil).FindByID), ctx, id)
}
