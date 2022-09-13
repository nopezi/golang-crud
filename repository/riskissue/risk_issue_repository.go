package riskissue

import (
	"riskmanagement/lib"
	models "riskmanagement/models/riskissue"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type RiskIssueDefinition interface {
	GetAll() (responses []models.RiskIssueResponse, err error)
	GetOne(id int64) (responses models.RiskIssueResponse, err error)
	Store(request *models.RiskIssueRequest) (responses bool, err error)
	Update(request *models.RiskIssueRequest) (responses bool, err error)
	Delete(id int64) (err error)
	WithTrx(trxHandle *gorm.DB) RiskIssueRepository
}

type RiskIssueRepository struct {
	db      lib.Database
	dbRaw   lib.Database
	logger  logger.Logger
	timeout time.Duration
}

// Delete implements RiskIssueDefinition
func (riskIssue RiskIssueRepository) Delete(id int64) (err error) {
	return riskIssue.db.DB.Where("id = ?", id).Delete(&models.RiskIssueResponse{}).Error
}

// GetAll implements RiskIssueDefinition
func (riskIssue RiskIssueRepository) GetAll() (responses []models.RiskIssueResponse, err error) {
	return responses, riskIssue.db.DB.Find(&responses).Error
}

// GetOne implements RiskIssueDefinition
func (riskIssue RiskIssueRepository) GetOne(id int64) (responses models.RiskIssueResponse, err error) {
	return responses, riskIssue.db.DB.Where("id = ?", id).Find(&responses).Error
}

// Store implements RiskIssueDefinition
func (riskIssue RiskIssueRepository) Store(request *models.RiskIssueRequest) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")
	return responses, riskIssue.db.DB.Save(&models.RiskIssueRequest{
		ID:             0,
		RiskTypeID:     request.RiskTypeID,
		RiskIssueCode:  request.RiskIssueCode,
		RiskIssue:      request.RiskIssue,
		Deskripsi:      request.Deskripsi,
		KategoriRisiko: request.KategoriRisiko,
		Status:         request.Status,
		CreatedAt:      &timeNow,
		UpdatedAt:      new(string),
	}).Error
}

// Update implements RiskIssueDefinition
func (riskIssue RiskIssueRepository) Update(request *models.RiskIssueRequest) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")

	return true, riskIssue.db.DB.Save(&models.RiskIssueRequest{
		ID:             request.ID,
		RiskTypeID:     request.RiskTypeID,
		RiskIssueCode:  request.RiskIssueCode,
		RiskIssue:      request.RiskIssue,
		Deskripsi:      request.Deskripsi,
		KategoriRisiko: request.KategoriRisiko,
		Status:         request.Status,
		CreatedAt:      request.CreatedAt,
		UpdatedAt:      &timeNow,
	}).Error
}

// WithTrx implements RiskIssueDefinition
func (riskIssue RiskIssueRepository) WithTrx(trxHandle *gorm.DB) RiskIssueRepository {
	if trxHandle == nil {
		riskIssue.logger.Zap.Error("transaction Database not found in gin context")
		return riskIssue
	}

	riskIssue.db.DB = trxHandle
	return riskIssue
}

func NewRiskIssueRepository(
	db lib.Database,
	dbRaw lib.Database,
	logger logger.Logger,
) RiskIssueDefinition {
	return RiskIssueRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: time.Second * 100,
	}
}
