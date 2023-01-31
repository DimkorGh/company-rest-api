package router

import (
	"github.com/gorilla/mux"

	"company-rest-api/internal/company/delivery"
	"company-rest-api/internal/core/auth"
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

	h.Router.HandleFunc("/company", h.compHandler.GetCompany).Methods("GET")
	h.Router.HandleFunc("/company", h.auth.Authenticate(h.compHandler.CreateCompany)).Methods("POST")
	h.Router.HandleFunc("/company", h.auth.Authenticate(h.compHandler.UpdateCompany)).Methods("PATCH")
	h.Router.HandleFunc("/company", h.auth.Authenticate(h.compHandler.DeleteCompany)).Methods("DELETE")
}
