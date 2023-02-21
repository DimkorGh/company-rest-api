package mocks

import (
	"net/http"
	"reflect"

	"github.com/golang/mock/gomock"
)

// MockResponseBuilderInt is a mock of ResponseBuilderInt interface.
type MockResponseBuilderInt struct {
	ctrl     *gomock.Controller
	recorder *MockResponseBuilderIntMockRecorder
}

// MockResponseBuilderIntMockRecorder is the mock recorder for MockResponseBuilderInt.
type MockResponseBuilderIntMockRecorder struct {
	mock *MockResponseBuilderInt
}

// NewMockResponseBuilderInt creates a new mock instance.
func NewMockResponseBuilderInt(ctrl *gomock.Controller) *MockResponseBuilderInt {
	mock := &MockResponseBuilderInt{ctrl: ctrl}
	mock.recorder = &MockResponseBuilderIntMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResponseBuilderInt) EXPECT() *MockResponseBuilderIntMockRecorder {
	return m.recorder
}

// SendResponse mocks base method.
func (m *MockResponseBuilderInt) SendResponse(w http.ResponseWriter, body interface{}, err error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendResponse", w, body, err)
}

// SendResponse indicates an expected call of SendResponse.
func (mr *MockResponseBuilderIntMockRecorder) SendResponse(w, body, err interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendResponse", reflect.TypeOf((*MockResponseBuilderInt)(nil).SendResponse), w, body, err)
}
