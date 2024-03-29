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
        "/api/v1/auth/apple-login": {
            "post": {
                "description": "Apple login account",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Apple login account",
                "operationId": "auth-apple-login",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The token you got from apple login response",
                        "name": "access_token",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Player ID",
                        "name": "player_id",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/auth/delete-image": {
            "delete": {
                "description": "Delete image",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Delete image",
                "operationId": "auth-delete-image",
                "parameters": [
                    {
                        "type": "string",
                        "description": "image filename",
                        "name": "filename",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/helper.HttpResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/auth/facebook-login": {
            "post": {
                "description": "Facebook login account",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Facebook login account",
                "operationId": "auth-facebook-login",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Facebook access token",
                        "name": "access_token",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Player ID",
                        "name": "player_id",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/auth/facebook-login-x": {
            "post": {
                "description": "Facebook login account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Facebook login account",
                "operationId": "auth-facebook-login-x",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/auth/forgot-password": {
            "post": {
                "description": "Forgot password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Forgot password",
                "operationId": "auth-forgot-password",
                "parameters": [
                    {
                        "description": "JSON Request Body",
                        "name": "RequestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.ForgotPasswordRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/auth/google-login": {
            "post": {
                "description": "Google login account",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Google login account",
                "operationId": "auth-google-login",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Google access token",
                        "name": "access_token",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Player ID",
                        "name": "player_id",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/auth/google-login-x": {
            "post": {
                "description": "Google login account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Google login account",
                "operationId": "auth-google-login-x",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/auth/login": {
            "post": {
                "description": "Login account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Login account",
                "operationId": "auth-login",
                "parameters": [
                    {
                        "description": "JSON Request Body",
                        "name": "RequestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/auth/register-user": {
            "post": {
                "description": "Register new user account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Register new user account",
                "operationId": "auth-register-user",
                "parameters": [
                    {
                        "description": "JSON Request Body",
                        "name": "RequestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.RegisterUserRequest"
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
        },
        "/api/v1/auth/registeruser": {
            "post": {
                "description": "Register new user account no otp",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Register new user account no otp",
                "operationId": "auth-register-user-nootp",
                "parameters": [
                    {
                        "description": "JSON Request Body",
                        "name": "RequestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.RegisterUserRequestNoOtp"
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
        },
        "/api/v1/auth/request-otp": {
            "post": {
                "description": "Request OTP",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Request OTP",
                "operationId": "auth-request-otp",
                "parameters": [
                    {
                        "description": "JSON Request Body",
                        "name": "RequestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.RequestOTPNoAuthRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/auth/upload-image": {
            "post": {
                "description": "Upload image",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Upload image",
                "operationId": "auth-upload-image",
                "parameters": [
                    {
                        "type": "file",
                        "description": "image file",
                        "name": "image",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/auth/verify-otp": {
            "post": {
                "description": "Verify OTP",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Verify OTP",
                "operationId": "auth-verify-otp",
                "parameters": [
                    {
                        "description": "JSON Request Body",
                        "name": "RequestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.VerifyOTPRequest"
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
        },
        "/api/v1/user/block/{user_id}": {
            "post": {
                "description": "Block a user",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Block a user",
                "operationId": "users-block-block-user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Auth Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "user id you want to block",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.UserBlock"
                        }
                    }
                }
            }
        },
        "/api/v1/user/change-password": {
            "patch": {
                "description": "Change password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Change password",
                "operationId": "users-change-password",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Auth Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "JSON Request Body",
                        "name": "RequestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.ChangePasswordRequest"
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
        },
        "/api/v1/user/delete/{id}": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Delete users",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Delete users",
                "operationId": "users-user-delete",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Auth Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
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
        },
        "/api/v1/user/edit-user-profile": {
            "patch": {
                "description": "Edit user profile",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Edit user profile",
                "operationId": "users-edit-user-profile",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Auth Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "JSON Request Body",
                        "name": "RequestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.EditUserRequest"
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
        },
        "/api/v1/user/my-profile": {
            "post": {
                "description": "Get my profile",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get my profile",
                "operationId": "users-my-profile",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Auth Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/helper.HttpResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.User"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v1/user/request-otp": {
            "post": {
                "description": "Request OTP",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Request OTP",
                "operationId": "users-request-otp",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Auth Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "JSON Request Body",
                        "name": "RequestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.RequestOTPRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/user/upload-profile-photo": {
            "post": {
                "description": "Upload profile photo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Upload profile photo",
                "operationId": "users-upload-profile-photo",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Auth Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "image file",
                        "name": "image",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/healthcheck": {
            "get": {
                "description": "Check API health status",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "healthcheck"
                ],
                "summary": "Check API health status",
                "operationId": "healthcheck-healthcheck",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.ChangePasswordRequest": {
            "type": "object",
            "properties": {
                "otp": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "repeat_password": {
                    "type": "string"
                }
            }
        },
        "controllers.EditUser": {
            "type": "object",
            "required": [
                "email",
                "name",
                "phone_number",
                "username"
            ],
            "properties": {
                "address": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "controllers.EditUserRequest": {
            "type": "object",
            "properties": {
                "user": {
                    "$ref": "#/definitions/controllers.EditUser"
                }
            }
        },
        "controllers.ForgotPasswordRequest": {
            "type": "object",
            "properties": {
                "otp": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "repeat_password": {
                    "type": "string"
                }
            }
        },
        "controllers.LoginRequest": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "player_id": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "controllers.RegisterUserRequest": {
            "type": "object",
            "properties": {
                "otp": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/controllers.User"
                }
            }
        },
        "controllers.RegisterUserRequestNoOtp": {
            "type": "object",
            "properties": {
                "user": {
                    "$ref": "#/definitions/controllers.User"
                }
            }
        },
        "controllers.RequestOTPNoAuthRequest": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                }
            }
        },
        "controllers.RequestOTPRequest": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                }
            }
        },
        "controllers.User": {
            "type": "object",
            "required": [
                "confirm_password",
                "email",
                "name",
                "password",
                "phone_number",
                "username"
            ],
            "properties": {
                "company": {
                    "type": "string"
                },
                "confirm_password": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "paket": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "controllers.VerifyOTPRequest": {
            "type": "object",
            "required": [
                "category",
                "code"
            ],
            "properties": {
                "category": {
                    "type": "string"
                },
                "code": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "helper.HttpResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object",
                    "additionalProperties": true
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "bank_accounts": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.UserBankAccount"
                    }
                },
                "company": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "is_partner": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "paket": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                },
                "player_id": {
                    "type": "string"
                },
                "profile_picture": {
                    "type": "string"
                },
                "status_active": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.UserBankAccount": {
            "type": "object",
            "properties": {
                "account_name": {
                    "type": "string"
                },
                "account_number": {
                    "type": "string"
                },
                "bank_id": {
                    "type": "integer"
                },
                "bank_name": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "models.UserBlock": {
            "type": "object",
            "properties": {
                "blocked_user_id": {
                    "type": "integer"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
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
