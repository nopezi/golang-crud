package activity

import (
	"fmt"
	models "riskmanagement/models/activity"
	repository "riskmanagement/repository/activity"

	"gitlab.com/golang-package-library/logger"
)

type ActivityDefinition interface {
	GetAll() (responses []models.ActivityResponse, err error)
	GetOne(id int64) (responses models.ActivityResponse, err error)
	Store(request *models.ActivityRequest) (err error)
	Update(request *models.ActivityRequest) (err error)
	Delete(id int64) (err error)
}

type ActivityService struct {
	logger     logger.Logger
	repository repository.ActivityDefinition
}

func NewActivityService(logger logger.Logger, repository repository.ActivityDefinition) ActivityDefinition {
	return ActivityService{
		logger:     logger,
		repository: repository,
	}
}

// Delete implements ActivityDefinition
func (activity ActivityService) Delete(id int64) (err error) {
	return activity.repository.Delete(id)
}

// GetAll implements ActivityDefinition
func (activity ActivityService) GetAll() (responses []models.ActivityResponse, err error) {
	return activity.repository.GetAll()
}

// GetOne implements ActivityDefinition
func (activity ActivityService) GetOne(id int64) (responses models.ActivityResponse, err error) {
	return activity.repository.GetOne(id)
}

// Store implements ActivityDefinition
func (activity ActivityService) Store(request *models.ActivityRequest) (err error) {
	fmt.Println("service =", request)
	_, err = activity.repository.Store(request)
	return err
}

// Update implements ActivityDefinition
func (activity ActivityService) Update(request *models.ActivityRequest) (err error) {
	_, err = activity.repository.Update(request)
	return err
}
