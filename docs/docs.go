// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/dump": {
            "post": {
                "description": "获取JSON",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "冻结程序"
                ],
                "summary": "冻结程序",
                "parameters": [
                    {
                        "description": "冻结所需参数",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/app.DumpParam"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "响应参数",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/restore": {
            "post": {
                "description": "获取JSON",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "恢复已冻结程序"
                ],
                "summary": "恢复已冻结程序",
                "parameters": [
                    {
                        "description": "恢复所需参数",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/app.RestoreParam"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "响应参数",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/startCmd": {
            "post": {
                "description": "获取JSON",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "根据命令参数启动程序"
                ],
                "summary": "根据命令参数启动程序",
                "parameters": [
                    {
                        "description": "启动所需参数",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/app.StartCmdModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "响应参数",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "app.DumpParam": {
            "type": "object",
            "properties": {
                "dir": {
                    "description": "冻结文件保存地址",
                    "type": "string"
                },
                "leaveRunning": {
                    "description": "在不停止程序时冻结当前程序进程",
                    "type": "boolean"
                },
                "pid": {
                    "description": "进程ID",
                    "type": "integer"
                },
                "shellJob": {
                    "description": "是否是一个命令行程序",
                    "type": "boolean"
                },
                "tcpClose": {
                    "description": "关闭已建立的TCP",
                    "type": "boolean"
                },
                "tcpEstablished": {
                    "description": "是否已建立TCP连接",
                    "type": "boolean"
                },
                "tcpSkipInFlight": {
                    "description": "是否跳过正在运行的TCP连接",
                    "type": "boolean"
                }
            }
        },
        "app.RestoreParam": {
            "type": "object",
            "properties": {
                "dir": {
                    "description": "恢复已冻结进程路径",
                    "type": "string"
                },
                "leaveRunning": {
                    "description": "在不停止程序时冻结当前程序进程",
                    "type": "boolean"
                },
                "rstSibling": {
                    "description": "rstSibling",
                    "type": "boolean"
                },
                "shellJob": {
                    "description": "是否是一个命令行程序",
                    "type": "boolean"
                },
                "tcpClose": {
                    "description": "关闭已建立的TCP",
                    "type": "boolean"
                },
                "tcpEstablished": {
                    "description": "是否已建立TCP连接",
                    "type": "boolean"
                },
                "tcpSkipInFlight": {
                    "description": "是否跳过正在运行的TCP连接",
                    "type": "boolean"
                }
            }
        },
        "app.StartCmdModel": {
            "type": "object",
            "properties": {
                "parameter": {
                    "description": "运行所需参数",
                    "type": "string"
                },
                "path": {
                    "description": "程序所在路径",
                    "type": "string"
                }
            }
        },
        "response.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "响应状态",
                    "type": "integer"
                },
                "data": {
                    "description": "返回数据",
                    "type": "object"
                },
                "msg": {
                    "description": "响应消息",
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "仿真热迁移系统",
	Description: "swagger Server api",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}