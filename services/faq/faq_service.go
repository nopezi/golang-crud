package repoFaq

import (
	"infolelang/lib"
	models "infolelang/models/faq"
	repository "infolelang/repository/faq"
)

type FaqDefinition interface {
	GetAll() (responses []models.FaqResponse, err error)
	GetOne(id int64) (responses models.FaqResponse, err error)
	Store(request *models.FaqRequest) (err error)
	Update(request *models.FaqRequest) (err error)
	Delete(id int64) (err error)
}
type FaqService struct {
	logger     lib.Logger
	repository repository.FaqDefinition
}

func NewFaqService(logger lib.Logger, repository repository.FaqDefinition) FaqDefinition {
	return FaqService{
		logger:     logger,
		repository: repository,
	}
}

// GetAll implements FaqDefinition
func (ap FaqService) GetAll() (responses []models.FaqResponse, err error) {
	return ap.repository.GetAll()
}

// GetOne implements FaqDefinition
func (ap FaqService) GetOne(id int64) (responses models.FaqResponse, err error) {
	return ap.repository.GetOne(id)
}

// Store implements FaqDefinition
func (ap FaqService) Store(request *models.FaqRequest) (err error) {
	_, err = ap.repository.Store(request)
	return err
}

// Update implements FaqDefinition
func (ap FaqService) Update(request *models.FaqRequest) (err error) {
	_, err = ap.repository.Update(request)
	return err
}

// Delete implements FaqDefinition
func (ap FaqService) Delete(id int64) (err error) {
	return ap.repository.Delete(id)
}
