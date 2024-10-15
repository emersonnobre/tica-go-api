package handlers

import (
	"github.com/emersonnobre/tica-api-go/internal/core/domain"
	"github.com/emersonnobre/tica-api-go/internal/core/usecases"
	"github.com/emersonnobre/tica-api-go/internal/delivery/fiber/util"
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

func (h *ProductHandler) Create(ctx *fiber.Ctx) error {
	var product domain.Product
	if err := ctx.BodyParser(&product); err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("Error parsing the request json")
	}
	response := h.createProductUseCase.Execute(product)
	if response.ErrorName != nil {
		return ctx.Status(util.CoreErrorToHttpError(*response.ErrorName)).SendString(*response.ErrorMessage)
	}
	return ctx.SendStatus(fiber.StatusCreated)
}
