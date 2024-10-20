package usecases

import (
	"log"
	"strconv"
	"strings"

	"github.com/emersonnobre/tica-api-go/src/internal/core/domain"
	"github.com/emersonnobre/tica-api-go/src/internal/core/repositories"
	"github.com/emersonnobre/tica-api-go/src/internal/core/usecases/types"
	"github.com/emersonnobre/tica-api-go/src/internal/core/usecases/types/requests"
)

type UpdateProductUseCase struct {
	repository         repositories.ProductRepository
	categoryRepository repositories.CategoryRepository
	employeeRepository repositories.EmployeeRepository
}

func NewUpdateProductUseCase(
	repository repositories.ProductRepository,
	categoryRepository repositories.CategoryRepository,
	employeeRepository repositories.EmployeeRepository,
) *UpdateProductUseCase {
	return &UpdateProductUseCase{
		repository:         repository,
		categoryRepository: categoryRepository,
		employeeRepository: employeeRepository,
	}
}

func (u *UpdateProductUseCase) Execute(productRequest requests.UpdateProductRequest) types.UseCaseResponse {
	if validationErrors := productRequest.ValidateObjectData(); validationErrors != nil {
		return types.NewErrorUseCaseResponse(types.GetValidationErrorName(), strings.Join(*validationErrors, "\n"))
	}

	product, _ := u.repository.GetById(productRequest.Id)
	if product == nil {
		return types.NewErrorUseCaseResponse(types.GetNotFoundErrorName(), "Produto não encontrado!")
	}

	if err := u.validateDatabaseRestrictions(productRequest, *product); err != nil {
		return types.NewErrorUseCaseResponse(types.GetValidationErrorName(), *err)
	}

	product = productRequest.MapForDomain()
	if err := u.repository.Update(product); err != nil {
		log.Println("ERROR:", err)
		return types.NewErrorUseCaseResponse(types.GetInternalErrorName(), "Erro ao atualizar produto!")
	}

	return types.NewUseCaseResponse(nil, nil, nil)
}

func (u *UpdateProductUseCase) validateDatabaseRestrictions(
	productRequest requests.UpdateProductRequest,
	productDB domain.Product,
) *string {
	if productRequest.Name != productDB.Name {
		if nameError := u.validateNameInDB(productRequest.Name); nameError != nil {
			return nameError
		}
	}

	if invalidForeignKeys := u.validateForeignKeys(productRequest); invalidForeignKeys != nil {
		errorMessage := strings.Join(*invalidForeignKeys, "\n")
		return &errorMessage
	}

	return nil
}

func (u *UpdateProductUseCase) validateNameInDB(name string) *string {
	filters := []repositories.Filter{*repositories.NewFilter("name", name, true, false)}

	count, _ := u.repository.GetCount(filters)
	if invalidName := count > 0; invalidName {
		message := "Já existe um produto com este nome!"
		return &message
	}
	return nil
}

func (u *UpdateProductUseCase) validateForeignKeys(product requests.UpdateProductRequest) *[]string {
	errors := []string{}

	filters := []repositories.Filter{
		*repositories.NewFilter("id", strconv.Itoa(product.CategoryId), false, false),
	}

	count, _ := u.categoryRepository.GetCount(filters)
	if invalidCategory := count == 0; invalidCategory {
		errors = append(errors, "Categoria inexistente!")
	}

	filters = []repositories.Filter{
		*repositories.NewFilter("id", strconv.Itoa(product.UpdatedBy), false, false),
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
