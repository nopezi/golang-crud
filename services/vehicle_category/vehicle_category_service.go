package repoVehicleCategory

import (
	"fmt"
	models "infolelang/models/vehicle_category"
	repository "infolelang/repository/vehicle_category"

	"gitlab.com/golang-package-library/logger"
)

type VehicleCategoryDefinition interface {
	GetAll() (responses []models.VehicleCategoryResponse, err error)
	GetOne(id int64) (responses models.VehicleCategoryResponse, err error)
	Store(request *models.VehicleCategoryRequest) (err error)
	Update(request *models.VehicleCategoryRequest) (err error)
	Delete(id int64) (err error)
}
type VehicleCategoryService struct {
	logger     logger.Logger
	repository repository.VehicleCategoryDefinition
}

func NewVehicleCategoryService(logger logger.Logger, repository repository.VehicleCategoryDefinition) VehicleCategoryDefinition {
	return VehicleCategoryService{
		logger:     logger,
		repository: repository,
	}
}

// GetAll implements VehicleCategoryDefinition
func (VehicleCategory VehicleCategoryService) GetAll() (responses []models.VehicleCategoryResponse, err error) {
	return VehicleCategory.repository.GetAll()
}

// GetOne implements VehicleCategoryDefinition
func (VehicleCategory VehicleCategoryService) GetOne(id int64) (responses models.VehicleCategoryResponse, err error) {
	return VehicleCategory.repository.GetOne(id)
}

// Store implements VehicleCategoryDefinition
func (VehicleCategory VehicleCategoryService) Store(request *models.VehicleCategoryRequest) (err error) {
	fmt.Println("service =", request)
	_, err = VehicleCategory.repository.Store(request)
	return err
}

// Update implements VehicleCategoryDefinition
func (VehicleCategory VehicleCategoryService) Update(request *models.VehicleCategoryRequest) (err error) {
	_, err = VehicleCategory.repository.Update(request)
	return err
}

// Delete implements VehicleCategoryDefinition
func (VehicleCategory VehicleCategoryService) Delete(id int64) (err error) {
	return VehicleCategory.repository.Delete(id)
}
