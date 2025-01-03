// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Developer",
            "email": ""
        },
        "license": {
            "name": "MIT",
            "url": ""
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/history": {
            "get": {
                "description": "Returns user's transaction history",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "history"
                ],
                "summary": "getHistory",
                "parameters": [
                    {
                        "minimum": 1,
                        "type": "integer",
                        "example": 1,
                        "description": "user id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "minimum": 0,
                        "type": "integer",
                        "example": 10,
                        "description": "pagination limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "minimum": 1,
                        "type": "integer",
                        "example": 1,
                        "description": "pagination page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "boolean",
                        "example": true,
                        "description": "descending sort",
                        "name": "desc",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "date",
                        "description": "sort by",
                        "name": "order_by",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.History"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    }
                }
            }
        },
        "/order": {
            "post": {
                "description": "Creates, commits or rollbacks order",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "order"
                ],
                "summary": "orderHandle",
                "parameters": [
                    {
                        "description": "order info",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.orderPostRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.emptyJSONResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    }
                }
            }
        },
        "/report": {
            "get": {
                "description": "Returns link to report file",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "report"
                ],
                "summary": "createReport",
                "parameters": [
                    {
                        "minimum": 1900,
                        "type": "integer",
                        "example": 2022,
                        "description": "year",
                        "name": "year",
                        "in": "query",
                        "required": true
                    },
                    {
                        "maximum": 12,
                        "minimum": 1,
                        "type": "integer",
                        "example": 10,
                        "description": "month",
                        "name": "month",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.reportGetResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    }
                }
            }
        },
        "/reports/{name}": {
            "get": {
                "description": "Returns report file",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "report"
                ],
                "summary": "getReport",
                "parameters": [
                    {
                        "type": "string",
                        "description": "file name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/user": {
            "get": {
                "description": "Returns user's balance",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "getByID",
                "parameters": [
                    {
                        "minimum": 1,
                        "type": "integer",
                        "example": 1,
                        "description": "user id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Balance"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    }
                }
            },
            "post": {
                "description": "Makes new replenishment",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "increaseAmount",
                "parameters": [
                    {
                        "description": "user id and amount",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.userPostRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.emptyJSONResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.Balance": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                }
            }
        },
        "entity.History": {
            "type": "object",
            "properties": {
                "orders": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Order"
                    }
                }
            }
        },
        "entity.MyTime": {
            "type": "object",
            "properties": {
                "time": {
                    "type": "string"
                }
            }
        },
        "entity.Order": {
            "type": "object",
            "properties": {
                "service": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "sum": {
                    "type": "string"
                },
                "time": {
                    "$ref": "#/definitions/entity.MyTime"
                }
            }
        },
        "v1.emptyJSONResponse": {
            "type": "object"
        },
        "v1.orderPostRequest": {
            "type": "object",
            "required": [
                "action",
                "order_id",
                "service_id",
                "sum",
                "user_id"
            ],
            "properties": {
                "action": {
                    "type": "string",
                    "enum": [
                        "create",
                        "approve",
                        "cancel"
                    ],
                    "example": "create"
                },
                "order_id": {
                    "type": "integer",
                    "minimum": 1,
                    "example": 1
                },
                "service_id": {
                    "type": "integer",
                    "minimum": 1,
                    "example": 1
                },
                "sum": {
                    "type": "string",
                    "example": "200"
                },
                "user_id": {
                    "type": "integer",
                    "minimum": 1,
                    "example": 1
                }
            }
        },
        "v1.reportGetResponse": {
            "type": "object",
            "properties": {
                "link": {
                    "type": "string"
                }
            }
        },
        "v1.response": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "v1.userPostRequest": {
            "type": "object",
            "required": [
                "amount",
                "id"
            ],
            "properties": {
                "amount": {
                    "type": "string",
                    "example": "200"
                },
                "id": {
                    "type": "integer",
                    "minimum": 1,
                    "example": 1
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "Balance API",
	Description:      "Service for interactions with user's money accounts",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
