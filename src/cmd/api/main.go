package api

import (
	"database/sql"

	"github.com/emersonnobre/tica-api-go/src/internal/core/usecases"
	"github.com/emersonnobre/tica-api-go/src/internal/delivery/fiber/handlers"
	mysql_repository "github.com/emersonnobre/tica-api-go/src/internal/repositories/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupAPI(app *fiber.App, connection *sql.DB) {
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost,http://localhost:3000",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Accept-Encoding, X-CSRF-Token, Authorization",
		AllowCredentials: true,
		ExposeHeaders:    "Content-Length",
	}))

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

	// product dependencies
	productRepository := mysql_repository.NewMySQLProductRepository(connection)
	createProductUseCase := usecases.NewCreateProductUseCase(productRepository, categoryRepository, employeeRepository)
	updateProductUseCase := usecases.NewUpdateProductUseCase(productRepository, categoryRepository, employeeRepository)
	getProductUseCase := usecases.NewGetProductUseCase(productRepository)
	productHandler := handlers.NewProductHandler(createProductUseCase, updateProductUseCase, getProductUseCase)
	productHandler.RegisterRoutes(app)
}
