package usecases

import (
	"github.com/emersonnobre/tica-api-go/src/internal/core/repositories"
	"github.com/emersonnobre/tica-api-go/src/internal/core/usecases/types"
	"github.com/emersonnobre/tica-api-go/src/internal/core/usecases/types/requests"
	"github.com/emersonnobre/tica-api-go/src/internal/core/usecases/types/responses"
	"github.com/gofiber/fiber/v2/log"
)

type GetSalesUseCase struct {
	repository repositories.SaleRepository
}

func NewGetSalesUseCase(repository repositories.SaleRepository) *GetSalesUseCase {
	return &GetSalesUseCase{repository: repository}
}

func (u *GetSalesUseCase) Execute(request *requests.GetSalesRequest) types.UseCaseResponse {
	sales, err := u.repository.Get(request.Limit(), request.Offset(), request.OrderBy(), request.Order())

	if err != nil {
		log.Error(err)
		return types.NewErrorUseCaseResponse(types.GetInternalErrorName(), "Erro ao obter vendas!")
	}

	salesResponse := []responses.SaleResponse{}
	for _, sale := range sales {
		response := responses.MapDomainToResponse(&sale)
		salesResponse = append(salesResponse, *response)
	}

	totalCount, _ := u.repository.GetCount()
	response := responses.NewPaginatedResponse(salesResponse, (request.Offset()/request.Limit())+1, request.Limit(), *totalCount)

	return types.NewSuccessUseCaseResponse(response)
}
