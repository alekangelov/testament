// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {
            "name": "MIT",
            "url": "/license.md"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/v1/auth/login": {
            "post": {
                "description": "Login via email and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "Login",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.LoginResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.GenericResponse"
                        }
                    }
                }
            }
        },
        "/auth/v1/refresh-token": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Refresh Token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Refresh Token",
                "parameters": [
                    {
                        "description": "RefreshToken",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.RefreshTokenRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.LoginResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.GenericResponse"
                        }
                    }
                }
            }
        },
        "/auth/v1/register": {
            "post": {
                "description": "Register via email and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Register",
                "parameters": [
                    {
                        "description": "Register",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.RegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.GenericResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.GenericResponse"
                        }
                    }
                }
            }
        },
        "/auth/v1/verify": {
            "post": {
                "description": "Verify some type of token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Verify",
                "parameters": [
                    {
                        "description": "Verify",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.VerifyRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.GenericResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.GenericResponse"
                        }
                    }
                }
            }
        },
        "/auth/v1/verify-token": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Verify Token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Verify Token",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.GenericResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.GenericResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.GenericResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "models.LoginRequest": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "maxLength": 255,
                    "minLength": 3
                },
                "password": {
                    "type": "string",
                    "maxLength": 255,
                    "minLength": 3
                }
            }
        },
        "models.LoginResponse": {
            "type": "object",
            "properties": {
                "refreshToken": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "models.RefreshTokenRequest": {
            "type": "object",
            "required": [
                "token"
            ],
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "models.RegisterRequest": {
            "type": "object",
            "required": [
                "email",
                "password",
                "passwordConfirmation"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "maxLength": 255,
                    "minLength": 3
                },
                "password": {
                    "type": "string",
                    "maxLength": 255,
                    "minLength": 3
                },
                "passwordConfirmation": {
                    "type": "string"
                }
            }
        },
        "models.VerifyRequest": {
            "type": "object",
            "required": [
                "code",
                "type"
            ],
            "properties": {
                "code": {
                    "type": "string",
                    "maxLength": 6,
                    "minLength": 6
                },
                "type": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Testament API",
	Description:      "This is the API for Testament",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
