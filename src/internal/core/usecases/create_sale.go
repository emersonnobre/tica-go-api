package usecases

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/emersonnobre/tica-api-go/src/internal/core/domain"
	domain_validators "github.com/emersonnobre/tica-api-go/src/internal/core/domain/validators"
	"github.com/emersonnobre/tica-api-go/src/internal/core/repositories"
	"github.com/emersonnobre/tica-api-go/src/internal/core/usecases/types"
	"github.com/emersonnobre/tica-api-go/src/internal/core/usecases/types/requests"
)

type CreateSaleUseCase struct {
	saleRepository     repositories.SaleRepository
	employeeRepository repositories.EmployeeRepository
	customerRepository repositories.CustomerRepository
	productRepository  repositories.ProductRepository
}

func NewCreateSaleUseCase(
	saleRepository repositories.SaleRepository,
	employeeRepository repositories.EmployeeRepository,
	customerRepository repositories.CustomerRepository,
	productRepository repositories.ProductRepository,
) *CreateSaleUseCase {
	return &CreateSaleUseCase{
		saleRepository:     saleRepository,
		employeeRepository: employeeRepository,
		customerRepository: customerRepository,
		productRepository:  productRepository,
	}
}

func (u *CreateSaleUseCase) Execute(request *requests.CreateSaleRequest) types.UseCaseResponse {
	validationErrors := request.ValidateObjectData()
	if validationErrors != nil {
		return types.NewErrorUseCaseResponse(types.GetValidationErrorName(), strings.Join(*validationErrors, "\n"))
	}

	sale := request.MapForDomain()
	if databaseErrors := u.validateForeignKeys(sale); databaseErrors != nil {
		message := strings.Join(*databaseErrors, "\n")
		return types.NewErrorUseCaseResponse(types.GetValidationErrorName(), message)
	}

	sale.CreatedAt = time.Now()

	for _, item := range sale.Items {
		p, _ := u.productRepository.GetById(item.Product.Id)

		if err := u.validateProductStock(*p, item.Quantity); err != nil {
			return types.NewErrorUseCaseResponse(types.GetValidationErrorName(), *err)
		}

		err := u.productRepository.UpdateStock(item.Product.Id, p.Stock-item.Quantity, sale.Employee.Id)
		if err != nil {
			return types.NewErrorUseCaseResponse(types.GetInternalErrorName(), "Erro ao registrar venda!")
		}
		sale.TotalPrice += p.SalePrice
	}
	if sale.Discount != nil {
		sale.TotalPrice -= *sale.Discount
	}

	if err := u.saleRepository.Create(sale); err != nil {
		log.Println("ERROR: ", err)
		return types.NewErrorUseCaseResponse(types.GetInternalErrorName(), "Erro ao registrar venda!")
	}

	return types.NewSuccessUseCaseResponse(nil)
}

func (u *CreateSaleUseCase) validateProductStock(product domain.Product, outflow int) *string {
	product.Stock -= outflow
	stockValidator := domain_validators.ProductStockValidator{}
	err := stockValidator.Validate(&product)
	if err != nil {
		msg := fmt.Sprintf("Quantidade em estoque insuficiente para: %s", product.Name)
		return &msg
	}
	return nil
}

func (u *CreateSaleUseCase) validateForeignKeys(sale *domain.Sale) *[]string {
	errors := []string{}

	filters := []repositories.Filter{
		*repositories.NewFilter("id", strconv.Itoa(sale.Employee.Id), false, false),
	}

	count, _ := u.employeeRepository.GetCount(filters)
	if invalidEmployee := count == 0; invalidEmployee {
		errors = append(errors, "Funcionário inexistente!")
	}

	filters = []repositories.Filter{
		*repositories.NewFilter("id", strconv.Itoa(sale.Customer.Id), false, false),
	}

	count, _ = u.customerRepository.GetCount(filters)
	if invalidCustomer := count == 0; invalidCustomer {
		errors = append(errors, "Cliente inexistente!")
	}

	for _, item := range sale.Items {
		filters = []repositories.Filter{
			*repositories.NewFilter("id", strconv.Itoa(item.Product.Id), false, false),
		}

		count, _ = u.productRepository.GetCount(filters)
		if invalidProduct := count == 0; invalidProduct {
			errors = append(errors, "Produto inexistente!")
		}
	}

	if len(errors) > 0 {
		return &errors
	}

	return nil
}
