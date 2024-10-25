package responses

import (
	"time"

	"github.com/emersonnobre/tica-api-go/src/internal/core/domain"
)

type SaleResponse struct {
	Id            int                   `json:"id"`
	TotalPrice    float32               `json:"total_price"`
	Discount      *float32              `json:"discount"`
	Comments      *string               `json:"comments"`
	TypeOfPayment int                   `json:"type_of_payment"`
	CreatedAt     time.Time             `json:"created_at"`
	Employee      *SaleEmployeeResponse `json:"employee"`
	Customer      *SaleCustomerResponse `json:"customer"`
}

type SaleCustomerResponse struct {
	Id   int     `json:"id"`
	Name string  `json:"name"`
	Cpf  *string `json:"cpf"`
}

type SaleEmployeeResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Cpf  string `json:"cpf"`
}

func MapDomainToResponse(sale *domain.Sale) *SaleResponse {
	customer := SaleCustomerResponse{
		Id:   sale.Customer.Id,
		Name: sale.Customer.Name,
		Cpf:  sale.Customer.Cpf,
	}

	employee := SaleEmployeeResponse{
		Id:   sale.Employee.Id,
		Name: sale.Employee.Name,
		Cpf:  sale.Employee.Cpf,
	}

	return &SaleResponse{
		Id:            sale.Id,
		TotalPrice:    sale.TotalPrice,
		Discount:      sale.Discount,
		Comments:      sale.Comments,
		TypeOfPayment: sale.TypeOfPayment,
		CreatedAt:     sale.CreatedAt,
		Employee:      &employee,
		Customer:      &customer,
	}
}
