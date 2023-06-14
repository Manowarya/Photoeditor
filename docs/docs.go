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
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/authorization": {
            "post": {
                "description": "Авторизация пользователя",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "AuthorizationUser",
                "parameters": [
                    {
                        "description": "User",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "int"
                        }
                    },
                    "400": {
                        "description": "Некорректный запрос",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Пользователя с такой почтой не существует",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "409": {
                        "description": "Неверный пароль",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/delete-template": {
            "post": {
                "description": "Удаление шаблона",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Templates"
                ],
                "summary": "DeleteTemplate",
                "parameters": [
                    {
                        "description": "TemplateID",
                        "name": "template",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/routes.TemplateID"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Шаблон удален",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "502": {
                        "description": "Ошибка сервера, попробуйте позже",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/get-templates/{id}": {
            "get": {
                "description": "Отправляет шаблон по индефикатору пользователя",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Templates"
                ],
                "summary": "GetTemplate",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Templates"
                        }
                    },
                    "502": {
                        "description": "Ошибка сервера, попробуйте позже",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "После ввода кода подтверждения регистрирует пользователя",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "RegisterUser",
                "parameters": [
                    {
                        "description": "User",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/routes.UserWithCode"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "int"
                        }
                    },
                    "400": {
                        "description": "Некорректный запрос",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "408": {
                        "description": "Неверный код, попробуйте заново",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Ошибка сервера, попробуйте позже",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/save-template": {
            "post": {
                "description": "Создание нового шаблона",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Templates"
                ],
                "summary": "SaveTemplate",
                "parameters": [
                    {
                        "description": "Template",
                        "name": "template",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Template"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Шаблон успешно создан",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "502": {
                        "description": "Ошибка сервера, попробуйте позже",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/update-template": {
            "post": {
                "description": "Изменение шаблона",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Templates"
                ],
                "summary": "UpdateTemplate",
                "parameters": [
                    {
                        "description": "Template",
                        "name": "template",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Template"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Шаблон изменен",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "502": {
                        "description": "Ошибка сервера, попробуйте позже",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/verification": {
            "post": {
                "description": "Отправляет код подтверждения при помощи smtp.gmail.com на почту пользователя",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "VerificationEmail",
                "parameters": [
                    {
                        "description": "User",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Код успешно отправлен",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Некорректный запрос",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "409": {
                        "description": "Пользователь с такой почтой уже существует",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Проверте корректность почты или попробуйте позже",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Template": {
            "type": "object",
            "properties": {
                "blur": {
                    "type": "number"
                },
                "bright": {
                    "type": "number"
                },
                "contrast": {
                    "type": "number"
                },
                "exposition": {
                    "type": "number"
                },
                "name": {
                    "type": "string"
                },
                "noise": {
                    "type": "number"
                },
                "saturation": {
                    "type": "number"
                },
                "tone": {
                    "type": "number"
                },
                "user_id": {
                    "type": "integer"
                },
                "vignette": {
                    "type": "number"
                }
            }
        },
        "model.Templates": {
            "type": "object",
            "properties": {
                "templates": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Template"
                    }
                }
            }
        },
        "model.User": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "routes.TemplateID": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "routes.UserWithCode": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "photofocus-production.up.railway.app",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "PhotoFocus API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
