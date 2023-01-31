package delivery

import (
	"net/http"

	"company-rest-api/internal/company/model"
	"company-rest-api/internal/company/repository"
	"company-rest-api/internal/company/service"
	"company-rest-api/internal/utils/parser"
	"company-rest-api/internal/utils/response"
)

type CompanyHandlerInt interface {
	GetCompany(w http.ResponseWriter, r *http.Request)
	CreateCompany(w http.ResponseWriter, r *http.Request)
	UpdateCompany(w http.ResponseWriter, r *http.Request)
	DeleteCompany(w http.ResponseWriter, r *http.Request)
}

type CompanyHandler struct {
	jsonParser      parser.JsonParserInt
	urlParamsParser parser.UrlParamsParserInt
	responseCreator response.ResponseBuilderInt
	companyService  service.CompanyServiceInt
}

func NewCompanyHandler(
	jsonParser parser.JsonParserInt,
	urlParamsParser parser.UrlParamsParserInt,
	responseBuilder response.ResponseBuilderInt,
	companyService service.CompanyServiceInt,
) *CompanyHandler {
	return &CompanyHandler{
		jsonParser:      jsonParser,
		urlParamsParser: urlParamsParser,
		responseCreator: responseBuilder,
		companyService:  companyService,
	}
}

func (ch *CompanyHandler) GetCompany(w http.ResponseWriter, r *http.Request) {
	var (
		err error
		res companyResponse
	)

	defer func() {
		if err != nil {
			res.ErrorMessage = err.Error()
		}
		ch.responseCreator.SendResponse(w, res, err)
	}()

	clientRequest := getCompanyRequest{}
	err = ch.urlParamsParser.ParseUrlParams(r, &clientRequest)
	if err != nil {
		return
	}

	companyDoc, err := ch.companyService.GetCompany(clientRequest.Id)
	if err != nil {
		return
	}

	res = ch.adaptToGetCompanyResponse(companyDoc)
}

func (ch *CompanyHandler) CreateCompany(w http.ResponseWriter, r *http.Request) {
	var (
		err error
		res companyResponse
	)

	defer func() {
		if err != nil {
			res.ErrorMessage = err.Error()
		}
		ch.responseCreator.SendResponse(w, res, err)
	}()

	clientRequest := createCompanyRequest{}
	err = ch.jsonParser.ParseJson(r, &clientRequest)
	if err != nil {
		return
	}

	domainData := ch.adaptToCompanyData(clientRequest)

	companyId, err := ch.companyService.CreateCompany(domainData)
	if err != nil {
		return
	}

	res.Id = companyId
}

func (ch *CompanyHandler) UpdateCompany(w http.ResponseWriter, r *http.Request) {
	var (
		err error
		res companyResponse
	)

	defer func() {
		if err != nil {
			res.ErrorMessage = err.Error()
		}
		ch.responseCreator.SendResponse(w, res, err)
	}()

	clientRequest := updateCompanyRequest{}
	err = ch.jsonParser.ParseJson(r, &clientRequest)
	if err != nil {
		return
	}

	domainData := ch.adaptToCompanyData(clientRequest)

	err = ch.companyService.UpdateCompany(domainData)
}

func (ch *CompanyHandler) DeleteCompany(w http.ResponseWriter, r *http.Request) {
	var (
		err error
		res companyResponse
	)

	defer func() {
		if err != nil {
			res.ErrorMessage = err.Error()
		}
		ch.responseCreator.SendResponse(w, res, err)
	}()

	clientRequest := deleteCompanyRequest{}
	err = ch.urlParamsParser.ParseUrlParams(r, &clientRequest)
	if err != nil {
		return
	}

	err = ch.companyService.DeleteCompany(clientRequest.Id)
}

func (ch *CompanyHandler) adaptToGetCompanyResponse(companyDoc *repository.CompanyDocument) companyResponse {
	return companyResponse{
		Id:          companyDoc.Id,
		Name:        companyDoc.Name,
		Description: companyDoc.Description,
		Amount:      companyDoc.Amount,
		Registered:  companyDoc.Registered,
		Type:        companyDoc.Type,
	}
}

func (ch *CompanyHandler) adaptToCompanyData(clientRequest interface{}) *model.CompanyEntity {
	if request, ok := clientRequest.(createCompanyRequest); ok {
		return &model.CompanyEntity{
			Name:        *request.Name,
			Description: *request.Description,
			Amount:      *request.Amount,
			Registered:  *request.Registered,
			Type:        *request.Type,
		}
	}

	request := clientRequest.(updateCompanyRequest)

	return &model.CompanyEntity{
		Id:          *request.Id,
		Name:        *request.Name,
		Description: *request.Description,
		Amount:      *request.Amount,
		Registered:  *request.Registered,
		Type:        *request.Type,
	}
}

type getCompanyRequest struct {
	Id string `schema:"id" validate:"required"`
}

type createCompanyRequest struct {
	Name        *string `json:"name" validate:"required,max=15"`
	Description *string `json:"description" validate:"max=3000"`
	Amount      *int    `json:"amount" validate:"required"`
	Registered  *bool   `json:"registered" validate:"required"`
	Type        *string `json:"type" validate:"oneof='Corporations' 'NonProfit' 'Cooperative' 'Sole Proprietorship',required"`
}

type updateCompanyRequest struct {
	Id          *string `json:"id" validate:"required"`
	Name        *string `json:"name,omitempty" validate:"omitempty,max=15"`
	Description *string `json:"description,omitempty" validate:"omitempty,max=3000"`
	Amount      *int    `json:"amount,omitempty"`
	Registered  *bool   `json:"registered,omitempty"`
	Type        *string `json:"type,omitempty" validate:"omitempty,oneof='Corporations' 'NonProfit' 'Cooperative' 'Sole Proprietorship' ''"`
}

type deleteCompanyRequest struct {
	Id string `schema:"id" validate:"required"`
}

type companyResponse struct {
	Id           string `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	Description  string `json:"description,omitempty"`
	Amount       int    `json:"amount,omitempty"`
	Registered   bool   `json:"registered,omitempty"`
	Type         string `json:"type,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}
