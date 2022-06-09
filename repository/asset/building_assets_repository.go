package asset

import (
	"infolelang/lib"
	models "infolelang/models/building_assets"
	"time"

	elastic "gitlab.com/golang-package-library/elasticsearch"
	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type BuildingAssetDefinition interface {
	GetAll() (responses []models.BuildingAssetsResponse, err error)
	GetOne(id int64) (responses models.BuildingAssetsResponse, err error)
	Store(request *models.BuildingAssets) (responses *models.BuildingAssets, err error)
	Update(request *models.BuildingAssetsRequest) (responses bool, err error)
	Delete(id int64) (err error)
	WithTrx(trxHandle *gorm.DB) BuildingAssetRepository
}
type BuildingAssetRepository struct {
	db      lib.Database
	db2     lib.Databases
	elastic elastic.Elasticsearch
	logger  logger.Logger
	timeout time.Duration
}

func NewBuildingAssetReporitory(
	db lib.Database,
	db2 lib.Databases,
	elastic elastic.Elasticsearch,
	logger logger.Logger) BuildingAssetDefinition {
	return BuildingAssetRepository{
		db:      db,
		db2:     db2,
		elastic: elastic,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// WithTrx implements BuildingAssetDefinition
func (buildingAsset BuildingAssetRepository) WithTrx(trxHandle *gorm.DB) BuildingAssetRepository {
	if trxHandle == nil {
		buildingAsset.logger.Zap.Error("transaction Database not found in gin context. ")
		return buildingAsset
	}
	buildingAsset.db.DB = trxHandle
	return buildingAsset
}

// GetAll implements BuildingAssetDefinition
func (buildingAsset BuildingAssetRepository) GetAll() (responses []models.BuildingAssetsResponse, err error) {
	return responses, buildingAsset.db.DB.Find(&responses).Error
}

// GetOne implements BuildingAssetDefinition
func (buildingAsset BuildingAssetRepository) GetOne(id int64) (responses models.BuildingAssetsResponse, err error) {
	return responses, buildingAsset.db.DB.Where("id = ?", id).Find(&responses).Error
}

// Store implements BuildingAssetDefinition
func (buildingAsset BuildingAssetRepository) Store(request *models.BuildingAssets) (responses *models.BuildingAssets, err error) {
	return request, buildingAsset.db.DB.Save(&request).Error
}

// Update implements BuildingAssetDefinition
func (buildingAsset BuildingAssetRepository) Update(request *models.BuildingAssetsRequest) (responses bool, err error) {
	return true, buildingAsset.db.DB.Save(&responses).Error
}

// Delete implements BuildingAssetDefinition
func (buildingAsset BuildingAssetRepository) Delete(id int64) (err error) {
	return buildingAsset.db.DB.Where("id = ?", id).Delete(&models.BuildingAssetsResponse{}).Error
}
