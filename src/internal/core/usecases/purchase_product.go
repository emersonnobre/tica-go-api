package usecases

import (
	"log"
	"time"

	"github.com/emersonnobre/tica-api-go/src/internal/core/domain"
	"github.com/emersonnobre/tica-api-go/src/internal/core/repositories"
	"github.com/emersonnobre/tica-api-go/src/internal/core/usecases/types"
	"github.com/emersonnobre/tica-api-go/src/internal/core/usecases/types/requests"
)

type PurchaseProductUseCase struct {
	productRepository     repositories.ProductRepository
	transactionRepository repositories.TransactionRepository
}

func NewPurchaseProductUseCase(
	productRepository repositories.ProductRepository,
	transactionRepository repositories.TransactionRepository,
) *PurchaseProductUseCase {
	return &PurchaseProductUseCase{
		productRepository:     productRepository,
		transactionRepository: transactionRepository,
	}
}

func (u *PurchaseProductUseCase) Execute(request *requests.PurchaseProductRequest) types.UseCaseResponse {
	product, err := u.productRepository.GetById(request.ProductId)

	if err != nil {
		log.Println("ERROR: ", err)
		return types.NewErrorUseCaseResponse(types.GetInternalErrorName(), "Erro ao registrar compra de produto!")
	}

	if product == nil {
		return types.NewErrorUseCaseResponse(types.GetNotFoundErrorName(), "Produto não encontrado!")
	}

	newTransaction := domain.NewTransaction("purchase", request.Quantity, 0, time.Now(), &domain.Employee{Id: request.CreatedBy}, product)
	_, err = u.transactionRepository.Create(newTransaction)

	if err != nil {
		log.Println("ERROR: ", err)
		return types.NewErrorUseCaseResponse(types.GetInternalErrorName(), "Erro ao registrar compra de produto!")
	}

	product.Update(product.Stock+request.Quantity, request.PurchasePrice, &domain.Employee{Id: request.CreatedBy})
	if err = u.productRepository.Update(product); err != nil {
		log.Println("ERROR: ", err)
		return types.NewErrorUseCaseResponse(types.GetInternalErrorName(), "Erro ao registrar compra de produto!")
	}

	return types.NewSuccessUseCaseResponse(nil)
}
