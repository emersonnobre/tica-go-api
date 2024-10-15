package fiberconfig

import (
	"database/sql"
	"log"

	"github.com/emersonnobre/tica-api-go/internal/core/usecases"
	"github.com/emersonnobre/tica-api-go/internal/database"
	"github.com/emersonnobre/tica-api-go/internal/delivery/fiber/handlers"
	mysql_repository "github.com/emersonnobre/tica-api-go/internal/repositories/mysql"
	"github.com/gofiber/fiber/v2"
)

type FiberSetup struct {
}

func NewFiberSetup() *FiberSetup {
	return &FiberSetup{}
}

func (f *FiberSetup) Execute() {
	app := fiber.New()

	mysqlConnection := database.NewMySQLDatabase()
	connection := activateDBConnection(&mysqlConnection)
	defer connection.Close()

	// category dependencies
	categoryRepository := mysql_repository.NewMySQLCategoryRepository(connection)
	createCategoryUseCase := usecases.NewCreateCategoryUseCase(categoryRepository)
	getCategoriesUseCase := usecases.NewGetCategoriesUseCase(categoryRepository)
	categoryHandler := handlers.NewCategoryHandler(createCategoryUseCase, getCategoriesUseCase)

	// product dependencies

	categoryHandler.RegisterRoutes(app)

	app.Listen(":3000")
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
