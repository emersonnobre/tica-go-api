package usecases

import (
	"log"

	"github.com/emersonnobre/tica-api-go/src/internal/core/repositories"
	"github.com/emersonnobre/tica-api-go/src/internal/core/usecases/types"
	"github.com/emersonnobre/tica-api-go/src/internal/core/usecases/types/responses"
)

type GetProductsUseCase struct {
	repository repositories.ProductRepository
}

func NewGetProductsUseCase(repository repositories.ProductRepository) *GetProductsUseCase {
	return &GetProductsUseCase{
		repository: repository,
	}
}

func (u *GetProductsUseCase) Execute(limit int, offset int, orderBy string, order string, filters []repositories.Filter) types.UseCaseResponse {
	products, err := u.repository.Get(limit, offset, orderBy, order, filters)

	if err != nil {
		log.Println("ERROR:", err)
		return types.NewErrorUseCaseResponse(types.GetInternalErrorName(), "Erro ao obter produtos!")
	}

	productsResponse := []responses.ProductResponse{}
	for _, obj := range products {
		productResponse := responses.NewProductResponse(obj.Id, obj.Stock, obj.Name, obj.Category, obj.CreatedAt, obj.IsFeedstock)
		productsResponse = append(productsResponse, *productResponse)
	}

	totalCount, _ := u.repository.GetCount(filters)
	response := responses.NewPaginatedResponse(productsResponse, (offset/limit)+1, limit, totalCount)
	return types.NewUseCaseResponse(response, nil, nil)
}
