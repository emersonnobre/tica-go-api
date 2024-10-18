package handlers

import (
	"net/http"
	"strconv"

	"github.com/emersonnobre/tica-api-go/src/internal/core/domain"
	"github.com/emersonnobre/tica-api-go/src/internal/core/usecases"
	"github.com/emersonnobre/tica-api-go/src/internal/delivery/fiber/util"
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

//	    CreateEmployee godoc
//
//		@Summary        Criar um novo funcionário
//		@Description    Cria um novo funcionário.
//		@Description    Campos obrigatórios: nome e CPF.
//		@Tags           employees
//		@Accept         json
//		@Produce        json
//		@Param          employee  body      domain.Employee  true    "Funcionário a ser criado"
//		@Success        201 	{string}	string	 	"Funcionário criado com sucesso"
//		@Failure        400 	{string}	string	 	"Erro de validação"
//		@Failure        500 	{string}	string	 	"Erro interno do sistema"
//		@Router         /employees [post]
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

//	    GetEmployeeById godoc
//
//		@Summary        Obter um funcionário
//		@Description    Obtém um funcionário por id.
//		@Tags           employees
//		@Produce        json
//		@Param          id  		path       integer true "Id do funcionário a ser obtido"
//		@Success        200 		{object}   domain.Employee 	"O funcionário encontrado"
//		@Failure        400 		{string}   string	 		"Erro de validação"
//		@Failure        404 		{string}   string	 		"Cliente não encontrado"
//		@Failure        500 		{string}   string	 		"Erro interno do sistema"
//		@Router         /employees/{id} [get]
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
