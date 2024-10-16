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
	updateCustomerUseCase *usecases.UpdateCustomerUseCase
	removeCustomerUseCase *usecases.RemoveCustomerUseCase
	getCustomersUseCase   *usecases.GetCustomersUseCase
}

func NewCustomerHandler(
	createCustomerUseCase *usecases.CreateCustomerUseCase,
	getCustomerUseCase *usecases.GetCustomerUseCase,
	updateCustomerUseCase *usecases.UpdateCustomerUseCase,
	removeCustomerUseCase *usecases.RemoveCustomerUseCase,
	getCustomersUseCase *usecases.GetCustomersUseCase,
) *CustomerHandler {
	return &CustomerHandler{
		createCustomerUseCase: createCustomerUseCase,
		getCustomerUseCase:    getCustomerUseCase,
		updateCustomerUseCase: updateCustomerUseCase,
		removeCustomerUseCase: removeCustomerUseCase,
		getCustomersUseCase:   getCustomersUseCase,
	}
}

func (h *CustomerHandler) RegisterRoutes(app *fiber.App) {
	group := app.Group("/customers")

	group.Post("/", h.Create)
	group.Get("/", h.Get)
	group.Get("/:id", h.GetById)
	group.Put("/:id", h.Update)
	group.Delete("/:id", h.Delete)
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

func (h *CustomerHandler) Get(ctx *fiber.Ctx) error {
	limit, _ := strconv.Atoi(ctx.Query("limit", "10"))
	offset, _ := strconv.Atoi(ctx.Query("offset", "0"))
	orderBy := ctx.Query("order_by", "name")
	order := ctx.Query("order", "asc")

	response := h.getCustomersUseCase.Execute(limit, offset, orderBy, order)
	if response.ErrorName != nil {
		return ctx.Status(util.CoreErrorToHttpError(*response.ErrorName)).SendString(*response.ErrorMessage)
	}
	return ctx.Status(fiber.StatusOK).JSON(response.Data)
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

func (h *CustomerHandler) Update(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("Erro ao interpretar o id!")
	}

	var customer domain.Customer
	if err := ctx.BodyParser(&customer); err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("Erro ao interpretar a requisição!")
	}

	customer.Id = id
	response := h.updateCustomerUseCase.Execute(customer)
	if response.ErrorName != nil {
		return ctx.Status(util.CoreErrorToHttpError(*response.ErrorName)).SendString(*response.ErrorMessage)
	}
	return ctx.SendStatus(fiber.StatusNoContent)
}

func (h *CustomerHandler) Delete(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("Erro ao interpretar o id!")
	}

	response := h.removeCustomerUseCase.Execute(id)
	if response.ErrorName != nil {
		return ctx.Status(util.CoreErrorToHttpError(*response.ErrorName)).SendString(*response.ErrorMessage)
	}
	return ctx.SendStatus(fiber.StatusNoContent)
}
