package repositories

import "github.com/emersonnobre/tica-api-go/internal/core/domain"

type CategoryRepository interface {
	Create(domain.Category) error
	GetAll() ([]domain.Category, error)
	GetByName(string) (*domain.Category, error)
}
