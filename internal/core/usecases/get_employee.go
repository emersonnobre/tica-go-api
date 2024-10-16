package usecases

import (
	"log"

	"github.com/emersonnobre/tica-api-go/internal/core/repositories"
	"github.com/emersonnobre/tica-api-go/internal/core/usecases/types"
)

type GetEmployeeUseCase struct {
	repository repositories.EmployeeRepository
}

func NewGetEmployeeUseCase(repository repositories.EmployeeRepository) *GetEmployeeUseCase {
	return &GetEmployeeUseCase{
		repository: repository,
	}
}

func (u *GetEmployeeUseCase) Execute(id int) types.UseCaseResponse {
	employee, err := u.repository.GetById(id)

	if err != nil {
		log.Println("ERROR:", err)
		message, errorName := "Erro ao obter funcionário!", types.GetInternalErrorName()
		return types.NewUseCaseResponse(nil, &errorName, &message)
	}

	if employee == nil {
		message, errorName := "Funcionário não encontrado!", types.GetNotFoundErrorName()
		return types.NewUseCaseResponse(nil, &errorName, &message)
	}

	return types.NewUseCaseResponse(employee, nil, nil)
}
