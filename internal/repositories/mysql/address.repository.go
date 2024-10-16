package mysql_repository

import (
	"database/sql"
	"fmt"

	"github.com/emersonnobre/tica-api-go/internal/core/domain"
)

type MySQLAddressRepository struct {
	db *sql.DB
}

func NewMySQLAddressRepository(db *sql.DB) *MySQLAddressRepository {
	return &MySQLAddressRepository{
		db: db,
	}
}

func (r *MySQLAddressRepository) Create(address domain.Address) error {
	stmt, err := r.db.Prepare("INSERT INTO addresses(street, neighborhood, cep, customer_id) VALUES(?, ?, ?, ?)")

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(address.Street, address.Neighborhood, address.Cep, address.CustomerId)
	if err != nil {
		return err
	}

	return nil
}

func (r *MySQLAddressRepository) GetByCustomerId(id int) (*domain.Address, error) {
	var address domain.Address
	query := fmt.Sprintf("SELECT id, street, neighborhood, cep, customer_id FROM addresses WHERE customer_id = %d", id)
	result := r.db.QueryRow(query)
	err := result.Scan(&address.Id, &address.Street, &address.Neighborhood, &address.Cep, &address.CustomerId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &address, nil
}

func (r *MySQLAddressRepository) Delete(id int) error {
	stmt, err := r.db.Prepare("DELETE FROM addresses WHERE id = ?)")

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
