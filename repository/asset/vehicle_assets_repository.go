package asset

import (
	"infolelang/lib"
	models "infolelang/models/vehicle_assets"
	"time"

	"gorm.io/gorm"
)

type VehicleAssetDefinition interface {
	GetAll() (responses []models.VehicleAssetsResponse, err error)
	GetOne(id int64) (responses models.VehicleAssetsResponse, err error)
	Store(request *models.VehicleAssets) (responses *models.VehicleAssets, err error)
	Update(request *models.VehicleAssetsRequest) (responses bool, err error)
	Delete(id int64) (err error)
	WithTrx(trxHandle *gorm.DB) VehicleAssetRepository
}
type VehicleAssetRepository struct {
	db      lib.Database
	db2     lib.Databases
	elastic lib.Elasticsearch
	logger  lib.Logger
	timeout time.Duration
}

func NewVehicleAssetReporitory(
	db lib.Database,
	db2 lib.Databases,
	elastic lib.Elasticsearch,
	logger lib.Logger) VehicleAssetDefinition {
	return VehicleAssetRepository{
		db:      db,
		db2:     db2,
		elastic: elastic,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// WithTrx implements VehicleAssetDefinition
func (vehicleAsset VehicleAssetRepository) WithTrx(trxHandle *gorm.DB) VehicleAssetRepository {
	if trxHandle == nil {
		vehicleAsset.logger.Zap.Error("transaction Database not found in gin context. ")
		return vehicleAsset
	}
	vehicleAsset.db.DB = trxHandle
	return vehicleAsset
}

// GetAll implements VehicleAssetDefinition
func (vehicleAsset VehicleAssetRepository) GetAll() (responses []models.VehicleAssetsResponse, err error) {
	return responses, vehicleAsset.db.DB.Find(&responses).Error
}

// GetOne implements VehicleAssetDefinition
func (vehicleAsset VehicleAssetRepository) GetOne(id int64) (responses models.VehicleAssetsResponse, err error) {
	return responses, vehicleAsset.db.DB.Where("id = ?", id).Find(&responses).Error
}

// Store implements VehicleAssetDefinition
func (vehicleAsset VehicleAssetRepository) Store(request *models.VehicleAssets) (responses *models.VehicleAssets, err error) {
	return request, vehicleAsset.db.DB.Save(&request).Error
}

// Update implements VehicleAssetDefinition
func (vehicleAsset VehicleAssetRepository) Update(request *models.VehicleAssetsRequest) (responses bool, err error) {
	return true, vehicleAsset.db.DB.Save(&responses).Error
}

// Delete implements VehicleAssetDefinition
func (vehicleAsset VehicleAssetRepository) Delete(id int64) (err error) {
	return vehicleAsset.db.DB.Where("id = ?", id).Delete(&models.VehicleAssetsResponse{}).Error
}
