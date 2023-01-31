package parser

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"

	"company-rest-api/internal/utils/validators"
)

func TestParseJson(t *testing.T) {
	type request struct {
		FirstField  string `json:"firstField" validate:"required"`
		SecondField string `json:"secondField" validate:"required"`
	}

	type args struct {
		httpRequest *http.Request
		parseData   interface{}
	}
	tests := []struct {
		name   string
		args   args
		field1 string
		field2 string
		err    error
	}{
		{
			name: "invalid json to request body",
			args: args{
				httpRequest: httptest.NewRequest(
					"POST",
					"http://test.com",
					io.NopCloser(
						bytes.NewReader([]byte(`<definitely>not a json</definitely>`)),
					),
				),
			},
			err: &JsonParserError{ErrorMessage: "invalid character '<' looking for beginning of value"},
		},
		{
			name: "required field missing",
			args: args{
				httpRequest: httptest.NewRequest(
					"POST",
					"http://test.com",
					io.NopCloser(
						bytes.NewReader([]byte(`{"firstField":"firstValue"}`))),
				),
				parseData: &request{},
			},
			err: &validators.StructValidatorError{ErrorMessage: "Key: 'request.SecondField' Error:Field validation for 'SecondField' failed on the 'required' tag"},
		},
		{
			name: "valid json parsing",
			args: args{
				httpRequest: httptest.NewRequest(
					"POST",
					"http://test.com",
					io.NopCloser(
						bytes.NewReader([]byte(`{"firstField":"firstValue","secondField":"secondValue"}`))),
				),
				parseData: &request{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			strValidator := validators.NewStructValidator(validator.New())
			parser := NewJsonParser(strValidator)

			err := parser.ParseJson(tt.args.httpRequest, tt.args.parseData)

			assert.Equal(t, tt.err, err)
		})
	}
}
