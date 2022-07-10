package repoVehicleTransmission

import (
	"fmt"
	models "infolelang/models/vehicle_transmission"
	repository "infolelang/repository/vehicle_transmission"

	"gitlab.com/golang-package-library/logger"
)

type VehicleTransmissionDefinition interface {
	GetAll() (responses []models.VehicleTransmissionResponse, err error)
	GetOne(id int64) (responses models.VehicleTransmissionResponse, err error)
	Store(request *models.VehicleTransmissionRequest) (err error)
	Update(request *models.VehicleTransmissionRequest) (err error)
	Delete(id int64) (err error)
}
type VehicleTransmissionService struct {
	logger     logger.Logger
	repository repository.VehicleTransmissionDefinition
}

func NewVehicleTransmissionService(logger logger.Logger, repository repository.VehicleTransmissionDefinition) VehicleTransmissionDefinition {
	return VehicleTransmissionService{
		logger:     logger,
		repository: repository,
	}
}

// GetAll implements VehicleTransmissionDefinition
func (VehicleTransmission VehicleTransmissionService) GetAll() (responses []models.VehicleTransmissionResponse, err error) {
	return VehicleTransmission.repository.GetAll()
}

// GetOne implements VehicleTransmissionDefinition
func (VehicleTransmission VehicleTransmissionService) GetOne(id int64) (responses models.VehicleTransmissionResponse, err error) {
	return VehicleTransmission.repository.GetOne(id)
}

// Store implements VehicleTransmissionDefinition
func (VehicleTransmission VehicleTransmissionService) Store(request *models.VehicleTransmissionRequest) (err error) {
	fmt.Println("service =", request)
	_, err = VehicleTransmission.repository.Store(request)
	return err
}

// Update implements VehicleTransmissionDefinition
func (VehicleTransmission VehicleTransmissionService) Update(request *models.VehicleTransmissionRequest) (err error) {
	_, err = VehicleTransmission.repository.Update(request)
	return err
}

// Delete implements VehicleTransmissionDefinition
func (VehicleTransmission VehicleTransmissionService) Delete(id int64) (err error) {
	return VehicleTransmission.repository.Delete(id)
}
