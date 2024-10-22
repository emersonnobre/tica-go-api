package usecases

import (
	"fmt"
	"log"
	"time"

	"github.com/emersonnobre/tica-api-go/src/internal/core/domain"
	"github.com/emersonnobre/tica-api-go/src/internal/core/repositories"
	"github.com/emersonnobre/tica-api-go/src/internal/core/usecases/types"
	"github.com/emersonnobre/tica-api-go/src/internal/core/usecases/types/requests"
)

type RegisterProductOutflowUseCase struct {
	productRepository     repositories.ProductRepository
	transactionRepository repositories.TransactionRepository
}

func NewRegisterProductOutflowUseCase(
	productRepository repositories.ProductRepository,
	transactionRepository repositories.TransactionRepository,
) *RegisterProductOutflowUseCase {
	return &RegisterProductOutflowUseCase{
		productRepository:     productRepository,
		transactionRepository: transactionRepository,
	}
}

func (u *RegisterProductOutflowUseCase) Execute(request *requests.ProductOutflow) types.UseCaseResponse {
	product, err := u.productRepository.GetById(request.ProductId)

	if err != nil {
		log.Println("ERROR: ", err)
		return types.NewErrorUseCaseResponse(types.GetInternalErrorName(), "Erro ao registrar saída do produto!")
	}

	if product == nil {
		return types.NewErrorUseCaseResponse(types.GetNotFoundErrorName(), "Produto não encontrado!")
	}

	if product.Stock-request.Quantity < 0 {
		return types.NewErrorUseCaseResponse(
			types.GetValidationErrorName(),
			fmt.Sprintf("Quantidade insuficiente de produtos para registrar a saída. Quantidade atual do produto em estoque: %d", product.Stock),
		)
	}

	transaction := domain.NewTransaction(request.Reason, request.Quantity, 1, time.Now(), &domain.Employee{Id: request.CreatedBy}, product)
	_, err = u.transactionRepository.Create(transaction)

	if err != nil {
		log.Println("ERROR: ", err)
		return types.NewErrorUseCaseResponse(types.GetInternalErrorName(), "Erro ao registrar saída do produto!")
	}

	product.Update(product.Stock-request.Quantity, product.PurchasePrice, &domain.Employee{Id: request.CreatedBy})
	err = u.productRepository.Update(product)

	if err != nil {
		log.Println("ERROR: ", err)
		return types.NewErrorUseCaseResponse(types.GetInternalErrorName(), "Erro ao registrar saída do produto!")
	}

	return types.NewSuccessUseCaseResponse(nil)
}
