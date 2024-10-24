package requests

import (
	"github.com/emersonnobre/tica-api-go/src/internal/core/domain"
	domain_validators "github.com/emersonnobre/tica-api-go/src/internal/core/domain/validators"
)

type CreateSaleRequest struct {
	Discount      *float32                `json:"discount"`
	Comments      *string                 `json:"comments"`
	TypeOfPayment int                     `json:"type_of_payment"`
	EmployeeId    int                     `json:"employee_id"`
	CustomerId    int                     `json:"customer_id"`
	Items         []CreateSaleItemRequest `json:"items"`
}

type CreateSaleItemRequest struct {
	Quantity  int `json:"quantity"`
	ProductId int `json:"product_id"`
}

func (r *CreateSaleRequest) MapForDomain() *domain.Sale {
	saleItems := make([]domain.SaleItem, 0)
	for _, item := range r.Items {
		saleItems = append(saleItems, domain.SaleItem{Quantity: item.Quantity, Product: &domain.Product{Id: item.ProductId}})
	}

	return &domain.Sale{
		Discount:      r.Discount,
		Comments:      r.Comments,
		TypeOfPayment: r.TypeOfPayment,
		Employee:      &domain.Employee{Id: r.EmployeeId},
		Customer:      &domain.Customer{Id: r.CustomerId},
		Items:         saleItems,
	}
}

func (r *CreateSaleRequest) ValidateObjectData() *[]string {
	validations := []domain_validators.SaleValidator{
		domain_validators.SaleTotalPriceValidator{},
		domain_validators.SaleCustomerValidator{},
		domain_validators.SaleEmployeeValidator{},
		domain_validators.SaleQuantityValidator{},
		domain_validators.SaleItemProductValidator{},
	}

	errors := []string{}
	for _, validation := range validations {
		if err := validation.Validate(r.MapForDomain()); err != nil {
			errors = append(errors, err.Error())
		}
	}

	if len(errors) > 0 {
		return &errors
	}

	return nil
}
