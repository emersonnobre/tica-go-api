package usecases

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/emersonnobre/tica-api-go/src/internal/core/repositories"
	"github.com/emersonnobre/tica-api-go/src/internal/core/usecases/types"
	"github.com/emersonnobre/tica-api-go/src/internal/core/usecases/types/requests"
)

type CreateProductUseCase struct {
	repository         repositories.ProductRepository
	categoryRepository repositories.CategoryRepository
	employeeRepository repositories.EmployeeRepository
}

func NewCreateProductUseCase(
	repository repositories.ProductRepository,
	categoryRepository repositories.CategoryRepository,
	employeeRepository repositories.EmployeeRepository,
) *CreateProductUseCase {
	return &CreateProductUseCase{
		repository:         repository,
		categoryRepository: categoryRepository,
		employeeRepository: employeeRepository,
	}
}

func (u *CreateProductUseCase) Execute(product requests.CreateProductRequest) types.UseCaseResponse {
	if validationErrors := product.ValidateObjectData(); validationErrors != nil {
		return types.NewErrorUseCaseResponse(types.GetValidationErrorName(), strings.Join(*validationErrors, "\n"))
	}

	if err := u.validateDatabaseRestrictions(product); err != nil {
		return types.NewErrorUseCaseResponse(types.GetValidationErrorName(), *err)
	}

	target := *product.MapForDomain()
	target.Active = true
	target.CreatedAt = time.Now()

	if err := u.repository.Create(target); err != nil {
		log.Println("ERROR:", err)
		return types.NewErrorUseCaseResponse(types.GetInternalErrorName(), "Erro ao criar produto!")
	}

	return types.NewUseCaseResponse(nil, nil, nil)
}

func (u *CreateProductUseCase) validateDatabaseRestrictions(product requests.CreateProductRequest) *string {
	if nameError := u.validateNameInDB(product.Name); nameError != nil {
		return nameError
	}

	if invalidForeignKeys := u.validateForeignKeys(product); invalidForeignKeys != nil {
		errorMessage := strings.Join(*invalidForeignKeys, "\n")
		return &errorMessage
	}

	return nil
}

func (u *CreateProductUseCase) validateNameInDB(name string) *string {
	filters := []repositories.Filter{*repositories.NewFilter("name", name, "string", false)}

	count, _ := u.repository.GetCount(filters)
	if invalidName := count > 0; invalidName {
		message := "Já existe um produto com este nome!"
		return &message
	}
	return nil
}

func (u *CreateProductUseCase) validateForeignKeys(product requests.CreateProductRequest) *[]string {
	errors := []string{}

	filters := []repositories.Filter{
		*repositories.NewFilter("id", strconv.Itoa(product.CategoryId), "integer", false),
	}

	count, _ := u.categoryRepository.GetCount(filters)
	if invalidCategory := count == 0; invalidCategory {
		errors = append(errors, "Categoria inexistente!")
	}

	filters = []repositories.Filter{
		*repositories.NewFilter("id", strconv.Itoa(product.CreatedBy), "integer", false),
	}

	count, _ = u.employeeRepository.GetCount(filters)
	if invalidEmployee := count == 0; invalidEmployee {
		errors = append(errors, "Funcionário inexistente!")
	}

	if len(errors) > 0 {
		return &errors
	}

	return nil
}
