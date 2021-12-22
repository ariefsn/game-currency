// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "url": "https://www.ariefsn.dev",
            "email": "ayiexz22@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/": {
            "get": {
                "description": "Ping provider service",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Root"
                ],
                "summary": "Ping provider service",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/swagger.ResponseSuccessText"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/swagger.ResponseError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/swagger.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/swagger.ResponseError"
                        }
                    }
                }
            }
        },
        "/convert": {
            "post": {
                "description": "Start a new convertion",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Convert"
                ],
                "summary": "Start convert",
                "parameters": [
                    {
                        "description": "Payload",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/swagger.PayloadConvert"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/swagger.ResponseConvert"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/swagger.ResponseError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/swagger.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/swagger.ResponseError"
                        }
                    }
                }
            }
        },
        "/currency": {
            "get": {
                "description": "Get currency list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Currency"
                ],
                "summary": "Currency list",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Name",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Skip",
                        "name": "skip",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Orders",
                        "name": "orders",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/swagger.ResponseCurrencyList"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/swagger.ResponseError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/swagger.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/swagger.ResponseError"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new currency",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Currency"
                ],
                "summary": "Create currency",
                "parameters": [
                    {
                        "description": "Payload",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/swagger.PayloadCurrency"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/swagger.ResponseSuccessText"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/swagger.ResponseError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/swagger.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/swagger.ResponseError"
                        }
                    }
                }
            }
        },
        "/currency/{id}": {
            "get": {
                "description": "Get currency detail",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Currency"
                ],
                "summary": "Currency detail",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/swagger.ResponseCurrency"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/swagger.ResponseError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/swagger.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/swagger.ResponseError"
                        }
                    }
                }
            },
            "put": {
                "description": "Update existing currency",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Currency"
                ],
                "summary": "Currency update",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Payload",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/swagger.PayloadCurrency"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/swagger.ResponseSuccessText"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/swagger.ResponseError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/swagger.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/swagger.ResponseError"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete existing currency",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Currency"
                ],
                "summary": "Currency delete",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/swagger.ResponseSuccessText"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/swagger.ResponseError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/swagger.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/swagger.ResponseError"
                        }
                    }
                }
            }
        },
        "/rate": {
            "get": {
                "description": "Get rate list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Rate"
                ],
                "summary": "Rate list",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Skip",
                        "name": "skip",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Orders",
                        "name": "orders",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/swagger.ResponseRateList"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/swagger.ResponseError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/swagger.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/swagger.ResponseError"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new rate",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Rate"
                ],
                "summary": "Create rate",
                "parameters": [
                    {
                        "description": "Payload",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/swagger.PayloadRate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/swagger.ResponseSuccessText"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/swagger.ResponseError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/swagger.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/swagger.ResponseError"
                        }
                    }
                }
            }
        },
        "/rate/{id}": {
            "get": {
                "description": "Get rate detail",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Rate"
                ],
                "summary": "Rate detail",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/swagger.ResponseRate"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/swagger.ResponseError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/swagger.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/swagger.ResponseError"
                        }
                    }
                }
            },
            "put": {
                "description": "Update existing rate",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Rate"
                ],
                "summary": "Rate update",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Payload",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/swagger.PayloadRate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/swagger.ResponseSuccessText"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/swagger.ResponseError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/swagger.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/swagger.ResponseError"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete existing rate",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Rate"
                ],
                "summary": "Rate delete",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/swagger.ResponseSuccessText"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/swagger.ResponseError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/swagger.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/swagger.ResponseError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "swagger.ListMap": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "type": "object"
                    }
                },
                "total": {
                    "type": "integer",
                    "example": 0
                }
            }
        },
        "swagger.PayloadConvert": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number",
                    "example": 10
                },
                "fromIdCurrency": {
                    "type": "string",
                    "example": "string"
                },
                "toIdCurrency": {
                    "type": "string",
                    "example": "string"
                }
            }
        },
        "swagger.PayloadCurrency": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string",
                    "example": "string"
                }
            }
        },
        "swagger.PayloadRate": {
            "type": "object",
            "properties": {
                "fromIdCurrency": {
                    "type": "string",
                    "example": "string"
                },
                "rate": {
                    "type": "number",
                    "example": 10
                },
                "toIdCurrency": {
                    "type": "string",
                    "example": "string"
                }
            }
        },
        "swagger.ResponseConvert": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "number",
                    "example": 0
                },
                "message": {
                    "type": "string",
                    "example": "some error message"
                },
                "success": {
                    "type": "boolean",
                    "example": true
                }
            }
        },
        "swagger.ResponseCurrency": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    },
                    "example": {
                        "createdAt": "datetime",
                        "id": "string",
                        "name": "string",
                        "updatedAt": "datetime"
                    }
                },
                "message": {
                    "type": "string",
                    "example": "some error message"
                },
                "success": {
                    "type": "boolean",
                    "example": true
                }
            }
        },
        "swagger.ResponseCurrencyList": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/swagger.ListMap"
                },
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean",
                    "example": true
                }
            }
        },
        "swagger.ResponseError": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object"
                },
                "message": {
                    "type": "string",
                    "example": "some error message"
                },
                "success": {
                    "type": "boolean",
                    "example": false
                }
            }
        },
        "swagger.ResponseRate": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    },
                    "example": {
                        "createdAt": "datetime",
                        "fromIdCurrency": "string",
                        "id": "string",
                        "rate": "float64",
                        "toIdCurrency": "string",
                        "updatedAt": "datetime"
                    }
                },
                "message": {
                    "type": "string",
                    "example": "some error message"
                },
                "success": {
                    "type": "boolean",
                    "example": true
                }
            }
        },
        "swagger.ResponseRateList": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/swagger.ListMap"
                },
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean",
                    "example": true
                }
            }
        },
        "swagger.ResponseSuccessText": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "string",
                    "example": "some data"
                },
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean",
                    "example": true
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "3.0",
	Host:        "",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register("swagger", &s{})
}
