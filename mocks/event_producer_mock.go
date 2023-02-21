package mocks

import (
	"reflect"

	"github.com/golang/mock/gomock"
)

// MockEventProducerInt is a mock of EventProducerInt interface.
type MockEventProducerInt struct {
	ctrl     *gomock.Controller
	recorder *MockEventProducerIntMockRecorder
}

// MockEventProducerIntMockRecorder is the mock recorder for MockEventProducerInt.
type MockEventProducerIntMockRecorder struct {
	mock *MockEventProducerInt
}

// NewMockEventProducerInt creates a new mock instance.
func NewMockEventProducerInt(ctrl *gomock.Controller) *MockEventProducerInt {
	mock := &MockEventProducerInt{ctrl: ctrl}
	mock.recorder = &MockEventProducerIntMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventProducerInt) EXPECT() *MockEventProducerIntMockRecorder {
	return m.recorder
}

// Initialize mocks base method.
func (m *MockEventProducerInt) Initialize() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Initialize")
}

// Initialize indicates an expected call of Initialize.
func (mr *MockEventProducerIntMockRecorder) Initialize() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Initialize", reflect.TypeOf((*MockEventProducerInt)(nil).Initialize))
}

// Produce mocks base method.
func (m *MockEventProducerInt) Produce(message []byte, topic string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Produce", message, topic)
}

// Produce indicates an expected call of Produce.
func (mr *MockEventProducerIntMockRecorder) Produce(message, topic interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Produce", reflect.TypeOf((*MockEventProducerInt)(nil).Produce), message, topic)
}
