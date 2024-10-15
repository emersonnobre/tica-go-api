package mysql_repository

import (
	"database/sql"
	"fmt"

	"github.com/emersonnobre/tica-api-go/internal/core/domain"
)

type MySQLCategoryRepository struct {
	db *sql.DB
}

func NewMySQLCategoryRepository(db *sql.DB) *MySQLCategoryRepository {
	return &MySQLCategoryRepository{
		db: db,
	}
}

func (r *MySQLCategoryRepository) Create(category domain.Category) error {
	stmt, err := r.db.Prepare("INSERT INTO category(description) VALUES(?)")

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(category.Description)
	if err != nil {
		return err
	}

	return nil
}

func (r *MySQLCategoryRepository) GetAll() ([]domain.Category, error) {
	result, err := r.db.Query("SELECT id, description FROM categories")
	if err != nil {
		return nil, err
	}
	defer result.Close()

	var categories []domain.Category
	for result.Next() {
		var category domain.Category
		err := result.Scan(&category.Id, &category.Description)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	err = result.Err()
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *MySQLCategoryRepository) GetByName(description string) (*domain.Category, error) {
	var category domain.Category
	query := fmt.Sprintf("SELECT id, description FROM categories where description like '%%%s%%'", description)
	result := r.db.QueryRow(query)
	err := result.Scan(&category.Id, &category.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &category, nil
}
