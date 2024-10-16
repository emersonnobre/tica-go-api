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

func (r *MySQLCustomerRepository) Update(customer domain.Customer) error {
	stmt, err := r.db.Prepare("UPDATE customers SET name = ?, phone = ?, email = ?, instagram = ?, birthday = ?, updated_at = ? WHERE id = ?")

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(customer.Name, customer.Phone, customer.Email, customer.Instagram, customer.Birthday, customer.UpdatedAt, customer.Id)
	if err != nil {
		return err
	}

	return nil
}

func (r *MySQLCustomerRepository) GetById(id int) (*domain.Customer, error) {
	var customer domain.Customer
	query := fmt.Sprintf("SELECT id, name, cpf, phone, email, instagram, birthday, active, created_at, updated_at FROM customers WHERE id = %d and active = TRUE", id)
	result := r.db.QueryRow(query)
	err := result.Scan(
		&customer.Id,
		&customer.Name,
		&customer.Cpf,
		&customer.Phone,
		&customer.Email,
		&customer.Instagram,
		&customer.Birthday,
		&customer.Active,
		&customer.CreatedAt,
		&customer.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	query = fmt.Sprintf("SELECT id, street, neighborhood, cep, customer_id FROM addresses WHERE customer_id = %d", customer.Id)
	results, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer results.Close()

	var addresses []domain.Address
	for results.Next() {
		var address domain.Address
		err := results.Scan(&address.Id, &address.Street, &address.Neighborhood, &address.Cep, &address.CustomerId)
		if err != nil {
			return nil, err
		}
		addresses = append(addresses, address)
	}
	err = result.Err()

	if err != nil {
		return nil, err
	}

	customer.Addresses = addresses
	return &customer, nil
}

func (r *MySQLCustomerRepository) GetByCPF(cpf string) (*domain.Customer, error) {
	var customer domain.Customer
	query := fmt.Sprintf("SELECT id, name, cpf, phone, email, instagram, birthday, active, created_at FROM customers WHERE cpf = %s and active = TRUE", cpf)
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

func (r *MySQLCustomerRepository) Delete(id int) error {
	stmt, err := r.db.Prepare("UPDATE customers SET active = FALSE, updated_at = NOW() WHERE id = ?")

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
