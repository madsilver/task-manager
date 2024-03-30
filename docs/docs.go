// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Rodrigo Prata",
            "email": "rbpsilver@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/tasks": {
            "get": {
                "description": "Find all tasks owned by a user. If the user is a manager, it returns all tasks of all users.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tasks"
                ],
                "summary": "Find tasks.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "x-user-id",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Role",
                        "name": "x-role",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/presenter.Task"
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/presenter.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/presenter.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new task.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tasks"
                ],
                "summary": "Create task.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "x-user-id",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Role",
                        "name": "x-role",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": " ",
                        "name": "Task",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/presenter.TaskCreate"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/presenter.Task"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/presenter.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/presenter.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/presenter.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/tasks/{id}": {
            "get": {
                "description": "Find task by ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tasks"
                ],
                "summary": "Find task.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "x-user-id",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Role",
                        "name": "x-role",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Task ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/presenter.Task"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/presenter.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/presenter.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete task by ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tasks"
                ],
                "summary": "Delete task.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "x-user-id",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Role",
                        "name": "x-role",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Task ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/presenter.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/presenter.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/presenter.ErrorResponse"
                        }
                    }
                }
            },
            "patch": {
                "description": "Update task by ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tasks"
                ],
                "summary": "Update task.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "x-user-id",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Role",
                        "name": "x-role",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": " ",
                        "name": "Task",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/presenter.TaskCreate"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "Task ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/presenter.Task"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/presenter.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/presenter.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/presenter.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/tasks/{id}/close": {
            "patch": {
                "description": "Close task by ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tasks"
                ],
                "summary": "Close task.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "x-user-id",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Role",
                        "name": "x-role",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Task ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/presenter.Task"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/presenter.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/presenter.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/presenter.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "presenter.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "presenter.Task": {
            "type": "object",
            "properties": {
                "date": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "summary": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "presenter.TaskCreate": {
            "type": "object",
            "properties": {
                "summary": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8000",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "Task Manager",
	Description:      "API Developer Practical Exercise",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}