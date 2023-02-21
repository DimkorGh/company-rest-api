package mocks

import (
	"reflect"

	"company-rest-api/internal/company/model"
	"company-rest-api/internal/company/repository"
	"github.com/golang/mock/gomock"
)

// MockCompanyRepositoryInt is a mock of CompanyRepositoryInt interface.
type MockCompanyRepositoryInt struct {
	ctrl     *gomock.Controller
	recorder *MockCompanyRepositoryIntMockRecorder
}

// MockCompanyRepositoryIntMockRecorder is the mock recorder for MockCompanyRepositoryInt.
type MockCompanyRepositoryIntMockRecorder struct {
	mock *MockCompanyRepositoryInt
}

// NewMockCompanyRepositoryInt creates a new mock instance.
func NewMockCompanyRepositoryInt(ctrl *gomock.Controller) *MockCompanyRepositoryInt {
	mock := &MockCompanyRepositoryInt{ctrl: ctrl}
	mock.recorder = &MockCompanyRepositoryIntMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCompanyRepositoryInt) EXPECT() *MockCompanyRepositoryIntMockRecorder {
	return m.recorder
}

// CreateCompany mocks base method.
func (m *MockCompanyRepositoryInt) CreateCompany(companyEntity *model.CompanyEntity) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCompany", companyEntity)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCompany indicates an expected call of CreateCompany.
func (mr *MockCompanyRepositoryIntMockRecorder) CreateCompany(companyEntity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCompany", reflect.TypeOf((*MockCompanyRepositoryInt)(nil).CreateCompany), companyEntity)
}

// DeleteCompany mocks base method.
func (m *MockCompanyRepositoryInt) DeleteCompany(companyId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCompany", companyId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCompany indicates an expected call of DeleteCompany.
func (mr *MockCompanyRepositoryIntMockRecorder) DeleteCompany(companyId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCompany", reflect.TypeOf((*MockCompanyRepositoryInt)(nil).DeleteCompany), companyId)
}

// GetCompany mocks base method.
func (m *MockCompanyRepositoryInt) GetCompany(companyId string) (*repository.CompanyDocument, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCompany", companyId)
	ret0, _ := ret[0].(*repository.CompanyDocument)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCompany indicates an expected call of GetCompany.
func (mr *MockCompanyRepositoryIntMockRecorder) GetCompany(companyId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCompany", reflect.TypeOf((*MockCompanyRepositoryInt)(nil).GetCompany), companyId)
}

// UpdateCompany mocks base method.
func (m *MockCompanyRepositoryInt) UpdateCompany(companyEntity *model.CompanyEntity) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCompany", companyEntity)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCompany indicates an expected call of UpdateCompany.
func (mr *MockCompanyRepositoryIntMockRecorder) UpdateCompany(companyEntity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCompany", reflect.TypeOf((*MockCompanyRepositoryInt)(nil).UpdateCompany), companyEntity)
}
