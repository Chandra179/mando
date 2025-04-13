// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/skills/add": {
            "post": {
                "description": "Add one or more skills",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "skills"
                ],
                "summary": "Add new skills",
                "parameters": [
                    {
                        "description": "Skills to add",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/skills.AddSkillRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "message: Skills added successfully",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "error: error message",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "skills.AddSkillRequest": {
            "description": "Model for adding skills",
            "type": "object",
            "required": [
                "skills"
            ],
            "properties": {
                "skills": {
                    "description": "List of skills to add",
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "Go"
                    ]
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api",
	Schemes:          []string{"http"},
	Title:            "Mando API",
	Description:      "This is the Mando server API documentation",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
