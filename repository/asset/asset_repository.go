package asset

import (
	"infolelang/lib"
	models "infolelang/models/assets"
	"time"

	elastic "gitlab.com/golang-package-library/elasticsearch"
	"gorm.io/gorm"
)

type AssetDefinition interface {
	WithTrx(trxHandle *gorm.DB) AssetRepository
	GetAll() (responses []models.AssetsResponse, err error)
	GetOne(id int64) (responses models.AssetsResponse, err error)
	Store(request *models.Assets) (responses *models.Assets, err error)
	Update(request *models.AssetsRequest) (responses bool, err error)
	Delete(id int64) (err error)
	StoreElastic(request *models.AssetsResponse) (response bool, err error)
}
type AssetRepository struct {
	db      lib.Database
	db2     lib.Databases
	elastic elastic.Elasticsearch
	logger  lib.Logger
	timeout time.Duration
}

func NewAssetReporitory(
	db lib.Database,
	db2 lib.Databases,
	elastic elastic.Elasticsearch,
	logger lib.Logger) AssetDefinition {
	return AssetRepository{
		db:      db,
		db2:     db2,
		elastic: elastic,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// WithTrx implements AssetDefinition
func (asset AssetRepository) WithTrx(trxHandle *gorm.DB) AssetRepository {
	if trxHandle == nil {
		asset.logger.Zap.Error("transaction Database not found in gin context. ")
		return asset
	}
	asset.db.DB = trxHandle
	return asset
}

// GetAll implements AssetDefinition
func (asset AssetRepository) GetAll() (responses []models.AssetsResponse, err error) {
	return responses, asset.db.DB.Find(&responses).Error
}

// GetOne implements AssetDefinition
func (asset AssetRepository) GetOne(id int64) (responses models.AssetsResponse, err error) {
	return responses, asset.db.DB.Where("id = ?", id).Find(&responses).Error
}

// Store implements AssetDefinition
func (asset AssetRepository) Store(request *models.Assets) (responses *models.Assets, err error) {
	return request, asset.db.DB.Save(&request).Error
}

// Update implements AssetDefinition
func (asset AssetRepository) Update(request *models.AssetsRequest) (responses bool, err error) {
	return true, asset.db.DB.Save(&responses).Error
}

// Delete implements AssetDefinition
func (asset AssetRepository) Delete(id int64) (err error) {
	return asset.db.DB.Where("id = ?", id).Delete(&models.AssetsResponse{}).Error
}

func (asset AssetRepository) StoreElastic(request *models.AssetsResponse) (response bool, err error) {
	return true, err
}
