package repositories

import "github.com/emersonnobre/tica-api-go/src/internal/core/domain"

type TransactionRepository interface {
	Create(*domain.Transaction) (*int, error)
}
