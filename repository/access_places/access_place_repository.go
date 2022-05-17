package repository

import (
	"database/sql"
	"infolelang/lib"
	models "infolelang/models/access_places"
	"time"

	"gorm.io/gorm"
)

type AccessPlaceDefinition interface {
	GetAll() (responses []models.AccessPlacesResponse, err error)
	GetOne(id int64) (responses models.AccessPlacesResponse, err error)
	Store(request *models.AccessPlacesRequest) (responses bool, err error)
	Update(request *models.AccessPlacesRequest) (responses bool, err error)
	Delete(id int64) (err error)
	WithTrx(trxHandle *gorm.DB) AccessPlaceRepository
}
type AccessPlaceRepository struct {
	db      lib.Database
	dbRaw   *sql.DB
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
func (ap AccessPlaceRepository) GetOne(id int64) (responses models.AccessPlacesResponse, err error) {
	return responses, ap.db.DB.Where("id = ?", id).Find(&responses).Error
}

// Store implements AccessPlaceDefinition
func (ap AccessPlaceRepository) Store(request *models.AccessPlacesRequest) (responses bool, err error) {
	return responses, ap.db.DB.Save(&responses).Error
}

// Update implements AccessPlaceDefinition
func (ap AccessPlaceRepository) Update(request *models.AccessPlacesRequest) (responses bool, err error) {
	return true, ap.db.DB.Save(&responses).Error
}

// Delete implements AccessPlaceDefinition
func (ap AccessPlaceRepository) Delete(id int64) (err error) {
	return ap.db.DB.Where("id = ?", id).Delete(&models.AccessPlacesResponse{}).Error
}
