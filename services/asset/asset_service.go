package repoAsset

import (
	"infolelang/lib"
	models "infolelang/models/assets"
	repository "infolelang/repository/asset"
)

type AssetDefinition interface {
	GetAll() (responses []models.AssetsResponse, err error)
	GetOne(id int64) (responses models.AssetsResponse, err error)
	Store(request *models.AssetsRequest) (err error)
	Update(request *models.AssetsRequest) (err error)
	Delete(id int64) (err error)
}
type AssetService struct {
	logger     lib.Logger
	repository repository.AssetDefinition
}

func NewAssetService(logger lib.Logger, repository repository.AssetDefinition) AssetDefinition {
	return AssetService{
		logger:     logger,
		repository: repository,
	}
}

// GetAll implements AssetDefinition
func (asset AssetService) GetAll() (responses []models.AssetsResponse, err error) {
	return asset.repository.GetAll()
}

// GetOne implements AssetDefinition
func (asset AssetService) GetOne(id int64) (responses models.AssetsResponse, err error) {
	return asset.repository.GetOne(id)
}

// Store implements AssetDefinition
func (asset AssetService) Store(request *models.AssetsRequest) (err error) {
	_, err = asset.repository.Store(request)
	return err
}

// Update implements AssetDefinition
func (asset AssetService) Update(request *models.AssetsRequest) (err error) {
	_, err = asset.repository.Update(request)
	return err
}

// Delete implements AssetDefinition
func (asset AssetService) Delete(id int64) (err error) {
	return asset.repository.Delete(id)
}
