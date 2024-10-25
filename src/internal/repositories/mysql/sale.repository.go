package mysql_repository

import (
	"database/sql"
	"fmt"

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

func (r *MySQLSaleRepository) Get(limit int, offset int, orderBy string, order string) ([]domain.Sale, error) {
	query := fmt.Sprintf(`
		SELECT s.id, s.total_price, s.discount, s.comments, s.type_of_payment_id, s.created_at, e.id, e.name, e.cpf, c.id, c.name, c.cpf
		FROM sale s 
		INNER JOIN employees e on s.employee_id = e.id
		INNER JOIN customers c on s.customer_id = c.id
		ORDER BY %s %s 
		LIMIT %d OFFSET %d
	`, orderBy, order, limit, offset)
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	sales := []domain.Sale{}
	for rows.Next() {
		sale := domain.Sale{}
		sale.Employee = &domain.Employee{}
		sale.Customer = &domain.Customer{}
		err := rows.Scan(
			&sale.Id,
			&sale.TotalPrice,
			&sale.Discount,
			&sale.Comments,
			&sale.TypeOfPayment,
			&sale.CreatedAt,
			&sale.Employee.Id,
			&sale.Employee.Name,
			&sale.Employee.Cpf,
			&sale.Customer.Id,
			&sale.Customer.Name,
			&sale.Customer.Cpf)
		if err != nil {
			return nil, err
		}
		sales = append(sales, sale)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return sales, nil
}

func (r *MySQLSaleRepository) GetCount() (*int, error) {
	count := 0
	row := r.db.QueryRow("SELECT COUNT(id) FROM sale")
	err := row.Scan(&count)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return &count, nil
}

func (r *MySQLSaleRepository) GetByCustomer(id int) ([]domain.Sale, error) {
	query := fmt.Sprintf(`
		SELECT s.id, s.total_price, s.discount, s.comments, s.type_of_payment_id, s.created_at, e.id, e.name, e.cpf, c.id, c.name, c.cpf
		FROM sale s
		INNER JOIN employees e on s.employee_id = e.id
		INNER JOIN customers c on s.customer_id = c.id
		where customer_id = %d
	`, id)
	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	sales := []domain.Sale{}
	for rows.Next() {
		sale := domain.Sale{}
		sale.Customer = &domain.Customer{}
		sale.Employee = &domain.Employee{}
		err := rows.Scan(
			&sale.Id,
			&sale.TotalPrice,
			&sale.Discount,
			&sale.Comments,
			&sale.TypeOfPayment,
			&sale.CreatedAt,
			&sale.Employee.Id,
			&sale.Employee.Name,
			&sale.Employee.Cpf,
			&sale.Customer.Id,
			&sale.Customer.Name,
			&sale.Customer.Cpf,
		)

		if err != nil {
			return nil, err
		}

		sales = append(sales, sale)
	}

	if rows.Err() != nil && rows.Err() != sql.ErrNoRows {
		return nil, err
	}

	return sales, nil
}
