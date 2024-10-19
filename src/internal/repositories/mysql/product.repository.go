package mysql_repository

import (
	"database/sql"
	"fmt"

	"github.com/emersonnobre/tica-api-go/src/internal/core/domain"
	"github.com/emersonnobre/tica-api-go/src/internal/core/repositories"
	"github.com/emersonnobre/tica-api-go/src/internal/repositories/mysql/util"
)

type MySQLProductRepository struct {
	db *sql.DB
}

func NewMySQLProductRepository(db *sql.DB) *MySQLProductRepository {
	return &MySQLProductRepository{
		db: db,
	}
}

func (r *MySQLProductRepository) Create(product domain.Product) error {
	stmt, err := r.db.Prepare("INSERT INTO products(name, purchase_price, sale_price, stock, barcode, category_id, active, created_at, created_by, is_feedstock) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.Name, product.PurchasePrice, product.SalePrice, product.Stock, product.Barcode, product.Category.Id, product.Active, product.CreatedAt, product.CreatedBy.Id, product.IsFeedstock)
	if err != nil {
		return err
	}

	return nil
}

func (r *MySQLProductRepository) GetCount(filters []repositories.Filter) (int, error) {
	var count int
	query := fmt.Sprintf("SELECT COUNT(*) FROM products %s", util.BuildConditionsString(filters))
	row := r.db.QueryRow(query)
	if err := row.Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}
