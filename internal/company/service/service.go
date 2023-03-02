package service

import (
	"encoding/json"
	"time"

	"company-rest-api/internal/company/event_message"
	"company-rest-api/internal/company/model"
	"company-rest-api/internal/company/repository"
	"company-rest-api/internal/core/eventProducer"
	"company-rest-api/internal/core/log"
)

type CompanyServiceInt interface {
	GetCompany(companyId string) (*repository.CompanyDocument, error)
	CreateCompany(companyEntity *model.CompanyEntity) (string, error)
	UpdateCompany(companyEntity *model.CompanyEntity) error
	DeleteCompany(companyId string) error
}

type CompanyService struct {
	logger            log.LoggerInt
	companyRepository repository.CompanyRepositoryInt
	eventProducer     eventProducer.EventProducerInt
}

func NewCompanyService(
	logger log.LoggerInt,
	companyRepository repository.CompanyRepositoryInt,
	eventProducer eventProducer.EventProducerInt,
) *CompanyService {
	return &CompanyService{
		logger:            logger,
		companyRepository: companyRepository,
		eventProducer:     eventProducer,
	}
}

func (cr *CompanyService) GetCompany(companyId string) (*repository.CompanyDocument, error) {
	companyDoc, err := cr.companyRepository.GetCompany(companyId)
	if err != nil {
		return nil, err
	}

	message := event_message.CompanyEventMessage{
		Type:      "get",
		Uuid:      companyDoc.Id,
		Timestamp: time.Now().String(),
	}

	jsonMessage, err := json.Marshal(message)
	if err != nil {
		cr.logger.Errorf("Error while marshalling company event message: %s", err.Error())

		return companyDoc, nil
	}

	go cr.eventProducer.Produce(jsonMessage, "company")

	return companyDoc, nil
}

func (cr *CompanyService) CreateCompany(companyEntity *model.CompanyEntity) (string, error) {
	companyUuid, err := cr.companyRepository.CreateCompany(companyEntity)
	if err != nil {
		return "", err
	}

	message := event_message.CompanyEventMessage{
		Type:      "create",
		Uuid:      companyUuid,
		Timestamp: time.Now().String(),
	}

	jsonMessage, err := json.Marshal(message)
	if err != nil {
		cr.logger.Errorf("Error while marshalling company event message: %s", err.Error())

		return companyUuid, nil
	}

	go cr.eventProducer.Produce(jsonMessage, "company")

	return companyUuid, nil
}

func (cr *CompanyService) UpdateCompany(companyEntity *model.CompanyEntity) error {
	err := cr.companyRepository.UpdateCompany(companyEntity)
	if err != nil {
		return err
	}

	message := event_message.CompanyEventMessage{
		Type:      "update",
		Uuid:      companyEntity.Id,
		Timestamp: time.Now().String(),
	}

	jsonMessage, err := json.Marshal(message)
	if err != nil {
		cr.logger.Errorf("Error while marshalling company event message: %s", err.Error())

		return nil
	}

	go cr.eventProducer.Produce(jsonMessage, "company")

	return nil
}

func (cr *CompanyService) DeleteCompany(companyId string) error {
	err := cr.companyRepository.DeleteCompany(companyId)
	if err != nil {
		return err
	}

	message := event_message.CompanyEventMessage{
		Type:      "delete",
		Uuid:      companyId,
		Timestamp: time.Now().String(),
	}

	jsonMessage, err := json.Marshal(message)
	if err != nil {
		cr.logger.Errorf("Error while marshalling company event message: %s", err.Error())

		return nil
	}

	go cr.eventProducer.Produce(jsonMessage, "company")

	return nil
}
