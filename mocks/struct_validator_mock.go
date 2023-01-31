package mocks

import (
	"github.com/golang/mock/gomock"
	"reflect"
)

// MockStructValidatorInt is a mock of StructValidatorInt interface.
type MockStructValidatorInt struct {
	ctrl     *gomock.Controller
	recorder *MockStructValidatorIntMockRecorder
}

// MockStructValidatorIntMockRecorder is the mock recorder for MockStructValidatorInt.
type MockStructValidatorIntMockRecorder struct {
	mock *MockStructValidatorInt
}

// NewMockStructValidatorInt creates a new mock instance.
func NewMockStructValidatorInt(ctrl *gomock.Controller) *MockStructValidatorInt {
	mock := &MockStructValidatorInt{ctrl: ctrl}
	mock.recorder = &MockStructValidatorIntMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStructValidatorInt) EXPECT() *MockStructValidatorIntMockRecorder {
	return m.recorder
}

// Validate mocks base method.
func (m *MockStructValidatorInt) Validate(structForCheck interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate", structForCheck)
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate.
func (mr *MockStructValidatorIntMockRecorder) Validate(structForCheck interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockStructValidatorInt)(nil).Validate), structForCheck)
}
