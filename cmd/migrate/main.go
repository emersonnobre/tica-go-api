package main

import (
	"fmt"
	"log"

	"github.com/emersonnobre/tica-api-go/internal/database"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	// cfg := mysqlDriver.Config{
	// 	Addr:         os.Getenv("MYSQL_HOST"),
	// 	DBName: os.Getenv("MYSQL_DATABASE"),
	// 	User:         os.Getenv("MYSQL_USER"),
	// 	Passwd:     os.Getenv("MYSQL_PASSWORD"),
	// }

	mysqlConn := database.NewMySQLDatabase()
	db, err := mysqlConn.Connect()
	if err != nil {
		log.Fatal(err)
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{})

	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations",
		"mysql",
		driver,
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(m)
	m.Up()
}
