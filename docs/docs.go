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
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/admin/add-category": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Adds a new category",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "api/admin"
                ],
                "summary": "AddCategory",
                "operationId": "adds category",
                "parameters": [
                    {
                        "description": "category info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/server.Category"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.ItemProcessedResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.Error"
                        }
                    }
                }
            }
        },
        "/api/admin/add-product": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Adds a new product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "api/admin"
                ],
                "summary": "AddProduct",
                "operationId": "adds product",
                "parameters": [
                    {
                        "description": "product info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/server.ProductInfoDescription"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.ItemProcessedResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.Error"
                        }
                    }
                }
            }
        },
        "/api/admin/all-orders": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "get orders by status",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "api/admin"
                ],
                "summary": "Get Orders",
                "operationId": "get orders",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Last item created at for pagination",
                        "name": "created_at",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Sort by orders status",
                        "name": "order_status",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/server.Order"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.Error"
                        }
                    }
                }
            }
        },
        "/api/admin/delete-category": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Deletes category",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "api/admin"
                ],
                "summary": "DeleteCategory",
                "operationId": "deletes category",
                "parameters": [
                    {
                        "description": "category info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.DeleteCategoryInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.ItemProcessedResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.Error"
                        }
                    }
                }
            }
        },
        "/api/admin/delete-product": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Deletes product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "api/admin"
                ],
                "summary": "DeleteProduct",
                "operationId": "delete product",
                "parameters": [
                    {
                        "description": "product id",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.ProductIdInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.ItemProcessedResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.Error"
                        }
                    }
                }
            }
        },
        "/api/admin/new-moderator": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "registers a new moderator",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "api/admin"
                ],
                "summary": "CreateModerator",
                "operationId": "create account for moderator",
                "parameters": [
                    {
                        "description": "account info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/server.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.Error"
                        }
                    }
                }
            }
        },
        "/api/admin/update-category": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Updates category",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "api/admin"
                ],
                "summary": "UpdateCategory",
                "operationId": "updates category",
                "parameters": [
                    {
                        "description": "category info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/server.Category"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.ItemProcessedResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.Error"
                        }
                    }
                }
            }
        },
        "/api/admin/update-product": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Updates product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "api/admin"
                ],
                "summary": "UpdateProduct",
                "operationId": "updates product",
                "parameters": [
                    {
                        "description": "product info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/server.ProductInfoDescription"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.ItemProcessedResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.Error"
                        }
                    }
                }
            }
        },
        "/api/all-categories": {
            "get": {
                "description": "get all categories",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "api"
                ],
                "summary": "GetCategories",
                "operationId": "get categories",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.allCategoriesResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.Error"
                        }
                    }
                }
            }
        },
        "/api/client/logout": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "ends session",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "api/client"
                ],
                "summary": "Logout",
                "operationId": "logout from account",
                "parameters": [
                    {
                        "description": "account info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/server.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.Error"
                        }
                    }
                }
            }
        },
        "/api/client/refresh": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Gets a new access using refreshToken",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "api/client"
                ],
                "summary": "Refresh",
                "operationId": "refreshToken from account",
                "parameters": [
                    {
                        "description": "refresh token",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.RefreshTokensInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.Error"
                        }
                    }
                }
            }
        },
        "/api/get-products": {
            "get": {
                "description": "get products from certain category with params",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "api"
                ],
                "summary": "GetProducts",
                "operationId": "gets products",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Category",
                        "name": "category",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Size",
                        "name": "size",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Type",
                        "name": "type",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Search",
                        "name": "search",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Price",
                        "name": "price",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "characteristic",
                        "name": "characteristic",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Created At",
                        "name": "created_at",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/server.Product"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.Error"
                        }
                    }
                }
            }
        },
        "/api/product": {
            "get": {
                "description": "get product full info by its id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "api"
                ],
                "summary": "GetProductById",
                "operationId": "gets full product info",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Product id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/server.ProductInfoDescription"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.Error"
                        }
                    }
                }
            }
        },
        "/auth/sign-in": {
            "post": {
                "description": "signs in account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "SignIn",
                "operationId": "sign-in account",
                "parameters": [
                    {
                        "description": "sign in account info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/server.SignInInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.Error"
                        }
                    }
                }
            }
        },
        "/auth/sign-up": {
            "post": {
                "description": "registers a new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "SignUp",
                "operationId": "create account",
                "parameters": [
                    {
                        "description": "account info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/server.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.DeleteCategoryInput": {
            "type": "object",
            "properties": {
                "category_title": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                }
            }
        },
        "handler.Error": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "handler.ItemProcessedResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "handler.ProductIdInput": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "handler.RefreshTokensInput": {
            "type": "object",
            "properties": {
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "handler.allCategoriesResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/server.Category"
                    }
                }
            }
        },
        "server.Category": {
            "type": "object",
            "required": [
                "category_title",
                "img_url"
            ],
            "properties": {
                "category_title": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "img_url": {
                    "type": "string"
                }
            }
        },
        "server.Order": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "delivery_type_title": {
                    "type": "string"
                },
                "order_comment": {
                    "type": "string"
                },
                "order_status": {
                    "type": "string"
                },
                "order_sum_price": {
                    "type": "number"
                },
                "payment_type_title": {
                    "type": "string"
                },
                "user_email": {
                    "type": "string"
                },
                "user_firstname": {
                    "type": "string"
                },
                "user_lastname": {
                    "type": "string"
                },
                "user_middle_name": {
                    "type": "string"
                },
                "user_phone_number": {
                    "type": "string"
                }
            }
        },
        "server.Product": {
            "type": "object",
            "properties": {
                "amount_in_stock": {
                    "type": "number",
                    "example": 120
                },
                "article": {
                    "type": "string",
                    "example": "000123"
                },
                "category_title": {
                    "type": "string",
                    "example": "Одноразові стакани"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer",
                    "example": 5
                },
                "img_url": {
                    "type": "string",
                    "example": "https://google-images.com/some-img123"
                },
                "packages_in_box": {
                    "type": "integer",
                    "example": 50
                },
                "price": {
                    "type": "number",
                    "example": 3.75
                },
                "product_title": {
                    "type": "string",
                    "example": "Стакан одноразовий Крафт 110мл"
                },
                "type_title": {
                    "type": "string",
                    "example": "Стакан"
                },
                "units_in_package": {
                    "type": "integer",
                    "example": 30
                }
            }
        },
        "server.ProductInfo": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string",
                    "example": "Білий"
                },
                "info_title": {
                    "type": "string",
                    "example": "Колір"
                },
                "product_id": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "server.ProductInfoDescription": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/server.ProductInfo"
                    }
                },
                "info": {
                    "$ref": "#/definitions/server.Product"
                }
            }
        },
        "server.SignInInput": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "server.User": {
            "type": "object",
            "required": [
                "email",
                "name",
                "password",
                "phone_number"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                },
                "role_id": {
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
	Host:             "localhost:8000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "AllianceCup API",
	Description:      "API Server for AllianceCup Application",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
