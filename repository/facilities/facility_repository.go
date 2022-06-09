package facilities

import (
	"infolelang/lib"
	models "infolelang/models/facilities"
	"time"

	elastic "gitlab.com/golang-package-library/elasticsearch"
	"gorm.io/gorm"
)

type FacilitiesDefinition interface {
	GetAll() (responses []models.FacilitiesResponse, err error)
	GetOne(id int64) (responses models.FacilitiesResponse, err error)
	Store(request *models.FacilitiesRequest) (responses bool, err error)
	Update(request *models.FacilitiesRequest) (responses bool, err error)
	Delete(id int64) (err error)
	WithTrx(trxHandle *gorm.DB) FacilitiesRepository
}
type FacilitiesRepository struct {
	db      lib.Database
	db2     lib.Databases
	elastic elastic.Elasticsearch
	logger  lib.Logger
	timeout time.Duration
}

func NewFacilitiesReporitory(
	db lib.Database,
	db2 lib.Databases,
	elastic elastic.Elasticsearch,
	logger lib.Logger) FacilitiesDefinition {
	return FacilitiesRepository{
		db:      db,
		db2:     db2,
		elastic: elastic,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// WithTrx implements FacilitiesDefinition
func (Facilities FacilitiesRepository) WithTrx(trxHandle *gorm.DB) FacilitiesRepository {
	if trxHandle == nil {
		Facilities.logger.Zap.Error("transaction Database not found in gin context. ")
		return Facilities
	}
	Facilities.db.DB = trxHandle
	return Facilities
}

// GetAll implements FacilitiesDefinition
func (Facilities FacilitiesRepository) GetAll() (responses []models.FacilitiesResponse, err error) {
	return responses, Facilities.db.DB.Find(&responses).Error
}

// GetOne implements FacilitiesDefinition
func (Facilities FacilitiesRepository) GetOne(id int64) (responses models.FacilitiesResponse, err error) {
	return responses, Facilities.db.DB.Where("id = ?", id).Find(&responses).Error
}

// Store implements FacilitiesDefinition
func (Facilities FacilitiesRepository) Store(request *models.FacilitiesRequest) (responses bool, err error) {
	return responses, Facilities.db.DB.Save(&responses).Error
}

// Update implements FacilitiesDefinition
func (Facilities FacilitiesRepository) Update(request *models.FacilitiesRequest) (responses bool, err error) {
	return true, Facilities.db.DB.Save(&responses).Error
}

// Delete implements FacilitiesDefinition
func (Facilities FacilitiesRepository) Delete(id int64) (err error) {
	return Facilities.db.DB.Where("id = ?", id).Delete(&models.FacilitiesResponse{}).Error
}
