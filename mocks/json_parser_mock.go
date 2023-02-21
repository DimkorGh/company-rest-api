package mocks

import (
	"net/http"
	"reflect"

	"github.com/golang/mock/gomock"
)

// MockJsonParserInt is a mock of JsonParserInt interface.
type MockJsonParserInt struct {
	ctrl     *gomock.Controller
	recorder *MockJsonParserIntMockRecorder
}

// MockJsonParserIntMockRecorder is the mock recorder for MockJsonParserInt.
type MockJsonParserIntMockRecorder struct {
	mock *MockJsonParserInt
}

// NewMockJsonParserInt creates a new mock instance.
func NewMockJsonParserInt(ctrl *gomock.Controller) *MockJsonParserInt {
	mock := &MockJsonParserInt{ctrl: ctrl}
	mock.recorder = &MockJsonParserIntMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJsonParserInt) EXPECT() *MockJsonParserIntMockRecorder {
	return m.recorder
}

// ParseJson mocks base method.
func (m *MockJsonParserInt) ParseJson(r *http.Request, data interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseJson", r, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// ParseJson indicates an expected call of ParseJson.
func (mr *MockJsonParserIntMockRecorder) ParseJson(r, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseJson", reflect.TypeOf((*MockJsonParserInt)(nil).ParseJson), r, data)
}
