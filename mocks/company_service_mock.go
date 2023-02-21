package mocks

import (
	"reflect"

	"company-rest-api/internal/company/model"
	"company-rest-api/internal/company/repository"
	"github.com/golang/mock/gomock"
)

// MockCompanyServiceInt is a mock of CompanyServiceInt interface.
type MockCompanyServiceInt struct {
	ctrl     *gomock.Controller
	recorder *MockCompanyServiceIntMockRecorder
}

// MockCompanyServiceIntMockRecorder is the mock recorder for MockCompanyServiceInt.
type MockCompanyServiceIntMockRecorder struct {
	mock *MockCompanyServiceInt
}

// NewMockCompanyServiceInt creates a new mock instance.
func NewMockCompanyServiceInt(ctrl *gomock.Controller) *MockCompanyServiceInt {
	mock := &MockCompanyServiceInt{ctrl: ctrl}
	mock.recorder = &MockCompanyServiceIntMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCompanyServiceInt) EXPECT() *MockCompanyServiceIntMockRecorder {
	return m.recorder
}

// CreateCompany mocks base method.
func (m *MockCompanyServiceInt) CreateCompany(companyEntity *model.CompanyEntity) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCompany", companyEntity)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCompany indicates an expected call of CreateCompany.
func (mr *MockCompanyServiceIntMockRecorder) CreateCompany(companyEntity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCompany", reflect.TypeOf((*MockCompanyServiceInt)(nil).CreateCompany), companyEntity)
}

// DeleteCompany mocks base method.
func (m *MockCompanyServiceInt) DeleteCompany(companyId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCompany", companyId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCompany indicates an expected call of DeleteCompany.
func (mr *MockCompanyServiceIntMockRecorder) DeleteCompany(companyId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCompany", reflect.TypeOf((*MockCompanyServiceInt)(nil).DeleteCompany), companyId)
}

// GetCompany mocks base method.
func (m *MockCompanyServiceInt) GetCompany(companyId string) (*repository.CompanyDocument, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCompany", companyId)
	ret0, _ := ret[0].(*repository.CompanyDocument)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCompany indicates an expected call of GetCompany.
func (mr *MockCompanyServiceIntMockRecorder) GetCompany(companyId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCompany", reflect.TypeOf((*MockCompanyServiceInt)(nil).GetCompany), companyId)
}

// UpdateCompany mocks base method.
func (m *MockCompanyServiceInt) UpdateCompany(companyEntity *model.CompanyEntity) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCompany", companyEntity)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCompany indicates an expected call of UpdateCompany.
func (mr *MockCompanyServiceIntMockRecorder) UpdateCompany(companyEntity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCompany", reflect.TypeOf((*MockCompanyServiceInt)(nil).UpdateCompany), companyEntity)
}
