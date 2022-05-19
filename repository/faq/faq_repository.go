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
func (faq FaqRepository) WithTrx(trxHandle *gorm.DB) FaqRepository {
	if trxHandle == nil {
		faq.logger.Zap.Error("transaction Database not found in gin context. ")
		return faq
	}
	faq.db.DB = trxHandle
	return faq
}

// GetAll implements FaqDefinition
func (faq FaqRepository) GetAll() (responses []models.FaqResponse, err error) {
	return responses, faq.db.DB.Find(&responses).Error
}

// GetOne implements FaqDefinition
func (faq FaqRepository) GetOne(id int64) (responses models.FaqResponse, err error) {
	return responses, faq.db.DB.Where("id = ?", id).Find(&responses).Error
}

// Store implements FaqDefinition
func (faq FaqRepository) Store(request *models.FaqRequest) (responses bool, err error) {
	return responses, faq.db.DB.Save(&responses).Error
}

// Update implements FaqDefinition
func (faq FaqRepository) Update(request *models.FaqRequest) (responses bool, err error) {
	return true, faq.db.DB.Save(&responses).Error
}

// Delete implements FaqDefinition
func (faq FaqRepository) Delete(id int64) (err error) {
	return faq.db.DB.Where("id = ?", id).Delete(&models.FaqResponse{}).Error
}
