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
            "name": "Saeed, Danny",
            "email": "mzahry36@gmail.com, dankeshavarz1075@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/client/": {
            "get": {
                "description": "Retrieve client details using their phone number. The phone number is provided in the request body. The response may include either a ` + "`" + `Client` + "`" + ` or a ` + "`" + `VIPClient` + "`" + ` object in the ` + "`" + `data` + "`" + ` field.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "client"
                ],
                "summary": "Get client information by phone number",
                "parameters": [
                    {
                        "description": "phone_number",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "VIP client retrieved successfully",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.SuccessResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.VIPClient"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Client not found",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "error message"
                },
                "success": {
                    "type": "boolean",
                    "example": false
                }
            }
        },
        "dto.LoginRequest": {
            "type": "object",
            "required": [
                "phone_number"
            ],
            "properties": {
                "phone_number": {
                    "type": "string"
                }
            }
        },
        "dto.SuccessResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string",
                    "example": "message"
                },
                "success": {
                    "type": "boolean",
                    "example": true
                }
            }
        },
        "models.AddressOfClient": {
            "type": "object",
            "properties": {
                "client_id": {
                    "type": "integer"
                },
                "province": {
                    "type": "string"
                },
                "remain_address": {
                    "type": "string"
                }
            }
        },
        "models.Client": {
            "type": "object",
            "properties": {
                "addresses": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.AddressOfClient"
                    }
                },
                "client_id": {
                    "type": "integer"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                },
                "referral_code": {
                    "type": "string"
                },
                "timestamp": {
                    "type": "string"
                },
                "wallet_balance": {
                    "type": "number"
                }
            }
        },
        "models.VIPClient": {
            "type": "object",
            "properties": {
                "client": {
                    "$ref": "#/definitions/models.Client"
                },
                "expiration_time": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8082",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Pissaze",
	Description:      "API for managing products, and users in a hardware shopping site.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
