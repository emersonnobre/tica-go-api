package domain_validators

import (
	"errors"
	"strings"
	"time"

	"github.com/emersonnobre/tica-api-go/src/internal/core/domain"
)

type ProductValidator interface {
	Validate(p *domain.Product) error
}

type ProductNameValidator struct{}

func (v ProductNameValidator) Validate(p *domain.Product) error {
	if strings.TrimSpace(p.Name) == "" {
		return errors.New("nome do produto é obrigatório")
	}
	if len(p.Name) > 255 {
		return errors.New("nome do produto não pode ter mais que 255 caracteres")
	}
	return nil
}

type ProductPurchasePriceValidator struct{}

func (v ProductPurchasePriceValidator) Validate(p *domain.Product) error {
	if p.PurchasePrice <= 0 {
		return errors.New("preço de compra do produto deve ser maior que zero")
	}
	return nil
}

type ProductSalePriceValidator struct{}

func (v ProductSalePriceValidator) Validate(p *domain.Product) error {
	if p.SalePrice <= 0 {
		return errors.New("preço de venda do produto deve ser maior que zero")
	}
	return nil
}

type ProductStockValidator struct{}

func (v ProductStockValidator) Validate(p *domain.Product) error {
	if p.Stock < 0 {
		return errors.New("estoque do produto deve ser maior ou igual a zero")
	}
	return nil
}

type ProductBarcodeValidator struct{}

func (v ProductBarcodeValidator) Validate(p *domain.Product) error {
	if strings.TrimSpace(p.Barcode) == "" {
		return errors.New("código de barras do produto é obrigatório")
	}
	if len(p.Barcode) > 255 {
		return errors.New("código de barras do produto não pode ter mais que 255 caracteres")
	}
	return nil
}

type ProductCategoryValidator struct{}

func (v ProductCategoryValidator) Validate(p *domain.Product) error {
	if p.Category == nil || p.Category.Id == 0 {
		return errors.New("categoria do produto é obrigatória")
	}
	return nil
}

type ProductCreatedAtValidator struct{}

func (v ProductCreatedAtValidator) Validate(p *domain.Product) error {
	if p.CreatedAt.After(time.Now()) {
		return errors.New("data de criação do produto não pode ser maior que a data atual")
	}
	return nil
}

type ProductCreatedByValidator struct{}

func (v ProductCreatedByValidator) Validate(p *domain.Product) error {
	if p.CreatedBy == nil || p.CreatedBy.Id == 0 {
		return errors.New("usuário de criação do produto é obrigatório")
	}
	return nil
}

type ProductUpdatedAtValidator struct{}

func (v ProductUpdatedAtValidator) Validate(p *domain.Product) error {
	if p.UpdatedAt.After(time.Now()) {
		return errors.New("data de atualização do produto não pode ser maior que a data atual")
	}
	return nil
}

type ProductUpdatedByValidator struct{}

func (v ProductUpdatedByValidator) Validate(p *domain.Product) error {
	if p.UpdatedBy == nil || p.UpdatedBy.Id == 0 {
		return errors.New("usuário de atualização do produto é obrigatório")
	}
	return nil
}
