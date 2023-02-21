package response

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"

	"company-rest-api/internal/core/database"
	"company-rest-api/internal/utils/parser"
	"company-rest-api/internal/utils/validators"
)

type ResponseBuilderInt interface {
	SendResponse(w http.ResponseWriter, body interface{}, err error)
}

type ResponseBuilder struct{}

func NewResponseBuilder() *ResponseBuilder {
	return &ResponseBuilder{}
}

func (rb *ResponseBuilder) SendResponse(w http.ResponseWriter, body interface{}, err error) {
	statusCode := rb.getStatusCode(err)

	err = rb.writeResponse(w, statusCode, body)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"body":       body,
			"statusCode": statusCode,
		}).Errorf("Error while writing response: %s", err.Error())
	}
}

func (rb *ResponseBuilder) writeResponse(w http.ResponseWriter, statusCode int, jsonResponse interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	return json.NewEncoder(w).Encode(jsonResponse)
}

func (rb *ResponseBuilder) getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch err.(type) {
	case *parser.JsonParserError,
		*validators.StructValidatorError,
		*database.DuplicateKeyError:
		return http.StatusBadRequest
	case *database.DatabaseError:
		return http.StatusInternalServerError
	case *database.NoDocumentsFoundError:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
