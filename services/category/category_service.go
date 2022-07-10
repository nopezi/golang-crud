package category

import (
	"fmt"
	models "infolelang/models/categories"
	repository "infolelang/repository/categories"

	"gitlab.com/golang-package-library/logger"
)

type CategoryDefinition interface {
	GetAll() (responses []models.CategoryResponse, err error)
	GetOne(id int64) (responses models.CategoryResponse, err error)
	Store(request *models.CategoryRequest) (err error)
	Update(request *models.CategoryRequest) (err error)
	Delete(id int64) (err error)
}
type CategoryService struct {
	logger     logger.Logger
	repository repository.CategoryDefinition
}

func NewCategoryService(logger logger.Logger, repository repository.CategoryDefinition) CategoryDefinition {
	return CategoryService{
		logger:     logger,
		repository: repository,
	}
}

// GetAll implements CategoryDefinition
func (category CategoryService) GetAll() (responses []models.CategoryResponse, err error) {
	return category.repository.GetAll()
}

// GetOne implements CategoryDefinition
func (category CategoryService) GetOne(id int64) (responses models.CategoryResponse, err error) {
	return category.repository.GetOne(id)
}

// Store implements CategoryDefinition
func (category CategoryService) Store(request *models.CategoryRequest) (err error) {
	fmt.Println("service =", request)
	_, err = category.repository.Store(request)
	return err
}

// Update implements CategoryDefinition
func (category CategoryService) Update(request *models.CategoryRequest) (err error) {
	_, err = category.repository.Update(request)
	return err
}

// Delete implements CategoryDefinition
func (category CategoryService) Delete(id int64) (err error) {
	return category.repository.Delete(id)
}
