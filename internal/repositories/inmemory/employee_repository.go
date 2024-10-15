package inmemory

import (
	"context"

	"github.com/emersonnobre/tica-api-go/internal/core/domain"
)

type InMemoryEmployeeRepository struct {
	employees []domain.Employee
}

func NewInMemoryEmployeeRepository() *InMemoryEmployeeRepository {
	return &InMemoryEmployeeRepository{
		employees: make([]domain.Employee, 0),
	}
}

func (r *InMemoryEmployeeRepository) Create(ctx context.Context, employee *domain.Employee) error {
	r.employees = append(r.employees, *employee)
	return nil
}

func (r *InMemoryEmployeeRepository) GetById(ctx context.Context, id string) *domain.Employee {
	for _, employee := range r.employees {
		if employee.Id == id {
			return &employee
		}
	}
	return nil
}
