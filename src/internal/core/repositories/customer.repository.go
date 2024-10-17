package repositories

import "github.com/emersonnobre/tica-api-go/src/internal/core/domain"

type CustomerRepository interface {
	Create(domain.Customer) (*int, error)
	Update(domain.Customer) error
	Get(limit int, offset int, orderBy string, order string, filters []Filter) ([]domain.Customer, error)
	GetCount(filters []Filter) (int, error)
	GetById(int) (*domain.Customer, error)
	GetByCPF(string) (*domain.Customer, error)
	Delete(int) error
}
