package subactivity

import (
	models "riskmanagement/models/subactivity"
	repository "riskmanagement/repository/subactivity"

	"gitlab.com/golang-package-library/logger"
)

type SubActivityDefinition interface {
	GetAll() (responses []models.SubActivityResponse, err error)
	GetOne(id int64) (responses models.SubActivityResponse, err error)
	Store(request *models.SubActivityRequest) (err error)
	Update(request *models.SubActivityRequest) (err error)
	Delete(id int64) (err error)
}

type SubActivityService struct {
	logger     logger.Logger
	repository repository.SubActivityDefinition
}

func NewSubActivityService(logger logger.Logger, repository repository.SubActivityDefinition) SubActivityDefinition {
	return SubActivityService{
		logger:     logger,
		repository: repository,
	}
}

// Delete implements SubAvtivityDefinition
func (subactivity SubActivityService) Delete(id int64) (err error) {
	return subactivity.repository.Delete(id)
}

// GetAll implements SubAvtivityDefinition
func (subactivity SubActivityService) GetAll() (responses []models.SubActivityResponse, err error) {
	return subactivity.repository.GetAll()
}

// GetOne implements SubAvtivityDefinition
func (subactivity SubActivityService) GetOne(id int64) (responses models.SubActivityResponse, err error) {
	return subactivity.repository.GetOne(id)
}

// Store implements SubAvtivityDefinition
func (subactivity SubActivityService) Store(request *models.SubActivityRequest) (err error) {
	_, err = subactivity.repository.Store(request)
	return err
}

// Update implements SubAvtivityDefinition
func (subactivity SubActivityService) Update(request *models.SubActivityRequest) (err error) {
	_, err = subactivity.repository.Update(request)
	return err
}
