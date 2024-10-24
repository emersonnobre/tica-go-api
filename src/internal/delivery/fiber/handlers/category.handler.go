package handlers

import (
	"net/http"

	_ "github.com/emersonnobre/tica-api-go/src/docs"
	"github.com/emersonnobre/tica-api-go/src/internal/core/domain"
	"github.com/emersonnobre/tica-api-go/src/internal/core/usecases"
	"github.com/emersonnobre/tica-api-go/src/internal/delivery/fiber/util"
	"github.com/gofiber/fiber/v2"
)

type CategoryHandler struct {
	createCategoryUseCase *usecases.CreateCategoryUseCase
	getCategoriesUseCase  *usecases.GetCategoriesUseCase
}

func NewCategoryHandler(createCategoryUseCase *usecases.CreateCategoryUseCase, getCategoriesUseCase *usecases.GetCategoriesUseCase) *CategoryHandler {
	return &CategoryHandler{
		createCategoryUseCase: createCategoryUseCase,
		getCategoriesUseCase:  getCategoriesUseCase,
	}
}

func (h *CategoryHandler) RegisterRoutes(app *fiber.App) {
	group := app.Group("/categories")

	group.Post("/", h.Create)
	group.Get("/", h.Get)
}

// CreateCategory godoc
//
//	@Summary        Criar uma categoria
//	@Description    Cria uma nova categoria de produtos.
//	@Tags           categories
//	@Accept         json
//	@Produce        json
//	@Param          category  body      domain.Category  true    "Categoria a ser criada"
//	@Success        201 	{string}	string	 	"Categoria criada com sucesso"
//	@Failure        400 	{string}	string	 	"Erro de validação"
//	@Failure        500 	{string}	string	 	"Erro interno do sistema"
//	@Router         /categories [post]
func (h *CategoryHandler) Create(ctx *fiber.Ctx) error {
	var category domain.Category
	if err := ctx.BodyParser(&category); err != nil {
		ctx.SendString("Error parsing the request json")
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	response := h.createCategoryUseCase.Execute(category)
	if response.ErrorName != nil {
		ctx.SendString(*response.ErrorMessage)
		return ctx.SendStatus(util.CoreErrorToHttpError(*response.ErrorName))
	}
	return ctx.SendStatus(http.StatusCreated)
}

// GetAllCategories godoc
//
//		@Summary        Obter todas as categorias
//		@Description    Obtém todas as categorias sem filtro ou ordenação.
//		@Tags           categories
//		@Produce        json
//		@Success        200 	{array}		domain.Category 	"Uma lista de categorias"
//	    @Failure        500 	{string} 	string 				"Erro interno do sistema"
//		@Router         /categories [get]
func (h *CategoryHandler) Get(ctx *fiber.Ctx) error {
	response := h.getCategoriesUseCase.Execute()
	if response.ErrorName != nil {
		return ctx.Status(util.CoreErrorToHttpError(*response.ErrorName)).SendString(*response.ErrorMessage)
	}
	return ctx.Status(fiber.StatusOK).JSON(response.Data)
}
