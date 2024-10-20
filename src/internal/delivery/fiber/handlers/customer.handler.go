package handlers

import (
	"net/http"
	"strconv"

	"github.com/emersonnobre/tica-api-go/src/internal/core/domain"
	"github.com/emersonnobre/tica-api-go/src/internal/core/repositories"
	"github.com/emersonnobre/tica-api-go/src/internal/core/usecases"
	"github.com/emersonnobre/tica-api-go/src/internal/delivery/fiber/util"
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

//	    CreateCustomer godoc
//
//		@Summary        Criar um novo cliente
//		@Description    Cria um novo cliente.
//		@Description    Campos obrigatórios: nome.
//		@Description    Campos opcionais: CPF, telefone, e-mail, instagram e data de nascimento.
//		@Description    Uma lista de endereços também pode ser cadastrada para o cliente.
//		@Description    Campos obrigatórios: Rua e bairro.
//		@Description    Campos opcionais: CEP.
//		@Tags           customers
//		@Accept         json
//		@Produce        json
//		@Param          customer  body      domain.Customer  true    "Cliente a ser criado"
//		@Success        201 	{string}	string	 	"Cliente criado com sucesso"
//		@Failure        400 	{string}	string	 	"Erro de validação"
//		@Failure        500 	{string}	string	 	"Erro interno do sistema"
//		@Router         /customers [post]
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

//	    GetCustomers godoc
//
//		@Summary        Obter uma lista de clientes paginada, ordenada e filtrada
//		@Description    Obtém uma lista de clientes paginada.
//		@Description    Filtros disponíveis: name (nome) e cpf.
//		@Description    Campos disponíveis para ordenação (em inglês): name, created_at e updated_at (orderBy)
//		@Description    Para ordenação, pode ser utilizado o mecanismo ascendente e descendente (ASC e DESC) (order)
//		@Description    offset: utilizado para paginação, define a quantidade de itens a serem "pulados".
//		@Description    limit: utilizado para paginação, define a quantidade máxima de itens a serem obtidos.
//		@Tags           customers
//		@Accept         json
//		@Produce        json
//		@Param          limit  		query      integer false "Limite de itens a serem obtidos"
//		@Param          offset  	query      integer false "Quantidade de itens a serem pulados"
//		@Param          order_by  	query      string  false "Nome do campo para ordenação"
//		@Param          order  		query      string  false "ASC ou DESC para ordenação"
//		@Param          name  		query      string  false "Nome para filtro"
//		@Param          cpf  		query      string  false "CPF para filtro"
//		@Success        200 		{array}	   domain.Customer	 	"Uma lista de clientes"
//		@Failure        500 		{string}   string	 	"Erro interno do sistema"
//		@Router         /customers [get]
func (h *CustomerHandler) Get(ctx *fiber.Ctx) error {
	limit, _ := strconv.Atoi(ctx.Query("limit", "10"))
	offset, _ := strconv.Atoi(ctx.Query("offset", "0"))
	orderBy := ctx.Query("order_by", "name")
	order := ctx.Query("order", "asc")
	name := ctx.Query("name")
	cpf := ctx.Query("cpf")

	var filters []repositories.Filter = []repositories.Filter{
		*repositories.NewFilter("active", "TRUE", false, false),
	}

	if name != "" {
		filters = append(filters, *repositories.NewFilter("name", name, true, true))
	}

	if cpf != "" {
		filters = append(filters, *repositories.NewFilter("cpf", cpf, true, true))
	}

	response := h.getCustomersUseCase.Execute(limit, offset, orderBy, order, filters)
	if response.ErrorName != nil {
		return ctx.Status(util.CoreErrorToHttpError(*response.ErrorName)).SendString(*response.ErrorMessage)
	}
	return ctx.Status(fiber.StatusOK).JSON(response.Data)
}

//	 	GetCustomerById godoc
//
//		@Summary        Obter um cliente
//		@Description    Obtém um cliente por id.
//		@Tags           customers
//		@Produce        json
//		@Param          id  		path      	integer true  "Id do cliente"
//		@Success        200 		{object}   domain.Customer	"O cliente encontrado"
//		@Failure        400 		{string}   string	 		"Erro de validação"
//		@Failure        404 		{string}   string	 		"Cliente não encontrado"
//		@Failure        500 		{string}   string	 		"Erro interno do sistema"
//		@Router         /customers/{id} [get]
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

//	    UpdateCustomer godoc
//
//		@Summary        Atualizar cliente
//		@Description    Atualiza um cliente.
//		@Description    Campos obrigatórios: nome.
//		@Description    Campos opcionais: CPF, telefone, e-mail, instagram e data de nascimento.
//		@Description    Os endereços também podem ser atualizados. Para criar um endereço, envie um objeto com id vazio. Para deletar um endereço existente, não o envie na lista.
//		@Description    Campos obrigatórios: Rua e bairro.
//		@Description    Campos opcionais: CEP.
//		@Tags           customers
//		@Accept         json
//		@Produce        json
//		@Param          id        path   	integer  		 true    "Id do cliente a ser atualizado"
//		@Param          customer  body      domain.Customer  true    "Cliente a ser atualizado"
//		@Success        204 	 	"Cliente atualizado com sucesso"
//		@Failure        400 	{string}	string	 	"Erro de validação"
//		@Failure        404 	{string}    string	 	"Cliente não encontrado"
//		@Failure        500 	{string}	string	 	"Erro interno do sistema"
//		@Router         /customers [put]
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

//	 	DeleteCustomer godoc
//
//		@Summary        Deleta um cliente
//		@Description    Deleta um cliente por id.
//		@Tags           customers
//		@Produce        json
//		@Param          id  		path      	integer true  "Id do cliente"
//		@Success        204
//		@Failure        400 		{string}   string	 		"Erro de validação"
//		@Failure        404 		{string}   string	 		"Cliente não encontrado"
//		@Failure        500 		{string}   string	 		"Erro interno do sistema"
//		@Router         /customers/{id} [delete]
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
