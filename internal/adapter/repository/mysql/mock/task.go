// Code generated by MockGen. DO NOT EDIT.
// Source: internal/adapter/repository/mysql/task.go

// Package mock_mysql is a generated GoMock package.
package mock_mysql

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockDB is a mock of DB interface.
type MockDB struct {
	ctrl     *gomock.Controller
	recorder *MockDBMockRecorder
}

// MockDBMockRecorder is the mock recorder for MockDB.
type MockDBMockRecorder struct {
	mock *MockDB
}

// NewMockDB creates a new mock instance.
func NewMockDB(ctrl *gomock.Controller) *MockDB {
	mock := &MockDB{ctrl: ctrl}
	mock.recorder = &MockDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDB) EXPECT() *MockDBMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockDB) Delete(query string, args any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", query, args)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockDBMockRecorder) Delete(query, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockDB)(nil).Delete), query, args)
}

// Query mocks base method.
func (m *MockDB) Query(query string, args any, fn func(func(...any) error) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Query", query, args, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// Query indicates an expected call of Query.
func (mr *MockDBMockRecorder) Query(query, args, fn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockDB)(nil).Query), query, args, fn)
}

// QueryRow mocks base method.
func (m *MockDB) QueryRow(query string, args any, fn func(func(...any) error) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryRow", query, args, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// QueryRow indicates an expected call of QueryRow.
func (mr *MockDBMockRecorder) QueryRow(query, args, fn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryRow", reflect.TypeOf((*MockDB)(nil).QueryRow), query, args, fn)
}

// Save mocks base method.
func (m *MockDB) Save(query string, args ...any) (any, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Save", varargs...)
	ret0, _ := ret[0].(any)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save.
func (mr *MockDBMockRecorder) Save(query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockDB)(nil).Save), varargs...)
}

// Update mocks base method.
func (m *MockDB) Update(query string, args ...any) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Update", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockDBMockRecorder) Update(query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockDB)(nil).Update), varargs...)
}
