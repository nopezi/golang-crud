package asset

import (
	"infolelang/lib"
	models "infolelang/models/assets"
	"time"

	"gorm.io/gorm"
)

type AssetImageDefinition interface {
	GetAll() (responses []models.AssetImagesResponse, err error)
	GetOne(id int64) (responses models.AssetImagesResponse, err error)
	Store(request *models.AssetImagesRequest) (responses bool, err error)
	Update(request *models.AssetImagesRequest) (responses bool, err error)
	Delete(id int64) (err error)
	WithTrx(trxHandle *gorm.DB) AssetImageRepository
}
type AssetImageRepository struct {
	db      lib.Database
	db2     lib.Databases
	elastic lib.Elasticsearch
	logger  lib.Logger
	timeout time.Duration
}

func NewAssetImageReporitory(
	db lib.Database,
	db2 lib.Databases,
	elastic lib.Elasticsearch,
	logger lib.Logger) AssetImageDefinition {
	return AssetImageRepository{
		db:      db,
		db2:     db2,
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
func (AssetImage AssetImageRepository) Store(request *models.AssetImagesRequest) (responses bool, err error) {
	return responses, AssetImage.db.DB.Save(&responses).Error
}

// Update implements AssetImageDefinition
func (AssetImage AssetImageRepository) Update(request *models.AssetImagesRequest) (responses bool, err error) {
	return true, AssetImage.db.DB.Save(&responses).Error
}

// Delete implements AssetImageDefinition
func (AssetImage AssetImageRepository) Delete(id int64) (err error) {
	return AssetImage.db.DB.Where("id = ?", id).Delete(&models.AssetImagesResponse{}).Error
}
