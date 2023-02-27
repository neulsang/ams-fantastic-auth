// Code generated by swaggo/swag. DO NOT EDIT
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
            "name": "API Support",
            "url": "https://github.com/neulsang",
            "email": "dgkwon90@gmail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/auth/login": {
            "post": {
                "description": "Login.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Login.",
                "parameters": [
                    {
                        "description": "Login infomation",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Login"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Token"
                        }
                    }
                }
            }
        },
        "/api/v1/auth/logout": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Logout.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Logout.",
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/auth/refresh": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Request a new access token.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Request a new access token.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Token"
                        }
                    }
                }
            }
        },
        "/api/v1/auth/register": {
            "post": {
                "description": "Create a new user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Create a new user.",
                "parameters": [
                    {
                        "description": "users infomation",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/model.UserResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/users": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get all exists users information.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Get all exists users information.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.UserResponse"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/users/me": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get my user information..",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Get my user information.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.UserResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/users/{id}": {
            "get": {
                "description": "Get user information by given ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Get user information by given ID.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id of the user",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.UserResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete user information by given ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Delete user information by given ID.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id of the user",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "patch": {
                "description": "Update user information.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Update user information.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id of the user",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "users infomation",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Login": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "maxLength": 255,
                    "example": "dgkwon90@naver.com"
                },
                "password": {
                    "type": "string",
                    "maxLength": 255,
                    "minLength": 8,
                    "example": "test1234"
                }
            }
        },
        "model.QnA": {
            "type": "object",
            "properties": {
                "answer": {
                    "type": "string",
                    "maxLength": 255,
                    "example": "blue"
                },
                "question": {
                    "type": "string",
                    "maxLength": 255,
                    "example": "What is my favorite color?"
                }
            }
        },
        "model.Token": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "model.User": {
            "type": "object",
            "properties": {
                "birth_date": {
                    "type": "string",
                    "example": "1990-07-29"
                },
                "email": {
                    "type": "string",
                    "maxLength": 255,
                    "example": "dgkwon90@naver.com"
                },
                "gender": {
                    "type": "string",
                    "enum": [
                        "male",
                        " female",
                        " other"
                    ],
                    "example": "male"
                },
                "name": {
                    "type": "string",
                    "maxLength": 255,
                    "example": "권대근"
                },
                "nick_name": {
                    "type": "string",
                    "maxLength": 36,
                    "example": "dgkwon90"
                },
                "password": {
                    "type": "string",
                    "maxLength": 255,
                    "example": "test1234"
                },
                "qna": {
                    "$ref": "#/definitions/model.QnA"
                }
            }
        },
        "model.UserResponse": {
            "type": "object",
            "properties": {
                "birthDate": {
                    "type": "string",
                    "example": "1990-07-29"
                },
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string",
                    "maxLength": 255,
                    "example": "dgkwon90@naver.com"
                },
                "gender": {
                    "type": "string",
                    "enum": [
                        "male",
                        " female",
                        " other"
                    ],
                    "example": "male"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string",
                    "maxLength": 255,
                    "example": "권대근"
                },
                "nick_name": {
                    "type": "string",
                    "maxLength": 36,
                    "example": "dgkwon90"
                },
                "qna": {
                    "$ref": "#/definitions/model.QnA"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "AMS Fantastic Auth Swagger API",
	Description:      "This is a Test auth api server",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
