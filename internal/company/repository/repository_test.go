package repository_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"company-rest-api/internal/company/model"
	"company-rest-api/internal/company/repository"
	"company-rest-api/mocks"
)

func TestNewCompanyRepository(t *testing.T) {
	t.Run("constructor return new struct", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		db := mocks.NewMockDatabaseInt(ctrl)

		result := repository.NewCompanyRepository(db)

		assert.Implements(t, new(repository.CompanyRepositoryInt), result)
	})
}

func TestAddCompany(t *testing.T) {
	type args struct {
		companyEntity *model.CompanyEntity
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "create company with no errors",
			args: args{
				companyEntity: &model.CompanyEntity{
					Name:        "company name",
					Description: "corporate company",
					Amount:      700,
					Registered:  true,
					Type:        "Corporations",
				},
			},
			wantErr: false,
		},
		{
			name: "error returned from database",
			args: args{
				companyEntity: &model.CompanyEntity{
					Name:        "company name",
					Description: "corporate company",
					Amount:      700,
					Registered:  true,
					Type:        "Corporations",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			db := mocks.NewMockDatabaseInt(ctrl)
			if tt.wantErr {
				db.
					EXPECT().
					InsertOne(gomock.Any(), gomock.Any()).
					Return(errors.New(""))
			} else {
				db.
					EXPECT().
					InsertOne(gomock.Any(), gomock.Any()).
					Return(nil)
			}

			companyRepo := repository.NewCompanyRepository(db)
			result, err := companyRepo.CreateCompany(tt.args.companyEntity)

			if tt.wantErr {
				assert.Equal(t, tt.wantErr, err != nil)
				assert.Equal(t, "", result)
			} else {
				_, uuidErr := uuid.Parse(result)
				assert.Equal(t, uuidErr, nil)
				assert.Equal(t, nil, err)
			}
		})
	}
}

func TestGetCompany(t *testing.T) {
	type args struct {
		companyId string
	}
	tests := []struct {
		name    string
		args    args
		want    *repository.CompanyDocument
		wantErr bool
	}{
		{
			name: "get company with no errors",
			args: args{
				companyId: "12345",
			},
			want: &repository.CompanyDocument{
				Id:          "",
				Name:        "",
				Description: "",
				Amount:      0,
				Registered:  false,
				Type:        "",
			},
			wantErr: false,
		},
		{
			name: "error returned from database",
			args: args{
				companyId: "12345",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			db := mocks.NewMockDatabaseInt(ctrl)
			if tt.wantErr {
				db.
					EXPECT().
					FindOne(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(errors.New("database error"))
			} else {
				db.
					EXPECT().
					FindOne(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(nil)
			}

			companyRepo := repository.NewCompanyRepository(db)

			result, err := companyRepo.GetCompany(tt.args.companyId)

			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestUpdateCompany(t *testing.T) {
	type args struct {
		companyEntity *model.CompanyEntity
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "update company with no errors",
			args: args{
				companyEntity: &model.CompanyEntity{},
			},
			wantErr: false,
		},
		{
			name: "error returned from database",
			args: args{
				companyEntity: &model.CompanyEntity{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			db := mocks.NewMockDatabaseInt(ctrl)
			if tt.wantErr {
				db.
					EXPECT().
					UpdateOne(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(errors.New(""))
			} else {
				db.
					EXPECT().
					UpdateOne(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(nil)
			}

			companyRepo := repository.NewCompanyRepository(db)

			err := companyRepo.UpdateCompany(tt.args.companyEntity)

			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestDeleteCompany(t *testing.T) {
	type args struct {
		companyId string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "delete company with no errors",
			args: args{
				companyId: "12345",
			},
			wantErr: false,
		},
		{
			name: "error returned from database",
			args: args{
				companyId: "12345",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			db := mocks.NewMockDatabaseInt(ctrl)
			if tt.wantErr {
				db.
					EXPECT().
					DeleteOne(gomock.Any(), gomock.Any()).
					Return(errors.New("database error"))
			} else {
				db.
					EXPECT().
					DeleteOne(gomock.Any(), gomock.Any()).
					Return(nil)
			}

			companyRepo := repository.NewCompanyRepository(db)

			err := companyRepo.DeleteCompany(tt.args.companyId)

			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}
