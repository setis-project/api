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
            "name": "API Support",
            "email": "setisproject@gmail.com"
        },
        "license": {
            "name": "GPL-3.0 License",
            "url": "https://github.com/setis-project/api/blob/main/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/admin/account/login": {
            "post": {
                "security": [
                    {
                        "securitydefinitions.apikey": []
                    }
                ],
                "description": "login an admin",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "admin login",
                "parameters": [
                    {
                        "description": "account email",
                        "name": "email",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "account password",
                        "name": "password",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "400": {
                        "description": "Execution error",
                        "schema": {
                            "$ref": "#/definitions/models.ApiError"
                        }
                    },
                    "404": {
                        "description": "Execution error",
                        "schema": {
                            "$ref": "#/definitions/models.ApiError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.ApiError": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "models.ApiRequestError": {
            "type": "object",
            "properties": {
                "field": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "models.ApiRequestErrors": {
            "type": "object",
            "properties": {
                "errors": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.ApiRequestError"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "Setis Project API",
	Description:      "This is the Setis Project's API.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}