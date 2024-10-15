package inmemory

import (
	"github.com/emersonnobre/tica-api-go/internal/core/domain"
)

type InMemoryCategoryRepository struct {
	categories []domain.Category
}

func NewInMemoryCategoryRepository() *InMemoryCategoryRepository {
	return &InMemoryCategoryRepository{
		categories: make([]domain.Category, 0),
	}
}

func (r *InMemoryCategoryRepository) Create(category domain.Category) error {
	r.categories = append(r.categories, category)
	return nil
}

func (r *InMemoryCategoryRepository) GetAll() ([]domain.Category, error) {
	return r.categories, nil
}

func (r *InMemoryCategoryRepository) GetByName(description string) (*domain.Category, error) {
	for _, category := range r.categories {
		if category.Description == description {
			return &category, nil
		}
	}
	return nil, nil
}
