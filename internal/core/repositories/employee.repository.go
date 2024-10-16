package repositories

import (
	"github.com/emersonnobre/tica-api-go/internal/core/domain"
)

type EmployeeRepository interface {
	Create(domain.Employee) error
	GetByCPF(string) (*domain.Employee, error)
	GetById(int) (*domain.Employee, error)
}