// Code generated by MockGen. DO NOT EDIT.
// Source: /home/saimon/Documents/Estudos/shortify/config/mediator/mediator.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockHandler is a mock of Handler interface.
type MockHandler struct {
	ctrl     *gomock.Controller
	recorder *MockHandlerMockRecorder
}

// MockHandlerMockRecorder is the mock recorder for MockHandler.
type MockHandlerMockRecorder struct {
	mock *MockHandler
}

// NewMockHandler creates a new mock instance.
func NewMockHandler(ctrl *gomock.Controller) *MockHandler {
	mock := &MockHandler{ctrl: ctrl}
	mock.recorder = &MockHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHandler) EXPECT() *MockHandlerMockRecorder {
	return m.recorder
}

// Notify mocks base method.
func (m *MockHandler) Notify(T any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Notify", T)
	ret0, _ := ret[0].(error)
	return ret0
}

// Notify indicates an expected call of Notify.
func (mr *MockHandlerMockRecorder) Notify(T interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Notify", reflect.TypeOf((*MockHandler)(nil).Notify), T)
}