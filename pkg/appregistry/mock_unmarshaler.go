// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/appregistry/unmarshaller.go

// Package appregistry is a generated GoMock package.
package appregistry

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockblobUnmarshaler is a mock of blobUnmarshaler interface
type MockblobUnmarshaler struct {
	ctrl     *gomock.Controller
	recorder *MockblobUnmarshalerMockRecorder
}

// MockblobUnmarshalerMockRecorder is the mock recorder for MockblobUnmarshaler
type MockblobUnmarshalerMockRecorder struct {
	mock *MockblobUnmarshaler
}

// NewMockblobUnmarshaler creates a new mock instance
func NewMockblobUnmarshaler(ctrl *gomock.Controller) *MockblobUnmarshaler {
	mock := &MockblobUnmarshaler{ctrl: ctrl}
	mock.recorder = &MockblobUnmarshalerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockblobUnmarshaler) EXPECT() *MockblobUnmarshalerMockRecorder {
	return m.recorder
}

// Unmarshal mocks base method
func (m *MockblobUnmarshaler) Unmarshal(in []byte) (*Manifest, error) {
	ret := m.ctrl.Call(m, "Unmarshal", in)
	ret0, _ := ret[0].(*Manifest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Unmarshal indicates an expected call of Unmarshal
func (mr *MockblobUnmarshalerMockRecorder) Unmarshal(in interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unmarshal", reflect.TypeOf((*MockblobUnmarshaler)(nil).Unmarshal), in)
}
