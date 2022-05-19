package access_places

import (
	"infolelang/lib"
	models "infolelang/models/faq"
	"time"

	"gorm.io/gorm"
)

type FaqDefinition interface {
	GetAll() (responses []models.FaqResponse, err error)
	GetOne(id int64) (responses models.FaqResponse, err error)
	Store(request *models.FaqRequest) (responses bool, err error)
	Update(request *models.FaqRequest) (responses bool, err error)
	Delete(id int64) (err error)
	WithTrx(trxHandle *gorm.DB) FaqRepository
}
type FaqRepository struct {
	db      lib.Database
	db2     lib.Databases
	elastic lib.Elasticsearch
	logger  lib.Logger
	timeout time.Duration
}

func NewFaqReporitory(
	db lib.Database,
	db2 lib.Databases,
	elastic lib.Elasticsearch,
	logger lib.Logger) FaqDefinition {
	return FaqRepository{
		db:      db,
		db2:     db2,
		elastic: elastic,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// WithTrx implements FaqDefinition
func (ap FaqRepository) WithTrx(trxHandle *gorm.DB) FaqRepository {
	if trxHandle == nil {
		ap.logger.Zap.Error("transaction Database not found in gin context. ")
		return ap
	}
	ap.db.DB = trxHandle
	return ap
}

// GetAll implements FaqDefinition
func (ap FaqRepository) GetAll() (responses []models.FaqResponse, err error) {
	return responses, ap.db.DB.Find(&responses).Error
}

// GetOne implements FaqDefinition
func (ap FaqRepository) GetOne(id int64) (responses models.FaqResponse, err error) {
	return responses, ap.db.DB.Where("id = ?", id).Find(&responses).Error
}

// Store implements FaqDefinition
func (ap FaqRepository) Store(request *models.FaqsRequest) (responses bool, err error) {
	return responses, ap.db.DB.Save(&responses).Error
}

// Update implements FaqDefinition
func (ap FaqRepository) Update(request *models.FaqsRequest) (responses bool, err error) {
	return true, ap.db.DB.Save(&responses).Error
}

// Delete implements FaqDefinition
func (ap FaqRepository) Delete(id int64) (err error) {
	return ap.db.DB.Where("id = ?", id).Delete(&models.FaqResponse{}).Error
}
