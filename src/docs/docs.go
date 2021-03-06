// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate_swagger = `{
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
        "/company": {
            "get": {
                "description": "Get all companies",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "company"
                ],
                "summary": "Get companies",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseModel"
                        }
                    }
                }
            },
            "post": {
                "description": "Enpoint to create a company",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "company"
                ],
                "summary": "Create company",
                "parameters": [
                    {
                        "description": "create company",
                        "name": "company",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/command.CommandRegisterCompany"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseModel"
                        }
                    }
                }
            }
        },
        "/publication": {
            "get": {
                "description": "Get all publications",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "publication"
                ],
                "summary": "Get publications",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseModel"
                        }
                    }
                }
            },
            "post": {
                "description": "Enpoint to create a publication",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "publication"
                ],
                "summary": "Create publication",
                "parameters": [
                    {
                        "description": "create publication",
                        "name": "publication",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/command.CommandCreatePublication"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseModel"
                        }
                    }
                }
            }
        },
        "/user": {
            "get": {
                "description": "Get all users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Get users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseModel"
                        }
                    }
                }
            },
            "post": {
                "description": "Enpoint to create a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Create user",
                "parameters": [
                    {
                        "description": "create user",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/command.CommandRegisterUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseModel"
                        }
                    }
                }
            },
            "delete": {
                "description": "Enpoint to delete a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Delete user",
                "parameters": [
                    {
                        "description": "delete user",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/command.CommandDeleteUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseModel"
                        }
                    }
                }
            },
            "patch": {
                "description": "Enpoint to update a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Update user",
                "parameters": [
                    {
                        "description": "update user",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/command.CommandEditUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseModel"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "Enpoint to login a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Login of user",
                "parameters": [
                    {
                        "description": "login user",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/command.CommandLoginUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseModel"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "command.CommandCreatePublication": {
            "type": "object",
            "required": [
                "Content",
                "Description",
                "Title",
                "Type"
            ],
            "properties": {
                "Content": {
                    "type": "string",
                    "minLength": 90
                },
                "Description": {
                    "type": "string",
                    "minLength": 20
                },
                "Title": {
                    "type": "string",
                    "minLength": 5
                },
                "Type": {
                    "type": "string"
                },
                "WiterUserId": {
                    "type": "integer"
                }
            }
        },
        "command.CommandDeleteUser": {
            "type": "object",
            "required": [
                "Email",
                "Password"
            ],
            "properties": {
                "Email": {
                    "type": "string",
                    "minLength": 5
                },
                "Password": {
                    "type": "string",
                    "minLength": 6
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "command.CommandEditUser": {
            "type": "object",
            "properties": {
                "Email": {
                    "type": "string"
                },
                "Name": {
                    "type": "string"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "command.CommandLoginUser": {
            "type": "object",
            "properties": {
                "Email": {
                    "type": "string"
                },
                "Password": {
                    "type": "string",
                    "minLength": 6
                }
            }
        },
        "command.CommandRegisterCompany": {
            "type": "object",
            "required": [
                "Email",
                "Name",
                "Owner",
                "Phone"
            ],
            "properties": {
                "Email": {
                    "type": "string",
                    "minLength": 5
                },
                "Name": {
                    "type": "string",
                    "minLength": 5
                },
                "Owner": {
                    "type": "string",
                    "minLength": 5
                },
                "Phone": {
                    "type": "string",
                    "minLength": 7
                }
            }
        },
        "command.CommandRegisterUser": {
            "type": "object",
            "required": [
                "CompanyId",
                "Email",
                "Name",
                "Password",
                "Role"
            ],
            "properties": {
                "CompanyId": {
                    "type": "integer"
                },
                "Email": {
                    "type": "string",
                    "minLength": 5
                },
                "Name": {
                    "type": "string",
                    "minLength": 5
                },
                "Password": {
                    "type": "string",
                    "minLength": 8
                },
                "Role": {
                    "type": "string"
                }
            }
        },
        "response.ResponseModel": {
            "type": "object",
            "properties": {
                "data": {},
                "error": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo_swagger holds exported Swagger Info so clients can modify it
var SwaggerInfo_swagger = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate_swagger,
}

func init() {
	swag.Register(SwaggerInfo_swagger.InstanceName(), SwaggerInfo_swagger)
}
