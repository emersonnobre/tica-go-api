package usecases

import (
	"log"
	"time"

	"github.com/emersonnobre/tica-api-go/internal/core/domain"
	"github.com/emersonnobre/tica-api-go/internal/core/repositories"
	"github.com/emersonnobre/tica-api-go/internal/core/usecases/types"
)

type CreateCustomerUseCase struct {
	repository repositories.CustomerRepository
}

func NewCreateCustomerUseCase(repository repositories.CustomerRepository) *CreateCustomerUseCase {
	return &CreateCustomerUseCase{
		repository: repository,
	}
}

func (u *CreateCustomerUseCase) Execute(customer domain.Customer) types.UseCaseResponse {
	if customer.Cpf != nil {
		customerInDB, err := u.repository.GetByCPF(*customer.Cpf)

		if err != nil {
			log.Println("ERROR:", err)
			message, errorName := "Erro ao criar cliente!", types.GetInternalErrorName()
			return types.NewUseCaseResponse(nil, &errorName, &message)
		}

		if customerInDB != nil {
			message, errorName := "Já existe um cliente com este CPF!", types.GetValidationErrorName()
			return types.NewUseCaseResponse(nil, &errorName, &message)
		}
	}

	customer.Id = 0
	customer.Active = true
	customer.CreatedAt = time.Now().Format("2006-01-02 15:04:05")

	if err := u.repository.Create(customer); err != nil {
		log.Println("ERROR:", err)
		message, errorName := "Erro ao criar cliente!", types.GetInternalErrorName()
		return types.NewUseCaseResponse(nil, &errorName, &message)
	}

	return types.NewUseCaseResponse(nil, nil, nil)
}
