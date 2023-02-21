package dependencies

import (
	"context"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"

	"company-rest-api/internal/company/delivery"
	"company-rest-api/internal/company/repository"
	"company-rest-api/internal/company/service"
	"company-rest-api/internal/core/auth"
	"company-rest-api/internal/core/database"
	"company-rest-api/internal/core/env"
	"company-rest-api/internal/core/eventProducer"
	"company-rest-api/internal/core/log"
	"company-rest-api/internal/core/router"
	"company-rest-api/internal/utils/parser"
	"company-rest-api/internal/utils/response"
	"company-rest-api/internal/utils/validators"
)

type Container struct {
	ctx         context.Context
	HttpHandler *router.HttpHandler
}

func NewContainer(ctx context.Context) *Container {
	return &Container{
		ctx: ctx,
	}
}

func (c *Container) Initialize() {
	env.NewEnvironment().Initialize()
	log.NewLogger().Initialize()

	db := database.NewDatabase(c.ctx, os.Getenv("DATABASE_NAME"))
	db.Connect()

	eventProd := eventProducer.NewEventProducer()
	eventProd.Initialize()

	goValidator := validator.New()
	authMiddleware := auth.NewAuthenticator(os.Getenv("API_SECRET"))
	muxRouter := mux.NewRouter()

	structValidator := validators.NewStructValidator(goValidator)
	respBuilder := response.NewResponseBuilder()
	jsonParser := parser.NewJsonParser(structValidator)
	urlParParser := parser.NewUrlParamsParser(structValidator)

	compRepo := repository.NewCompanyRepository(db)
	compService := service.NewCompanyService(compRepo, eventProd)
	companyHandler := delivery.NewCompanyHandler(jsonParser, urlParParser, respBuilder, compService)

	c.HttpHandler = router.NewHttpHandler(muxRouter, authMiddleware, companyHandler)
	c.HttpHandler.InitializeRouter()
}
