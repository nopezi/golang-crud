package asset

import (
	"infolelang/lib"
	models "infolelang/models/assets"
	"time"

	elastic "gitlab.com/golang-package-library/elasticsearch"
	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type AssetImageDefinition interface {
	GetAll() (responses []models.AssetImagesResponse, err error)
	GetOne(id int64) (responses models.AssetImagesResponse, err error)
	Store(request *models.AssetImages, tx *gorm.DB) (responses *models.AssetImages, err error)
	Update(request *models.AssetImagesRequest) (responses bool, err error)
	Delete(id int64) (err error)
	DeleteAssetID(id int64, tx *gorm.DB) (err error)
	WithTrx(trxHandle *gorm.DB) AssetImageRepository
}
type AssetImageRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	elastic elastic.Elasticsearch
	logger  logger.Logger
	timeout time.Duration
}

func NewAssetImageReporitory(
	db lib.Database,
	dbRaw lib.Databases,
	elastic elastic.Elasticsearch,
	logger logger.Logger) AssetImageDefinition {
	return AssetImageRepository{
		db:      db,
		dbRaw:   dbRaw,
		elastic: elastic,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// WithTrx implements AssetImageDefinition
func (AssetImage AssetImageRepository) WithTrx(trxHandle *gorm.DB) AssetImageRepository {
	if trxHandle == nil {
		AssetImage.logger.Zap.Error("transaction Database not found in gin context. ")
		return AssetImage
	}
	AssetImage.db.DB = trxHandle
	return AssetImage
}

// GetAll implements AssetImageDefinition
func (AssetImage AssetImageRepository) GetAll() (responses []models.AssetImagesResponse, err error) {
	return responses, AssetImage.db.DB.Find(&responses).Error
}

// GetOne implements AssetImageDefinition
func (AssetImage AssetImageRepository) GetOne(id int64) (responses models.AssetImagesResponse, err error) {
	return responses, AssetImage.db.DB.Where("id = ?", id).Find(&responses).Error
}

// Store implements AssetImageDefinition
func (AssetImage AssetImageRepository) Store(request *models.AssetImages, tx *gorm.DB) (responses *models.AssetImages, err error) {
	return request, tx.Save(&request).Error
}

// Update implements AssetImageDefinition
func (AssetImage AssetImageRepository) Update(request *models.AssetImagesRequest) (responses bool, err error) {
	return true, AssetImage.db.DB.Save(&request).Error
}

// Delete implements AssetImageDefinition
func (AssetImage AssetImageRepository) Delete(id int64) (err error) {
	return AssetImage.db.DB.Where("id = ?", id).Delete(&models.AssetImagesResponse{}).Error
}

// DeleteAssetID implements AssetImageDefinition
func (AssetImage AssetImageRepository) DeleteAssetID(id int64, tx *gorm.DB) (err error) {
	return tx.Where("asset_id = ?", id).Delete(&models.AssetImagesResponse{}).Error
}
