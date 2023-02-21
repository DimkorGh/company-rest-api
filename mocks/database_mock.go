package mocks

import (
	"reflect"

	"github.com/golang/mock/gomock"
	"go.mongodb.org/mongo-driver/bson"
)

// MockDatabaseInt is a mock of DatabaseInt interface.
type MockDatabaseInt struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseIntMockRecorder
}

// MockDatabaseIntMockRecorder is the mock recorder for MockDatabaseInt.
type MockDatabaseIntMockRecorder struct {
	mock *MockDatabaseInt
}

// NewMockDatabaseInt creates a new mock instance.
func NewMockDatabaseInt(ctrl *gomock.Controller) *MockDatabaseInt {
	mock := &MockDatabaseInt{ctrl: ctrl}
	mock.recorder = &MockDatabaseIntMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDatabaseInt) EXPECT() *MockDatabaseIntMockRecorder {
	return m.recorder
}

// Connect mocks base method.
func (m *MockDatabaseInt) Connect() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Connect")
}

// Connect indicates an expected call of Connect.
func (mr *MockDatabaseIntMockRecorder) Connect() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockDatabaseInt)(nil).Connect))
}

// DeleteOne mocks base method.
func (m *MockDatabaseInt) DeleteOne(collectionName string, filter bson.D) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOne", collectionName, filter)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOne indicates an expected call of DeleteOne.
func (mr *MockDatabaseIntMockRecorder) DeleteOne(collectionName, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOne", reflect.TypeOf((*MockDatabaseInt)(nil).DeleteOne), collectionName, filter)
}

// FindOne mocks base method.
func (m *MockDatabaseInt) FindOne(collectionName string, filter bson.D, data interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOne", collectionName, filter, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// FindOne indicates an expected call of FindOne.
func (mr *MockDatabaseIntMockRecorder) FindOne(collectionName, filter, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOne", reflect.TypeOf((*MockDatabaseInt)(nil).FindOne), collectionName, filter, data)
}

// InsertOne mocks base method.
func (m *MockDatabaseInt) InsertOne(collectionName string, data interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertOne", collectionName, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertOne indicates an expected call of InsertOne.
func (mr *MockDatabaseIntMockRecorder) InsertOne(collectionName, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertOne", reflect.TypeOf((*MockDatabaseInt)(nil).InsertOne), collectionName, data)
}

// UpdateOne mocks base method.
func (m *MockDatabaseInt) UpdateOne(collectionName string, filter bson.M, data interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOne", collectionName, filter, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateOne indicates an expected call of UpdateOne.
func (mr *MockDatabaseIntMockRecorder) UpdateOne(collectionName, filter, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOne", reflect.TypeOf((*MockDatabaseInt)(nil).UpdateOne), collectionName, filter, data)
}
