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
	categoryHandler.RegisterRoutes(app)

	// employee dependencies
	employeeRepository := mysql_repository.NewMySQLEmployeeRepository(connection)
	createEmployeeUseCase := usecases.NewCreateEmployeeUseCase(employeeRepository)
	getEmployeeUseCase := usecases.NewGetEmployeeUseCase(employeeRepository)
	employeeHandler := handlers.NewEmployeeHandler(createEmployeeUseCase, getEmployeeUseCase)
	employeeHandler.RegisterRoutes(app)

	// customer dependencies
	customerRepository := mysql_repository.NewMySQLCustomerRepository(connection)
	createCustomerUseCase := usecases.NewCreateCustomerUseCase(customerRepository)
	getCustomerUseCase := usecases.NewGetCustomerUseCase(customerRepository)
	customerHandler := handlers.NewCustomerHandler(createCustomerUseCase, getCustomerUseCase)
	customerHandler.RegisterRoutes(app)

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
