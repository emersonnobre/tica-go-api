package usecases

import (
	"log"
	"strconv"

	"github.com/emersonnobre/tica-api-go/internal/core/repositories"
	"github.com/emersonnobre/tica-api-go/internal/core/usecases/types"
)

type RemoveCustomerUseCase struct {
	repository repositories.CustomerRepository
}

func NewRemoveCustomerUseCase(repository repositories.CustomerRepository) *RemoveCustomerUseCase {
	return &RemoveCustomerUseCase{
		repository: repository,
	}
}

func (u *RemoveCustomerUseCase) Execute(id int) types.UseCaseResponse {
	var filters []repositories.Filter = []repositories.Filter{
		*repositories.NewFilter("active", "TRUE", "boolean", false),
		*repositories.NewFilter("id", strconv.Itoa(id), "integer", false),
	}

	count, _ := u.repository.GetCount(filters)
	if count == 0 {
		message, errorName := "Cliente não encontrado!", types.GetNotFoundErrorName()
		return types.NewUseCaseResponse(nil, &errorName, &message)
	}

	if err := u.repository.Delete(id); err != nil {
		log.Println("ERROR:", err)
		message, errorName := "Erro ao remover cliente!", types.GetInternalErrorName()
		return types.NewUseCaseResponse(nil, &errorName, &message)
	}

	return types.NewUseCaseResponse(nil, nil, nil)
}
