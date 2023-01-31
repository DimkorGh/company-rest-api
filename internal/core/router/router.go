package router

import (
	"github.com/gorilla/mux"
	"xmAssesment/internal/core/auth"
	"xmAssesment/internal/core/injector"
)

type HttpHandler struct {
	Router        *mux.Router
	injector      *injector.Injector
	authenticator *auth.Authenticator
}

func NewHttpHandler(
	router *mux.Router,
	injector *injector.Injector,
	authenticator *auth.Authenticator,
) *HttpHandler {
	return &HttpHandler{
		Router:        router,
		injector:      injector,
		authenticator: authenticator,
	}
}

func (httpHandler *HttpHandler) InitializeRouter() {
	httpHandler.injector.InitializeDependencies()

	httpHandler.Router.HandleFunc("/token", httpHandler.authenticator.GenerateToken).Methods("GET")
	httpHandler.Router.HandleFunc("/company", httpHandler.authenticator.Authenticate(httpHandler.injector.CompanyPostHandler.CreateCompany)).Methods("POST")
	httpHandler.Router.HandleFunc("/company", httpHandler.injector.CompanyGetHandler.GetCompany).Methods("GET")
	httpHandler.Router.HandleFunc("/company", httpHandler.authenticator.Authenticate(httpHandler.injector.CompanyUpdateHandler.UpdateCompany)).Methods("PATCH")
	httpHandler.Router.HandleFunc("/company", httpHandler.authenticator.Authenticate(httpHandler.injector.CompanyDeleteHandler.DeleteCompany)).Methods("DELETE")
}
