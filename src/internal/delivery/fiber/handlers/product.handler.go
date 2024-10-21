package handlers

import (
	"strconv"

	"github.com/emersonnobre/tica-api-go/src/internal/core/usecases"
	"github.com/emersonnobre/tica-api-go/src/internal/core/usecases/types/requests"
	"github.com/emersonnobre/tica-api-go/src/internal/delivery/fiber/util"
	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	createProductUseCase *usecases.CreateProductUseCase
	updateProductUseCase *usecases.UpdateProductUseCase
	getProductUseCase    *usecases.GetProductUseCase
	removeProductUseCase *usecases.RemoveProductUseCase
}

func NewProductHandler(
	createProductUseCase *usecases.CreateProductUseCase,
	updateProductUseCase *usecases.UpdateProductUseCase,
	getProductUseCase *usecases.GetProductUseCase,
	removeProductUseCase *usecases.RemoveProductUseCase,
) *ProductHandler {
	return &ProductHandler{
		createProductUseCase: createProductUseCase,
		updateProductUseCase: updateProductUseCase,
		getProductUseCase:    getProductUseCase,
		removeProductUseCase: removeProductUseCase,
	}
}

func (h *ProductHandler) RegisterRoutes(app *fiber.App) {
	group := app.Group("/products")

	group.Post("/", h.Create)
	group.Put("/:id", h.Update)
	group.Get("/:id", h.GetById)
	group.Delete("/:id", h.Delete)
}

//	    CreateProduct godoc
//
//		@Summary        Criar um novo produto
//		@Description    Cria um novo produto.
//		@Description    Campos obrigatórios: nome e CPF.
//		@Tags           products
//		@Accept         json
//		@Produce        json
//		@Param          employee  body      requests.CreateProductRequest  true    "Produto a ser criado"
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
		return ctx.Status(fiber.StatusBadRequest).SendString("Erro ao interpretar requisição!")
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
//		@Tags           products
//		@Accept         json
//		@Produce        json
//		@Param          id  		path       integer true "Id do produto a ser obtido"
//		@Success        200		{object}	domain.Product  	"Produto encontrado"
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
