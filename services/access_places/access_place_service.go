package repoAccessPlace

import (
	"infolelang/lib"
	models "infolelang/models/access_places"
	repository "infolelang/repository/access_places"
)

type AccessPlaceDefinition interface {
	GetAll() (responses []models.AccessPlacesResponse, err error)
	GetOne(id int64) (responses models.AccessPlacesResponse, err error)
	Store(request *models.AccessPlacesRequest) (err error)
	Update(request *models.AccessPlacesRequest) (err error)
	Delete(id int64) (err error)
}
type AccessPlaceService struct {
	logger     lib.Logger
	repository repository.AccessPlaceDefinition
}

func NewAccessPlaceService(logger lib.Logger, repository repository.AccessPlaceDefinition) AccessPlaceDefinition {
	return AccessPlaceService{
		logger:     logger,
		repository: repository,
	}
}

// GetAll implements AccessPlaceDefinition
func (ap AccessPlaceService) GetAll() (responses []models.AccessPlacesResponse, err error) {
	return ap.repository.GetAll()
}

// GetOne implements AccessPlaceDefinition
func (ap AccessPlaceService) GetOne(id int64) (responses models.AccessPlacesResponse, err error) {
	return ap.repository.GetOne(id)
}

// Store implements AccessPlaceDefinition
func (ap AccessPlaceService) Store(request *models.AccessPlacesRequest) (err error) {
	_, err = ap.repository.Store(request)
	return err
}

// Update implements AccessPlaceDefinition
func (ap AccessPlaceService) Update(request *models.AccessPlacesRequest) (err error) {
	_, err = ap.repository.Update(request)
	return err
}

// Delete implements AccessPlaceDefinition
func (ap AccessPlaceService) Delete(id int64) (err error) {
	return ap.repository.Delete(id)
}
