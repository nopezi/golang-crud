package riskindicator

import (
	"riskmanagement/lib"
	models "riskmanagement/models/riskindicator"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type RiskIndicatorDefinition interface {
	GetAll() (responses []models.RiskIndicatorResponse, err error)
	GetOne(id int64) (responses models.RiskIndicatorResponse, err error)
	Store(request *models.RiskIndicatorRequest) (responses bool, err error)
	Update(request *models.RiskIndicatorRequest) (responses bool, err error)
	Delete(id int64) (err error)
	WithTrx(trxHandle *gorm.DB) RiskIndicatorRepository
}

type RiskIndicatorRepository struct {
	db      lib.Database
	dbRaw   lib.Database
	logger  logger.Logger
	timeout time.Duration
}

// Delete implements RiskIndicatorDefinition
func (riskIndicator RiskIndicatorRepository) Delete(id int64) (err error) {
	return riskIndicator.db.DB.Where("id = ?", id).Delete(&models.RiskIndicatorResponse{}).Error
}

// GetAll implements RiskIndicatorDefinition
func (riskIndicator RiskIndicatorRepository) GetAll() (responses []models.RiskIndicatorResponse, err error) {
	return responses, riskIndicator.db.DB.Find(&responses).Error
}

// GetOne implements RiskIndicatorDefinition
func (riskIndicator RiskIndicatorRepository) GetOne(id int64) (responses models.RiskIndicatorResponse, err error) {
	return responses, riskIndicator.db.DB.Where("id = ?", id).Find(&responses).Error
}

// Store implements RiskIndicatorDefinition
func (riskIndicator RiskIndicatorRepository) Store(request *models.RiskIndicatorRequest) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")
	return responses, riskIndicator.db.DB.Save(&models.RiskIndicatorRequest{
		IndicatorCode: request.IndicatorCode,
		Name:          request.Name,
		CreatedAt:     &timeNow,
	}).Error
}

// Update implements RiskIndicatorDefinition
func (riskIndicator RiskIndicatorRepository) Update(request *models.RiskIndicatorRequest) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")
	return true, riskIndicator.db.DB.Save(&models.RiskIndicatorRequest{
		ID:            request.ID,
		IndicatorCode: request.IndicatorCode,
		Name:          request.Name,
		CreatedAt:     request.CreatedAt,
		UpdatedAt:     &timeNow,
	}).Error
}

// WithTrx implements RiskIndicatorDefinition
func (riskIndicator RiskIndicatorRepository) WithTrx(trxHandle *gorm.DB) RiskIndicatorRepository {
	if trxHandle == nil {
		riskIndicator.logger.Zap.Error("transaction Database not found in gin context")
		return riskIndicator
	}

	riskIndicator.db.DB = trxHandle
	return riskIndicator
}

func NewRiskIndicatorRepository(
	db lib.Database,
	dbRaw lib.Database,
	logger logger.Logger,
) RiskIndicatorDefinition {
	return RiskIndicatorRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: time.Second * 100,
	}
}
