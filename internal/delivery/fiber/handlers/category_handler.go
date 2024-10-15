package handlers

import (
	"net/http"

	"github.com/emersonnobre/tica-api-go/internal/core/domain"
	"github.com/emersonnobre/tica-api-go/internal/core/usecases"
	"github.com/emersonnobre/tica-api-go/internal/delivery/fiber/util"
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

func (h *CategoryHandler) Get(ctx *fiber.Ctx) error {
	response := h.getCategoriesUseCase.Execute()
	if response.ErrorName != nil {
		return ctx.Status(util.CoreErrorToHttpError(*response.ErrorName)).SendString(*response.ErrorMessage)
	}
	return ctx.Status(fiber.StatusOK).JSON(response.Data)
}
