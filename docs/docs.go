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
        "/account/delete": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "用于删除账号",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "删除账号"
                ],
                "summary": "删除账号",
                "parameters": [
                    {
                        "type": "string",
                        "description": "删除账号参数",
                        "name": "username",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "删除账号成功",
                        "schema": {
                            "$ref": "#/definitions/response.Information"
                        }
                    },
                    "401": {
                        "description": "删除账号失败",
                        "schema": {
                            "$ref": "#/definitions/response.Information"
                        }
                    },
                    "402": {
                        "description": "请输入账号",
                        "schema": {
                            "$ref": "#/definitions/response.Information"
                        }
                    }
                }
            }
        },
        "/account/login": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "用于用户登录",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "登录"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "description": "登录参数",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/account.LoginBinder"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "登陆成功",
                        "schema": {
                            "$ref": "#/definitions/response.SuccessLogin"
                        }
                    },
                    "401": {
                        "description": "账号或者密码错误",
                        "schema": {
                            "$ref": "#/definitions/response.Information"
                        }
                    },
                    "402": {
                        "description": "请输入账号或者密码",
                        "schema": {
                            "$ref": "#/definitions/response.Information"
                        }
                    },
                    "403": {
                        "description": "该用户未注册",
                        "schema": {
                            "$ref": "#/definitions/response.Information"
                        }
                    },
                    "404": {
                        "description": "更新时间失败",
                        "schema": {
                            "$ref": "#/definitions/response.Information"
                        }
                    }
                }
            }
        },
        "/account/logout": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "用于登录用户注销",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "注销"
                ],
                "summary": "用户注销",
                "parameters": [
                    {
                        "type": "string",
                        "description": "注销参数",
                        "name": "email",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "注销成功",
                        "schema": {
                            "$ref": "#/definitions/response.Information"
                        }
                    },
                    "401": {
                        "description": "注销失败",
                        "schema": {
                            "$ref": "#/definitions/response.Information"
                        }
                    },
                    "402": {
                        "description": "您还未登录",
                        "schema": {
                            "$ref": "#/definitions/response.Information"
                        }
                    },
                    "403": {
                        "description": "请输入邮箱账号",
                        "schema": {
                            "$ref": "#/definitions/response.Information"
                        }
                    }
                }
            }
        },
        "/account/mail": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "新用户发送验证码用于注册账号",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "验证码"
                ],
                "summary": "发送验证码邮件",
                "parameters": [
                    {
                        "description": "发送邮件参数",
                        "name": "object",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/account.MailBinder"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "发送验证码成功",
                        "schema": {
                            "$ref": "#/definitions/response.Information"
                        }
                    },
                    "401": {
                        "description": "邮件发送失败",
                        "schema": {
                            "$ref": "#/definitions/response.Information"
                        }
                    },
                    "402": {
                        "description": "请输入邮箱或者密码",
                        "schema": {
                            "$ref": "#/definitions/response.Information"
                        }
                    }
                }
            }
        },
        "/account/modify": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "用于修改用户账号信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "修改用户账号信息"
                ],
                "summary": "修改用户账号信息",
                "parameters": [
                    {
                        "description": "修改用户账号信息参数",
                        "name": "object",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/account.ModifyBinder"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "修改信息成功",
                        "schema": {
                            "$ref": "#/definitions/response.Information"
                        }
                    },
                    "401": {
                        "description": "修改信息失败",
                        "schema": {
                            "$ref": "#/definitions/response.Information"
                        }
                    },
                    "402": {
                        "description": "请输入信息",
                        "schema": {
                            "$ref": "#/definitions/response.Information"
                        }
                    }
                }
            }
        },
        "/account/modify/information": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "用于修改用户详细信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "修改用户详细信息"
                ],
                "summary": "修改用户详细信息",
                "parameters": [
                    {
                        "description": "修改用户详细信息参数",
                        "name": "object",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/account.ModifyInformationBinder"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "修改信息成功",
                        "schema": {
                            "$ref": "#/definitions/response.Information"
                        }
                    },
                    "401": {
                        "description": "修改信息失败",
                        "schema": {
                            "$ref": "#/definitions/response.Information"
                        }
                    },
                    "402": {
                        "description": "请输入信息",
                        "schema": {
                            "$ref": "#/definitions/response.Information"
                        }
                    }
                }
            }
        },
        "/account/register": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "用于新用户注册账号使用",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "注册"
                ],
                "summary": "新用户注册",
                "parameters": [
                    {
                        "description": "注册参数",
                        "name": "object",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/account.RegisterBinder"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "注册成功",
                        "schema": {
                            "$ref": "#/definitions/response.Information"
                        }
                    },
                    "401": {
                        "description": "账号已注册",
                        "schema": {
                            "$ref": "#/definitions/response.Information"
                        }
                    },
                    "402": {
                        "description": "验证码错误",
                        "schema": {
                            "$ref": "#/definitions/response.Information"
                        }
                    },
                    "403": {
                        "description": "请输入完整的信息",
                        "schema": {
                            "$ref": "#/definitions/response.Information"
                        }
                    }
                }
            }
        },
        "/image/index": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "用于显示图片详细信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "显示图片详细信息"
                ],
                "summary": "显示图片详细信息",
                "responses": {
                    "200": {
                        "description": "获取图片详细信息成功",
                        "schema": {
                            "$ref": "#/definitions/response.Information"
                        }
                    },
                    "401": {
                        "description": "获取图片详细信息失败",
                        "schema": {
                            "$ref": "#/definitions/response.Information"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "account.LoginBinder": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "account.MailBinder": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                }
            }
        },
        "account.ModifyBinder": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "is_superuser": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "account.ModifyInformationBinder": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "age": {
                    "type": "string"
                },
                "avatar": {
                    "type": "string"
                },
                "biography": {
                    "type": "string"
                },
                "brith_date": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "nickname": {
                    "type": "string"
                },
                "posts": {
                    "type": "string"
                },
                "sex": {
                    "type": "string"
                },
                "style": {
                    "type": "string"
                }
            }
        },
        "account.RegisterBinder": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "response.Information": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "msg": {
                    "type": "string"
                }
            }
        },
        "response.SuccessLogin": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "msg": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "图片摄影创作社交网站",
	Description:      "在这里你可以获取想要的照片，并对照片进行二次创作，快来开始使用吧！",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
