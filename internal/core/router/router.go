package router

import (
	"company-rest-api/internal/company/delivery"
	"company-rest-api/internal/core/auth"
	"net/http"

	_ "company-rest-api/docs"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

type HttpHandler struct {
	Router      *mux.Router
	auth        *auth.Authenticator
	compHandler delivery.CompanyHandlerInt
}

func NewHttpHandler(
	router *mux.Router,
	auth *auth.Authenticator,
	compHandler delivery.CompanyHandlerInt,
) *HttpHandler {
	return &HttpHandler{
		Router:      router,
		auth:        auth,
		compHandler: compHandler,
	}
}

func (h *HttpHandler) InitializeRouter() {
	h.Router.HandleFunc("/token", h.auth.GenerateToken).Methods("GET")

	h.Router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	h.Router.HandleFunc("/company", h.compHandler.GetCompany).Methods(http.MethodGet)
	h.Router.HandleFunc("/company", h.auth.Authenticate(h.compHandler.CreateCompany)).Methods(http.MethodPost)
	h.Router.HandleFunc("/company", h.auth.Authenticate(h.compHandler.UpdateCompany)).Methods(http.MethodPatch)
	h.Router.HandleFunc("/company", h.auth.Authenticate(h.compHandler.DeleteCompany)).Methods(http.MethodDelete)
}
