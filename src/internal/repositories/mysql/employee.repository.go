package mysql_repository

import (
	"database/sql"
	"fmt"

	"github.com/emersonnobre/tica-api-go/src/internal/core/domain"
	"github.com/emersonnobre/tica-api-go/src/internal/core/repositories"
	"github.com/emersonnobre/tica-api-go/src/internal/repositories/mysql/util"
)

type MySQLEmployeeRepository struct {
	db *sql.DB
}

func NewMySQLEmployeeRepository(db *sql.DB) *MySQLEmployeeRepository {
	return &MySQLEmployeeRepository{
		db: db,
	}
}

func (r *MySQLEmployeeRepository) Create(employee domain.Employee) error {
	stmt, err := r.db.Prepare("INSERT INTO employees(name, cpf, created_at) VALUES(?, ?, ?)")

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(employee.Name, employee.Cpf, employee.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (r *MySQLEmployeeRepository) GetById(id int) (*domain.Employee, error) {
	var employee domain.Employee
	query := fmt.Sprintf("SELECT id, name, cpf, created_at FROM employees WHERE id = %d", id)
	result := r.db.QueryRow(query)
	err := result.Scan(&employee.Id, &employee.Name, &employee.Cpf, &employee.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &employee, nil
}

func (r *MySQLEmployeeRepository) GetByCPF(cpf string) (*domain.Employee, error) {
	var employee domain.Employee
	query := fmt.Sprintf("SELECT id, name, cpf, created_at FROM employees WHERE cpf = %s", cpf)
	result := r.db.QueryRow(query)
	err := result.Scan(&employee.Id, &employee.Name, &employee.Cpf, &employee.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &employee, nil
}

func (r *MySQLEmployeeRepository) GetCount(filters []repositories.Filter) (int, error) {
	var count int
	query := fmt.Sprintf("SELECT COUNT(*) FROM employees %s", util.BuildConditionsString(filters))
	row := r.db.QueryRow(query)
	if err := row.Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}
