package repoKpknl

import (
	models "infolelang/models/kpknl"
	repository "infolelang/repository/kpknl"

	"gitlab.com/golang-package-library/logger"
)

type KpknlDefinition interface {
	GetAll() (responses []models.KpknlResponse, err error)
	GetOne(id int64) (responses models.KpknlResponse, err error)
	Store(request *models.KpknlRequest) (status bool, err error)
	Update(request *models.KpknlRequest) (status bool, err error)
	Delete(id int64) (err error)
}
type KpknlService struct {
	logger     logger.Logger
	repository repository.KpknlDefinition
}

func NewKpknlService(logger logger.Logger, repository repository.KpknlDefinition) KpknlDefinition {
	return KpknlService{
		logger:     logger,
		repository: repository,
	}
}

// GetAll implements KpknlDefinition
func (kpknl KpknlService) GetAll() (responses []models.KpknlResponse, err error) {
	return kpknl.repository.GetAll()
}

// GetOne implements KpknlDefinition
func (kpknl KpknlService) GetOne(id int64) (responses models.KpknlResponse, err error) {
	return kpknl.repository.GetOne(id)
}

// Store implements KpknlDefinition
func (kpknl KpknlService) Store(request *models.KpknlRequest) (status bool, err error) {
	status, err = kpknl.repository.Store(request)
	return status, err
}

// Update implements KpknlDefinition
func (kpknl KpknlService) Update(request *models.KpknlRequest) (status bool, err error) {
	status, err = kpknl.repository.Update(request)
	return status, err
}

// Delete implements KpknlDefinition
func (kpknl KpknlService) Delete(id int64) (err error) {
	return kpknl.repository.Delete(id)
}
