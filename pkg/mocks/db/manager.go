// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/db/manager.go

// Package mock_db is a generated GoMock package.
package mock_db

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	db "github.com/mangalores/go-api-skeleton/pkg/db"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Handle mocks base method.
func (m *MockRepository) Handle(q db.QueryObject) db.QueryObject {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Handle", q)
	ret0, _ := ret[0].(db.QueryObject)
	return ret0
}

// Handle indicates an expected call of Handle.
func (mr *MockRepositoryMockRecorder) Handle(q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*MockRepository)(nil).Handle), q)
}

// Supports mocks base method.
func (m *MockRepository) Supports(t interface{}) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Supports", t)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Supports indicates an expected call of Supports.
func (mr *MockRepositoryMockRecorder) Supports(t interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Supports", reflect.TypeOf((*MockRepository)(nil).Supports), t)
}