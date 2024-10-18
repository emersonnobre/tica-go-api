package requests

import (
	"github.com/emersonnobre/tica-api-go/src/internal/core/domain"
	domain_validators "github.com/emersonnobre/tica-api-go/src/internal/core/domain/validators"
)

type CreateProductRequest struct {
	Name          string  `json:"name"`
	PurchasePrice float32 `json:"purchase_price"`
	SalePrice     float32 `json:"sale_price"`
	Stock         int     `json:"stock"`
	CategoryId    int     `json:"category_id"`
	CreatedBy     int     `json:"created_by"`
	IsFeedstock   bool    `json:"is_feedstock"`
}

func (r *CreateProductRequest) MapForDomain() *domain.Product {
	return &domain.Product{
		Name:          r.Name,
		PurchasePrice: r.PurchasePrice,
		SalePrice:     r.SalePrice,
		Stock:         r.Stock,
		Category:      &domain.Category{Id: r.CategoryId},
		CreatedBy:     &domain.Employee{Id: r.CreatedBy},
		IsFeedstock:   r.IsFeedstock,
	}
}

func (r *CreateProductRequest) Validate() *[]string {
	validations := []domain_validators.ProductValidator{
		domain_validators.ProductNameValidator{},
		domain_validators.ProductPurchasePriceValidator{},
		domain_validators.ProductSalePriceValidator{},
		domain_validators.ProductStockValidator{},
		domain_validators.ProductCategoryValidator{},
		domain_validators.ProductCreatedByValidator{},
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
