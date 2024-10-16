package usecases

import (
	"log"

	"github.com/emersonnobre/tica-api-go/src/internal/core/repositories"
	"github.com/emersonnobre/tica-api-go/src/internal/core/usecases/types"
	"github.com/emersonnobre/tica-api-go/src/internal/core/usecases/types/responses"
)

type GetCustomersUseCase struct {
	repository repositories.CustomerRepository
}

func NewGetCustomersUseCase(repository repositories.CustomerRepository) *GetCustomersUseCase {
	return &GetCustomersUseCase{
		repository: repository,
	}
}

func (u *GetCustomersUseCase) Execute(limit int, offset int, orderBy string, order string, filters []repositories.Filter) types.UseCaseResponse {
	customers, err := u.repository.Get(limit, offset, orderBy, order, filters)

	if err != nil {
		log.Println("ERROR:", err)
		message, errorName := "Erro ao buscar clientes!", types.GetInternalErrorName()
		return types.NewUseCaseResponse(nil, &errorName, &message)
	}

	totalCount, _ := u.repository.GetCount(filters)
	response := responses.NewPaginatedResponse(customers, (offset/limit)+1, limit, totalCount)
	return types.NewUseCaseResponse(response, nil, nil)
}
