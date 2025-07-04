// Code generated by MockGen. DO NOT EDIT.
// Source: internal/core/ports/repositories/student_repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	domain "github.com/PyMarcus/FreelaIF/users-api/internal/core/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockIStudentRepository is a mock of IStudentRepository interface.
type MockIStudentRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIStudentRepositoryMockRecorder
}

// MockIStudentRepositoryMockRecorder is the mock recorder for MockIStudentRepository.
type MockIStudentRepositoryMockRecorder struct {
	mock *MockIStudentRepository
}

// NewMockIStudentRepository creates a new mock instance.
func NewMockIStudentRepository(ctrl *gomock.Controller) *MockIStudentRepository {
	mock := &MockIStudentRepository{ctrl: ctrl}
	mock.recorder = &MockIStudentRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIStudentRepository) EXPECT() *MockIStudentRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIStudentRepository) Create(student *domain.StudentEntity) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", student)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockIStudentRepositoryMockRecorder) Create(student interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIStudentRepository)(nil).Create), student)
}

// GetStudentByEmail mocks base method.
func (m *MockIStudentRepository) GetStudentByEmail(email string) (*domain.StudentEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStudentByEmail", email)
	ret0, _ := ret[0].(*domain.StudentEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStudentByEmail indicates an expected call of GetStudentByEmail.
func (mr *MockIStudentRepositoryMockRecorder) GetStudentByEmail(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStudentByEmail", reflect.TypeOf((*MockIStudentRepository)(nil).GetStudentByEmail), email)
}

// GetStudentByUserName mocks base method.
func (m *MockIStudentRepository) GetStudentByUserName(username string) (*domain.StudentEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStudentByUserName", username)
	ret0, _ := ret[0].(*domain.StudentEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStudentByUserName indicates an expected call of GetStudentByUserName.
func (mr *MockIStudentRepositoryMockRecorder) GetStudentByUserName(username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStudentByUserName", reflect.TypeOf((*MockIStudentRepository)(nil).GetStudentByUserName), username)
}

// ListAll mocks base method.
func (m *MockIStudentRepository) ListAll() ([]*domain.StudentEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAll")
	ret0, _ := ret[0].([]*domain.StudentEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAll indicates an expected call of ListAll.
func (mr *MockIStudentRepositoryMockRecorder) ListAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAll", reflect.TypeOf((*MockIStudentRepository)(nil).ListAll))
}
