package repositories

import "github.com/emersonnobre/tica-api-go/internal/core/domain"

type ProductRepository interface {
	Create(domain.Product) error
	GetByName(string) (*domain.Product, error)
}
