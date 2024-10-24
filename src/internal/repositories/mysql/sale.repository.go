package mysql_repository

import (
	"database/sql"

	"github.com/emersonnobre/tica-api-go/src/internal/core/domain"
)

type MySQLSaleRepository struct {
	db *sql.DB
}

func NewMySQLSaleRepository(db *sql.DB) *MySQLSaleRepository {
	return &MySQLSaleRepository{db: db}
}

func (r *MySQLSaleRepository) Create(sale *domain.Sale) error {
	stmt, err := r.db.Prepare(`
		INSERT INTO sale(total_price, 
						  discount, 
						  comments, 
						  type_of_payment_id, 
						  created_at, 
						  employee_id, 
						  customer_id
						) VALUES(?, ?, ?, ?, ?, ?, ?)`)

	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(sale.TotalPrice, sale.Discount, sale.Comments, sale.TypeOfPayment, sale.CreatedAt, sale.Employee.Id, sale.Customer.Id)
	if err != nil {
		return err
	}

	lastId, _ := result.LastInsertId()
	lastId32 := int32(lastId)

	for _, item := range sale.Items {
		stmt, err = r.db.Prepare(`
		INSERT INTO sale_item(quantity, 
						  product_id, 
						  sale_id
						) VALUES(?, ?, ?)`)

		if err != nil {
			return err
		}

		_, err = stmt.Exec(item.Quantity, item.Product.Id, lastId32)
		if err != nil {
			return err
		}
	}

	return nil
}
