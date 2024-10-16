package usecases

import (
	"log"
	"strings"

	"github.com/emersonnobre/tica-api-go/internal/core/domain"
	"github.com/emersonnobre/tica-api-go/internal/core/repositories"
	"github.com/emersonnobre/tica-api-go/internal/core/usecases/types"
)

type CreateAddressUseCase struct {
	repository repositories.AddressRepository
}

func NewCreateAddressUseCase(repository repositories.AddressRepository) *CreateAddressUseCase {
	return &CreateAddressUseCase{
		repository: repository,
	}
}

func (u *CreateAddressUseCase) Execute(address domain.Address) types.UseCaseResponse {
	validationErrors := make([]string, 0)

	if address.Street == "" {
		validationErrors = append(validationErrors, "Informe a rua do endereço!")
	}

	if address.Neighborhood == "" {
		validationErrors = append(validationErrors, "Informe o bairro do endereço!")
	}

	if len(validationErrors) > 0 {
		errorName, message := types.GetValidationErrorName(), strings.Join(validationErrors, "\n")
		return types.NewUseCaseResponse(nil, &errorName, &message)
	}

	address.Id = 0

	if err := u.repository.Create(address); err != nil {
		log.Println("ERROR:", err)
		message, errorName := "Erro ao criar endereço!", types.GetInternalErrorName()
		return types.NewUseCaseResponse(nil, &errorName, &message)
	}

	return types.NewUseCaseResponse(nil, nil, nil)
}
