package usecases

import (
	"log"
	"time"

	"github.com/emersonnobre/tica-api-go/src/internal/core/repositories"
	"github.com/emersonnobre/tica-api-go/src/internal/core/usecases/types"
)

type GetProductUseCase struct {
	repository repositories.ProductRepository
}

func NewGetProductUseCase(repository repositories.ProductRepository) *GetProductUseCase {
	return &GetProductUseCase{repository: repository}
}

func (u *GetProductUseCase) Execute(id int) types.UseCaseResponse {
	product, err := u.repository.GetById(id)

	if err != nil {
		log.Println("ERROR: ", err)
		return types.NewErrorUseCaseResponse(types.GetInternalErrorName(), "Erro ao obter produto!")
	}

	if product == nil {
		return types.NewErrorUseCaseResponse(types.GetNotFoundErrorName(), "Produto não encontrado!")
	}

	loc, _ := time.LoadLocation("America/Campo_Grande")
	product.CreatedAt = product.CreatedAt.In(loc)
	return types.NewUseCaseResponse(product, nil, nil)
}
