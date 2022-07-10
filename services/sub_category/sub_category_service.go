package repoSubCategory

import (
	models "infolelang/models/sub_categories"
	repository "infolelang/repository/sub_categories"

	"gitlab.com/golang-package-library/logger"
)

type SubCategoryDefinition interface {
	GetAll(category_id int) (responses []models.SubCategoriesResponse, err error)
	GetOne(id int64) (responses models.SubCategoriesResponse, err error)
	Store(request *models.SubCategoriesRequest) (err error)
	Update(request *models.SubCategoriesRequest) (err error)
	Delete(id int64) (err error)
}
type SubCategoryService struct {
	logger     logger.Logger
	repository repository.SubCategoryDefinition
}

func NewSubCategoryService(logger logger.Logger, repository repository.SubCategoryDefinition) SubCategoryDefinition {
	return SubCategoryService{
		logger:     logger,
		repository: repository,
	}
}

// GetAll implements SubCategoryDefinition
func (subCategory SubCategoryService) GetAll(category_id int) (responses []models.SubCategoriesResponse, err error) {
	return subCategory.repository.GetAll(category_id)
}

// GetOne implements SubCategoryDefinition
func (subCategory SubCategoryService) GetOne(id int64) (responses models.SubCategoriesResponse, err error) {
	return subCategory.repository.GetOne(id)
}

// Store implements SubCategoryDefinition
func (subCategory SubCategoryService) Store(request *models.SubCategoriesRequest) (err error) {
	_, err = subCategory.repository.Store(request)
	return err
}

// Update implements SubCategoryDefinition
func (subCategory SubCategoryService) Update(request *models.SubCategoriesRequest) (err error) {
	_, err = subCategory.repository.Update(request)
	return err
}

// Delete implements SubCategoryDefinition
func (subCategory SubCategoryService) Delete(id int64) (err error) {
	return subCategory.repository.Delete(id)
}
