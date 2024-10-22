// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/categories": {
            "get": {
                "description": "Obtém todas as categorias sem filtro ou ordenação.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "categories"
                ],
                "summary": "Obter todas as categorias",
                "responses": {
                    "200": {
                        "description": "Uma lista de categorias",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.Category"
                            }
                        }
                    },
                    "500": {
                        "description": "Erro interno do sistema",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Cria uma nova categoria de produtos.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "categories"
                ],
                "summary": "Criar uma categoria",
                "parameters": [
                    {
                        "description": "Categoria a ser criada",
                        "name": "category",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Category"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Categoria criada com sucesso",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Erro de validação",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Erro interno do sistema",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/customers": {
            "get": {
                "description": "Obtém uma lista de clientes paginada.\nFiltros disponíveis: name (nome) e cpf.\nCampos disponíveis para ordenação (em inglês): name, created_at e updated_at (orderBy)\nPara ordenação, pode ser utilizado o mecanismo ascendente e descendente (ASC e DESC) (order)\noffset: utilizado para paginação, define a quantidade de itens a serem \"pulados\".\nlimit: utilizado para paginação, define a quantidade máxima de itens a serem obtidos.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "customers"
                ],
                "summary": "Obter uma lista de clientes paginada, ordenada e filtrada",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Limite de itens a serem obtidos",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Quantidade de itens a serem pulados",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Nome do campo para ordenação",
                        "name": "order_by",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "ASC ou DESC para ordenação",
                        "name": "order",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Nome para filtro",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "CPF para filtro",
                        "name": "cpf",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Uma lista de clientes",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.Customer"
                            }
                        }
                    },
                    "500": {
                        "description": "Erro interno do sistema",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "Atualiza um cliente.\nCampos obrigatórios: nome.\nCampos opcionais: CPF, telefone, e-mail, instagram e data de nascimento.\nOs endereços também podem ser atualizados. Para criar um endereço, envie um objeto com id vazio. Para deletar um endereço existente, não o envie na lista.\nCampos obrigatórios: Rua e bairro.\nCampos opcionais: CEP.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "customers"
                ],
                "summary": "Atualizar cliente",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Id do cliente a ser atualizado",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Cliente a ser atualizado",
                        "name": "customer",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Customer"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Cliente atualizado com sucesso"
                    },
                    "400": {
                        "description": "Erro de validação",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Cliente não encontrado",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Erro interno do sistema",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Cria um novo cliente.\nCampos obrigatórios: nome.\nCampos opcionais: CPF, telefone, e-mail, instagram e data de nascimento.\nUma lista de endereços também pode ser cadastrada para o cliente.\nCampos obrigatórios: Rua e bairro.\nCampos opcionais: CEP.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "customers"
                ],
                "summary": "Criar um novo cliente",
                "parameters": [
                    {
                        "description": "Cliente a ser criado",
                        "name": "customer",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Customer"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Cliente criado com sucesso",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Erro de validação",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Erro interno do sistema",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/customers/{id}": {
            "get": {
                "description": "Obtém um cliente por id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "customers"
                ],
                "summary": "Obter um cliente",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Id do cliente",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "O cliente encontrado",
                        "schema": {
                            "$ref": "#/definitions/domain.Customer"
                        }
                    },
                    "400": {
                        "description": "Erro de validação",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Cliente não encontrado",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Erro interno do sistema",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deleta um cliente por id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "customers"
                ],
                "summary": "Deleta um cliente",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Id do cliente",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Erro de validação",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Cliente não encontrado",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Erro interno do sistema",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/employees": {
            "post": {
                "description": "Cria um novo funcionário.\nCampos obrigatórios: nome e CPF.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "employees"
                ],
                "summary": "Criar um novo funcionário",
                "parameters": [
                    {
                        "description": "Funcionário a ser criado",
                        "name": "employee",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Employee"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Funcionário criado com sucesso",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Erro de validação",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Erro interno do sistema",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/employees/{id}": {
            "get": {
                "description": "Obtém um funcionário por id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "employees"
                ],
                "summary": "Obter um funcionário",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Id do funcionário a ser obtido",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "O funcionário encontrado",
                        "schema": {
                            "$ref": "#/definitions/domain.Employee"
                        }
                    },
                    "400": {
                        "description": "Erro de validação",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Cliente não encontrado",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Erro interno do sistema",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/products": {
            "get": {
                "description": "Obtém uma lista de produtos paginada.\nRequisitos funcionais relacionados: 2B.\nFiltros disponíveis: name (nome), is_feedstock (se é matéria prima), category_id (id da categoria).\nCampos disponíveis para ordenação (em inglês): name e created_at (orderBy)\nPara ordenação, pode ser utilizado o mecanismo ascendente e descendente (ASC e DESC) (order)\noffset: utilizado para paginação, define a quantidade de itens a serem \"pulados\".\nlimit: utilizado para paginação, define a quantidade máxima de itens a serem obtidos.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Obter uma lista de produtos paginada, ordenada e filtrada",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Limite de itens a serem obtidos",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Quantidade de itens a serem pulados",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Nome do campo para ordenação (name, created_at)",
                        "name": "order_by",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "ASC ou DESC para ordenação",
                        "name": "order",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Nome para filtro",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Se é matéria prima ou não para filtro (True ou False)",
                        "name": "is_feedstock",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Id da categoria do produto para filtro",
                        "name": "category_id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Uma lista de produtos",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/responses.ProductResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Erro de validação",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Erro interno do sistema",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Cria um novo produto.\nRequisitos funcionais relacionados: 2A.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Criar um novo produto",
                "parameters": [
                    {
                        "description": "Produto a ser criado",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.CreateProductRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Produto criado com sucesso",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Erro de validação",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Erro interno do sistema",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/products/{id}": {
            "get": {
                "description": "Obtém um produto pelo id.\nRequisitos funcionais relacionados: 2G.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Obter um produto",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Id do produto a ser obtido",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Produto encontrado",
                        "schema": {
                            "$ref": "#/definitions/domain.Product"
                        }
                    },
                    "404": {
                        "description": "Produto não encontrado",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Erro interno do sistema",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "Atualiza um produto pelo id.\nRequisitos funcionais relacionados: 2C.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Atualizar produto",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Id do produto a ser atualizado",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Informações para atualizar",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.UpdateProductRequest"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Produto atualizado com sucesso"
                    },
                    "400": {
                        "description": "Erro de validação",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Erro interno do sistema",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deleta um produto pelo id.\nRequisitos funcionais relacionados: 2D.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Deleta um produto",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Id do produto a ser deletado",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Produto deletado com sucesso"
                    },
                    "400": {
                        "description": "Erro de validação",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Produto não encontrado",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Erro interno do sistema",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/products/{id}/outflow": {
            "post": {
                "description": "Registra a saída manual de um produto e atualiza o estoque.\nRequisitos funcionais relacionados: 2F, 2F.a.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Registrar a saída manual de um produto",
                "parameters": [
                    {
                        "description": "Informações da saída para registro",
                        "name": "outflow",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.ProductOutflow"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Saída registrada com sucesso",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Erro de validação",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Produto não encontrado",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Erro interno do sistema",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/products/{id}/purchase": {
            "post": {
                "description": "Registra a compra de um produto e atualiza o estoque.\nRequisitos funcionais relacionados: 2E, 2E.a.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Registrar a compra de um produto",
                "parameters": [
                    {
                        "description": "Informações da compra para registro",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.PurchaseProductRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Compra registrada com sucesso",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Erro de validação",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Erro interno do sistema",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.Address": {
            "type": "object",
            "properties": {
                "cep": {
                    "type": "string"
                },
                "customer_id": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "neighborhood": {
                    "type": "string"
                },
                "street": {
                    "type": "string"
                }
            }
        },
        "domain.Category": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                }
            }
        },
        "domain.Customer": {
            "type": "object",
            "properties": {
                "addresses": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/domain.Address"
                    }
                },
                "birthday": {
                    "type": "string"
                },
                "cpf": {
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
                "instagram": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "domain.Employee": {
            "type": "object",
            "properties": {
                "cpf": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "domain.Product": {
            "type": "object",
            "properties": {
                "barcode": {
                    "type": "string"
                },
                "category": {
                    "$ref": "#/definitions/domain.Category"
                },
                "created_at": {
                    "type": "string"
                },
                "created_by": {
                    "$ref": "#/definitions/domain.Employee"
                },
                "id": {
                    "type": "integer"
                },
                "is_feedstock": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "purchase_price": {
                    "type": "number"
                },
                "sale_price": {
                    "type": "number"
                },
                "stock": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                },
                "updated_by": {
                    "$ref": "#/definitions/domain.Employee"
                }
            }
        },
        "requests.CreateProductRequest": {
            "type": "object",
            "properties": {
                "category_id": {
                    "type": "integer"
                },
                "created_by": {
                    "type": "integer"
                },
                "is_feedstock": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "purchase_price": {
                    "type": "number"
                },
                "sale_price": {
                    "type": "number"
                },
                "stock": {
                    "type": "integer"
                }
            }
        },
        "requests.ProductOutflow": {
            "type": "object",
            "properties": {
                "created_by": {
                    "type": "integer"
                },
                "product_id": {
                    "type": "integer"
                },
                "quantity": {
                    "type": "integer"
                },
                "reason": {
                    "type": "string"
                }
            }
        },
        "requests.PurchaseProductRequest": {
            "type": "object",
            "properties": {
                "created_by": {
                    "type": "integer"
                },
                "product_id": {
                    "type": "integer"
                },
                "purchase_price": {
                    "type": "number"
                },
                "quantity": {
                    "type": "integer"
                }
            }
        },
        "requests.UpdateProductRequest": {
            "type": "object",
            "properties": {
                "category_id": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "is_feedstock": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "purchase_price": {
                    "type": "number"
                },
                "sale_price": {
                    "type": "number"
                },
                "stock": {
                    "type": "integer"
                },
                "updated_by": {
                    "type": "integer"
                }
            }
        },
        "responses.ProductResponse": {
            "type": "object",
            "properties": {
                "category": {
                    "$ref": "#/definitions/domain.Category"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "is_feedstock": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "stock": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3000",
	BasePath:         "/",
	Schemes:          []string{"http", "https"},
	Title:            "Fiber Example API",
	Description:      "This is a sample swagger for Fiber",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
