package facility

import (
	models "infolelang/models/facilities"
	repository "infolelang/repository/facilities"

	"gitlab.com/golang-package-library/logger"
)

type FacilitiesDefinition interface {
	GetAll() (responses []models.FacilitiesResponse, err error)
	GetOne(id int64) (responses models.FacilitiesResponse, err error)
	Store(request *models.FacilitiesRequest) (status bool, err error)
	Update(request *models.FacilitiesRequest) (status bool, err error)
	Delete(id int64) (err error)
}
type FacilityService struct {
	logger     logger.Logger
	repository repository.FacilitiesDefinition
}

func NewFacilityService(logger logger.Logger, repository repository.FacilitiesDefinition) FacilitiesDefinition {
	return FacilityService{
		logger:     logger,
		repository: repository,
	}
}

// GetAll implements FacilitiesDefinition
func (Facility FacilityService) GetAll() (responses []models.FacilitiesResponse, err error) {
	return Facility.repository.GetAll()
}

// GetOne implements FacilitiesDefinition
func (Facility FacilityService) GetOne(id int64) (responses models.FacilitiesResponse, err error) {
	return Facility.repository.GetOne(id)
}

// Store implements FacilitiesDefinition
func (Facility FacilityService) Store(request *models.FacilitiesRequest) (status bool, err error) {
	status, err = Facility.repository.Store(request)
	return status, err
}

// Update implements FacilitiesDefinition
func (Facility FacilityService) Update(request *models.FacilitiesRequest) (status bool, err error) {
	status, err = Facility.repository.Update(request)
	return status, err
}

// Delete implements FacilitiesDefinition
func (Facility FacilityService) Delete(id int64) (err error) {
	return Facility.repository.Delete(id)
}
