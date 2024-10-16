package mysql_repository

import (
	"database/sql"
	"fmt"

	"github.com/emersonnobre/tica-api-go/internal/core/domain"
)

type MySQLCustomerRepository struct {
	db *sql.DB
}

func NewMySQLCustomerRepository(db *sql.DB) *MySQLCustomerRepository {
	return &MySQLCustomerRepository{
		db: db,
	}
}

func (r *MySQLCustomerRepository) Create(customer domain.Customer) error {
	stmt, err := r.db.Prepare("INSERT INTO customers(name, cpf, phone, email, instagram, birthday, active, created_at) VALUES(?, ?, ?, ?, ?, ?, ?, ?)")

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(customer.Name, customer.Cpf, customer.Phone, customer.Email, customer.Instagram, customer.Birthday, customer.Active, customer.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (r *MySQLCustomerRepository) GetById(id int) (*domain.Customer, error) {
	var customer domain.Customer
	query := fmt.Sprintf("SELECT id, name, cpf, phone, email, instagram, birthday, active, created_at FROM customers WHERE id = %d", id)
	result := r.db.QueryRow(query)
	err := result.Scan(&customer.Id, &customer.Name, &customer.Cpf, &customer.Phone, &customer.Email, &customer.Instagram, &customer.Birthday, &customer.Active, &customer.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &customer, nil
}

func (r *MySQLCustomerRepository) GetByCPF(cpf string) (*domain.Customer, error) {
	var customer domain.Customer
	query := fmt.Sprintf("SELECT id, name, cpf, phone, email, instagram, birthday, active, created_at FROM customers WHERE cpf = %s", cpf)
	result := r.db.QueryRow(query)
	err := result.Scan(&customer.Id, &customer.Name, &customer.Cpf, &customer.Phone, &customer.Email, &customer.Instagram, &customer.Birthday, &customer.Active, &customer.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &customer, nil
}