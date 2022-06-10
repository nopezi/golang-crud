package postalcode

import (
	models "infolelang/models/postalcode"
	repository "infolelang/repository/postalcode"

	"gitlab.com/golang-package-library/logger"
)

type PostalcodeDefinition interface {
	GetAll() (responses []models.PostalcodeResponse, err error)
	GetOne(id int64) (responses models.PostalcodeResponse, err error)
	FindPostalCode(id string) (responses models.PostalcodeResponse, err error)
	Store(request *models.PostalcodeRequest) (status bool, err error)
	Update(request *models.PostalcodeRequest) (status bool, err error)
	Delete(id int64) (err error)
}
type PostalcodeService struct {
	logger     logger.Logger
	repository repository.PostalcodeDefinition
}

func NewPostalcodeService(logger logger.Logger, repository repository.PostalcodeDefinition) PostalcodeDefinition {
	return PostalcodeService{
		logger:     logger,
		repository: repository,
	}
}

// GetAll implements PostalcodeDefinition
func (Postalcode PostalcodeService) GetAll() (responses []models.PostalcodeResponse, err error) {
	return Postalcode.repository.GetAll()
}

// GetOne implements PostalcodeDefinition
func (Postalcode PostalcodeService) GetOne(id int64) (responses models.PostalcodeResponse, err error) {
	return Postalcode.repository.GetOne(id)
}

// GetOne implements PostalcodeDefinition
func (Postalcode PostalcodeService) FindPostalCode(id string) (responses models.PostalcodeResponse, err error) {
	return Postalcode.repository.FindPostalCode(id)
}

// Store implements PostalcodeDefinition
func (Postalcode PostalcodeService) Store(request *models.PostalcodeRequest) (status bool, err error) {
	status, err = Postalcode.repository.Store(request)
	return status, err
}

// Update implements PostalcodeDefinition
func (Postalcode PostalcodeService) Update(request *models.PostalcodeRequest) (status bool, err error) {
	status, err = Postalcode.repository.Update(request)
	return status, err
}

// Delete implements PostalcodeDefinition
func (Postalcode PostalcodeService) Delete(id int64) (err error) {
	return Postalcode.repository.Delete(id)
}
