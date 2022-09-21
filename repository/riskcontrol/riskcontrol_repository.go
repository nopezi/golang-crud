package riskcontrol

import (
	"riskmanagement/lib"
	models "riskmanagement/models/riskcontrol"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

var (
	timeNow = lib.GetTimeNow("timestime")
)

type RiskControlDefinition interface {
	GetAll() (responses []models.RiskControlResponse, err error)
	GetOne(id int64) (responses models.RiskControlResponse, err error)
	Store(request *models.RiskControlRequest) (responses bool, err error)
	Update(request *models.RiskControlRequest) (responses bool, err error)
	Delete(id int64) (err error)
	WithTrx(trxHandle *gorm.DB) RiskControlRepository
}

type RiskControlRepository struct {
	db      lib.Database
	dbRaw   lib.Database
	logger  logger.Logger
	timeout time.Duration
}

func NewRiskControlRepository(
	db lib.Database,
	dbRaw lib.Database,
	logger logger.Logger,
) RiskControlDefinition {
	return RiskControlRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// Delete implements RiskControlDefinition
func (riskControl RiskControlRepository) Delete(id int64) (err error) {
	return riskControl.db.DB.Where("id = ?", id).Delete(&models.RiskControlResponse{}).Error
}

// GetAll implements RiskControlDefinition
func (riskControl RiskControlRepository) GetAll() (responses []models.RiskControlResponse, err error) {
	return responses, riskControl.db.DB.Find(&responses).Error
}

// GetOne implements RiskControlDefinition
func (riskControl RiskControlRepository) GetOne(id int64) (responses models.RiskControlResponse, err error) {
	return responses, riskControl.db.DB.Where("id = ?", id).Find(&responses).Error
}

// Store implements RiskControlDefinition
func (riskControl RiskControlRepository) Store(request *models.RiskControlRequest) (responses bool, err error) {
	return responses, riskControl.db.DB.Save(&models.RiskControlRequest{
		Kode:        request.Kode,
		RiskControl: request.RiskControl,
		CreatedAt:   &timeNow,
	}).Error
}

// Update implements RiskControlDefinition
func (riskControl RiskControlRepository) Update(request *models.RiskControlRequest) (responses bool, err error) {
	return responses, riskControl.db.DB.Save(&models.RiskControlRequest{
		ID:          request.ID,
		Kode:        request.Kode,
		RiskControl: request.RiskControl,
		CreatedAt:   request.CreatedAt,
		UpdatedAt:   &timeNow,
	}).Error
}

// WithTrx implements RiskControlDefinition
func (riskControl RiskControlRepository) WithTrx(trxHandle *gorm.DB) RiskControlRepository {
	if trxHandle == nil {
		riskControl.logger.Zap.Error("transaction Database not found in gin context")
		return riskControl
	}

	riskControl.db.DB = trxHandle
	return riskControl
}
