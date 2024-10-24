package domain_validators

import (
	"errors"

	"github.com/emersonnobre/tica-api-go/src/internal/core/domain"
)

type SaleValidator interface {
	Validate(p *domain.Sale) error
}

type SaleTotalPriceValidator struct{}

func (v SaleTotalPriceValidator) Validate(sale *domain.Sale) error {
	if sale.Discount != nil && *sale.Discount < 0 {
		return errors.New("desconto deve ser maior ou igual a zero")
	}
	if sale.Comments != nil && len(*sale.Comments) > 255 {
		return errors.New("o campo de observações da venda deve ter no máximo 255 caracteres")
	}
	return nil
}

type SaleEmployeeValidator struct{}

func (v SaleEmployeeValidator) Validate(sale *domain.Sale) error {
	if sale.Employee == nil || sale.Employee.Id == 0 {
		return errors.New("funcionário da venda é obrigatório")
	}
	return nil
}

type SaleCustomerValidator struct{}

func (v SaleCustomerValidator) Validate(sale *domain.Sale) error {
	if sale.Customer == nil || sale.Customer.Id == 0 {
		return errors.New("cliente da venda é obrigatório")
	}
	return nil
}

type SaleQuantityValidator struct{}

func (v SaleQuantityValidator) Validate(sale *domain.Sale) error {
	for _, item := range sale.Items {
		if item.Quantity <= 0 {
			return errors.New("quantidade de produtos deve ser maior que zero")
		}
	}
	return nil
}

type SaleItemProductValidator struct{}

func (v SaleItemProductValidator) Validate(sale *domain.Sale) error {
	for _, item := range sale.Items {
		if item.Product == nil || item.Product.Id == 0 {
			return errors.New("produto da venda é obrigatório")
		}
	}
	return nil
}
