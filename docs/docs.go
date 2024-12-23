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
        "/addproduct": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Adding a product with a name and price",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ProductsActions"
                ],
                "summary": "Add new product",
                "parameters": [
                    {
                        "description": "Product",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/productsActions.CredentialsForAddProduct"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Новый товар успешно добавлен"
                    }
                }
            }
        },
        "/delorrestoreproduct": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Allows you to remove or mark the deletion",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ProductsActions"
                ],
                "summary": "Delete or restore product",
                "parameters": [
                    {
                        "description": "Product",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/productsActions.CredentialsForDelProduct"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Пометка удаления успешно изменена"
                    }
                }
            }
        },
        "/getallproducts": {
            "get": {
                "description": "Returns a list of all products",
                "tags": [
                    "ProductsActions"
                ],
                "summary": "Get products",
                "responses": {}
            }
        },
        "/login": {
            "post": {
                "description": "Login User with name and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Autentification"
                ],
                "summary": "Login User",
                "parameters": [
                    {
                        "description": "User",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/userActions.Credentials"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Успешный вход пользователя ",
                        "schema": {
                            "$ref": "#/definitions/userActions.TokenResponse"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "Register User with name and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Registration"
                ],
                "summary": "Register User",
                "parameters": [
                    {
                        "description": "User",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/userActions.Credentials"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Успешная регистрация нового пользователя"
                    }
                }
            }
        }
    },
    "definitions": {
        "productsActions.CredentialsForAddProduct": {
            "type": "object",
            "properties": {
                "Название": {
                    "type": "string"
                },
                "Описание": {
                    "type": "string"
                },
                "цена": {
                    "type": "number"
                }
            }
        },
        "productsActions.CredentialsForDelProduct": {
            "type": "object",
            "properties": {
                "Название": {
                    "type": "string"
                },
                "Пометка удаления": {
                    "type": "boolean"
                }
            }
        },
        "userActions.Credentials": {
            "type": "object",
            "properties": {
                "Логин": {
                    "type": "string"
                },
                "Пароль": {
                    "type": "string"
                }
            }
        },
        "userActions.TokenResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "IndustrialBack",
	Description:      "This is a project for industrial technogies",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
