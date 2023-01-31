package response

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"company-rest-api/internal/core/database"
	"company-rest-api/internal/utils/parser"
	"company-rest-api/internal/utils/validators"
)

func TestSendResponse(t *testing.T) {
	type args struct {
		w    *httptest.ResponseRecorder
		body interface{}
		err  error
	}
	tests := []struct {
		name           string
		args           args
		wantStatusCode int
		wantBody       string
	}{
		{
			name: "response with no error and status 200 OK",
			args: args{
				w: httptest.NewRecorder(),
			},
			wantStatusCode: http.StatusOK,
		},
		{
			name: "response with no error and json body",
			args: args{
				w: httptest.NewRecorder(),
				body: struct {
					Name string `json:"name"`
				}{
					Name: "company name",
				},
				err: &database.NoDocumentsFoundError{},
			},
			wantStatusCode: http.StatusNotFound,
			wantBody:       `{"name":"company name"}` + "\n",
		},
		{
			name: "response with JsonParserError",
			args: args{
				w:   httptest.NewRecorder(),
				err: &parser.JsonParserError{},
			},
			wantStatusCode: http.StatusBadRequest,
		},
		{
			name: "response with StructValidatorError",
			args: args{
				w:   httptest.NewRecorder(),
				err: &validators.StructValidatorError{},
			},
			wantStatusCode: http.StatusBadRequest,
		},
		{
			name: "response with DuplicateKeyError",
			args: args{
				w:   httptest.NewRecorder(),
				err: &database.DuplicateKeyError{},
			},
			wantStatusCode: http.StatusBadRequest,
		},
		{
			name: "response with DatabaseError",
			args: args{
				w:   httptest.NewRecorder(),
				err: &database.DatabaseError{},
			},
			wantStatusCode: http.StatusInternalServerError,
		},
		{
			name: "response with NoDocumentsFoundError",
			args: args{
				w:   httptest.NewRecorder(),
				err: &database.NoDocumentsFoundError{},
			},
			wantStatusCode: http.StatusNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			responseBuilder := &ResponseBuilder{}
			responseBuilder.SendResponse(tt.args.w, tt.args.body, tt.args.err)

			if tt.wantBody != "" {
				assert.Equal(t, tt.wantBody, tt.args.w.Body.String())
			}

			assert.Equal(t, tt.wantStatusCode, tt.args.w.Code)
			assert.Equal(t, "application/json", tt.args.w.Header().Get("Content-Type"))
		})
	}
}
