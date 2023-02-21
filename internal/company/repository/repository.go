package repository

import (
	"encoding/json"
	"errors"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"

	"company-rest-api/internal/company/model"
	"company-rest-api/internal/core/database"
)

const companyCollection = "company"

type CompanyDocument struct {
	Id          string `bson:"_id"`
	Name        string `bson:"name,omitempty"`
	Description string `bson:"description,omitempty"`
	Amount      int    `bson:"amount,omitempty"`
	Registered  bool   `bson:"registered,omitempty"`
	Type        string `bson:"type,omitempty"`
}

type CompanyRepositoryInt interface {
	GetCompany(companyId string) (*CompanyDocument, error)
	CreateCompany(companyEntity *model.CompanyEntity) (string, error)
	UpdateCompany(companyEntity *model.CompanyEntity) error
	DeleteCompany(companyId string) error
}

type CompanyRepository struct {
	database database.DatabaseInt
}

func NewCompanyRepository(database database.DatabaseInt) *CompanyRepository {
	return &CompanyRepository{
		database: database,
	}
}

func (cr *CompanyRepository) GetCompany(companyId string) (*CompanyDocument, error) {
	filter := bson.D{{Key: "_id", Value: companyId}}

	companyDoc := &CompanyDocument{}

	err := cr.database.FindOne(companyCollection, filter, companyDoc)
	if err != nil {
		return nil, err
	}

	return companyDoc, nil
}

func (cr *CompanyRepository) CreateCompany(companyEntity *model.CompanyEntity) (string, error) {
	companyUuid := uuid.New().String()

	companyDoc := cr.build(companyEntity, companyUuid)

	err := cr.database.InsertOne(companyCollection, companyDoc)
	if err != nil {
		return "", err
	}

	return companyUuid, nil
}

func (cr *CompanyRepository) UpdateCompany(companyEntity *model.CompanyEntity) error {
	marshalledDomainData, err := json.Marshal(&companyEntity)
	if err != nil {
		logrus.Errorf("Error while marshalling company update domain data to json: %s", err.Error())
		return errors.New("internal error")
	}

	companyDoc := CompanyDocument{}
	err = json.Unmarshal([]byte(marshalledDomainData), &companyDoc)
	if err != nil {
		logrus.Errorf("Error while unmarshalling company update domain data json to company doc: %s", err.Error())
		return errors.New("internal error")
	}

	filter := bson.M{"_id": companyEntity.Id}
	update := bson.M{"$set": &companyDoc}

	err = cr.database.UpdateOne(companyCollection, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (cr *CompanyRepository) DeleteCompany(companyId string) error {
	filter := bson.D{{Key: "_id", Value: companyId}}

	err := cr.database.DeleteOne(companyCollection, filter)
	if err != nil {
		return err
	}

	return nil
}

func (cr *CompanyRepository) build(domainData *model.CompanyEntity, uuid string) *CompanyDocument {
	return &CompanyDocument{
		Id:          uuid,
		Name:        domainData.Name,
		Description: domainData.Description,
		Amount:      domainData.Amount,
		Registered:  domainData.Registered,
		Type:        domainData.Type,
	}
}
