package main

import (
	"database/sql"
	"log"

	"github.com/emersonnobre/tica-api-go/src/internal/database"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env.development")
	mysqlConn := database.NewMySQLDatabase()
	db, err := mysqlConn.Connect()
	if err != nil {
		log.Fatal(err, err)
	}

	seedCategories(db)
	seedEmployees(db)
	seedProducts(db)
	seedCustomers(db)
	seedAddresses(db)
}

func seedCategories(db *sql.DB) {
	row := db.QueryRow("SELECT COUNT(*) FROM categories")
	count := 0

	if err := row.Scan(&count); err != nil {
		log.Fatal("Error seeding categories table", err)
	}

	if count > 0 {
		return
	}

	_, err := db.Exec(`
		INSERT INTO categories (description)
		VALUES ('louças'),
			   ('americanos'),
			   ('decorações');
				
	`)

	if err != nil {
		log.Fatal("Error seeding categories table", err)
	}
}

func seedEmployees(db *sql.DB) {
	row := db.QueryRow("SELECT COUNT(*) FROM employees")
	count := 0

	if err := row.Scan(&count); err != nil {
		log.Fatal("Error seeding employees table", err)
	}

	if count > 0 {
		return
	}

	_, err := db.Exec(`
		INSERT INTO employees (name, cpf, created_at)
		VALUES ('Emerson Nobre', '00000000000', NOW()),
		 	   ('Murilo Silva', '11111111111', NOW()),
		 	   ('Diogo Menezes', '22222222222', NOW());
				
	`)

	if err != nil {
		log.Fatal("Error seeding employees table", err)
	}
}

func seedProducts(db *sql.DB) {
	row := db.QueryRow("SELECT COUNT(*) FROM products")
	count := 0

	if err := row.Scan(&count); err != nil {
		log.Fatal("Error seeding products table", err)
	}

	if count > 0 {
		return
	}

	_, err := db.Exec(`
		INSERT INTO products (name, purchase_price, sale_price, stock, barcode, category_id, active, created_at, created_by, is_feedstock)
		VALUES ('Prato marfim raso', 19.99, 29.99, 10, "barcode teste", 1, 1, NOW(), 1, 0),
		 	   ('Prato azul fundo', 19.99, 29.99, 14, "barcode teste", 1, 1, NOW(), 2, 0),
		 	   ('Americano indiano com flores', 9.99, 39.99, 30, "barcode teste", 2, 1, NOW(), 1, 0),
		 	   ('Americano animalesco de onça', 10.99, 49.99, 30, "barcode teste", 2, 1, NOW(), 3, 0),
		 	   ('Flores de mesa', 17.99, 42.99, 21, "barcode teste", 3, 1, NOW(), 3, 0),
		 	   ('Vaso azul profundo', 20.99, 49.99, 5, "barcode teste", 3, 1, NOW(), 2, 0);
				
	`)

	if err != nil {
		log.Fatal("Error seeding products table", err)
	}
}

func seedCustomers(db *sql.DB) {
	row := db.QueryRow("SELECT COUNT(*) FROM customers")
	count := 0

	if err := row.Scan(&count); err != nil {
		log.Fatal("Error seeding customers table", err)
	}

	if count > 0 {
		return
	}

	_, err := db.Exec(`
		INSERT INTO customers (name, phone, cpf, email, instagram, birthday, created_at, active)
		VALUES ('Leilane Beatriz', '999999999', '44444444444', 'leilane@example.com', 'leilane', '2000-01-01', NOW(), 1),
		 	   ('Abida Moreira', '999999999', '55555555555', 'abida@example.com', 'abida', '2000-08-24', NOW(), 1);
				
	`)

	if err != nil {
		log.Fatal("Error seeding customers table", err)
	}
}

func seedAddresses(db *sql.DB) {
	row := db.QueryRow("SELECT COUNT(*) FROM addresses")
	count := 0

	if err := row.Scan(&count); err != nil {
		log.Fatal("Error seeding addresses table", err)
	}

	if count > 0 {
		return
	}

	_, err := db.Exec(`
		INSERT INTO addresses (street, neighborhood, cep, customer_id)
		VALUES ('Rua santo angelo, 22', 'Vila Macoré', '9999999', 1),
		 	   ('Rua santo miguelito, 672', 'Vila Amonbarto', '8888888', 2);
				
	`)

	if err != nil {
		log.Fatal("Error seeding addresses table", err)
	}
}
