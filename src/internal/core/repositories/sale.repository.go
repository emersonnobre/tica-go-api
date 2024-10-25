package repositories

import "github.com/emersonnobre/tica-api-go/src/internal/core/domain"

type SaleRepository interface {
	Create(*domain.Sale) error
	Get(limit int, offset int, orderBy string, order string) ([]domain.Sale, error)
	GetCount() (*int, error)
	GetByCustomer(id int) ([]domain.Sale, error)
}
