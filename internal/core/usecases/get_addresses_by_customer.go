package usecases

import (
	"log"

	"github.com/emersonnobre/tica-api-go/internal/core/repositories"
	"github.com/emersonnobre/tica-api-go/internal/core/usecases/types"
)

type GetAddressesByCustomerUseCase struct {
	repository repositories.AddressRepository
}

func NewGetAddressesByCustomerUseCase(repository repositories.AddressRepository) *GetAddressesByCustomerUseCase {
	return &GetAddressesByCustomerUseCase{
		repository: repository,
	}
}

func (u *GetAddressesByCustomerUseCase) Execute(id int) types.UseCaseResponse {
	addresses, err := u.repository.GetByCustomerId(id)
	if err != nil {
		log.Println(err)
		message, errorName := "Erro ao obter endereços do cliente!", types.GetInternalErrorName()
		return types.NewUseCaseResponse(nil, &errorName, &message)
	}
	return types.NewUseCaseResponse(addresses, nil, nil)
}
