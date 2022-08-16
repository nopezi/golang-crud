package subactivity

import (
	"riskmanagement/lib"
	models "riskmanagement/models/subactivity"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type SubActivityDefinition interface {
	GetAll() (responses []models.SubActivityResponse, err error)
	GetOne(id int64) (responses models.SubActivityResponse, err error)
	Store(request *models.SubActivityRequest) (responses bool, err error)
	Update(request *models.SubActivityRequest) (responses bool, err error)
	Delete(id int64) (err error)
	WithTrx(trxHandle *gorm.DB) SubActivityRepository
}

type SubActivityRepository struct {
	db      lib.Database
	dbRaw   lib.Database
	logger  logger.Logger
	timeout time.Duration
}

func NewSubActivityRepository(db lib.Database, dbRaw lib.Database, logger logger.Logger) SubActivityDefinition {
	return SubActivityRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// Delete implements SubActivityDefinition
func (subactivity SubActivityRepository) Delete(id int64) (err error) {
	return subactivity.db.DB.Where("id = ?", id).Delete(&models.SubActivityResponse{}).Error
}

// GetAll implements SubActivityDefinition
func (subactivity SubActivityRepository) GetAll() (responses []models.SubActivityResponse, err error) {
	return responses, subactivity.db.DB.Find(&responses).Error
}

// GetOne implements SubActivityDefinition
func (subactivity SubActivityRepository) GetOne(id int64) (responses models.SubActivityResponse, err error) {
	return responses, subactivity.db.DB.Where("id = ?", id).Find(&responses).Error
}

// Store implements SubActivityDefinition
func (subactivity SubActivityRepository) Store(request *models.SubActivityRequest) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")
	return responses, subactivity.db.DB.Save(&models.SubActivityRequest{
		ActivityID: request.ActivityID,
		Name:       request.Name,
		CreatedAt:  &timeNow,
	}).Error
}

// Update implements SubActivityDefinition
func (subactivity SubActivityRepository) Update(request *models.SubActivityRequest) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")
	return true, subactivity.db.DB.Save(&models.SubActivityRequest{
		ID:         request.ID,
		ActivityID: request.ActivityID,
		Name:       request.Name,
		CreatedAt:  request.CreatedAt,
		UpdatedAt:  &timeNow,
	}).Error
}

// WithTrx implements SubActivityDefinition
func (subactivity SubActivityRepository) WithTrx(trxHandle *gorm.DB) SubActivityRepository {
	if trxHandle == nil {
		subactivity.logger.Zap.Error("transaction Database not found in gin context")
		return subactivity
	}

	subactivity.db.DB = trxHandle
	return subactivity
}
