package usecases

import (
	"log"
	"strings"
	"time"

	"github.com/emersonnobre/tica-api-go/src/internal/core/repositories"
	"github.com/emersonnobre/tica-api-go/src/internal/core/usecases/types"
	"github.com/emersonnobre/tica-api-go/src/internal/core/usecases/types/requests"
)

type CreateProductUseCase struct {
	repository repositories.ProductRepository
}

func NewCreateProductUseCase(repository repositories.ProductRepository) *CreateProductUseCase {
	return &CreateProductUseCase{
		repository: repository,
	}
}

func (u *CreateProductUseCase) Execute(product requests.CreateProductRequest) types.UseCaseResponse {
	// productInDB, err := u.repository.GetByName(product.Name)
	// if err != nil {
	// 	log.Println("ERROR:", err)
	// 	message, errorName := "Erro ao criar produto!", types.GetInternalErrorName()
	// 	return types.NewUseCaseResponse(nil, &errorName, &message)
	// }

	// if productInDB != nil {
	// 	message, errorName := "Já existe um produto com este nome!", types.GetValidationErrorName()
	// 	return types.NewUseCaseResponse(nil, &errorName, &message)
	// }

	if validationErrors := product.Validate(); validationErrors != nil {
		return types.NewErrorUseCaseResponse(types.GetValidationErrorName(), strings.Join(*validationErrors, "\n"))
	}

	target := *product.MapForDomain()
	target.Active = true
	target.CreatedAt = time.Now()

	err := u.repository.Create(target)
	if err != nil {
		log.Println("ERROR:", err)
		message, errorName := "Erro ao criar produto!", types.GetInternalErrorName()
		return types.NewUseCaseResponse(nil, &errorName, &message)
	}

	return types.NewUseCaseResponse(nil, nil, nil)
}
