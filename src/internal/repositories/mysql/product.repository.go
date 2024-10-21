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
	product := domain.NewEmptyProduct()

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
					 p.is_feedstock,
					 p.category_id
		FROM products p
		WHERE p.id = %d
		and active = True
	`, id)
	row := r.db.QueryRow(query)
	if row == nil {
		return nil, nil
	}
	err := row.Scan(
		&product.Id,
		&product.Name,
		&product.PurchasePrice,
		&product.SalePrice,
		&product.Stock,
		&product.Barcode,
		&product.Active,
		&product.CreatedAt,
		&product.UpdatedAt,
		&product.IsFeedstock,
		&product.Category.Id,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	query = fmt.Sprintf("SELECT description FROM categories where id = %d", product.Category.Id)
	row = r.db.QueryRow(query)
	row.Scan(&product.Category.Description)

	return product, nil
}

func (r *MySQLProductRepository) Update(product *domain.Product) error {
	var isFeedstock string
	if product.IsFeedstock {
		isFeedstock = "True"
	} else {
		isFeedstock = "False"
	}
	_, err := r.db.Exec(fmt.Sprintf(`
		UPDATE products
		SET name = '%s',
		    purchase_price = %f,
				sale_price = %f,
				stock = %d,
				category_id = %d,
				updated_by = %d,
				updated_at = '%s',
				is_feedstock = %s
		WHERE id = %d
	`, product.Name,
		product.PurchasePrice,
		product.SalePrice,
		product.Stock,
		product.Category.Id,
		product.UpdatedBy.Id,
		product.UpdatedAt.Format("2006-01-02 15:04:05"),
		isFeedstock,
		product.Id,
	))

	if err != nil {
		if err != sql.ErrNoRows {
			return err
		}
	}

	return nil
}

func (r *MySQLProductRepository) Delete(id int) error {
	result, err := r.db.Exec(fmt.Sprintf("UPDATE products SET active = False WHERE id = %d AND active = True", id))
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (r *MySQLProductRepository) Get(limit int, offset int, orderBy string, order string, filters []repositories.Filter) ([]domain.Product, error) {
	query := fmt.Sprintf(`
		SELECT id, name, stock, category_id, created_at, is_feedstock 
		FROM products
		%s 
		ORDER BY %s %s 
		LIMIT %d OFFSET %d
	`, util.BuildConditionsString(filters), orderBy, order, limit, offset)

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []domain.Product
	for rows.Next() {
		product := domain.NewEmptyProduct()
		err := rows.Scan(&product.Id, &product.Name, &product.Stock, &product.Category.Id, &product.CreatedAt, &product.IsFeedstock)
		if err != nil {
			return nil, err
		}

		query = fmt.Sprintf("SELECT description FROM categories WHERE id = %d", product.Category.Id)
		row := r.db.QueryRow(query)
		row.Scan(&product.Category.Description)
		products = append(products, *product)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
