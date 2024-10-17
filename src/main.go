package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/emersonnobre/tica-api-go/src/cmd/api"
	_ "github.com/emersonnobre/tica-api-go/src/docs"
	"github.com/emersonnobre/tica-api-go/src/internal/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @host localhost:3000
// @BasePath /
// @schemes http https
func main() {
	var env string
	if len(os.Args) > 1 {
		env = os.Args[1]
	}
	envFile := pickEnvironmentFile(env)
	godotenv.Load(envFile)

	mysqlConnection := database.NewMySQLDatabase()
	connection := activateDBConnection(&mysqlConnection)
	defer connection.Close()

	app := fiber.New()
	api.SetupAPI(app, connection)
	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Listen(":3000")
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

func activateDBConnection(db *database.MySQLDatabase) *sql.DB {
	connection, err := db.Connect()
	if err != nil {
		log.Fatal("Error connecting database!", err)
	}
	if err = connection.Ping(); err != nil {
		log.Fatal(err)
	}
	return connection
}
