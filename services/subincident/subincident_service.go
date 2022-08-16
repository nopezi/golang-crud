package subincident

import (
	models "riskmanagement/models/subincident"
	repository "riskmanagement/repository/subincident"

	"gitlab.com/golang-package-library/logger"
)

type SubIncidentDefinition interface {
	GetAll() (responses []models.SubIncidentResponse, err error)
	GetOne(id int64) (responses models.SubIncidentResponse, err error)
	Store(request *models.SubIncidentRequest) (err error)
	Update(request *models.SubIncidentRequest) (err error)
	Delete(id int64) (err error)
}

type SubIncidentService struct {
	logger     logger.Logger
	repository repository.SubIncidentDefinition
}

func NewSubIncidentService(logger logger.Logger, repository repository.SubIncidentDefinition) SubIncidentDefinition {
	return SubIncidentService{
		logger:     logger,
		repository: repository,
	}
}

// Delete implements SubIncidentDefinition
func (subIncident SubIncidentService) Delete(id int64) (err error) {
	return subIncident.repository.Delete(id)
}

// GetAll implements SubIncidentDefinition
func (subIncident SubIncidentService) GetAll() (responses []models.SubIncidentResponse, err error) {
	return subIncident.repository.GetAll()
}

// GetOne implements SubIncidentDefinition
func (subIncident SubIncidentService) GetOne(id int64) (responses models.SubIncidentResponse, err error) {
	return subIncident.repository.GetOne(id)
}

// Store implements SubIncidentDefinition
func (subIncident SubIncidentService) Store(request *models.SubIncidentRequest) (err error) {
	_, err = subIncident.repository.Store(request)
	return err
}

// Update implements SubIncidentDefinition
func (subIncident SubIncidentService) Update(request *models.SubIncidentRequest) (err error) {
	_, err = subIncident.repository.Update(request)
	return err
}
