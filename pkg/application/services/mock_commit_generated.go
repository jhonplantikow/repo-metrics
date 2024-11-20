// Code generated by MockGen. DO NOT EDIT.
// Source: commit.go

// Package services is a generated GoMock package.
package services

import (
	models "repo-metrics/pkg/application/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCommit is a mock of Commit interface.
type MockCommit struct {
	ctrl     *gomock.Controller
	recorder *MockCommitMockRecorder
}

// MockCommitMockRecorder is the mock recorder for MockCommit.
type MockCommitMockRecorder struct {
	mock *MockCommit
}

// NewMockCommit creates a new mock instance.
func NewMockCommit(ctrl *gomock.Controller) *MockCommit {
	mock := &MockCommit{ctrl: ctrl}
	mock.recorder = &MockCommitMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCommit) EXPECT() *MockCommitMockRecorder {
	return m.recorder
}

// ReadAll mocks base method.
func (m *MockCommit) ReadAll() ([]models.Commit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadAll")
	ret0, _ := ret[0].([]models.Commit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadAll indicates an expected call of ReadAll.
func (mr *MockCommitMockRecorder) ReadAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadAll", reflect.TypeOf((*MockCommit)(nil).ReadAll))
}
