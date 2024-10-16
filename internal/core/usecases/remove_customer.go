package usecases

import (
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
	if err := u.repository.Delete(id); err != nil {
		log.Println("ERROR:", err)
		message, errorName := "Erro ao remover cliente!", types.GetInternalErrorName()
		return types.NewUseCaseResponse(nil, &errorName, &message)
	}

	return types.NewUseCaseResponse(nil, nil, nil)
}
