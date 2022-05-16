package repository

import (
	"eform-gateway/lib"
	models "eform-gateway/models/access_places"
	"time"

	"gorm.io/gorm"
)

type AccessPlaceDefinition interface {
	GetAll() (responses []models.AccessPlacesResponse, err error)
	GetOne(id models.AccessPlacesRequest) (responses models.AccessPlacesResponse, err error)
	Store(request models.AccessPlacesRequest) (responses bool, err error)
	Update(request models.AccessPlacesRequest) (responses bool, err error)
	Delete(id models.AccessPlacesRequest) (responses bool, err error)
	WithTrx(trxHandle *gorm.DB) AccessPlaceRepository
}
type AccessPlaceRepository struct {
	db      lib.Database
	elastic lib.Elasticsearch
	logger  lib.Logger
	timeout time.Duration
}

func NewAccessPlaceReporitory(
	db lib.Database,
	elastic lib.Elasticsearch,
	logger lib.Logger) AccessPlaceDefinition {
	return AccessPlaceRepository{
		db:      db,
		elastic: elastic,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// WithTrx implements AccessPlaceDefinition
func (ap AccessPlaceRepository) WithTrx(trxHandle *gorm.DB) AccessPlaceRepository {
	if trxHandle == nil {
		ap.logger.Zap.Error("transaction Database not found in gin context. ")
		return ap
	}
	ap.db.DB = trxHandle
	return ap
}

// GetAll implements AccessPlaceDefinition
func (ap AccessPlaceRepository) GetAll() (responses []models.AccessPlacesResponse, err error) {
	return responses, ap.db.DB.Find(&responses).Error
}

// GetOne implements AccessPlaceDefinition
func (ap AccessPlaceRepository) GetOne(id models.AccessPlacesRequest) (responses models.AccessPlacesResponse, err error) {
	return responses, ap.db.DB.Where("id = ?", id.ID).Find(&responses).Error
}

// Store implements AccessPlaceDefinition
func (ap AccessPlaceRepository) Store(request models.AccessPlacesRequest) (responses bool, err error) {
	return responses, ap.db.DB.Save(&responses).Error
}

// Update implements AccessPlaceDefinition
func (ap AccessPlaceRepository) Update(request models.AccessPlacesRequest) (responses bool, err error) {
	return true, ap.db.DB.Save(&responses).Error
}

// Delete implements AccessPlaceDefinition
func (ap AccessPlaceRepository) Delete(id models.AccessPlacesRequest) (responses bool, err error) {
	return true, ap.db.DB.Where("id = ?", id.ID).Delete(&models.AccessPlacesResponse{}).Error
}
