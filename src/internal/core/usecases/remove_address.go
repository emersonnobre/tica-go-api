package usecases

import (
	"log"

	"github.com/emersonnobre/tica-api-go/src/internal/core/repositories"
	"github.com/emersonnobre/tica-api-go/src/internal/core/usecases/types"
)

type RemoveAddressUseCase struct {
	repository repositories.AddressRepository
}

func NewRemoveAddressUseCase(repository repositories.AddressRepository) *RemoveAddressUseCase {
	return &RemoveAddressUseCase{
		repository: repository,
	}
}

func (u *RemoveAddressUseCase) Execute(id int) types.UseCaseResponse {
	if err := u.repository.Delete(id); err != nil {
		log.Println("ERROR:", err)
		message, errorName := "Erro ao remover endereço!", types.GetInternalErrorName()
		return types.NewUseCaseResponse(nil, &errorName, &message)
	}

	return types.NewUseCaseResponse(nil, nil, nil)
}
