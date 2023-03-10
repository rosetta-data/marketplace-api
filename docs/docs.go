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
        "contact": {
            "name": "Evgeniy Dammer",
            "email": "evgeniydammer@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/categories/": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Create category method.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "categories"
                ],
                "summary": "Create category method.",
                "parameters": [
                    {
                        "description": "Category data",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/category.CreateCategoryInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Category ID",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    }
                }
            },
            "patch": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Update category method.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "categories"
                ],
                "summary": "Update category method.",
                "parameters": [
                    {
                        "description": "Category data",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/category.UpdateCategoryInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/http.StatusResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/categories/{org_id}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Get all categories method.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "categories"
                ],
                "summary": "Get all categories method.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Organization ID",
                        "name": "org_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Category List",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/category.Category"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/categories/{org_id}/{id}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Get category by id method.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "categories"
                ],
                "summary": "Get category by id method.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Organization ID",
                        "name": "org_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Category ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Category data",
                        "schema": {
                            "$ref": "#/definitions/category.Category"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Delete category method.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "categories"
                ],
                "summary": "Delete category method.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Organization ID",
                        "name": "org_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Category ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/http.StatusResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "category.Category": {
            "type": "object",
            "required": [
                "nameen",
                "nameru",
                "nametm",
                "nametr",
                "organization"
            ],
            "properties": {
                "id": {
                    "description": "Category ID",
                    "type": "string"
                },
                "level": {
                    "description": "Depth level",
                    "type": "integer"
                },
                "nameen": {
                    "description": "Name English",
                    "type": "string"
                },
                "nameru": {
                    "description": "Name Russian",
                    "type": "string"
                },
                "nametm": {
                    "description": "Name Turkmen",
                    "type": "string"
                },
                "nametr": {
                    "description": "Name Turkish",
                    "type": "string"
                },
                "organization": {
                    "description": "Organization ID",
                    "type": "string"
                },
                "parent": {
                    "description": "Parent category ID",
                    "type": "string"
                }
            }
        },
        "category.CreateCategoryInput": {
            "type": "object",
            "required": [
                "nameen",
                "nameru",
                "nametm",
                "nametr",
                "organization"
            ],
            "properties": {
                "level": {
                    "description": "Depth level",
                    "type": "integer"
                },
                "nameen": {
                    "description": "Name English",
                    "type": "string"
                },
                "nameru": {
                    "description": "Name Russian",
                    "type": "string"
                },
                "nametm": {
                    "description": "Name Turkmen",
                    "type": "string"
                },
                "nametr": {
                    "description": "Name Turkish",
                    "type": "string"
                },
                "organization": {
                    "description": "Organization ID",
                    "type": "string"
                },
                "parent": {
                    "description": "Parent category ID",
                    "type": "string"
                }
            }
        },
        "category.UpdateCategoryInput": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "Category ID",
                    "type": "string"
                },
                "level": {
                    "description": "Depth level",
                    "type": "integer"
                },
                "nameen": {
                    "description": "Name English",
                    "type": "string"
                },
                "nameru": {
                    "description": "Name Russian",
                    "type": "string"
                },
                "nametm": {
                    "description": "Name Turkmen",
                    "type": "string"
                },
                "nametr": {
                    "description": "Name Turkish",
                    "type": "string"
                },
                "organisation": {
                    "description": "Organization ID",
                    "type": "string"
                },
                "parent": {
                    "description": "Parent category ID",
                    "type": "string"
                }
            }
        },
        "http.ErrorResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "description": "Error message",
                    "type": "string"
                }
            }
        },
        "http.StatusResponse": {
            "type": "object",
            "properties": {
                "status": {
                    "description": "Status message",
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo_swagger holds exported Swagger Info so clients can modify it
var SwaggerInfo_swagger = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Emenu API",
	Description:      "Emenu API service",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate_swagger,
}

func init() {
	swag.Register(SwaggerInfo_swagger.InstanceName(), SwaggerInfo_swagger)
}
