package delivery_test

import (
	"context"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mongodb"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"

	"company-rest-api/internal/company/delivery"
	"company-rest-api/internal/company/repository"
	"company-rest-api/internal/company/service"
	"company-rest-api/internal/core/auth"
	"company-rest-api/internal/core/database"
	"company-rest-api/internal/core/env"
	"company-rest-api/internal/core/eventProducer"
	"company-rest-api/internal/core/router"
	"company-rest-api/internal/utils/parser"
	"company-rest-api/internal/utils/response"
	"company-rest-api/internal/utils/validators"
	"company-rest-api/mocks"
)

const (
	companyEndpointUri = "/company"
	mockSecretKey      = "testKey"
	mockBearerToken    = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjIzMDc5NjMzMjV9.7HVTktBEaFUnlDq7tG2pXXb9mDfJFw68s6gcgQeByzk"
)

func TestNewCompanyRepository(t *testing.T) {
	t.Run("constructor return new struct", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		jsonParser := mocks.NewMockJsonParserInt(ctrl)
		urlParamParser := mocks.NewMockUrlParamsParserInt(ctrl)
		responseBuilder := mocks.NewMockResponseBuilderInt(ctrl)
		companyService := mocks.NewMockCompanyServiceInt(ctrl)

		result := delivery.NewCompanyHandler(jsonParser, urlParamParser, responseBuilder, companyService)

		assert.Implements(t, new(delivery.CompanyHandlerInt), result)
	})
}

func TestGetCompany(t *testing.T) {
	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}
	tests := []struct {
		name               string
		args               args
		expectedStatusCode int
		expectedBody       string
	}{
		{
			name: "valid company retrieval",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, companyEndpointUri+"?id=40ed576f-e56d-4867-aca3-ac1f30cc1409", nil),
			},
			expectedStatusCode: http.StatusOK,
			expectedBody:       `{"id":"40ed576f-e56d-4867-aca3-ac1f30cc1409","name":"randomCompanyName","description":"randomDescription","amount":21,"registered":true,"type":"Corporations"}`,
		},
		{
			name: "missing id from url params",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, companyEndpointUri, nil),
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedBody:       `{"errorMessage":"Key: 'getCompanyRequest.Id' Error:Field validation for 'Id' failed on the 'required' tag"}` + "\n",
		},
		{
			name: "invalid id to url params",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, companyEndpointUri+"?id=12345", nil),
			},
			expectedStatusCode: http.StatusNotFound,
			expectedBody:       `{"errorMessage":"mongo: no documents in result"}` + "\n",
		},
	}
	for _, tt := range tests {
		if testing.Short() {
			t.Skip("this is integration test so skip it while running unit tests")
		}

		t.Run(tt.name, func(t *testing.T) {
			initializeTestEnv()

			m := getMigration()
			_ = m.Up()

			createTestEndpoint(tt.args.w, tt.args.r)

			_ = m.Down()

			assert.Equal(t, tt.expectedStatusCode, tt.args.w.Code)
			assert.JSONEq(t, tt.expectedBody, tt.args.w.Body.String())
		})
	}
}

func TestCreateCompany(t *testing.T) {
	mockHeader := http.Header{}
	mockHeader.Set("Authorization", "Bearer "+mockBearerToken)

	type req struct {
		method  string
		body    string
		headers http.Header
	}
	type args struct {
		w *httptest.ResponseRecorder
		r req
	}
	tests := []struct {
		name               string
		args               args
		expectedStatusCode int
		expectedBody       string
	}{
		{
			name: "valid company creation",
			args: args{
				w: httptest.NewRecorder(),
				r: req{
					method: http.MethodPost,
					body: `{
								"name": "` + time.Now().Format("20060102150405") + `",
								"description": "new company",
								"amount": 700,
								"registered": true,
								"type": "Corporations"
							}`,
					headers: mockHeader,
				},
			},
			expectedStatusCode: http.StatusOK,
			expectedBody:       `{"id":"`,
		},
		{
			name: "unauthorized",
			args: args{
				w: httptest.NewRecorder(),
				r: req{
					method: http.MethodPost,
					body: `{
								"name": "` + time.Now().Format("01022006150405") + `",
								"description": "new company",
								"amount": 700,
								"registered": true,
								"type": "Corporations"
							}`,
					headers: nil,
				},
			},
			expectedStatusCode: http.StatusUnauthorized,
			expectedBody:       "",
		},
		{
			name: "invalid json",
			args: args{
				w: httptest.NewRecorder(),
				r: req{
					method: http.MethodPost,
					body: `{
								idpe "Corporations"
							}`,
					headers: mockHeader,
				},
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedBody:       `{"errorMessage":"Error while parsing json body : invalid character 'i' looking for beginning of object key string"}`,
		},
		{
			name: "required name field missing from client's request",
			args: args{
				w: httptest.NewRecorder(),
				r: req{
					method: http.MethodPost,
					body: `{
								"description": "new company",
								"amount": 700,
								"registered": true,
								"type": "Corporations"
							}`,
					headers: mockHeader,
				},
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedBody:       `{"errorMessage":"Key: 'createCompanyRequest.Name' Error:Field validation for 'Name' failed on the 'required' tag"}` + "\n",
		},
		{
			name: "company name with more than 15 characters",
			args: args{
				w: httptest.NewRecorder(),
				r: req{
					method: http.MethodPost,
					body: `{
								"name": "companyWithALongName",
								"description": "new company",
								"amount": 700,
								"registered": true,
								"type": "Corporations"
							}`,
					headers: mockHeader,
				},
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedBody:       `{"errorMessage":"Key: 'createCompanyRequest.Name' Error:Field validation for 'Name' failed on the 'max' tag"}` + "\n",
		},
		{
			name: "non required description field missing",
			args: args{
				w: httptest.NewRecorder(),
				r: req{
					method: http.MethodPost,
					body: `{
								"name": "` + time.Now().Format("200602011504") + `",
								"description": "company descr",
								"amount": 700,
								"registered": true,
								"type": "Corporations"
							}`,
					headers: mockHeader,
				},
			},
			expectedStatusCode: http.StatusOK,
			expectedBody:       `{"id":"`,
		},
		{
			name: "required amount field missing from client's request",
			args: args{
				w: httptest.NewRecorder(),
				r: req{
					method: http.MethodPost,
					body: `{
								"name": "company name",
								"description": "company descr",
								"registered": true,
								"type": "Corporations"
							}`,
					headers: mockHeader,
				},
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedBody:       `{"errorMessage":"Key: 'createCompanyRequest.Amount' Error:Field validation for 'Amount' failed on the 'required' tag"}` + "\n",
		},
		{
			name: "required registered field missing from client's request",
			args: args{
				w: httptest.NewRecorder(),
				r: req{
					method: http.MethodPost,
					body: `{
								"name": "company name",
								"description": "company descr",
								"amount": 700,
								"type": "Corporations"
							}`,
					headers: mockHeader,
				},
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedBody:       `{"errorMessage":"Key: 'createCompanyRequest.Registered' Error:Field validation for 'Registered' failed on the 'required' tag"}` + "\n",
		},
		{
			name: "required type field missing from client's request",
			args: args{
				w: httptest.NewRecorder(),
				r: req{
					method: http.MethodPost,
					body: `{
								"name": "company name",
								"description": "company descr",
								"amount": 700,
								"registered": true
							}`,
					headers: mockHeader,
				},
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedBody:       `{"errorMessage":"Key: 'createCompanyRequest.Type' Error:Field validation for 'Type' failed on the 'oneof' tag"}` + "\n",
		},
		{
			name: "invalid value to type field",
			args: args{
				w: httptest.NewRecorder(),
				r: req{
					method: http.MethodPost,
					body: `{
							"name": "company name",
							"description": "company descr",
							"amount": 700,
							"registered": true,
							"type": "Corporate"
						}`,
					headers: mockHeader,
				},
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedBody:       `{"errorMessage":"Key: 'createCompanyRequest.Type' Error:Field validation for 'Type' failed on the 'oneof' tag"}` + "\n",
		},
	}
	for _, tt := range tests {
		if testing.Short() {
			t.Skip("this is integration test so skip it while running unit tests")
		}

		t.Run(tt.name, func(t *testing.T) {
			initializeTestEnv()

			testReq := httptest.NewRequest(tt.args.r.method, companyEndpointUri, nil)
			testReq.Header = tt.args.r.headers
			testReq.Body = io.NopCloser(strings.NewReader(tt.args.r.body))

			createTestEndpoint(tt.args.w, testReq)

			assert.Equal(t, tt.expectedStatusCode, tt.args.w.Code)
			assert.Contains(t, tt.args.w.Body.String(), tt.expectedBody)
		})
	}
}

func TestUpdateCompany(t *testing.T) {
	mockHeader := http.Header{}
	mockHeader.Set("Authorization", "Bearer "+mockBearerToken)

	type req struct {
		method  string
		body    string
		headers http.Header
	}
	type args struct {
		w *httptest.ResponseRecorder
		r req
	}
	tests := []struct {
		name               string
		args               args
		expectedStatusCode int
		expectedBody       string
	}{
		{
			name: "valid company update",
			args: args{
				w: httptest.NewRecorder(),
				r: req{
					method: http.MethodPatch,
					body: `{
								"id": "40ed576f-e56d-4867-aca3-ac1f30cc1409",
								"name": "new company",
								"description": "new type",
								"amount": 700,
								"registered": true,
								"type": "NonProfit"
							}`,
					headers: mockHeader,
				},
			},
			expectedStatusCode: http.StatusOK,
			expectedBody:       "{}" + "\n",
		},
		{
			name: "unauthorized",
			args: args{
				w: httptest.NewRecorder(),
				r: req{
					method: http.MethodPatch,
					body: `{
								"id": "40ed576f-e56d-4867-aca3-ac1f30cc1409",
								"name": "new company",
								"description": "new type",
								"amount": 700,
								"registered": true,
								"type": "NonProfit"
							}`,
					headers: nil,
				},
			},
			expectedStatusCode: http.StatusUnauthorized,
			expectedBody:       "",
		},
		{
			name: "invalid json",
			args: args{
				w: httptest.NewRecorder(),
				r: req{
					method: http.MethodPatch,
					body: `{
								idpe "Corporations"
							}`,
					headers: mockHeader,
				},
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedBody:       `{"errorMessage":"Error while parsing json body : invalid character 'i' looking for beginning of object key string"}` + "\n",
		},
		{
			name: "invalid company id",
			args: args{
				w: httptest.NewRecorder(),
				r: req{
					method: http.MethodPatch,
					body: `{
								"id": "12345",
								"name": "new company",
								"description": "new type",
								"amount": 700,
								"registered": true,
								"type": "Corporations"
							}`,
					headers: mockHeader,
				},
			},
			expectedStatusCode: http.StatusNotFound,
			expectedBody:       `{"errorMessage":"company not found"}` + "\n",
		},
		{
			name: "company name with more than 15 characters",
			args: args{
				w: httptest.NewRecorder(),
				r: req{
					method: http.MethodPatch,
					body: `{
								"id": "2e554444-d802-4a72-81c1-d9a6f6f531f1",
								"name": "companyWithALongName"
							}`,
					headers: mockHeader,
				},
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedBody:       `{"errorMessage":"Key: 'updateCompanyRequest.Name' Error:Field validation for 'Name' failed on the 'max' tag"}` + "\n",
		},
		{
			name: "required id field missing from client's request",
			args: args{
				w: httptest.NewRecorder(),
				r: req{
					method: http.MethodPatch,
					body: `{
								"name": "company name",
								"description": "new type",
								"amount": 700,
								"registered": true
							}`,
					headers: mockHeader,
				},
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedBody:       `{"errorMessage":"Key: 'updateCompanyRequest.Id' Error:Field validation for 'Id' failed on the 'required' tag"}` + "\n",
		},
		{
			name: "invalid input value to type field",
			args: args{
				w: httptest.NewRecorder(),
				r: req{
					method: http.MethodPatch,
					body: `{
								"id": "2e554444-d802-4a72-81c1-d9a6f6f531f1",
								"name": "new company",
								"description": "new type",
							  	"amount": 700,
								"registered": true,
								"type": "Corporate"
							}`,
					headers: mockHeader,
				},
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedBody:       `{"errorMessage":"Key: 'updateCompanyRequest.Type' Error:Field validation for 'Type' failed on the 'oneof' tag"}` + "\n",
		},
	}
	for _, tt := range tests {
		if testing.Short() {
			t.Skip("this is integration test so skip it while running unit tests")
		}

		t.Run(tt.name, func(t *testing.T) {
			initializeTestEnv()

			testReq := httptest.NewRequest(tt.args.r.method, companyEndpointUri, nil)
			testReq.Header = tt.args.r.headers
			testReq.Body = io.NopCloser(strings.NewReader(tt.args.r.body))

			m := getMigration()
			_ = m.Up()

			createTestEndpoint(tt.args.w, testReq)

			_ = m.Down()

			assert.Equal(t, tt.expectedStatusCode, tt.args.w.Code)
			assert.Equal(t, tt.expectedBody, tt.args.w.Body.String())
		})
	}
}

func TestDeleteCompany(t *testing.T) {
	mockHeader := http.Header{}
	mockHeader.Set("Authorization", "Bearer "+mockBearerToken)

	type req struct {
		method  string
		uri     string
		headers http.Header
	}
	type args struct {
		w *httptest.ResponseRecorder
		r req
	}
	tests := []struct {
		name               string
		args               args
		expectedStatusCode int
		expectedBody       string
	}{
		{
			name: "valid company deletion",
			args: args{
				w: httptest.NewRecorder(),
				r: req{
					method:  http.MethodDelete,
					uri:     "?id=40ed576f-e56d-4867-aca3-ac1f30cc1409",
					headers: mockHeader,
				},
			},
			expectedStatusCode: http.StatusOK,
			expectedBody:       "{}\n",
		},
		{
			name: "unauthorized",
			args: args{
				w: httptest.NewRecorder(),
				r: req{
					method:  http.MethodDelete,
					uri:     "?id=40ed576f-e56d-4867-aca3-ac1f30cc1409",
					headers: nil,
				},
			},
			expectedStatusCode: http.StatusUnauthorized,
			expectedBody:       "",
		},
		{
			name: "company with specific id not found",
			args: args{
				w: httptest.NewRecorder(),
				r: req{
					method:  http.MethodDelete,
					uri:     "?id=12345",
					headers: mockHeader,
				},
			},
			expectedStatusCode: http.StatusNotFound,
			expectedBody:       `{"errorMessage":"No company with this id"}` + "\n",
		},
		{
			name: "missing id from url params",
			args: args{
				w: httptest.NewRecorder(),
				r: req{
					method:  http.MethodDelete,
					uri:     "",
					headers: mockHeader,
				},
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedBody:       `{"errorMessage":"Key: 'deleteCompanyRequest.Id' Error:Field validation for 'Id' failed on the 'required' tag"}` + "\n",
		},
	}
	for _, tt := range tests {
		if testing.Short() {
			t.Skip("this is integration test so skip it while running unit tests")
		}

		t.Run(tt.name, func(t *testing.T) {
			initializeTestEnv()

			testReq := httptest.NewRequest(tt.args.r.method, companyEndpointUri+tt.args.r.uri, nil)
			testReq.Header = tt.args.r.headers

			m := getMigration()
			_ = m.Up()

			createTestEndpoint(tt.args.w, testReq)

			_ = m.Down()

			assert.Equal(t, tt.expectedStatusCode, tt.args.w.Code)
			assert.Equal(t, tt.expectedBody, tt.args.w.Body.String())
		})
	}
}

func createTestEndpoint(w *httptest.ResponseRecorder, r *http.Request) {
	muxRouter := mux.NewRouter()
	authMiddleware := auth.NewAuthenticator(mockSecretKey)

	db := database.NewDatabase(context.Background(), os.Getenv("DATABASE_NAME"))
	db.Connect()

	eventProd := eventProducer.NewEventProducer()
	eventProd.Initialize()

	goValidator := validator.New()

	structValidator := validators.NewStructValidator(goValidator)
	respBuilder := response.NewResponseBuilder()
	jsonParser := parser.NewJsonParser(structValidator)
	urlParParser := parser.NewUrlParamsParser(structValidator)

	compRepo := repository.NewCompanyRepository(db)
	compService := service.NewCompanyService(compRepo, eventProd)
	companyHandler := delivery.NewCompanyHandler(jsonParser, urlParParser, respBuilder, compService)

	testRouter := router.NewHttpHandler(muxRouter, authMiddleware, companyHandler)
	testRouter.InitializeRouter()

	testRouter.Router.ServeHTTP(w, r)
}

func initializeTestEnv() {
	environment := env.NewEnvironment()
	environment.InitializeTestEnv()
}

func getMigration() *migrate.Migrate {
	m, err := migrate.New(
		"file://../../../migrations",
		"mongodb://"+
			os.Getenv("MONGO_USERNAME")+":"+
			os.Getenv("MONGO_PASSWORD")+"@"+
			os.Getenv("MONGO_ADDRESS")+"/"+
			os.Getenv("DATABASE_NAME")+
			"?authSource=admin",
	)
	if err != nil {
		log.Fatal(err.Error())
	}

	return m
}
