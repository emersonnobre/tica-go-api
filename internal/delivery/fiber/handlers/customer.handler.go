package handlers

import (
	"net/http"
	"strconv"

	"github.com/emersonnobre/tica-api-go/internal/core/domain"
	"github.com/emersonnobre/tica-api-go/internal/core/usecases"
	"github.com/emersonnobre/tica-api-go/internal/delivery/fiber/util"
	"github.com/gofiber/fiber/v2"
)

type CustomerHandler struct {
	createCustomerUseCase *usecases.CreateCustomerUseCase
	getCustomerUseCase    *usecases.GetCustomerUseCase
}

func NewCustomerHandler(createCustomerUseCase *usecases.CreateCustomerUseCase, getCustomerUseCase *usecases.GetCustomerUseCase) *CustomerHandler {
	return &CustomerHandler{
		createCustomerUseCase: createCustomerUseCase,
		getCustomerUseCase:    getCustomerUseCase,
	}
}

func (h *CustomerHandler) RegisterRoutes(app *fiber.App) {
	group := app.Group("/customers")

	group.Post("/", h.Create)
	group.Get("/:id", h.GetById)
}

func (h *CustomerHandler) Create(ctx *fiber.Ctx) error {
	var customer domain.Customer
	if err := ctx.BodyParser(&customer); err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("Erro ao interpretar a requisição!")
	}
	response := h.createCustomerUseCase.Execute(customer)
	if response.ErrorName != nil {
		return ctx.Status(util.CoreErrorToHttpError(*response.ErrorName)).SendString(*response.ErrorMessage)
	}
	return ctx.SendStatus(http.StatusCreated)
}

func (h *CustomerHandler) GetById(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("Erro ao interpretar o id!")
	}

	response := h.getCustomerUseCase.Execute(id)
	if response.ErrorName != nil {
		return ctx.Status(util.CoreErrorToHttpError(*response.ErrorName)).SendString(*response.ErrorMessage)
	}
	return ctx.Status(fiber.StatusOK).JSON(response.Data)
}
