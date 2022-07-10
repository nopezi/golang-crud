package repoVehicleCapacity

import (
	"fmt"
	models "infolelang/models/vehicle_capacity"
	repository "infolelang/repository/vehicle_capacity"

	"gitlab.com/golang-package-library/logger"
)

type VehicleCapacityDefinition interface {
	GetAll() (responses []models.VehicleCapacityResponse, err error)
	GetOne(id int64) (responses models.VehicleCapacityResponse, err error)
	Store(request *models.VehicleCapacityRequest) (err error)
	Update(request *models.VehicleCapacityRequest) (err error)
	Delete(id int64) (err error)
}
type VehicleCapacityService struct {
	logger     logger.Logger
	repository repository.VehicleCapacityDefinition
}

func NewVehicleCapacityService(logger logger.Logger, repository repository.VehicleCapacityDefinition) VehicleCapacityDefinition {
	return VehicleCapacityService{
		logger:     logger,
		repository: repository,
	}
}

// GetAll implements VehicleCapacityDefinition
func (VehicleCapacity VehicleCapacityService) GetAll() (responses []models.VehicleCapacityResponse, err error) {
	return VehicleCapacity.repository.GetAll()
}

// GetOne implements VehicleCapacityDefinition
func (VehicleCapacity VehicleCapacityService) GetOne(id int64) (responses models.VehicleCapacityResponse, err error) {
	return VehicleCapacity.repository.GetOne(id)
}

// Store implements VehicleCapacityDefinition
func (VehicleCapacity VehicleCapacityService) Store(request *models.VehicleCapacityRequest) (err error) {
	fmt.Println("service =", request)
	_, err = VehicleCapacity.repository.Store(request)
	return err
}

// Update implements VehicleCapacityDefinition
func (VehicleCapacity VehicleCapacityService) Update(request *models.VehicleCapacityRequest) (err error) {
	_, err = VehicleCapacity.repository.Update(request)
	return err
}

// Delete implements VehicleCapacityDefinition
func (VehicleCapacity VehicleCapacityService) Delete(id int64) (err error) {
	return VehicleCapacity.repository.Delete(id)
}
