package repositories

import "github.com/emersonnobre/tica-api-go/src/internal/core/domain"

type CategoryRepository interface {
	Create(domain.Category) error
	GetAll() ([]domain.Category, error)
	GetByName(string) (*domain.Category, error)
	GetCount(filters []Filter) (int, error)
}
