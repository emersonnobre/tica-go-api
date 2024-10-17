package repositories

import "github.com/emersonnobre/tica-api-go/internal/core/domain"

type CustomerRepository interface {
	Create(domain.Customer) (*int, error)
	Update(domain.Customer) error
	Get(int, int, string, string) ([]domain.Customer, error)
	GetCount(where string) (int, error)
	GetById(int) (*domain.Customer, error)
	GetByCPF(string) (*domain.Customer, error)
	Delete(int) error
}
