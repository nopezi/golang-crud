package access_places

import (
	"infolelang/lib"
	models "infolelang/models/kpknl"
	"time"

	"gorm.io/gorm"
)

type KpknlDefinition interface {
	GetAll() (responses []models.KpknlResponse, err error)
	GetOne(id int64) (responses models.KpknlResponse, err error)
	Store(request *models.KpknlRequest) (responses bool, err error)
	Update(request *models.KpknlRequest) (responses bool, err error)
	Delete(id int64) (err error)
	WithTrx(trxHandle *gorm.DB) KpknlRepository
}
type KpknlRepository struct {
	db      lib.Database
	db2     lib.Databases
	elastic lib.Elasticsearch
	logger  lib.Logger
	timeout time.Duration
}

func NewKpknlReporitory(
	db lib.Database,
	db2 lib.Databases,
	elastic lib.Elasticsearch,
	logger lib.Logger) KpknlDefinition {
	return KpknlRepository{
		db:      db,
		db2:     db2,
		elastic: elastic,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// WithTrx implements KpknlDefinition
func (kpknl KpknlRepository) WithTrx(trxHandle *gorm.DB) KpknlRepository {
	if trxHandle == nil {
		kpknl.logger.Zap.Error("transaction Database not found in gin context. ")
		return kpknl
	}
	kpknl.db.DB = trxHandle
	return kpknl
}

// GetAll implements KpknlDefinition
func (kpknl KpknlRepository) GetAll() (responses []models.KpknlResponse, err error) {
	return responses, kpknl.db.DB.Find(&responses).Error
}

// GetOne implements KpknlDefinition
func (kpknl KpknlRepository) GetOne(id int64) (responses models.KpknlResponse, err error) {
	return responses, kpknl.db.DB.Where("id = ?", id).Find(&responses).Error
}

// Store implements KpknlDefinition
func (kpknl KpknlRepository) Store(request *models.KpknlRequest) (responses bool, err error) {
	return responses, kpknl.db.DB.Save(&responses).Error
}

// Update implements KpknlDefinition
func (kpknl KpknlRepository) Update(request *models.KpknlRequest) (responses bool, err error) {
	return true, kpknl.db.DB.Save(&responses).Error
}

// Delete implements KpknlDefinition
func (kpknl KpknlRepository) Delete(id int64) (err error) {
	return kpknl.db.DB.Where("id = ?", id).Delete(&models.KpknlResponse{}).Error
}
