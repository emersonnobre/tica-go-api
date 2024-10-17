package usecases

import (
	"log"

	"github.com/emersonnobre/tica-api-go/src/internal/core/domain"
	"github.com/emersonnobre/tica-api-go/src/internal/core/repositories"
	"github.com/emersonnobre/tica-api-go/src/internal/core/usecases/types"
)

type CreateCategoryUseCase struct {
	repository repositories.CategoryRepository
}

func NewCreateCategoryUseCase(repository repositories.CategoryRepository) *CreateCategoryUseCase {
	return &CreateCategoryUseCase{
		repository: repository,
	}
}

func (u *CreateCategoryUseCase) Execute(category domain.Category) types.UseCaseResponse {
	categoryInDB, err := u.repository.GetByName(category.Description)

	if err != nil {
		log.Println("ERROR:", err)
		message, errorName := "Error creating new category!", types.GetInternalErrorName()
		return types.NewUseCaseResponse(nil, &errorName, &message)
	}

	if categoryInDB != nil {
		message, errorName := "A category with this name already exists!", types.GetValidationErrorName()
		return types.NewUseCaseResponse(nil, &errorName, &message)
	}

	if err := u.repository.Create(category); err != nil {
		log.Println("ERROR:", err)
		message, errorName := "Error creating new category!", types.GetInternalErrorName()
		return types.NewUseCaseResponse(nil, &errorName, &message)
	}

	return types.NewUseCaseResponse(nil, nil, nil)
}
