package usecases

import (
	"fmt"
	"log"

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
	count, _ := u.repository.GetCount(fmt.Sprintf("WHERE active = TRUE AND id = %d", id))
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
