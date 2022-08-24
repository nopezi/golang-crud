package briefing

import (
	"riskmanagement/lib"
	models "riskmanagement/models/briefing"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type BriefingMateriDefinition interface {
	GetAll() (responses []models.BriefingMateriResponse, err error)
	GetOne(id int64) (responses models.BriefingMateriResponse, err error)
	Store(request *models.BriefingMateri, tx *gorm.DB) (responses *models.BriefingMateri, err error)
	Update(request *models.BriefingMateriRequest) (responses bool, err error)
	Delete(id int64) (err error)
	DeleteBriefingID(id int64, tx *gorm.DB) (err error)
	WithTrx(trxHandle *gorm.DB) BriefingMateriRepository
}

type BriefingMateriRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

func NewBriefingMateriRepository(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger,
) BriefingMateriDefinition {
	return BriefingMateriRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// Delete implements BriefingMateriDefinition
func (BriefingMateri BriefingMateriRepository) Delete(id int64) (err error) {
	return BriefingMateri.db.DB.Where("id = ?", id).Delete(&models.BriefingMateriResponse{}).Error
}

// DeleteBriefingID implements BriefingMateriDefinition
func (BriefingMateri BriefingMateriRepository) DeleteBriefingID(id int64, tx *gorm.DB) (err error) {
	return tx.Where("briefing_id = ?", id).Delete(&models.BriefingMateriResponse{}).Error
}

// GetAll implements BriefingMateriDefinition
func (BriefingMateri BriefingMateriRepository) GetAll() (responses []models.BriefingMateriResponse, err error) {
	return responses, BriefingMateri.db.DB.Find(&responses).Error
}

// GetOne implements BriefingMateriDefinition
func (BriefingMateri BriefingMateriRepository) GetOne(id int64) (responses models.BriefingMateriResponse, err error) {
	return responses, BriefingMateri.db.DB.Where("id = ?", id).Find(&responses).Error
}

// Store implements BriefingMateriDefinition
func (BriefingMateri BriefingMateriRepository) Store(request *models.BriefingMateri, tx *gorm.DB) (responses *models.BriefingMateri, err error) {
	return request, tx.Save(&request).Error
}

// Update implements BriefingMateriDefinition
func (BriefingMateri BriefingMateriRepository) Update(request *models.BriefingMateriRequest) (responses bool, err error) {
	return true, BriefingMateri.db.DB.Save(&request).Error
}

// WithTrx implements BriefingMateriDefinition
func (BriefingMateri BriefingMateriRepository) WithTrx(trxHandle *gorm.DB) BriefingMateriRepository {
	if trxHandle == nil {
		BriefingMateri.logger.Zap.Error("transacton Database not found in gin context.")
		return BriefingMateri
	}
	BriefingMateri.db.DB = trxHandle
	return BriefingMateri
}
