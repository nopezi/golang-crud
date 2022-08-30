package coaching

import (
	"riskmanagement/lib"
	models "riskmanagement/models/coaching"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type CoachingDefinition interface {
	WithTrx(trxHandle *gorm.DB) CoachingRepository
	GetAll() (responses []models.CoachingResponse, err error)
	GetOne(id int64) (responses models.CoachingResponse, err error)
	Store(request *models.Coaching, tx *gorm.DB) (responses *models.Coaching, err error)
	Delete(request *models.CoachingUpdateDelete, include []string, tx *gorm.DB) (responses bool, err error)
	DeleteCoachingActivity(id int64, tx *gorm.DB) (err error)
	UpdateAllCoaching(request *models.CoachingUpdateActivity, include []string, tx *gorm.DB) (responses bool, err error)
}

type CoachingRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

func NewCoachingRepository(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger,
) CoachingDefinition {
	return CoachingRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// Delete implements CoachingDefinition
func (coaching CoachingRepository) Delete(request *models.CoachingUpdateDelete, include []string, tx *gorm.DB) (responses bool, err error) {
	return true, tx.Save(&request).Error
}

// DeleteCoachingActivity implements CoachingDefinition
func (coaching CoachingRepository) DeleteCoachingActivity(id int64, tx *gorm.DB) (err error) {
	return tx.Where("id = ?", id).Delete(&models.CoachingActivityRequest{}).Error
}

// GetAll implements CoachingDefinition
func (coaching CoachingRepository) GetAll() (responses []models.CoachingResponse, err error) {
	return responses, coaching.db.DB.Find(&responses).Error
}

// GetOne implements CoachingDefinition
func (coaching CoachingRepository) GetOne(id int64) (responses models.CoachingResponse, err error) {
	err = coaching.db.DB.Raw(`
	SELECT
		coa.id,
		coa.no_pelaporan,
		coa.unit_kerja,
		coa.peserta,
		coa.jumlah_peserta,
		coa.activity_id,
		coa.sub_activity_id,
		coa.maker_id,
		coa.maker_desc,
		coa.maker_date,
		coa.last_maker_id,
		coa.last_maker_desc,
		coa.last_maker_date,
		coa.status,
		coa.action,
		coa.deleted,
		coa.created_at,
		coa.updated_at
	FROM coaching coa
	WHERE coa.id = ?`, id).Find(&responses).Error

	if err != nil {
		coaching.logger.Zap.Error(err)
		return responses, err
	}
	return responses, err
}

// Store implements CoachingDefinition
func (coaching CoachingRepository) Store(request *models.Coaching, tx *gorm.DB) (responses *models.Coaching, err error) {
	return request, tx.Save(&request).Error
}

// UpdateAllCoaching implements CoachingDefinition
func (coaching CoachingRepository) UpdateAllCoaching(request *models.CoachingUpdateActivity, include []string, tx *gorm.DB) (responses bool, err error) {
	return true, tx.Save(&request).Error
}

// WithTrx implements CoachingDefinition
func (coaching CoachingRepository) WithTrx(trxHandle *gorm.DB) CoachingRepository {
	if trxHandle == nil {
		coaching.logger.Zap.Error("transaction Database not found in gin context")
		return coaching
	}
	coaching.db.DB = trxHandle
	return coaching
}
