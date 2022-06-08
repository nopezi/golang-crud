package asset

import (
	"infolelang/lib"
	models "infolelang/models/assets"
	"time"

	"gorm.io/gorm"
)

type AssetAccessPlaceDefinition interface {
	GetAll() (responses []models.AssetAccessPlacesResponse, err error)
	GetOne(id int64) (responses models.AssetAccessPlacesResponse, err error)
	Store(request *models.AssetAccessPlaces) (responses bool, err error)
	Update(request *models.AssetAccessPlacesRequest) (responses bool, err error)
	Delete(id int64) (err error)
	WithTrx(trxHandle *gorm.DB) AssetAccessPlaceRepository
}
type AssetAccessPlaceRepository struct {
	db      lib.Database
	db2     lib.Databases
	elastic lib.Elasticsearch
	logger  lib.Logger
	timeout time.Duration
}

func NewAssetAccessPlaceReporitory(
	db lib.Database,
	db2 lib.Databases,
	elastic lib.Elasticsearch,
	logger lib.Logger) AssetAccessPlaceDefinition {
	return AssetAccessPlaceRepository{
		db:      db,
		db2:     db2,
		elastic: elastic,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// WithTrx implements AssetAccessPlaceDefinition
func (assetAccessPlace AssetAccessPlaceRepository) WithTrx(trxHandle *gorm.DB) AssetAccessPlaceRepository {
	if trxHandle == nil {
		assetAccessPlace.logger.Zap.Error("transaction Database not found in gin context. ")
		return assetAccessPlace
	}
	assetAccessPlace.db.DB = trxHandle
	return assetAccessPlace
}

// GetAll implements AssetAccessPlaceDefinition
func (assetAccessPlace AssetAccessPlaceRepository) GetAll() (responses []models.AssetAccessPlacesResponse, err error) {
	return responses, assetAccessPlace.db.DB.Find(&responses).Error
}

// GetOne implements AssetAccessPlaceDefinition
func (assetAccessPlace AssetAccessPlaceRepository) GetOne(id int64) (responses models.AssetAccessPlacesResponse, err error) {
	return responses, assetAccessPlace.db.DB.Where("id = ?", id).Find(&responses).Error
}

// Store implements AssetAccessPlaceDefinition
func (assetAccessPlace AssetAccessPlaceRepository) Store(request *models.AssetAccessPlaces) (responses bool, err error) {
	return responses, assetAccessPlace.db.DB.Save(&request).Error
}

// Update implements AssetAccessPlaceDefinition
func (assetAccessPlace AssetAccessPlaceRepository) Update(request *models.AssetAccessPlacesRequest) (responses bool, err error) {
	return true, assetAccessPlace.db.DB.Save(&request).Error
}

// Delete implements AssetAccessPlaceDefinition
func (assetAccessPlace AssetAccessPlaceRepository) Delete(id int64) (err error) {
	return assetAccessPlace.db.DB.Where("id = ?", id).Delete(&models.AssetAccessPlacesResponse{}).Error
}
