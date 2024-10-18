// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "url": "http://www.swagger.io/support",
            "email": "suporte@swagger.io"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/health": {
            "get": {
                "description": "Check up app",
                "produces": [
                    "application/json"
                ],
                "summary": "Get Health",
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/dto.Health"
                        }
                    }
                }
            }
        },
        "/url-record": {
            "post": {
                "description": "Create one urlRecord in database",
                "produces": [
                    "application/json"
                ],
                "summary": "Post url",
                "responses": {
                    "201": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/dto.UrlRecord"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.Health": {
            "type": "object",
            "properties": {
                "environment": {
                    "type": "string"
                },
                "time": {
                    "type": "string"
                },
                "up": {
                    "type": "boolean"
                }
            }
        },
        "dto.UrlRecord": {
            "type": "object",
            "properties": {
                "url": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Shortify",
	Description:      "Shortener api",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
