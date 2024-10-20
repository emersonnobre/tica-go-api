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

func (r *MySQLProductRepository) GetById(id int) (*domain.Product, error) {
	var product domain.Product
	query := fmt.Sprintf(`
		SELECT p.id,
					 p.name,
					 p.purchase_price,
					 p.sale_price,
					 p.stock,
					 p.barcode,
					 p.active,
					 p.created_at,
					 p.updated_at,
					 p.is_feedstock
		FROM products p
		WHERE p.id = %d
	`, id)
	row := r.db.QueryRow(query)
	if row == nil {
		return nil, nil
	}
	row.Scan(
		&product.Id,
		&product.Name,
		&product.PurchasePrice,
		&product.SalePrice,
		&product.Stock,
		&product.Barcode,
		&product.Active,
		&product.CreatedAt,
		&product.UpdatedAt,
		&product.IsFeedstock)
	return &product, nil
}

func (r *MySQLProductRepository) Update(product *domain.Product) error {
	_, err := r.db.Exec(fmt.Sprintf(`
		UPDATE products
		SET name = '%s',
		    purchase_price = %f,
				sale_price = %f,
				stock = %d,
				category_id = %d,
				updated_by = %d,
				updated_at = %s,
				is_feedstock = %t
		WHERE id = %d
	`, product.Name,
		product.PurchasePrice,
		product.SalePrice,
		product.Stock,
		product.Category.Id,
		product.UpdatedBy.Id,
		product.UpdatedAt,
		product.IsFeedstock,
		product.Id,
	))

	if err != nil {
		if err != sql.ErrNoRows {
			return err
		}
	}

	return nil
}
