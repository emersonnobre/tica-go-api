package handlers

import (
	"net/http"

	"github.com/emersonnobre/tica-api-go/src/internal/core/usecases"
	"github.com/emersonnobre/tica-api-go/src/internal/core/usecases/types/requests"
	"github.com/emersonnobre/tica-api-go/src/internal/delivery/fiber/util"
	"github.com/gofiber/fiber/v2"
)

type SaleHandler struct {
	createSaleUseCase *usecases.CreateSaleUseCase
}

func NewSaleHandler(createSaleUseCase *usecases.CreateSaleUseCase) *SaleHandler {
	return &SaleHandler{
		createSaleUseCase: createSaleUseCase,
	}
}

func (h *SaleHandler) RegisterRoutes(app *fiber.App) {
	group := app.Group("/sales")

	group.Post("/", h.Create)
}

//	    CreateSale godoc
//
//		@Summary        Registrar uma nova venda
//		@Description    Registra uma nova venda.
//		@Description    Requisitos funcionais relacionados: 3A, 3A.1.
//		@Description    Desconto (discount): Desconto total em cima da venda. É opcional.
//		@Description    Observações (comments): É opcional.
//		@Description    Tipo do pagamento (type_of_payment_id): Tipo do pagamento, 1, 2 ou 3. É obrigatório.
//		@Description    Funcionário da venda (employee_id): É obrigatório.
//		@Description    Cliente da venda (customer_id): É obrigatório.
//		@Description    Itens da venda (items): Os produtos da venda e a quantidade de cada um.
//		@Tags           sales
//		@Accept         json
//		@Produce        json
//		@Param          sale  body      requests.CreateSaleRequest  true    "Venda a ser registrada"
//		@Success        201 	{string}	string	 	"Venda registrada com sucesso"
//		@Failure        400 	{string}	string	 	"Erro de validação"
//		@Failure        500 	{string}	string	 	"Erro interno do sistema"
//		@Router         /sales [post]
func (h *SaleHandler) Create(ctx *fiber.Ctx) error {
	var request requests.CreateSaleRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("Erro ao interpretar a requisição!")
	}
	response := h.createSaleUseCase.Execute(&request)
	if response.ErrorName != nil {
		return ctx.Status(util.CoreErrorToHttpError(*response.ErrorName)).SendString(*response.ErrorMessage)
	}
	return ctx.SendStatus(http.StatusCreated)
}
