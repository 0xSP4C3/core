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
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "email": "p3nj@bumpto.space"
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
        "/v1/book": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update book.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Book"
                ],
                "summary": "update book",
                "parameters": [
                    {
                        "description": "Book ID",
                        "name": "id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "Title",
                        "name": "title",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "Author",
                        "name": "author",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "User ID",
                        "name": "user_id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "Book status",
                        "name": "book_status",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "integer"
                        }
                    },
                    {
                        "description": "Book attributes",
                        "name": "book_attrs",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.BookAttrs"
                        }
                    }
                ],
                "responses": {
                    "202": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create a new book.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Book"
                ],
                "summary": "create a new book",
                "parameters": [
                    {
                        "description": "Title",
                        "name": "title",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "Author",
                        "name": "author",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "User ID",
                        "name": "user_id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "Book attributes",
                        "name": "book_attrs",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.BookAttrs"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Book"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Delete book by given ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Book"
                ],
                "summary": "delete book by given ID",
                "parameters": [
                    {
                        "description": "Book ID",
                        "name": "id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
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
            }
        },
        "/v1/book/{id}": {
            "get": {
                "description": "Get book by given ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Book"
                ],
                "summary": "get book by given ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Book ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Book"
                        }
                    }
                }
            }
        },
        "/v1/books": {
            "get": {
                "description": "Get all exists books.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Books"
                ],
                "summary": "get all exists books",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Book"
                            }
                        }
                    }
                }
            }
        },
        "/v1/coin": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update coin.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Coin"
                ],
                "summary": "update coin",
                "parameters": [
                    {
                        "description": "Coin ID",
                        "name": "id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "Name",
                        "name": "name",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "Code",
                        "name": "code",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "Description",
                        "name": "description",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "Exchange ID",
                        "name": "exchange_id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "Coin Uri",
                        "name": "coin_uri",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CoinUri"
                        }
                    }
                ],
                "responses": {
                    "202": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create a new coin.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Coin"
                ],
                "summary": "create a new coin",
                "parameters": [
                    {
                        "description": "Name",
                        "name": "name",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "Code",
                        "name": "code",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "Exchange ID",
                        "name": "exchange_id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "Coin Uri",
                        "name": "coin_uri",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CoinUri"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Coin"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Delete coin by given ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Coin"
                ],
                "summary": "delete coin by given ID",
                "parameters": [
                    {
                        "description": "Coin ID",
                        "name": "id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
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
            }
        },
        "/v1/coin/{id}": {
            "get": {
                "description": "Get coin by given ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Coin"
                ],
                "summary": "get coin by given ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Coin ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Coin"
                        }
                    }
                }
            }
        },
        "/v1/coins": {
            "get": {
                "description": "Get all exists coins.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Coins"
                ],
                "summary": "get all exists coins",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Coin"
                            }
                        }
                    }
                }
            }
        },
        "/v1/exchange": {
            "get": {
                "description": "Get all exists exchanges.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "exchange"
                ],
                "summary": "get all exists exchanges",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Exchange"
                            }
                        }
                    }
                }
            }
        },
        "/v1/exchange/{id}": {
            "get": {
                "description": "Get exchange by given ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Exchange"
                ],
                "summary": "get exchange by given ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Exchange ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Exchange"
                        }
                    }
                }
            }
        },
        "/v1/token/renew": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Renew access and refresh tokens.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Token"
                ],
                "summary": "renew access and refresh tokens",
                "parameters": [
                    {
                        "description": "Refresh token",
                        "name": "refresh_token",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
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
        "/v1/user/sign/in": {
            "post": {
                "description": "Auth user and return access and refresh token.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "auth user and return access and refresh token",
                "parameters": [
                    {
                        "description": "User Email",
                        "name": "email",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "User Password",
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
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/user/sign/out": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "De-authorize user and delete refresh token from Redis.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "de-authorize user and delete refresh token from Redis",
                "responses": {
                    "204": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/user/sign/up": {
            "post": {
                "description": "Create a new user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "create a new user",
                "parameters": [
                    {
                        "description": "Email",
                        "name": "email",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "Password",
                        "name": "password",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "User role",
                        "name": "user_role",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Book": {
            "type": "object",
            "required": [
                "author",
                "book_attrs",
                "book_status",
                "id",
                "title",
                "user_id"
            ],
            "properties": {
                "author": {
                    "type": "string",
                    "maxLength": 255
                },
                "book_attrs": {
                    "$ref": "#/definitions/models.BookAttrs"
                },
                "book_status": {
                    "type": "integer"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "title": {
                    "type": "string",
                    "maxLength": 255
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "models.BookAttrs": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "picture": {
                    "type": "string"
                },
                "rating": {
                    "type": "integer",
                    "maximum": 10,
                    "minimum": 1
                }
            }
        },
        "models.Coin": {
            "type": "object",
            "required": [
                "code",
                "coin_uri",
                "exchange_id",
                "id",
                "is_deleted",
                "name"
            ],
            "properties": {
                "code": {
                    "type": "string",
                    "maxLength": 10,
                    "minLength": 2
                },
                "coin_uri": {
                    "$ref": "#/definitions/models.CoinUri"
                },
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string",
                    "maxLength": 255
                },
                "exchange_id": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "image_uri": {
                    "type": "string"
                },
                "is_deleted": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string",
                    "maxLength": 25,
                    "minLength": 2
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.CoinUri": {
            "type": "object",
            "required": [
                "coin_id",
                "uri"
            ],
            "properties": {
                "coin_id": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "uri": {
                    "type": "string"
                }
            }
        },
        "models.Exchange": {
            "type": "object",
            "required": [
                "id",
                "is_blocked",
                "is_enabled",
                "name"
            ],
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string",
                    "maxLength": 255
                },
                "id": {
                    "type": "string"
                },
                "is_blocked": {
                    "description": "We got blocked?",
                    "type": "boolean"
                },
                "is_enabled": {
                    "description": "Enable Crawling?",
                    "type": "boolean"
                },
                "name": {
                    "type": "string",
                    "maxLength": 25
                },
                "updated_at": {
                    "type": "string"
                },
                "uri": {
                    "type": "string"
                }
            }
        },
        "models.User": {
            "type": "object",
            "required": [
                "email",
                "id",
                "password_hash",
                "user_role",
                "user_status"
            ],
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string",
                    "maxLength": 255
                },
                "id": {
                    "type": "string"
                },
                "password_hash": {
                    "type": "string",
                    "maxLength": 255
                },
                "updated_at": {
                    "type": "string"
                },
                "user_role": {
                    "type": "string",
                    "maxLength": 25
                },
                "user_status": {
                    "type": "integer"
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
	Title:            "core API",
	Description:      "This is an auto-generated API Docs.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
