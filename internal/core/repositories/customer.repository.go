package repositories

import "github.com/emersonnobre/tica-api-go/internal/core/domain"

type CustomerRepository interface {
	Create(domain.Customer) error
	Update(domain.Customer) error
	GetById(int) (*domain.Customer, error)
	GetByCPF(string) (*domain.Customer, error)
}
