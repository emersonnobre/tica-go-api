package repositories

import "github.com/emersonnobre/tica-api-go/src/internal/core/domain"

type SaleRepository interface {
	Create(*domain.Sale) error
}
