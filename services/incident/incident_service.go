package incident

import (
	"fmt"
	models "riskmanagement/models/incident"
	repository "riskmanagement/repository/incident"

	"gitlab.com/golang-package-library/logger"
)

type IncidentDefinition interface {
	GetAll() (responses []models.IncidentResponse, err error)
	GetOne(id int64) (responses models.IncidentResponse, err error)
	Store(request *models.IncidentRequest) (err error)
	Update(request *models.IncidentRequest) (err error)
	Delete(id int64) (err error)
}

type IncidentService struct {
	logger     logger.Logger
	repository repository.IncidentDefinition
}

// Delete implements IncidentDefinition
func (incident IncidentService) Delete(id int64) (err error) {
	return incident.repository.Delete(id)
}

// GetAll implements IncidentDefinition
func (incident IncidentService) GetAll() (responses []models.IncidentResponse, err error) {
	return incident.repository.GetAll()
}

// GetOne implements IncidentDefinition
func (incident IncidentService) GetOne(id int64) (responses models.IncidentResponse, err error) {
	return incident.repository.GetOne(id)
}

// Store implements IncidentDefinition
func (incident IncidentService) Store(request *models.IncidentRequest) (err error) {
	fmt.Println("service = ", request)
	_, err = incident.repository.Store(request)
	return err
}

// Update implements IncidentDefinition
func (incident IncidentService) Update(request *models.IncidentRequest) (err error) {
	_, err = incident.repository.Update(request)
	return err
}

func NewIncidentService(
	logger logger.Logger,
	repository repository.IncidentDefinition,
) IncidentDefinition {
	return IncidentService{
		logger:     logger,
		repository: repository,
	}
}
