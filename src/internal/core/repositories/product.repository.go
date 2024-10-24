package repositories

import "github.com/emersonnobre/tica-api-go/src/internal/core/domain"

type ProductRepository interface {
	Create(domain.Product) error
	GetCount(filters []Filter) (int, error)
	GetById(id int) (*domain.Product, error)
	Update(*domain.Product) error
	Delete(id int) error
	Get(limit int, offset int, orderBy string, order string, filters []Filter) ([]domain.Product, error)
	UpdateStock(id, stock, employee_id int) error
}
