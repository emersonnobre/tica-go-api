package main

import (
	"fmt"
	"log"
	"os"

	"github.com/emersonnobre/tica-api-go/internal/database"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/joho/godotenv"
)

func main() {
	var env, action string
	if len(os.Args) > 1 {
		env = os.Args[1]
		action = os.Args[2]
	}

	envFile := pickEnvironmentFile(env)

	godotenv.Load(envFile)

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

	if action == "up" {
		m.Up()
	} else if action == "down" {
		m.Down()
	}
}

func pickEnvironmentFile(env string) string {
	switch env {
	case "development":
		return ".env.development"
	case "production":
		return ".env.production"
	default:
		return ".env.production"
	}
}
