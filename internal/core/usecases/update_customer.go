package usecases

import (
	"log"
	"time"

	"github.com/emersonnobre/tica-api-go/internal/core/domain"
	"github.com/emersonnobre/tica-api-go/internal/core/repositories"
	"github.com/emersonnobre/tica-api-go/internal/core/usecases/types"
)

type UpdateCustomerUseCase struct {
	repository repositories.CustomerRepository
}

func NewUpdateCustomerUseCase(repository repositories.CustomerRepository) *UpdateCustomerUseCase {
	return &UpdateCustomerUseCase{
		repository: repository,
	}
}

func (u *UpdateCustomerUseCase) Execute(customer domain.Customer) types.UseCaseResponse {
	customerInDB, err := u.repository.GetByCPF(*customer.Cpf)

	if err != nil {
		log.Println("ERROR:", err)
		message, errorName := "Erro ao atualizar cliente!", types.GetInternalErrorName()
		return types.NewUseCaseResponse(nil, &errorName, &message)
	}

	if customerInDB == nil {
		message, errorName := "Cliente não encontrado!", types.GetValidationErrorName()
		return types.NewUseCaseResponse(nil, &errorName, &message)
	}

	customerInDB.Name = customer.Name
	customerInDB.Phone = customer.Phone
	customerInDB.Email = customer.Email
	customerInDB.Instagram = customer.Instagram
	customerInDB.Birthday = customer.Birthday
	currentDate := time.Now().Format("2006-01-02 15:04:05")
	customerInDB.UpdatedAt = &currentDate

	if err := u.repository.Update(*customerInDB); err != nil {
		log.Println("ERROR:", err)
		message, errorName := "Erro ao atualizar cliente!", types.GetInternalErrorName()
		return types.NewUseCaseResponse(nil, &errorName, &message)
	}

	return types.NewUseCaseResponse(nil, nil, nil)
}
