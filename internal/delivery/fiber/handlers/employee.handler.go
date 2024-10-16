package handlers

import (
	"net/http"
	"strconv"

	"github.com/emersonnobre/tica-api-go/internal/core/domain"
	"github.com/emersonnobre/tica-api-go/internal/core/usecases"
	"github.com/emersonnobre/tica-api-go/internal/delivery/fiber/util"
	"github.com/gofiber/fiber/v2"
)

type EmployeeHandler struct {
	createEmployeeUseCase *usecases.CreateEmployeeUseCase
	getEmployeeUseCase    *usecases.GetEmployeeUseCase
}

func NewEmployeeHandler(
	createEmployeeUseCase *usecases.CreateEmployeeUseCase,
	getEmployeeUseCase *usecases.GetEmployeeUseCase) *EmployeeHandler {
	return &EmployeeHandler{
		createEmployeeUseCase: createEmployeeUseCase,
		getEmployeeUseCase:    getEmployeeUseCase,
	}
}

func (h *EmployeeHandler) RegisterRoutes(app *fiber.App) {
	group := app.Group("/employees")

	group.Post("/", h.Create)
	group.Get("/:id", h.GetById)
}

func (h *EmployeeHandler) Create(ctx *fiber.Ctx) error {
	var employee domain.Employee
	if err := ctx.BodyParser(&employee); err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("Error parsing the request json")
	}
	response := h.createEmployeeUseCase.Execute(employee)
	if response.ErrorName != nil {
		return ctx.Status(util.CoreErrorToHttpError(*response.ErrorName)).SendString(*response.ErrorMessage)
	}
	return ctx.SendStatus(http.StatusCreated)
}

func (h *EmployeeHandler) GetById(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("Erro ao interpretar o id!")
	}

	response := h.getEmployeeUseCase.Execute(id)
	if response.ErrorName != nil {
		return ctx.Status(util.CoreErrorToHttpError(*response.ErrorName)).SendString(*response.ErrorMessage)
	}
	return ctx.Status(fiber.StatusOK).JSON(response.Data)
}
