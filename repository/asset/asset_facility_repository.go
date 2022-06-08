package asset

import (
	"infolelang/lib"
	models "infolelang/models/assets"
	"time"

	"gorm.io/gorm"
)

type AssetFacilityDefinition interface {
	GetAll() (responses []models.AssetFacilitiesResponse, err error)
	GetOne(id int64) (responses models.AssetFacilitiesResponse, err error)
	Store(request *models.AssetFacilities) (responses bool, err error)
	Update(request *models.AssetFacilitiesRequest) (responses bool, err error)
	Delete(id int64) (err error)
	WithTrx(trxHandle *gorm.DB) AssetFacilityRepository
}
type AssetFacilityRepository struct {
	db      lib.Database
	db2     lib.Databases
	elastic lib.Elasticsearch
	logger  lib.Logger
	timeout time.Duration
}

func NewAssetFacilityReporitory(
	db lib.Database,
	db2 lib.Databases,
	elastic lib.Elasticsearch,
	logger lib.Logger) AssetFacilityDefinition {
	return AssetFacilityRepository{
		db:      db,
		db2:     db2,
		elastic: elastic,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// WithTrx implements AssetFacilityDefinition
func (AssetFacility AssetFacilityRepository) WithTrx(trxHandle *gorm.DB) AssetFacilityRepository {
	if trxHandle == nil {
		AssetFacility.logger.Zap.Error("transaction Database not found in gin context. ")
		return AssetFacility
	}
	AssetFacility.db.DB = trxHandle
	return AssetFacility
}

// GetAll implements AssetFacilityDefinition
func (AssetFacility AssetFacilityRepository) GetAll() (responses []models.AssetFacilitiesResponse, err error) {
	return responses, AssetFacility.db.DB.Find(&responses).Error
}

// GetOne implements AssetFacilityDefinition
func (AssetFacility AssetFacilityRepository) GetOne(id int64) (responses models.AssetFacilitiesResponse, err error) {
	return responses, AssetFacility.db.DB.Where("id = ?", id).Find(&responses).Error
}

// Store implements AssetFacilityDefinition
func (AssetFacility AssetFacilityRepository) Store(request *models.AssetFacilities) (responses bool, err error) {
	return responses, AssetFacility.db.DB.Save(&request).Error
}

// Update implements AssetFacilityDefinition
func (AssetFacility AssetFacilityRepository) Update(request *models.AssetFacilitiesRequest) (responses bool, err error) {
	return true, AssetFacility.db.DB.Save(&request).Error
}

// Delete implements AssetFacilityDefinition
func (AssetFacility AssetFacilityRepository) Delete(id int64) (err error) {
	return AssetFacility.db.DB.Where("id = ?", id).Delete(&models.AssetFacilitiesResponse{}).Error
}
