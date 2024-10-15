package usecases

import (
	"log"

	"github.com/emersonnobre/tica-api-go/internal/core/repositories"
	"github.com/emersonnobre/tica-api-go/internal/core/usecases/types"
)

type GetCategoriesUseCase struct {
	repository repositories.CategoryRepository
}

func NewGetCategoriesUseCase(repository repositories.CategoryRepository) *GetCategoriesUseCase {
	return &GetCategoriesUseCase{
		repository: repository,
	}
}

func (u *GetCategoriesUseCase) Execute() types.UseCaseResponse {
	categories, err := u.repository.GetAll()
	if err != nil {
		log.Println(err)
		message, errorName := "Error getting categories!", types.GetInternalErrorName()
		return types.NewUseCaseResponse(nil, &errorName, &message)
	}
	return types.NewUseCaseResponse(categories, nil, nil)
}
