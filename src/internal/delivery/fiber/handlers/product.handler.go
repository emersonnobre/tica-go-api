package handlers

import (
	"strconv"

	"github.com/emersonnobre/tica-api-go/src/internal/core/repositories"
	"github.com/emersonnobre/tica-api-go/src/internal/core/usecases"
	"github.com/emersonnobre/tica-api-go/src/internal/core/usecases/types/requests"
	"github.com/emersonnobre/tica-api-go/src/internal/delivery/fiber/util"
	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	createProductUseCase          *usecases.CreateProductUseCase
	updateProductUseCase          *usecases.UpdateProductUseCase
	getProductUseCase             *usecases.GetProductUseCase
	removeProductUseCase          *usecases.RemoveProductUseCase
	getProductsUseCase            *usecases.GetProductsUseCase
	purchaseProductUseCase        *usecases.PurchaseProductUseCase
	registerProductOutflowUseCase *usecases.RegisterProductOutflowUseCase
}

func NewProductHandler(
	createProductUseCase *usecases.CreateProductUseCase,
	updateProductUseCase *usecases.UpdateProductUseCase,
	getProductUseCase *usecases.GetProductUseCase,
	removeProductUseCase *usecases.RemoveProductUseCase,
	getProductsUseCase *usecases.GetProductsUseCase,
	purchaseProductUseCase *usecases.PurchaseProductUseCase,
	registerProductOutflowUseCase *usecases.RegisterProductOutflowUseCase,
) *ProductHandler {
	return &ProductHandler{
		createProductUseCase:          createProductUseCase,
		updateProductUseCase:          updateProductUseCase,
		getProductUseCase:             getProductUseCase,
		removeProductUseCase:          removeProductUseCase,
		getProductsUseCase:            getProductsUseCase,
		purchaseProductUseCase:        purchaseProductUseCase,
		registerProductOutflowUseCase: registerProductOutflowUseCase,
	}
}

func (h *ProductHandler) RegisterRoutes(app *fiber.App) {
	group := app.Group("/products")

	group.Post("/", h.Create)
	group.Get("/", h.Get)
	group.Put("/:id", h.Update)
	group.Get("/:id", h.GetById)
	group.Delete("/:id", h.Delete)
	group.Post("/:id/purchase", h.Purchase)
	group.Post("/:id/outflow", h.RegisterOutflow)
}

//	    CreateProduct godoc
//
//		@Summary        Criar um novo produto
//		@Description    Cria um novo produto.
//		@Description    Requisitos funcionais relacionados: 2A.
//		@Tags           products
//		@Accept         json
//		@Produce        json
//		@Param          product  body      requests.CreateProductRequest  true    "Produto a ser criado"
//		@Success        201 	{string}	string	 	"Produto criado com sucesso"
//		@Failure        400 	{string}	string	 	"Erro de validação"
//		@Failure        500 	{string}	string	 	"Erro interno do sistema"
//		@Router         /products [post]
func (h *ProductHandler) Create(ctx *fiber.Ctx) error {
	var product requests.CreateProductRequest
	if err := ctx.BodyParser(&product); err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("Error parsing the request json")
	}
	response := h.createProductUseCase.Execute(product)
	if response.ErrorName != nil {
		return ctx.Status(util.CoreErrorToHttpError(*response.ErrorName)).SendString(*response.ErrorMessage)
	}
	return ctx.SendStatus(fiber.StatusCreated)
}

//	    UpdateProduct godoc
//
//		@Summary        Atualizar produto
//		@Description    Atualiza um produto pelo id.
//		@Description    Requisitos funcionais relacionados: 2C.
//		@Tags           products
//		@Accept         json
//		@Produce        json
//		@Param          id  		path       integer true "Id do produto a ser atualizado"
//		@Param          product  body      requests.UpdateProductRequest  true    "Informações para atualizar"
//		@Success        204 	"Produto atualizado com sucesso"
//		@Failure        400 	{string}	string	 	"Erro de validação"
//		@Failure        500 	{string}	string	 	"Erro interno do sistema"
//		@Router         /products/{id} [put]
func (h *ProductHandler) Update(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("Erro ao interpretar o id!")
	}

	var product requests.UpdateProductRequest
	if err := ctx.BodyParser(&product); err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("Erro ao interpretar a requisição!")
	}

	product.Id = id
	response := h.updateProductUseCase.Execute(product)
	if response.ErrorName != nil {
		return ctx.Status(util.CoreErrorToHttpError(*response.ErrorName)).SendString(*response.ErrorMessage)
	}
	return ctx.SendStatus(fiber.StatusNoContent)
}

//	    GetProduct godoc
//
//		@Summary        Obter um produto
//		@Description    Obtém um produto pelo id.
//		@Description    Requisitos funcionais relacionados: 2G.
//		@Tags           products
//		@Accept         json
//		@Produce        json
//		@Param          id  		path       integer true "Id do produto a ser obtido"
//		@Success        200		{object}	domain.Product  			"Produto encontrado"
//		@Failure        404 	{string}	string	 					"Produto não encontrado"
//		@Failure        500 	{string}	string	 					"Erro interno do sistema"
//		@Router         /products/{id} [get]
func (h *ProductHandler) GetById(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("Erro ao interpretar o id!")
	}

	response := h.getProductUseCase.Execute(id)
	if response.ErrorName != nil {
		return ctx.Status(util.CoreErrorToHttpError(*response.ErrorName)).SendString(*response.ErrorMessage)
	}
	return ctx.Status(fiber.StatusOK).JSON(response.Data)
}

//	    DeleteProduct godoc
//
//		@Summary        Deleta um produto
//		@Description    Deleta um produto pelo id.
//		@Description    Requisitos funcionais relacionados: 2D.
//		@Tags           products
//		@Accept         json
//		@Produce        json
//		@Param          id  		path       integer true "Id do produto a ser deletado"
//		@Success        204	  	"Produto deletado com sucesso"
//		@Failure        404 	{string}	string	 					"Produto não encontrado"
//		@Failure        400 	{string}	string	 					"Erro de validação"
//		@Failure        500 	{string}	string	 					"Erro interno do sistema"
//		@Router         /products/{id} [delete]
func (h *ProductHandler) Delete(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("Erro ao interpretar o id!")
	}

	response := h.removeProductUseCase.Execute(id)
	if response.ErrorName != nil {
		return ctx.Status(util.CoreErrorToHttpError(*response.ErrorName)).SendString(*response.ErrorMessage)
	}
	return ctx.SendStatus(fiber.StatusNoContent)
}

//	    GetProducts godoc
//
//		@Summary        Obter uma lista de produtos paginada, ordenada e filtrada
//		@Description    Obtém uma lista de produtos paginada.
//		@Description    Requisitos funcionais relacionados: 2B.
//		@Description    Filtros disponíveis: name (nome), is_feedstock (se é matéria prima), category_id (id da categoria).
//		@Description    Campos disponíveis para ordenação (em inglês): name e created_at (orderBy)
//		@Description    Para ordenação, pode ser utilizado o mecanismo ascendente e descendente (ASC e DESC) (order)
//		@Description    offset: utilizado para paginação, define a quantidade de itens a serem "pulados".
//		@Description    limit: utilizado para paginação, define a quantidade máxima de itens a serem obtidos.
//		@Tags           products
//		@Accept         json
//		@Produce        json
//		@Param          limit  			query      integer false "Limite de itens a serem obtidos"
//		@Param          offset  		query      integer false "Quantidade de itens a serem pulados"
//		@Param          order_by  		query      string  false "Nome do campo para ordenação (name, created_at)"
//		@Param          order  			query      string  false "ASC ou DESC para ordenação"
//		@Param          name  			query      string  false "Nome para filtro"
//		@Param          is_feedstock  	query      string  false "Se é matéria prima ou não para filtro (True ou False)"
//		@Param          category_id  	query      integer false "Id da categoria do produto para filtro"
//		@Success        200 		{array}	   responses.ProductResponse	"Uma lista de produtos"
//		@Success        400 		{string}   string	 					"Erro de validação"
//		@Failure        500 		{string}   string	 					"Erro interno do sistema"
//		@Router         /products [get]
func (h *ProductHandler) Get(ctx *fiber.Ctx) error {
	limit, _ := strconv.Atoi(ctx.Query("limit", "10"))
	offset, _ := strconv.Atoi(ctx.Query("offset", "0"))
	orderBy := ctx.Query("order_by", "name")
	order := ctx.Query("order", "asc")
	name := ctx.Query("name")
	isFeedstock := ctx.Query("is_feedstock")
	categoryId := ctx.Query("category_id")

	var filters []repositories.Filter = []repositories.Filter{
		*repositories.NewFilter("active", "TRUE", false, false),
	}

	if name != "" {
		filters = append(filters, *repositories.NewFilter("name", name, true, true))
	}

	if isFeedstock != "" {
		if isFeedstock != "True" && isFeedstock != "False" {
			return ctx.Status(fiber.StatusBadRequest).SendString("is_feedstock deve ser True ou False!")
		}
		filters = append(filters, *repositories.NewFilter("is_feedstock", isFeedstock, false, false))
	}

	if categoryId != "" {
		_, err := strconv.Atoi(categoryId)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).SendString("category_id deve ser do tipo inteiro!")
		}
		filters = append(filters, *repositories.NewFilter("category_id", categoryId, false, false))
	}

	response := h.getProductsUseCase.Execute(limit, offset, orderBy, order, filters)
	if response.ErrorName != nil {
		return ctx.Status(util.CoreErrorToHttpError(*response.ErrorName)).SendString(*response.ErrorMessage)
	}
	return ctx.Status(fiber.StatusOK).JSON(response.Data)
}

//	    PurchaseProduct godoc
//
//		@Summary        Registrar a compra de um produto
//		@Description    Registra a compra de um produto e atualiza o estoque.
//		@Description    Requisitos funcionais relacionados: 2E, 2E.a.
//		@Tags           products
//		@Accept         json
//		@Produce        json
//		@Param          product  body      requests.PurchaseProductRequest  true    "Informações da compra para registro"
//		@Success        201 	{string}	string	 	"Compra registrada com sucesso"
//		@Failure        400 	{string}	string	 	"Erro de validação"
//		@Failure        500 	{string}	string	 	"Erro interno do sistema"
//		@Router         /products/{id}/purchase [post]
func (h *ProductHandler) Purchase(ctx *fiber.Ctx) error {
	var purchase requests.PurchaseProductRequest
	if err := ctx.BodyParser(&purchase); err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("Erro ao interpretar a requisição!")
	}
	response := h.purchaseProductUseCase.Execute(&purchase)
	if response.ErrorName != nil {
		return ctx.Status(util.CoreErrorToHttpError(*response.ErrorName)).SendString(*response.ErrorMessage)
	}
	return ctx.SendStatus(fiber.StatusCreated)
}

//	    RegisterProductOutflow godoc
//
//		@Summary        Registrar a saída manual de um produto
//		@Description    Registra a saída manual de um produto e atualiza o estoque.
//		@Description    Requisitos funcionais relacionados: 2F, 2F.a.
//		@Tags           products
//		@Accept         json
//		@Produce        json
//		@Param          outflow  body      requests.ProductOutflow  true    "Informações da saída para registro"
//		@Success        201 	{string}	string	 	"Saída registrada com sucesso"
//		@Failure        404 	{string}	string	 	"Produto não encontrado"
//		@Failure        400 	{string}	string	 	"Erro de validação"
//		@Failure        500 	{string}	string	 	"Erro interno do sistema"
//		@Router         /products/{id}/outflow [post]
func (h *ProductHandler) RegisterOutflow(ctx *fiber.Ctx) error {
	var outflow requests.ProductOutflow
	if err := ctx.BodyParser(&outflow); err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("Erro ao interpretar a requisição!")
	}
	response := h.registerProductOutflowUseCase.Execute(&outflow)
	if response.ErrorName != nil {
		return ctx.Status(util.CoreErrorToHttpError(*response.ErrorName)).SendString(*response.ErrorMessage)
	}
	return ctx.SendStatus(fiber.StatusCreated)
}
