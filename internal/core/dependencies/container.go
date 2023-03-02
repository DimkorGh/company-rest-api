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
	"company-rest-api/internal/core/config"
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
	Conf        *config.Config
	Logger      log.LoggerInt
	HttpHandler *router.HttpHandler
}

func NewContainer(ctx context.Context) *Container {
	return &Container{
		ctx: ctx,
	}
}

func (c *Container) Initialize() {
	env.NewEnvironment().Initialize()
	c.Conf = config.NewConfig().Load(os.Getenv("CONFIG_NAME"), os.Getenv("CONFIG_PATH"))

	c.Logger = log.NewLogger(c.Conf)
	c.Logger.Initialize()

	db := database.NewDatabase(c.ctx, c.Conf, c.Logger)
	db.Connect()

	eventProd := eventProducer.NewEventProducer(c.Conf, c.Logger)
	eventProd.Initialize()

	goValidator := validator.New()
	authMiddleware := auth.NewAuthenticator(c.Conf.Server.JwtSecretKey)
	muxRouter := mux.NewRouter()

	structValidator := validators.NewStructValidator(goValidator)
	respBuilder := response.NewResponseBuilder(c.Logger)
	jsonParser := parser.NewJsonParser(structValidator)
	urlParParser := parser.NewUrlParamsParser(structValidator)

	compRepo := repository.NewCompanyRepository(c.Logger, db)
	compService := service.NewCompanyService(c.Logger, compRepo, eventProd)
	companyHandler := delivery.NewCompanyHandler(jsonParser, urlParParser, respBuilder, compService)

	c.HttpHandler = router.NewHttpHandler(muxRouter, authMiddleware, companyHandler)
	c.HttpHandler.InitializeRouter()
}
