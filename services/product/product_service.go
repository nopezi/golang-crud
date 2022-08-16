package product

import (
	"fmt"
	models "riskmanagement/models/product"
	repository "riskmanagement/repository/product"

	"gitlab.com/golang-package-library/logger"
)

type ProductDefinition interface {
	GetAll() (responses []models.ProductResponse, err error)
	GetOne(id int64) (responses models.ProductResponse, err error)
	Store(request *models.ProductRequest) (err error)
	Update(request *models.ProductRequest) (err error)
	Delete(id int64) (err error)
}

type ProductService struct {
	logger     logger.Logger
	repository repository.ProductDefinition
}

func NewProductService(
	logger logger.Logger,
	repository repository.ProductDefinition,
) ProductDefinition {
	return ProductService{
		logger:     logger,
		repository: repository,
	}
}

// Delete implements ProductDefinition
func (product ProductService) Delete(id int64) (err error) {
	return product.repository.Delete(id)
}

// GetAll implements ProductDefinition
func (product ProductService) GetAll() (responses []models.ProductResponse, err error) {
	return product.repository.GetAll()
}

// GetOne implements ProductDefinition
func (product ProductService) GetOne(id int64) (responses models.ProductResponse, err error) {
	return product.repository.GetOne(id)
}

// Store implements ProductDefinition
func (product ProductService) Store(request *models.ProductRequest) (err error) {
	fmt.Println("service =", request)
	_, err = product.repository.Store(request)
	return err
}

// Update implements ProductDefinition
func (product ProductService) Update(request *models.ProductRequest) (err error) {
	_, err = product.repository.Update(request)
	return err
}
