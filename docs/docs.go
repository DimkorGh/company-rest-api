// Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Dimkor",
            "url": "https://github.com/DimkorGh"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/company": {
            "post": {
                "description": "Create a company with specific data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Company"
                ],
                "summary": "Create company",
                "parameters": [
                    {
                        "description": "Company data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/delivery.createCompanyRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/delivery.companyResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request - Company input data validation failed, Json parser error, company exists"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "delete": {
                "description": "Delete company",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Company"
                ],
                "summary": "Delete company",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Company ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/delivery.companyResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request - Company Id missing from url params"
                    },
                    "404": {
                        "description": "Not Found - No company with the specific id found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "patch": {
                "description": "Update all company data or specific fields",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Company"
                ],
                "summary": "Update company",
                "parameters": [
                    {
                        "description": "Company data for update",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/delivery.updateCompanyRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/delivery.companyResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request - Company input data validation failed, Json parser error"
                    },
                    "404": {
                        "description": "Not Found - No company with the specific id found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/company/{id}": {
            "get": {
                "description": "Return all data for a company",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Company"
                ],
                "summary": "Get company",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Company ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/delivery.companyResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request - Company Id missing from url params"
                    },
                    "404": {
                        "description": "Not Found - No company with the specific id found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/token": {
            "get": {
                "description": "Returns a jwt token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Generate jwt which lasts 5 minutes",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "delivery.companyResponse": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "errorMessage": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "registered": {
                    "type": "boolean"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "delivery.createCompanyRequest": {
            "type": "object",
            "required": [
                "amount",
                "name",
                "registered",
                "type"
            ],
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "description": {
                    "type": "string",
                    "maxLength": 3000
                },
                "name": {
                    "type": "string",
                    "maxLength": 15
                },
                "registered": {
                    "type": "boolean"
                },
                "type": {
                    "type": "string",
                    "enum": [
                        "Corporations",
                        "NonProfit",
                        "Cooperative",
                        "Sole Proprietorship"
                    ]
                }
            }
        },
        "delivery.updateCompanyRequest": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "description": {
                    "type": "string",
                    "maxLength": 3000
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string",
                    "maxLength": 15
                },
                "registered": {
                    "type": "boolean"
                },
                "type": {
                    "type": "string",
                    "enum": [
                        "Corporations",
                        "NonProfit",
                        "Cooperative",
                        "Sole Proprietorship",
                        ""
                    ]
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Bearer Token",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:9092",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Company REST API",
	Description:      "A simple REST API for a company entity",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
