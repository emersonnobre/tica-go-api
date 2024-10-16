package usecases

import (
	"log"

	"github.com/emersonnobre/tica-api-go/internal/core/repositories"
	"github.com/emersonnobre/tica-api-go/internal/core/usecases/types"
)

type GetCustomerUseCase struct {
	repository repositories.CustomerRepository
}

func NewGetCustomerUseCase(repository repositories.CustomerRepository) *GetCustomerUseCase {
	return &GetCustomerUseCase{
		repository: repository,
	}
}

func (u *GetCustomerUseCase) Execute(id int) types.UseCaseResponse {
	customer, err := u.repository.GetById(id)

	if err != nil {
		log.Println("ERROR:", err)
		message, errorName := "Erro ao buscar cliente!", types.GetInternalErrorName()
		return types.NewUseCaseResponse(nil, &errorName, &message)
	}

	if customer == nil {
		message, errorName := "Cliente não encontrado!", types.GetNotFoundErrorName()
		return types.NewUseCaseResponse(nil, &errorName, &message)
	}

	return types.NewUseCaseResponse(customer, nil, nil)
}
