package responses

import (
	"time"

	"github.com/emersonnobre/tica-api-go/src/internal/core/domain"
)

type ProductResponse struct {
	Id          int              `json:"id"`
	Name        string           `json:"name"`
	Stock       int              `json:"stock"`
	Category    *domain.Category `json:"category"`
	CreatedAt   time.Time        `json:"created_at"`
	IsFeedstock bool             `json:"is_feedstock"`
}

func NewProductResponse(id, stock int, name string, category *domain.Category, createdAt time.Time, isFeedstock bool) *ProductResponse {
	return &ProductResponse{
		Id:          id,
		Name:        name,
		Stock:       stock,
		Category:    category,
		CreatedAt:   createdAt,
		IsFeedstock: isFeedstock,
	}
}
