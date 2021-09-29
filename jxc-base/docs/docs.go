// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-12-08 09:06:25.258685 +0800 CST m=+0.049912508

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
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/deliverytype/del": {
            "post": {
                "description": "删除出库类型接口",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "删除出库类型接口",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "出库类型编号",
                        "name": "id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\\\"status\\\":1,\u003cbr\u003e\\\"msg\\\":\\\"成功\\\",\u003cbr\u003e\\\"data\\\":{},\u003cbr\u003e\\\"page\\\":null}",
                        "schema": {
                            "$ref": "#/definitions/model.DeliveryType"
                        }
                    }
                }
            }
        },
        "/deliverytype/list": {
            "post": {
                "description": "列出商户的所有出库类型接口",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "列出商户的所有出库类型接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "商户号",
                        "name": "shopId",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\\\"status\\\":1,\u003cbr\u003e\\\"msg\\\":\\\"成功\\\",\u003cbr\u003e\\\"data\\\":{},\u003cbr\u003e\\\"page\\\":null}",
                        "schema": {
                            "$ref": "#/definitions/model.DeliveryType"
                        }
                    }
                }
            }
        },
        "/deliverytype/save": {
            "post": {
                "description": "保存出库类型接口",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "保存出库类型接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "商户号",
                        "name": "shopId",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "出库类型",
                        "name": "deliveryType",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\\\"status\\\":1,\u003cbr\u003e\\\"msg\\\":\\\"成功\\\",\u003cbr\u003e\\\"data\\\":{},\u003cbr\u003e\\\"page\\\":null}",
                        "schema": {
                            "$ref": "#/definitions/model.DeliveryType"
                        }
                    }
                }
            }
        },
        "/deliverytype/update": {
            "post": {
                "description": "更新出库类型接口",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "更新出库类型接口",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "出库类型编号",
                        "name": "id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "出库类型",
                        "name": "deliveryType",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\\\"status\\\":1,\u003cbr\u003e\\\"msg\\\":\\\"成功\\\",\u003cbr\u003e\\\"data\\\":{},\u003cbr\u003e\\\"page\\\":null}",
                        "schema": {
                            "$ref": "#/definitions/model.DeliveryType"
                        }
                    }
                }
            }
        },
        "/entrytype/del": {
            "post": {
                "description": "删除入库类型接口",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "删除入库类型接口",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "入库类型编号",
                        "name": "id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\\\"status\\\":1,\u003cbr\u003e\\\"msg\\\":\\\"成功\\\",\u003cbr\u003e\\\"data\\\":{},\u003cbr\u003e\\\"page\\\":null}",
                        "schema": {
                            "$ref": "#/definitions/model.EntryType"
                        }
                    }
                }
            }
        },
        "/entrytype/list": {
            "post": {
                "description": "列出商户的所有入库类型接口",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "列出商户的所有入库类型接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "商户号",
                        "name": "shopId",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\\\"status\\\":1,\u003cbr\u003e\\\"msg\\\":\\\"成功\\\",\u003cbr\u003e\\\"data\\\":{},\u003cbr\u003e\\\"page\\\":null}",
                        "schema": {
                            "$ref": "#/definitions/model.EntryType"
                        }
                    }
                }
            }
        },
        "/entrytype/save": {
            "post": {
                "description": "保存入库类型接口",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "保存入库类型接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "商户号",
                        "name": "shopId",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "入库类型",
                        "name": "entryType",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\\\"status\\\":1,\u003cbr\u003e\\\"msg\\\":\\\"成功\\\",\u003cbr\u003e\\\"data\\\":{},\u003cbr\u003e\\\"page\\\":null}",
                        "schema": {
                            "$ref": "#/definitions/model.EntryType"
                        }
                    }
                }
            }
        },
        "/entrytype/update": {
            "post": {
                "description": "更新入库类型接口",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "更新入库类型接口",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "入库类型编号",
                        "name": "id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "入库类型",
                        "name": "entryType",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\\\"status\\\":1,\u003cbr\u003e\\\"msg\\\":\\\"成功\\\",\u003cbr\u003e\\\"data\\\":{},\u003cbr\u003e\\\"page\\\":null}",
                        "schema": {
                            "$ref": "#/definitions/model.EntryType"
                        }
                    }
                }
            }
        },
        "/storage/del": {
            "post": {
                "description": "删除仓库接口",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "删除仓库接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "商户账号",
                        "name": "shopId",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "自定义仓库编码",
                        "name": "storageId",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\\\"status\\\":1,\u003cbr\u003e\\\"msg\\\":\\\"成功\\\",\u003cbr\u003e\\\"data\\\":{},\u003cbr\u003e\\\"page\\\":null}",
                        "schema": {
                            "$ref": "#/definitions/model.ShopStorage"
                        }
                    }
                }
            }
        },
        "/storage/down": {
            "post": {
                "description": "仓库顺序号向下移一个位置接口",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "仓库顺序号向下移一个位置接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "商户号",
                        "name": "shopId",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "仓库编号",
                        "name": "storageId",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\\\"status\\\":1,\u003cbr\u003e\\\"msg\\\":\\\"成功\\\",\u003cbr\u003e\\\"data\\\":{},\u003cbr\u003e\\\"page\\\":null}",
                        "schema": {
                            "$ref": "#/definitions/model.ShopStorage"
                        }
                    }
                }
            }
        },
        "/storage/get": {
            "post": {
                "description": "获取仓库接口",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "获取仓库接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "商户号",
                        "name": "shopId",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "仓库编号",
                        "name": "storageId",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "仓库名称",
                        "name": "name",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\\\"status\\\":1,\u003cbr\u003e\\\"msg\\\":\\\"成功\\\",\u003cbr\u003e\\\"data\\\":{},\u003cbr\u003e\\\"page\\\":null}",
                        "schema": {
                            "$ref": "#/definitions/model.ShopStorage"
                        }
                    }
                }
            }
        },
        "/storage/last": {
            "post": {
                "description": "获取指定仓库最后修改信息接口",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "获取指定仓库最后修改信息接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "商户账号",
                        "name": "shopId",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "仓库编号",
                        "name": "storageId",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\\\"status\\\":1,\u003cbr\u003e\\\"msg\\\":\\\"成功\\\",\u003cbr\u003e\\\"data\\\":{},\u003cbr\u003e\\\"page\\\":null}",
                        "schema": {
                            "$ref": "#/definitions/model.StockDelivery"
                        }
                    }
                }
            }
        },
        "/storage/list": {
            "post": {
                "description": "列出商户的所有仓库接口",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "列出商户的所有仓库接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "商户号",
                        "name": "shopId",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\\\"status\\\":1,\u003cbr\u003e\\\"msg\\\":\\\"成功\\\",\u003cbr\u003e\\\"data\\\":{},\u003cbr\u003e\\\"page\\\":null}",
                        "schema": {
                            "$ref": "#/definitions/model.ShopStorage"
                        }
                    }
                }
            }
        },
        "/storage/save": {
            "post": {
                "description": "保存仓库接口",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "保存仓库接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "商户账号",
                        "name": "shopId",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "仓库名称",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "自定义仓库编码，不传则自动生成",
                        "name": "storageId",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "备注",
                        "name": "remark",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\\\"status\\\":1,\u003cbr\u003e\\\"msg\\\":\\\"成功\\\",\u003cbr\u003e\\\"data\\\":{},\u003cbr\u003e\\\"page\\\":null}",
                        "schema": {
                            "$ref": "#/definitions/model.ShopStorage"
                        }
                    }
                }
            }
        },
        "/storage/up": {
            "post": {
                "description": "仓库顺序号向上移一个位置接口",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "仓库顺序号向上移一个位置接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "商户号",
                        "name": "shopId",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "仓库编号",
                        "name": "storageId",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\\\"status\\\":1,\u003cbr\u003e\\\"msg\\\":\\\"成功\\\",\u003cbr\u003e\\\"data\\\":{},\u003cbr\u003e\\\"page\\\":null}",
                        "schema": {
                            "$ref": "#/definitions/model.ShopStorage"
                        }
                    }
                }
            }
        },
        "/storage/update": {
            "post": {
                "description": "更新仓库接口",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "更新仓库接口",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "仓库记录序列号，若不传时则不可修改自定义仓库编码",
                        "name": "id",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "商户账号",
                        "name": "shopId",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "自定义仓库编码",
                        "name": "storageId",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "仓库名称",
                        "name": "name",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "备注",
                        "name": "remark",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\\\"status\\\":1,\u003cbr\u003e\\\"msg\\\":\\\"成功\\\",\u003cbr\u003e\\\"data\\\":{},\u003cbr\u003e\\\"page\\\":null}",
                        "schema": {
                            "$ref": "#/definitions/model.ShopStorage"
                        }
                    }
                }
            }
        },
        "/unit/del": {
            "post": {
                "description": "删除单位接口",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "删除单位接口",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "单位编号",
                        "name": "id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\\\"status\\\":1,\u003cbr\u003e\\\"msg\\\":\\\"成功\\\",\u003cbr\u003e\\\"data\\\":{},\u003cbr\u003e\\\"page\\\":null}",
                        "schema": {
                            "$ref": "#/definitions/model.Unit"
                        }
                    }
                }
            }
        },
        "/unit/list": {
            "post": {
                "description": "列出商户的所有单位接口",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "列出商户的所有单位接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "商户号",
                        "name": "shopId",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\\\"status\\\":1,\u003cbr\u003e\\\"msg\\\":\\\"成功\\\",\u003cbr\u003e\\\"data\\\":{},\u003cbr\u003e\\\"page\\\":null}",
                        "schema": {
                            "$ref": "#/definitions/model.Unit"
                        }
                    }
                }
            }
        },
        "/unit/save": {
            "post": {
                "description": "保存单位接口",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "保存单位接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "商户号",
                        "name": "shopId",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "单位",
                        "name": "unit",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\\\"status\\\":1,\u003cbr\u003e\\\"msg\\\":\\\"成功\\\",\u003cbr\u003e\\\"data\\\":{},\u003cbr\u003e\\\"page\\\":null}",
                        "schema": {
                            "$ref": "#/definitions/model.Unit"
                        }
                    }
                }
            }
        },
        "/unit/update": {
            "post": {
                "description": "更新单位接口",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "更新单位接口",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "单位编号",
                        "name": "id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "单位",
                        "name": "unit",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\\\"status\\\":1,\u003cbr\u003e\\\"msg\\\":\\\"成功\\\",\u003cbr\u003e\\\"data\\\":{},\u003cbr\u003e\\\"page\\\":null}",
                        "schema": {
                            "$ref": "#/definitions/model.Unit"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.DeliveryType": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "shop_id": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "model.EntryType": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "shop_id": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "model.ShopStorage": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "remark": {
                    "type": "string"
                },
                "shop_id": {
                    "type": "string"
                },
                "sort_number": {
                    "type": "integer"
                },
                "storage_id": {
                    "type": "string"
                }
            }
        },
        "model.StockDelivery": {
            "type": "object",
            "properties": {
                "cost": {
                    "type": "integer"
                },
                "customer_id": {
                    "type": "string"
                },
                "customer_name": {
                    "type": "string"
                },
                "delivery_no": {
                    "type": "string"
                },
                "delivery_time": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "multi_storage": {
                    "type": "string"
                },
                "number": {
                    "type": "integer"
                },
                "operator": {
                    "type": "string"
                },
                "order_no": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "remark": {
                    "type": "string"
                },
                "shop_id": {
                    "type": "string"
                },
                "sku_id": {
                    "type": "string"
                },
                "storage_id": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                },
                "unit": {
                    "type": "string"
                }
            }
        },
        "model.Unit": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "shop_id": {
                    "type": "string"
                },
                "unit": {
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
	Version:     "1.0.0(jxc-base)",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
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