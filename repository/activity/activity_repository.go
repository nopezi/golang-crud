package activity

import (
	"fmt"
	"riskmanagement/lib"
	models "riskmanagement/models/activity"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type ActivityDefinition interface {
	GetAll() (responses []models.ActivityResponse, err error)
	GetOne(id int64) (responses models.ActivityResponse, err error)
	Store(request *models.ActivityRequest) (responses bool, err error)
	Update(request *models.ActivityRequest) (responses bool, err error)
	Delete(id int64) (err error)
	WithTrx(trxHandle *gorm.DB) ActivityRepository
}

type ActivityRepository struct {
	db      lib.Database
	dbRaw   lib.Database
	logger  logger.Logger
	timeout time.Duration
}

func NewActivityRepository(db lib.Database, dbRaw lib.Database, logger logger.Logger) ActivityDefinition {
	return ActivityRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// Delete implements ActicityDefinition
func (activity ActivityRepository) Delete(id int64) (err error) {
	return activity.db.DB.Where("id = ?", id).Delete(&models.ActivityResponse{}).Error
}

// GetAll implements ActicityDefinition
func (activity ActivityRepository) GetAll() (responses []models.ActivityResponse, err error) {
	return responses, activity.db.DB.Find(&responses).Error
}

// GetOne implements ActicityDefinition
func (activity ActivityRepository) GetOne(id int64) (responses models.ActivityResponse, err error) {
	return responses, activity.db.DB.Where("id = ?", id).Find(&responses).Error
}

// Store implements ActicityDefinition
func (activity ActivityRepository) Store(request *models.ActivityRequest) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")
	fmt.Println("repo = ", models.ActivityRequest{
		Name:         request.Name,
		KodeActivity: request.KodeActivity,
		CreateAt:     &timeNow,
	})
	err = activity.db.DB.Save(&models.ActivityRequest{
		Name:         request.Name,
		KodeActivity: request.KodeActivity,
		CreateAt:     &timeNow,
	}).Error

	fmt.Println(err)
	return true, err
}

// Update implements ActicityDefinition
func (activity ActivityRepository) Update(request *models.ActivityRequest) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")
	return true, activity.db.DB.Save(&models.ActivityRequest{
		ID:           request.ID,
		KodeActivity: request.KodeActivity,
		Name:         request.Name,
		CreateAt:     request.CreateAt,
		UpdateAt:     &timeNow,
	}).Error
}

// WithTrx implements ActicityDefinition
func (activity ActivityRepository) WithTrx(trxHandle *gorm.DB) ActivityRepository {
	if trxHandle == nil {
		activity.logger.Zap.Error("transaction Database not found in gin context")
		return activity
	}
	activity.db.DB = trxHandle
	return activity
}
