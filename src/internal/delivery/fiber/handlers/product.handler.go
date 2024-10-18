package handlers

import (
	"github.com/emersonnobre/tica-api-go/src/internal/core/usecases"
	"github.com/emersonnobre/tica-api-go/src/internal/core/usecases/types/requests"
	"github.com/emersonnobre/tica-api-go/src/internal/delivery/fiber/util"
	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	createProductUseCase *usecases.CreateProductUseCase
}

func NewProductHandler(createProductUseCase *usecases.CreateProductUseCase) *ProductHandler {
	return &ProductHandler{
		createProductUseCase: createProductUseCase,
	}
}

func (h *ProductHandler) RegisterRoutes(app *fiber.App) {
	group := app.Group("/products")

	group.Post("/", h.Create)
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
