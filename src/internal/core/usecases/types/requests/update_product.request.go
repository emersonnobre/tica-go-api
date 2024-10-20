package requests

import (
	"time"

	"github.com/emersonnobre/tica-api-go/src/internal/core/domain"
	domain_validators "github.com/emersonnobre/tica-api-go/src/internal/core/domain/validators"
)

type UpdateProductRequest struct {
	Id            int     `json:"id"`
	Name          string  `json:"name"`
	PurchasePrice float32 `json:"purchase_price"`
	SalePrice     float32 `json:"sale_price"`
	Stock         int     `json:"stock"`
	CategoryId    int     `json:"category_id"`
	UpdatedBy     int     `json:"updated_by"`
	IsFeedstock   bool    `json:"is_feedstock"`
}

func (r *UpdateProductRequest) MapForDomain() *domain.Product {
	now := time.Now()
	return &domain.Product{
		Id:            r.Id,
		Name:          r.Name,
		PurchasePrice: r.PurchasePrice,
		SalePrice:     r.SalePrice,
		Stock:         r.Stock,
		Category:      &domain.Category{Id: r.CategoryId},
		UpdatedBy:     &domain.Employee{Id: r.UpdatedBy},
		UpdatedAt:     &now,
		IsFeedstock:   r.IsFeedstock,
	}
}

func (r *UpdateProductRequest) ValidateObjectData() *[]string {
	validations := []domain_validators.ProductValidator{
		domain_validators.ProductNameValidator{},
		domain_validators.ProductPurchasePriceValidator{},
		domain_validators.ProductSalePriceValidator{},
		domain_validators.ProductStockValidator{},
		domain_validators.ProductCategoryValidator{},
		domain_validators.ProductUpdatedByValidator{},
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
