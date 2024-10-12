// Code generated by MockGen. DO NOT EDIT.
// Source: /home/saimon/Documents/Estudos/shortify/internal/usecase/urlRecord.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	dto "github.com/ribeirosaimon/shortify/internal/dto"
	entity "github.com/ribeirosaimon/shortify/internal/entity"
)

// MockUrlRecord is a mock of UrlRecord interface.
type MockUrlRecord struct {
	ctrl     *gomock.Controller
	recorder *MockUrlRecordMockRecorder
}

// MockUrlRecordMockRecorder is the mock recorder for MockUrlRecord.
type MockUrlRecordMockRecorder struct {
	mock *MockUrlRecord
}

// NewMockUrlRecord creates a new mock instance.
func NewMockUrlRecord(ctrl *gomock.Controller) *MockUrlRecord {
	mock := &MockUrlRecord{ctrl: ctrl}
	mock.recorder = &MockUrlRecordMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUrlRecord) EXPECT() *MockUrlRecordMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUrlRecord) Create(ctx context.Context, urlRecord *dto.UrlRecord) (*entity.UrlRecord, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, urlRecord)
	ret0, _ := ret[0].(*entity.UrlRecord)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockUrlRecordMockRecorder) Create(ctx, urlRecord interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUrlRecord)(nil).Create), ctx, urlRecord)
}
