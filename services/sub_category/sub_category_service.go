package repoSubCategory

import (
	"infolelang/lib"
	models "infolelang/models/sub_categories"
	repository "infolelang/repository/sub_categories"
)

type SubCategoryDefinition interface {
	GetAll() (responses []models.SubCategoriesResponse, err error)
	GetOne(id int64) (responses models.SubCategoriesResponse, err error)
	Store(request *models.SubCategoriesRequest) (err error)
	Update(request *models.SubCategoriesRequest) (err error)
	Delete(id int64) (err error)
}
type SubCategoryService struct {
	logger     lib.Logger
	repository repository.SubCategoryDefinition
}

func NewSubCategoryService(logger lib.Logger, repository repository.SubCategoryDefinition) SubCategoryDefinition {
	return SubCategoryService{
		logger:     logger,
		repository: repository,
	}
}

// GetAll implements SubCategoryDefinition
func (subCategory SubCategoryService) GetAll() (responses []models.SubCategoriesResponse, err error) {
	return subCategory.repository.GetAll()
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
