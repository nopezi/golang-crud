package riskindicator

import (
	"riskmanagement/lib"
	models "riskmanagement/models/riskindicator"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type RiskIndicatorDefinition interface {
	// GetAll() (responses []models.RiskIndicatorResponse, err error)
	GetAll() (responses []models.RiskIndicatorResponses, err error)
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
// func (riskIndicator RiskIndicatorRepository) GetAll() (responses []models.RiskIndicatorResponse, err error) {
// 	return responses, riskIndicator.db.DB.Find(&responses).Error
// }

func (riskIndicator RiskIndicatorRepository) GetAll() (responses []models.RiskIndicatorResponses, err error) {
	rows, err := riskIndicator.db.DB.Raw(`
		SELECT
			riskInd.id 'id',
			riskInd.risk_indicator_code 'risk_indicator_code',
			riskInd.risk_indicator 'risk_indicator',
			riskInd.activity_id 'activity_id',
			act.name 'activity',
			riskInd.product_id 'product_id',
			prod.name 'product',
			riskInd.created_at 'created_at',
			riskInd.updated_at 'updated_at'
		FROM risk_indicator riskInd
		JOIN activity act on act.id = riskInd.activity_id
		JOIN product prod on prod.id = riskInd.product_id
	`).Rows()

	defer rows.Close()

	var riskInd models.RiskIndicatorResponses

	for rows.Next() {
		riskIndicator.db.DB.ScanRows(rows, &riskInd)
		responses = append(responses, riskInd)
	}

	return responses, err

}

// GetOne implements RiskIndicatorDefinition
func (riskIndicator RiskIndicatorRepository) GetOne(id int64) (responses models.RiskIndicatorResponse, err error) {
	return responses, riskIndicator.db.DB.Where("id = ?", id).Find(&responses).Error
}

// Store implements RiskIndicatorDefinition
func (riskIndicator RiskIndicatorRepository) Store(request *models.RiskIndicatorRequest) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")
	return responses, riskIndicator.db.DB.Save(&models.RiskIndicatorRequest{
		RiskIndicatorCode: request.RiskIndicatorCode,
		RiskIndicator:     request.RiskIndicator,
		ActivityID:        request.ActivityID,
		ProductID:         request.ProductID,
		CreatedAt:         &timeNow,
	}).Error
}

// Update implements RiskIndicatorDefinition
func (riskIndicator RiskIndicatorRepository) Update(request *models.RiskIndicatorRequest) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")
	return true, riskIndicator.db.DB.Save(&models.RiskIndicatorRequest{
		ID:                request.ID,
		RiskIndicatorCode: request.RiskIndicatorCode,
		RiskIndicator:     request.RiskIndicator,
		ActivityID:        request.ActivityID,
		ProductID:         request.ProductID,
		CreatedAt:         request.CreatedAt,
		UpdatedAt:         &timeNow,
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
