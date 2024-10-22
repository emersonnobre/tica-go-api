package mysql_repository

import (
	"database/sql"

	"github.com/emersonnobre/tica-api-go/src/internal/core/domain"
)

type MySQLTransactionRepository struct {
	db *sql.DB
}

func NewMySQLTransactionRepository(db *sql.DB) *MySQLTransactionRepository {
	return &MySQLTransactionRepository{db: db}
}

func (r *MySQLTransactionRepository) Create(transaction *domain.Transaction) (*int, error) {
	stmt, err := r.db.Prepare(`
        INSERT INTO transactions(reason, quantity, type, created_at, product_id) 
        VALUES(?, ?, ?, ?, ?)
    `)

	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(transaction.Reason, transaction.Quantity, transaction.Type, transaction.CreatedAt, transaction.Product.Id)
	if err != nil {
		return nil, err
	}

	id, _ := result.LastInsertId()
	id32 := int(id)

	return &id32, nil
}
