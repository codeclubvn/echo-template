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
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
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
        "/v1/api/files": {
            "put": {
                "description": "Update",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "File"
                ],
                "summary": "Update",
                "parameters": [
                    {
                        "type": "string",
                        "description": "authorization token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/entity.SimpleResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Upload",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "File"
                ],
                "summary": "Upload",
                "parameters": [
                    {
                        "type": "string",
                        "description": "authorization token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/entity.SimpleResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "File"
                ],
                "summary": "Delete",
                "parameters": [
                    {
                        "type": "string",
                        "description": "authorization token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/entity.SimpleResponse"
                        }
                    }
                }
            }
        },
        "/v1/api/files/:id": {
            "get": {
                "description": "GetOne",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "File"
                ],
                "summary": "GetOne",
                "parameters": [
                    {
                        "type": "string",
                        "description": "authorization token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/entity.SimpleResponse"
                        }
                    }
                }
            }
        },
        "/v1/api/files/download/:id": {
            "get": {
                "description": "Download",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "File"
                ],
                "summary": "Download",
                "parameters": [
                    {
                        "type": "string",
                        "description": "authorization token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/entity.SimpleResponse"
                        }
                    }
                }
            }
        },
        "/v1/api/health": {
            "get": {
                "description": "Health",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Health"
                ],
                "summary": "Health",
                "parameters": [
                    {
                        "type": "string",
                        "description": "authorization token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/entity.SimpleResponse"
                        }
                    }
                }
            }
        },
        "/v1/api/image/upload": {
            "post": {
                "description": "Upload",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Image"
                ],
                "summary": "Upload",
                "parameters": [
                    {
                        "type": "string",
                        "description": "authorization token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/entity.SimpleResponse"
                        }
                    }
                }
            }
        },
        "/v1/api/posts": {
            "get": {
                "description": "GetList",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Post"
                ],
                "summary": "GetList",
                "parameters": [
                    {
                        "type": "string",
                        "description": "authorization token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/entity.SimpleResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Update",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Post"
                ],
                "summary": "Update",
                "parameters": [
                    {
                        "type": "string",
                        "description": "authorization token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/entity.SimpleResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Post"
                ],
                "summary": "Create",
                "parameters": [
                    {
                        "type": "string",
                        "description": "authorization token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/entity.SimpleResponse"
                        }
                    }
                }
            }
        },
        "/v1/api/posts/:id": {
            "get": {
                "description": "GetOne",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Post"
                ],
                "summary": "GetOne",
                "parameters": [
                    {
                        "type": "string",
                        "description": "authorization token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/entity.SimpleResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Post"
                ],
                "summary": "Delete",
                "parameters": [
                    {
                        "type": "string",
                        "description": "authorization token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/entity.SimpleResponse"
                        }
                    }
                }
            }
        },
        "/v1/api/request/call-back": {
            "post": {
                "description": "GoogleCallback",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "GoogleCallback",
                "parameters": [
                    {
                        "type": "string",
                        "description": "authorization token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/entity.SimpleResponse"
                        }
                    }
                }
            }
        },
        "/v1/api/request/google/login": {
            "post": {
                "description": "GoogleLogin",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "GoogleLogin",
                "parameters": [
                    {
                        "type": "string",
                        "description": "authorization token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "302": {
                        "description": "Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/api/request/login": {
            "post": {
                "description": "Login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "type": "string",
                        "description": "authorization token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/entity.LoginResponse"
                        }
                    }
                }
            }
        },
        "/v1/api/request/register": {
            "post": {
                "description": "Register",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Register",
                "parameters": [
                    {
                        "type": "string",
                        "description": "authorization token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/entity.SimpleResponse"
                        }
                    }
                }
            }
        },
        "/v1/api/users": {
            "get": {
                "description": "GetList",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "GetList",
                "parameters": [
                    {
                        "type": "string",
                        "description": "authorization token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/entity.SimpleResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Update",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Update",
                "parameters": [
                    {
                        "type": "string",
                        "description": "authorization token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/entity.SimpleResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Create",
                "parameters": [
                    {
                        "type": "string",
                        "description": "authorization token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/entity.SimpleResponse"
                        }
                    }
                }
            }
        },
        "/v1/api/users/:id": {
            "get": {
                "description": "GetOne",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "GetOne",
                "parameters": [
                    {
                        "type": "string",
                        "description": "authorization token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/entity.SimpleResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Delete",
                "parameters": [
                    {
                        "type": "string",
                        "description": "authorization token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/entity.SimpleResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.LoginResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "$ref": "#/definitions/entity.TokenResponse"
                },
                "user": {
                    "$ref": "#/definitions/entity.UserResponse"
                }
            }
        },
        "entity.SimpleResponse": {
            "type": "object",
            "properties": {
                "domain": {},
                "message": {
                    "description": "Code    int         ` + "`" + `json:\"code\"` + "`" + `",
                    "type": "string"
                }
            }
        },
        "entity.TokenResponse": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "expires_in": {
                    "type": "integer"
                },
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "entity.UserResponse": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "role_id": {
                    "type": "string"
                },
                "user_name": {
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
	Version:          "2.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Swagger Trail Backend API",
	Description:      "This is Trail Backend API docs",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
