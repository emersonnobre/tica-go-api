package usecases

import (
	"log"

	"github.com/emersonnobre/tica-api-go/src/internal/core/domain"
	"github.com/emersonnobre/tica-api-go/src/internal/core/repositories"
	"github.com/emersonnobre/tica-api-go/src/internal/core/usecases/types"
)

type CreateProductUseCase struct {
	repository repositories.ProductRepository
}

func NewCreateProductUseCase(repository repositories.ProductRepository) *CreateProductUseCase {
	return &CreateProductUseCase{
		repository: repository,
	}
}

func (u *CreateProductUseCase) Execute(product domain.Product) types.UseCaseResponse {
	productInDB, err := u.repository.GetByName(product.Name)
	if err != nil {
		log.Println("ERROR:", err)
		message, errorName := "Erro ao criar produto!", types.GetInternalErrorName()
		return types.NewUseCaseResponse(nil, &errorName, &message)
	}

	if productInDB != nil {
		message, errorName := "Já existe um produto com este nome!", types.GetValidationErrorName()
		return types.NewUseCaseResponse(nil, &errorName, &message)
	}

	err = u.repository.Create(product)
	if err != nil {
		log.Println("ERROR:", err)
		message, errorName := "Erro ao criar produto!", types.GetInternalErrorName()
		return types.NewUseCaseResponse(nil, &errorName, &message)
	}

	return types.NewUseCaseResponse(nil, nil, nil)
}
