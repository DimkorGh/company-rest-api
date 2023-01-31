package injector

import (
	"context"
	"github.com/go-playground/validator/v10"
	"os"
	"xmAssesment/internal/api/parsers"
	"xmAssesment/internal/api/response"
	"xmAssesment/internal/api/validators"
	"xmAssesment/internal/company/domain"
	"xmAssesment/internal/company/presentation"
	"xmAssesment/internal/company/service"
	"xmAssesment/internal/core/database"
	"xmAssesment/internal/core/eventProducer"
)

type Injector struct {
	ctx                  context.Context
	CompanyPostHandler   presentation.CreateCompanyControllerInt
	CompanyGetHandler    presentation.GetCompanyControllerInt
	CompanyUpdateHandler presentation.UpdateCompanyControllerInt
	CompanyDeleteHandler presentation.DeleteCompanyControllerInt
}

func NewInjector(ctx context.Context) *Injector {
	return &Injector{
		ctx: ctx,
	}
}

func (injector *Injector) InitializeDependencies() {
	db := database.NewDatabase(os.Getenv("DATABASE_NAME"), injector.ctx)
	db.Connect()

	eventProd := eventProducer.NewEventProducer()
	eventProd.Initialize()

	structValidator := validators.NewStructValidator(validator.New())
	responseCreator := response.NewResponseBuilder()
	jsonParser := parsers.NewJsonParser(structValidator)
	urlParamsParser := parsers.NewUrlParamsParser(structValidator)

	injector.CompanyPostHandler = presentation.NewCreateCompanyController(
		jsonParser,
		responseCreator,
		service.NewCreateCompanyService(
			domain.NewCompanyRepository(db),
			eventProd,
		),
	)

	injector.CompanyGetHandler = presentation.NewGetCompanyController(
		urlParamsParser,
		responseCreator,
		service.NewGetCompanyService(
			domain.NewCompanyRepository(db),
		),
	)

	injector.CompanyUpdateHandler = presentation.NewUpdateCompanyController(
		jsonParser,
		responseCreator,
		service.NewUpdateCompanyService(
			domain.NewCompanyRepository(db),
			eventProd,
		),
	)

	injector.CompanyDeleteHandler = presentation.NewDeleteCompanyController(
		urlParamsParser,
		responseCreator,
		service.NewDeleteCompanyService(
			domain.NewCompanyRepository(db),
			eventProd,
		),
	)
}
