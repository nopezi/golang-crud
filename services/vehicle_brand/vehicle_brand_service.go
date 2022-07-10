package repoVehicleBrand

import (
	"fmt"
	models "infolelang/models/vehicle_brand"
	repository "infolelang/repository/vehicle_brand"

	"gitlab.com/golang-package-library/logger"
)

type VehicleBrandDefinition interface {
	GetAll() (responses []models.VehicleBrandResponse, err error)
	GetOne(id int64) (responses models.VehicleBrandResponse, err error)
	Store(request *models.VehicleBrandRequest) (err error)
	Update(request *models.VehicleBrandRequest) (err error)
	Delete(id int64) (err error)
}
type VehicleBrandService struct {
	logger     logger.Logger
	repository repository.VehicleBrandDefinition
}

func NewVehicleBrandService(logger logger.Logger, repository repository.VehicleBrandDefinition) VehicleBrandDefinition {
	return VehicleBrandService{
		logger:     logger,
		repository: repository,
	}
}

// GetAll implements VehicleBrandDefinition
func (VehicleBrand VehicleBrandService) GetAll() (responses []models.VehicleBrandResponse, err error) {
	return VehicleBrand.repository.GetAll()
}

// GetOne implements VehicleBrandDefinition
func (VehicleBrand VehicleBrandService) GetOne(id int64) (responses models.VehicleBrandResponse, err error) {
	return VehicleBrand.repository.GetOne(id)
}

// Store implements VehicleBrandDefinition
func (VehicleBrand VehicleBrandService) Store(request *models.VehicleBrandRequest) (err error) {
	fmt.Println("service =", request)
	_, err = VehicleBrand.repository.Store(request)
	return err
}

// Update implements VehicleBrandDefinition
func (VehicleBrand VehicleBrandService) Update(request *models.VehicleBrandRequest) (err error) {
	_, err = VehicleBrand.repository.Update(request)
	return err
}

// Delete implements VehicleBrandDefinition
func (VehicleBrand VehicleBrandService) Delete(id int64) (err error) {
	return VehicleBrand.repository.Delete(id)
}
