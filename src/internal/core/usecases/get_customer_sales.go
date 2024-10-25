package usecases

import (
	"github.com/emersonnobre/tica-api-go/src/internal/core/repositories"
	"github.com/emersonnobre/tica-api-go/src/internal/core/usecases/types"
	"github.com/emersonnobre/tica-api-go/src/internal/core/usecases/types/responses"
	"github.com/gofiber/fiber/v2/log"
)

type GetCustomerSalesUseCase struct {
	repository repositories.SaleRepository
}

func NewGetCustomerSalesUseCase(repository repositories.SaleRepository) *GetCustomerSalesUseCase {
	return &GetCustomerSalesUseCase{repository: repository}
}

func (u *GetCustomerSalesUseCase) Execute(id int) types.UseCaseResponse {
	sales, err := u.repository.GetByCustomer(id)
	if err != nil {
		log.Error(err)
		return types.NewErrorUseCaseResponse(types.GetInternalErrorName(), "Erro ao obter as vendas!")
	}
	salesResponse := []responses.SaleResponse{}
	for _, sale := range sales {
		salesResponse = append(salesResponse, *responses.MapDomainToResponse(&sale))
	}
	return types.NewSuccessUseCaseResponse(salesResponse)
}
