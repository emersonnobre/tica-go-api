package api

import (
	"database/sql"

	"github.com/emersonnobre/tica-api-go/src/internal/core/usecases"
	"github.com/emersonnobre/tica-api-go/src/internal/delivery/fiber/handlers"
	mysql_repository "github.com/emersonnobre/tica-api-go/src/internal/repositories/mysql"
	"github.com/gofiber/fiber/v2"
)

func SetupAPI(app *fiber.App, connection *sql.DB) {
	setupDependencies(app, connection)
}

func setupDependencies(app *fiber.App, connection *sql.DB) {
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
	getCustomersUseCase := usecases.NewGetCustomersUseCase(customerRepository)
	updateCustomerUseCase := usecases.NewUpdateCustomerUseCase(customerRepository, createAddressUseCase, removeAddressUseCase)
	removeCustomerUseCase := usecases.NewRemoveCustomerUseCase(customerRepository)
	customerHandler := handlers.NewCustomerHandler(createCustomerUseCase, getCustomerUseCase, updateCustomerUseCase, removeCustomerUseCase, getCustomersUseCase)
	customerHandler.RegisterRoutes(app)
}
