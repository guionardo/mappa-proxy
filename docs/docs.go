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
        "contact": {
            "name": "Guionardo Furlan",
            "email": "guionardo@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/": {
            "get": {
                "description": "route: /",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "mappa-proxy"
                ],
                "summary": "Index handler",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.IndexResponse"
                        }
                    }
                }
            }
        },
        "/admin/cache/reset": {
            "post": {
                "description": "Publica mensagem em chat do Telegram",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "Clear cache",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.ReplyMessage"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.ReplyMessage"
                        }
                    }
                }
            }
        },
        "/admin/routes": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "List registered routes",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.ReplyMessage"
                        }
                    }
                }
            }
        },
        "/admin/stats": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "Data statistics",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.StatsResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.ReplyMessage"
                        }
                    }
                }
            }
        },
        "/hc": {
            "get": {
                "description": "Service healthcheck",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "mappa-proxy"
                ],
                "summary": "Healthcheck handler",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.HealthCheckResponse"
                        }
                    }
                }
            }
        },
        "/mappa/conquistas/{cod_secao}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "mappa"
                ],
                "summary": "Lista de conquistas da secão",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Código Seção",
                        "name": "cod_secao",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Data de início do período (YYYY-MM-DD) padrão 1 ano atrás",
                        "name": "desde",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/responses.FullConquistaResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.ReplyMessage"
                        }
                    }
                }
            }
        },
        "/mappa/escotista/{userId}": {
            "get": {
                "description": "Informações do escotista, associado e grupo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "mappa"
                ],
                "summary": "Detalhes do escotista",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.MappaDetalhesResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.ReplyMessage"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/handlers.ReplyMessage"
                        }
                    }
                }
            }
        },
        "/mappa/especialidades": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "db"
                ],
                "summary": "Lista de especialidades e items",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.MappaEspecialidadeResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.ReplyMessage"
                        }
                    }
                }
            }
        },
        "/mappa/login": {
            "post": {
                "description": "User login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "mappa"
                ],
                "summary": "Mappa Login handler",
                "parameters": [
                    {
                        "description": "Login request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.MappaLoginResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.ReplyMessage"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/handlers.ReplyMessage"
                        }
                    }
                }
            }
        },
        "/mappa/marcacoes/{cod_secao}": {
            "get": {
                "description": "Lista de marcações da sessão",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "mappa"
                ],
                "summary": "MappaMarcacoes handler",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Código Seção",
                        "name": "cod_secao",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.MappaMarcacaoResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.ReplyMessage"
                        }
                    }
                }
            }
        },
        "/mappa/progressoes/{ramo}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "db"
                ],
                "summary": "Lista de progressões do ramo",
                "parameters": [
                    {
                        "enum": [
                            "L",
                            "E",
                            "S",
                            "P"
                        ],
                        "type": "string",
                        "description": "Ramo",
                        "name": "ramo",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/responses.MappaProgressaoResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.ReplyMessage"
                        }
                    }
                }
            }
        },
        "/mappa/secoes/{userId}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "mappa"
                ],
                "summary": "Seções do escotista",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.MappaSecaoResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.ReplyMessage"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/handlers.ReplyMessage"
                        }
                    }
                }
            }
        },
        "/tg/pub": {
            "post": {
                "description": "Publica mensagem em chat do Telegram",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "telegram"
                ],
                "summary": "Telegram Publisher handler",
                "parameters": [
                    {
                        "description": "Bot request data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.BotRequestData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.ReplyMessage"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.ReplyMessage"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handlers.BotRequestData": {
            "type": "object",
            "properties": {
                "cId": {
                    "description": "Chat ID",
                    "type": "integer"
                },
                "mId": {
                    "description": "Message ID to respond to",
                    "type": "integer"
                },
                "msg": {
                    "description": "Message",
                    "type": "string"
                }
            }
        },
        "handlers.ReplyMessage": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "infra.MemoryStatus": {
            "type": "object",
            "properties": {
                "alloc": {
                    "description": "Alloc is bytes of allocated heap objects.",
                    "type": "integer"
                },
                "heap_alloc": {
                    "description": "HeapAlloc is bytes of allocated heap objects.",
                    "type": "integer"
                },
                "num_gc": {
                    "description": "NumGC is the number of completed GC cycles.",
                    "type": "integer"
                },
                "sys": {
                    "description": "Sys is the total bytes of memory obtained from the OS.",
                    "type": "integer"
                },
                "total_alloc": {
                    "description": "TotalAlloc is cumulative bytes allocated for heap objects.",
                    "type": "integer"
                }
            }
        },
        "requests.LoginRequest": {
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
        "responses.FullConquistaResponse": {
            "type": "object",
            "properties": {
                "associado": {
                    "$ref": "#/definitions/responses.MappaAssociadoResponse"
                },
                "codigoAssociado": {
                    "type": "integer"
                },
                "codigoEscotistaUltimaAlteracao": {
                    "type": "integer"
                },
                "codigoEspecialidade": {
                    "type": "integer"
                },
                "dataConquista": {
                    "type": "string"
                },
                "especialidade": {
                    "$ref": "#/definitions/responses.MappaEspecialidadeResponse"
                },
                "numeroNivel": {
                    "type": "integer"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "responses.HealthCheckResponse": {
            "type": "object",
            "properties": {
                "mappa_server": {
                    "$ref": "#/definitions/responses.MappaServerResponse"
                },
                "memory": {
                    "$ref": "#/definitions/infra.MemoryStatus"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "responses.IndexResponse": {
            "type": "object",
            "properties": {
                "app": {
                    "type": "string"
                },
                "build_time": {
                    "type": "string"
                },
                "running_by": {
                    "type": "string"
                },
                "version": {
                    "type": "string"
                }
            }
        },
        "responses.MappaAssociadoResponse": {
            "type": "object",
            "properties": {
                "codigo": {
                    "type": "integer"
                },
                "codigoCategoria": {
                    "type": "integer"
                },
                "codigoEquipe": {
                    "type": "integer"
                },
                "codigoFoto": {
                    "type": "integer"
                },
                "codigoRamo": {
                    "type": "integer"
                },
                "codigoRamoAdulto": {
                    "type": "integer"
                },
                "codigoSegundaCategoria": {
                    "type": "integer"
                },
                "codigoTerceiraCategoria": {
                    "type": "integer"
                },
                "dataAcompanhamento": {
                    "type": "string"
                },
                "dataNascimento": {
                    "type": "string"
                },
                "dataValidade": {
                    "type": "string"
                },
                "linhaFormacao": {
                    "type": "string"
                },
                "nome": {
                    "type": "string"
                },
                "nomeAbreviado": {
                    "type": "string"
                },
                "numeroDigito": {
                    "type": "integer"
                },
                "sexo": {
                    "type": "string"
                },
                "username": {
                    "type": "integer"
                }
            }
        },
        "responses.MappaDetalhesResponse": {
            "type": "object",
            "properties": {
                "associado": {
                    "$ref": "#/definitions/responses.MappaAssociadoResponse"
                },
                "escotista": {
                    "$ref": "#/definitions/responses.MappaEscotistaResponse"
                },
                "grupo": {
                    "$ref": "#/definitions/responses.MappaGrupoResponse"
                }
            }
        },
        "responses.MappaEscotistaResponse": {
            "type": "object",
            "properties": {
                "ativo": {
                    "$ref": "#/definitions/types.Bool"
                },
                "codigo": {
                    "type": "integer"
                },
                "codigoAssociado": {
                    "type": "integer"
                },
                "codigoFoto": {
                    "type": "integer"
                },
                "codigoGrupo": {
                    "type": "integer"
                },
                "codigoRegiao": {
                    "type": "string"
                },
                "nomeCompleto": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "responses.MappaEspecialidadeItemResponse": {
            "type": "object",
            "properties": {
                "codigoEspecialidade": {
                    "description": "Id                  int    ` + "`" + `json:\"id\"` + "`" + `",
                    "type": "integer"
                },
                "descricao": {
                    "type": "string"
                },
                "numero": {
                    "type": "integer"
                }
            }
        },
        "responses.MappaEspecialidadeResponse": {
            "type": "object",
            "properties": {
                "codigo": {
                    "type": "integer"
                },
                "descricao": {
                    "type": "string"
                },
                "itens": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/responses.MappaEspecialidadeItemResponse"
                    }
                },
                "prerequisito": {
                    "type": "string"
                },
                "ramoConhecimento": {
                    "type": "string"
                }
            }
        },
        "responses.MappaGrupoResponse": {
            "type": "object",
            "properties": {
                "codigo": {
                    "type": "integer"
                },
                "codigoModalidade": {
                    "type": "integer"
                },
                "codigoRegiao": {
                    "type": "string"
                },
                "nome": {
                    "type": "string"
                }
            }
        },
        "responses.MappaLoginResponse": {
            "type": "object",
            "properties": {
                "created": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "ttl": {
                    "type": "integer"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "responses.MappaMarcacaoResponse": {
            "type": "object",
            "properties": {
                "codigoAssociado": {
                    "type": "integer"
                },
                "codigoAtividade": {
                    "type": "integer"
                },
                "codigoUltimoEscotista": {
                    "type": "integer"
                },
                "dataAtividade": {
                    "type": "string"
                },
                "dataHoraAtualizacao": {
                    "type": "string"
                },
                "dataStatusEscotista": {
                    "type": "string"
                },
                "segmento": {
                    "type": "string"
                }
            }
        },
        "responses.MappaProgressaoResponse": {
            "type": "object",
            "properties": {
                "codigo": {
                    "type": "integer"
                },
                "codigoCaminho": {
                    "type": "integer"
                },
                "codigoCompetencia": {
                    "type": "integer"
                },
                "codigoDesenvolvimento": {
                    "type": "integer"
                },
                "codigoRegiao": {
                    "type": "string"
                },
                "codigoUeb": {
                    "type": "string"
                },
                "descricao": {
                    "type": "string"
                },
                "numeroGrupo": {
                    "type": "integer"
                },
                "ordenacao": {
                    "type": "integer"
                },
                "segmento": {
                    "type": "string"
                }
            }
        },
        "responses.MappaSecaoResponse": {
            "type": "object",
            "properties": {
                "codigo": {
                    "type": "integer"
                },
                "codigoGrupo": {
                    "type": "integer"
                },
                "codigoRegiao": {
                    "type": "string"
                },
                "codigoTipoSecao": {
                    "type": "integer"
                },
                "nome": {
                    "type": "string"
                },
                "subsecoes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/responses.MappaSubSecaoResponse"
                    }
                }
            }
        },
        "responses.MappaServerResponse": {
            "type": "object",
            "properties": {
                "mappa_server_url": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "status_code": {
                    "type": "integer"
                }
            }
        },
        "responses.MappaSubSecaoResponse": {
            "type": "object",
            "properties": {
                "associados": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/responses.MappaAssociadoResponse"
                    }
                },
                "codigo": {
                    "type": "integer"
                },
                "codigoLider": {
                    "type": "integer"
                },
                "codigoSecao": {
                    "type": "integer"
                },
                "codigoViceLider": {
                    "type": "integer"
                },
                "nome": {
                    "type": "string"
                }
            }
        },
        "responses.StatsResponse": {
            "type": "object",
            "properties": {
                "associados": {
                    "type": "integer"
                },
                "escotistas": {
                    "type": "integer"
                },
                "grupos": {
                    "type": "integer"
                },
                "secoes": {
                    "type": "integer"
                }
            }
        },
        "types.Bool": {
            "type": "object"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.5.4",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "mappa-proxy",
	Description:      "Proxy and data analysis for Mappa",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
