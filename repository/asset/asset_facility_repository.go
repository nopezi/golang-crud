package asset

import (
	"infolelang/lib"
	models "infolelang/models/assets"
	"time"

	elastic "gitlab.com/golang-package-library/elasticsearch"
	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type AssetFacilityDefinition interface {
	GetAll() (responses []models.AssetFacilitiesResponse, err error)
	GetOne(id int64) (responses models.AssetFacilitiesResponse, err error)
	Store(request *models.AssetFacilities) (responses *models.AssetFacilities, err error)
	Update(request *models.AssetFacilitiesRequest) (responses bool, err error)
	Delete(id int64) (err error)
	WithTrx(trxHandle *gorm.DB) AssetFacilityRepository
}
type AssetFacilityRepository struct {
	db      lib.Database
	db2     lib.Databases
	elastic elastic.Elasticsearch
	logger  logger.Logger
	timeout time.Duration
}

func NewAssetFacilityReporitory(
	db lib.Database,
	db2 lib.Databases,
	elastic elastic.Elasticsearch,
	logger logger.Logger) AssetFacilityDefinition {
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
func (AssetFacility AssetFacilityRepository) Store(request *models.AssetFacilities) (responses *models.AssetFacilities, err error) {
	return request, AssetFacility.db.DB.Save(&request).Error
}

// Update implements AssetFacilityDefinition
func (AssetFacility AssetFacilityRepository) Update(request *models.AssetFacilitiesRequest) (responses bool, err error) {
	return true, AssetFacility.db.DB.Save(&request).Error
}

// Delete implements AssetFacilityDefinition
func (AssetFacility AssetFacilityRepository) Delete(id int64) (err error) {
	return AssetFacility.db.DB.Where("id = ?", id).Delete(&models.AssetFacilitiesResponse{}).Error
}
