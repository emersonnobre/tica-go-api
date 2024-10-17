package usecases

import (
	"log"
	"time"

	"github.com/emersonnobre/tica-api-go/src/internal/core/domain"
	"github.com/emersonnobre/tica-api-go/src/internal/core/repositories"
	"github.com/emersonnobre/tica-api-go/src/internal/core/usecases/types"
)

type CreateEmployeeUseCase struct {
	repository repositories.EmployeeRepository
}

func NewCreateEmployeeUseCase(repository repositories.EmployeeRepository) *CreateEmployeeUseCase {
	return &CreateEmployeeUseCase{
		repository: repository,
	}
}

func (u *CreateEmployeeUseCase) Execute(employee domain.Employee) types.UseCaseResponse {
	employeeExists, err := u.repository.GetByCPF(employee.Cpf)

	if err != nil {
		log.Println("ERROR:", err)
		message, errorName := "Error creating new employee!", types.GetInternalErrorName()
		return types.NewUseCaseResponse(nil, &errorName, &message)
	}

	if employeeExists != nil {
		message, errorName := "Já existe um funcionário com este CPF!", types.GetValidationErrorName()
		return types.NewUseCaseResponse(nil, &errorName, &message)
	}

	employee.Id = 0
	employee.CreatedAt = time.Now().Format("2006-01-02 15:04:05")

	if err := u.repository.Create(employee); err != nil {
		log.Println("ERROR:", err)
		message, errorName := "Error creating new employee!", types.GetInternalErrorName()
		return types.NewUseCaseResponse(nil, &errorName, &message)
	}

	return types.NewUseCaseResponse(employee, nil, nil)
}
