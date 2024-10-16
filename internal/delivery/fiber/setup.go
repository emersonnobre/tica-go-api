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

	addressRepository := mysql_repository.NewMySQLAddressRepository(connection)
	createAddressUseCase := usecases.NewCreateAddressUseCase(addressRepository)
	removeAddressUseCase := usecases.NewRemoveAddressUseCase(addressRepository)

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
	createCustomerUseCase := usecases.NewCreateCustomerUseCase(customerRepository, createAddressUseCase)
	getCustomerUseCase := usecases.NewGetCustomerUseCase(customerRepository)
	updateCustomerUseCase := usecases.NewUpdateCustomerUseCase(customerRepository, createAddressUseCase, removeAddressUseCase)
	removeCustomerUseCase := usecases.NewRemoveCustomerUseCase(customerRepository)
	customerHandler := handlers.NewCustomerHandler(createCustomerUseCase, getCustomerUseCase, updateCustomerUseCase, removeCustomerUseCase)
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
