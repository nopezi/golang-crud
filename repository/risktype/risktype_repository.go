package risktype

import (
	"riskmanagement/lib"
	models "riskmanagement/models/risktype"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type RiskTypeDefinition interface {
	GetAll() (responses []models.RiskTypeResponse, err error)
	GetOne(id int64) (responses models.RiskTypeResponse, err error)
	Store(request *models.RiskTypeRequest) (responses bool, err error)
	Update(request *models.RiskTypeRequest) (responses bool, err error)
	Delete(id int64) (err error)
	WithTrx(trxHandle *gorm.DB) RiskTypeRepository
}

type RiskTypeRepository struct {
	db      lib.Database
	dbRaw   lib.Database
	logger  logger.Logger
	timeout time.Duration
}

// Delete implements RiskTypeDefinition
func (riskType RiskTypeRepository) Delete(id int64) (err error) {
	return riskType.db.DB.Where("id = ?", id).Delete(&models.RiskTypeResponse{}).Error
}

// GetAll implements RiskTypeDefinition
func (riskType RiskTypeRepository) GetAll() (responses []models.RiskTypeResponse, err error) {
	return responses, riskType.db.DB.Find(&responses).Error
}

// GetOne implements RiskTypeDefinition
func (riskType RiskTypeRepository) GetOne(id int64) (responses models.RiskTypeResponse, err error) {
	return responses, riskType.db.DB.Where("id = ?", id).Find(&responses).Error
}

// Store implements RiskTypeDefinition
func (riskType RiskTypeRepository) Store(request *models.RiskTypeRequest) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")
	return responses, riskType.db.DB.Save(&models.RiskTypeRequest{
		RiskTypeCode: request.RiskTypeCode,
		RiskType:     request.RiskType,
		CreatedAt:    &timeNow,
	}).Error
}

// Update implements RiskTypeDefinition
func (riskType RiskTypeRepository) Update(request *models.RiskTypeRequest) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")
	return true, riskType.db.DB.Save(&models.RiskTypeRequest{
		ID:           request.ID,
		RiskTypeCode: request.RiskTypeCode,
		RiskType:     request.RiskType,
		CreatedAt:    request.CreatedAt,
		UpdatedAt:    &timeNow,
	}).Error
}

// WithTrx implements RiskTypeDefinition
func (riskType RiskTypeRepository) WithTrx(trxHandle *gorm.DB) RiskTypeRepository {
	if trxHandle == nil {
		riskType.logger.Zap.Error("transaction Database not found in gin context")
		return riskType
	}

	riskType.db.DB = trxHandle
	return riskType
}

func NewRiskTypeRepository(
	db lib.Database,
	dbRaw lib.Database,
	logger logger.Logger,
) RiskTypeDefinition {
	return RiskTypeRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: time.Second * 100,
	}
}
