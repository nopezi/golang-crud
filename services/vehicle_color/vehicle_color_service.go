package repoVehicleColor

import (
	"fmt"
	models "infolelang/models/vehicle_color"
	repository "infolelang/repository/vehicle_color"

	"gitlab.com/golang-package-library/logger"
)

type VehicleColorDefinition interface {
	GetAll() (responses []models.VehicleColorResponse, err error)
	GetOne(id int64) (responses models.VehicleColorResponse, err error)
	Store(request *models.VehicleColorRequest) (err error)
	Update(request *models.VehicleColorRequest) (err error)
	Delete(id int64) (err error)
}
type VehicleColorService struct {
	logger     logger.Logger
	repository repository.VehicleColorDefinition
}

func NewVehicleColorService(logger logger.Logger, repository repository.VehicleColorDefinition) VehicleColorDefinition {
	return VehicleColorService{
		logger:     logger,
		repository: repository,
	}
}

// GetAll implements VehicleColorDefinition
func (VehicleColor VehicleColorService) GetAll() (responses []models.VehicleColorResponse, err error) {
	return VehicleColor.repository.GetAll()
}

// GetOne implements VehicleColorDefinition
func (VehicleColor VehicleColorService) GetOne(id int64) (responses models.VehicleColorResponse, err error) {
	return VehicleColor.repository.GetOne(id)
}

// Store implements VehicleColorDefinition
func (VehicleColor VehicleColorService) Store(request *models.VehicleColorRequest) (err error) {
	fmt.Println("service =", request)
	_, err = VehicleColor.repository.Store(request)
	return err
}

// Update implements VehicleColorDefinition
func (VehicleColor VehicleColorService) Update(request *models.VehicleColorRequest) (err error) {
	_, err = VehicleColor.repository.Update(request)
	return err
}

// Delete implements VehicleColorDefinition
func (VehicleColor VehicleColorService) Delete(id int64) (err error) {
	return VehicleColor.repository.Delete(id)
}
