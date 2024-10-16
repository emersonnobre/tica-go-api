package usecases

import (
	"log"
	"time"

	"github.com/emersonnobre/tica-api-go/internal/core/domain"
	"github.com/emersonnobre/tica-api-go/internal/core/repositories"
	"github.com/emersonnobre/tica-api-go/internal/core/usecases/types"
)

type CreateCustomerUseCase struct {
	repository           repositories.CustomerRepository
	createAddressUseCase *CreateAddressUseCase
}

func NewCreateCustomerUseCase(repository repositories.CustomerRepository, createAddressUseCase *CreateAddressUseCase) *CreateCustomerUseCase {
	return &CreateCustomerUseCase{
		repository:           repository,
		createAddressUseCase: createAddressUseCase,
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

		if customerInDB != nil && customerInDB.Active {
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

	createdCustomer, err := u.repository.GetByCPF(*customer.Cpf)

	if err != nil {
		log.Println("ERROR:", err)
		message, errorName := "Erro ao criar cliente!", types.GetInternalErrorName()
		return types.NewUseCaseResponse(nil, &errorName, &message)
	}

	if len(customer.Addresses) > 0 {
		for _, address := range customer.Addresses {
			address.CustomerId = createdCustomer.Id
			response := u.createAddressUseCase.Execute(address)
			if response.ErrorName != nil {
				return response
			}
		}
	}

	return types.NewUseCaseResponse(nil, nil, nil)
}
