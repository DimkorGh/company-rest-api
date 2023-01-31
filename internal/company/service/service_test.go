package service_test

import (
	"company-rest-api/internal/company/repository"
	"company-rest-api/internal/company/service"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"company-rest-api/internal/company/model"
	"company-rest-api/mocks"
)

func TestGetCompany(t *testing.T) {
	ctrl := gomock.NewController(t)

	type args struct {
		companyId string
	}
	tests := []struct {
		name                string
		args                args
		databaseThrowsError bool
		want                *repository.CompanyDocument
		wantErr             bool
	}{
		{
			name: "GetCompany function returns error",
			args: args{
				companyId: "12345",
			},
			databaseThrowsError: true,
			want:                nil,
			wantErr:             true,
		},
		{
			name: "get company without errors",
			args: args{
				companyId: "12345",
			},
			databaseThrowsError: false,
			want: &repository.CompanyDocument{
				Id:          "6507c3f0-6364-44fd-9a80-1190d9dc0446",
				Name:        "company name",
				Description: "corporate company",
				Amount:      700,
				Registered:  true,
				Type:        "Corporations",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			companyRepository := mocks.NewMockCompanyRepositoryInt(ctrl)
			if tt.databaseThrowsError {
				companyRepository.
					EXPECT().
					GetCompany(gomock.Any()).
					Return(nil, errors.New("database error"))
			} else {
				companyRepository.
					EXPECT().
					GetCompany(gomock.Any()).
					Return(tt.want, nil)
			}
			eventProducerMock := mocks.NewMockEventProducerInt(ctrl)

			companyService := service.NewCompanyService(companyRepository, eventProducerMock)
			result, err := companyService.GetCompany(tt.args.companyId)

			assert.Equal(t, tt.want, result)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestCreateCompany(t *testing.T) {
	ctrl := gomock.NewController(t)

	type args struct {
		clientRequest *model.CompanyEntity
	}
	tests := []struct {
		name                string
		args                args
		databaseThrowsError bool
		wantUuid            bool
		wantErr             bool
	}{
		{
			name: "CreateCompany function returns error",
			args: args{
				clientRequest: &model.CompanyEntity{},
			},
			databaseThrowsError: true,
			wantUuid:            false,
			wantErr:             true,
		},
		{
			name: "create company without errors",
			args: args{
				clientRequest: &model.CompanyEntity{},
			},
			databaseThrowsError: false,
			wantUuid:            true,
			wantErr:             false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			companyRepository := mocks.NewMockCompanyRepositoryInt(ctrl)
			if tt.databaseThrowsError {
				companyRepository.
					EXPECT().
					CreateCompany(gomock.Any()).
					Return("", errors.New("database error"))
			} else {
				companyRepository.
					EXPECT().
					CreateCompany(gomock.Any()).
					Return("466b525d-128d-4a52-83ae-8966fc9e6b1d", nil)
			}

			eventProducerMock := mocks.NewMockEventProducerInt(ctrl)

			companyService := service.NewCompanyService(companyRepository, eventProducerMock)
			result, err := companyService.CreateCompany(tt.args.clientRequest)

			assert.Equal(t, tt.wantErr, err != nil)

			if tt.wantUuid {
				_, uuidErr := uuid.Parse(result)
				assert.Equal(t, uuidErr, nil)
			} else {
				assert.Equal(t, result, "")
			}
		})
	}
}

func TestUpdateCompany(t *testing.T) {
	ctrl := gomock.NewController(t)

	type args struct {
		domainData *model.CompanyEntity
	}
	tests := []struct {
		name                string
		args                args
		databaseThrowsError bool
		wantUuid            bool
		wantErr             bool
	}{
		{
			name: "UpdateCompany function returns error",
			args: args{
				domainData: &model.CompanyEntity{},
			},
			databaseThrowsError: true,
			wantUuid:            false,
			wantErr:             true,
		},
		{
			name: "update company without errors",
			args: args{
				domainData: &model.CompanyEntity{
					Id:     "6507c3f0-6364-44fd-9a80-1190d9dc0446",
					Amount: 700,
				},
			},
			databaseThrowsError: false,
			wantUuid:            true,
			wantErr:             false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			companyRepository := mocks.NewMockCompanyRepositoryInt(ctrl)
			if tt.databaseThrowsError {
				companyRepository.
					EXPECT().
					UpdateCompany(gomock.Any()).
					Return(errors.New("database error"))
			} else {
				companyRepository.
					EXPECT().
					UpdateCompany(gomock.Any()).
					Return(nil)
			}
			eventProducerMock := mocks.NewMockEventProducerInt(ctrl)

			companyService := service.NewCompanyService(companyRepository, eventProducerMock)
			err := companyService.UpdateCompany(tt.args.domainData)

			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestDeleteCompany(t *testing.T) {
	ctrl := gomock.NewController(t)

	type args struct {
		companyId string
	}
	tests := []struct {
		name                string
		args                args
		databaseThrowsError bool
		wantErr             bool
	}{
		{
			name: "DeleteCompany function returns error",
			args: args{
				companyId: "12345",
			},
			databaseThrowsError: true,
			wantErr:             true,
		},
		{
			name: "delete company without errors",
			args: args{
				companyId: "12345",
			},
			databaseThrowsError: false,
			wantErr:             false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			companyRepository := mocks.NewMockCompanyRepositoryInt(ctrl)
			if tt.databaseThrowsError {
				companyRepository.
					EXPECT().
					DeleteCompany(gomock.Any()).
					Return(errors.New("database error"))
			} else {
				companyRepository.
					EXPECT().
					DeleteCompany(gomock.Any()).
					Return(nil)
			}

			eventProducerMock := mocks.NewMockEventProducerInt(ctrl)

			companyService := service.NewCompanyService(companyRepository, eventProducerMock)
			err := companyService.DeleteCompany(tt.args.companyId)

			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}
