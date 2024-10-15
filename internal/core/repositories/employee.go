package repositories

import (
	"context"

	"github.com/emersonnobre/tica-api-go/internal/core/domain"
)

type EmployeeRepository interface {
	Create(ctx context.Context, employee *domain.Employee) error
	GetById(ctx context.Context, id string) *domain.Employee
}
