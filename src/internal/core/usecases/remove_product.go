package usecases

import (
	"database/sql"

	"github.com/emersonnobre/tica-api-go/src/internal/core/repositories"
	"github.com/emersonnobre/tica-api-go/src/internal/core/usecases/types"
)

type RemoveProductUseCase struct {
	repository repositories.ProductRepository
}

func NewRemoveProductUseCase(repository repositories.ProductRepository) *RemoveProductUseCase {
	return &RemoveProductUseCase{repository: repository}
}

func (u *RemoveProductUseCase) Execute(id int) types.UseCaseResponse {
	err := u.repository.Delete(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return types.NewErrorUseCaseResponse(types.GetNotFoundErrorName(), "Produto não encontrado!")
		}
		return types.NewErrorUseCaseResponse(types.GetInternalErrorName(), "Erro ao remover produto!")
	}

	return types.NewUseCaseResponse(nil, nil, nil)
}
