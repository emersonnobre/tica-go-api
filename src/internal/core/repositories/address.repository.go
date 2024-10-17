package repositories

import "github.com/emersonnobre/tica-api-go/src/internal/core/domain"

type AddressRepository interface {
	Create(domain.Address) error
	Delete(int) error
	GetByCustomerId(int) ([]domain.Address, error)
}
