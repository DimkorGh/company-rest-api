package parser

import (
	"encoding/json"
	"net/http"

	"company-rest-api/internal/utils/validators"
)

type JsonParserInt interface {
	ParseJson(r *http.Request, data interface{}) error
}

type JsonParser struct {
	structValidator validators.StructValidatorInt
}

func NewJsonParser(structValidator validators.StructValidatorInt) *JsonParser {
	return &JsonParser{
		structValidator: structValidator,
	}
}

func (jp *JsonParser) ParseJson(r *http.Request, data interface{}) error {
	requestBody := json.NewDecoder(r.Body)
	if err := requestBody.Decode(&data); err != nil {
		return &JsonParserError{err.Error()}
	}

	if err := jp.structValidator.Validate(data); err != nil {
		return err
	}

	return nil
}

type JsonParserError struct {
	ErrorMessage string
}

func (jpe *JsonParserError) Error() string {
	return "Error while parsing json body : " + jpe.ErrorMessage
}
