package incident

import (
	"riskmanagement/lib"
	models "riskmanagement/models/incident"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type IncidentDefinition interface {
	GetAll() (responses []models.IncidentResponse, err error)
	GetOne(id int64) (responses models.IncidentResponse, err error)
	Store(request *models.IncidentRequest) (responses bool, err error)
	Update(request *models.IncidentRequest) (responses bool, err error)
	Delete(id int64) (err error)
	WithTrx(trxHandle *gorm.DB) IncidentRepository
}

type IncidentRepository struct {
	db      lib.Database
	dbRaw   lib.Database
	logger  logger.Logger
	timeout time.Duration
}

// Delete implements IncidentDefinition
func (incident IncidentRepository) Delete(id int64) (err error) {
	return incident.db.DB.Where("id = ?", id).Delete(&models.IncidentResponse{}).Error
}

// GetAll implements IncidentDefinition
func (incident IncidentRepository) GetAll() (responses []models.IncidentResponse, err error) {
	return responses, incident.db.DB.Find(&responses).Error
}

// GetOne implements IncidentDefinition
func (incident IncidentRepository) GetOne(id int64) (responses models.IncidentResponse, err error) {
	return responses, incident.db.DB.Where("id = ?", id).Find(&responses).Error
}

// Store implements IncidentDefinition
func (incident IncidentRepository) Store(request *models.IncidentRequest) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")
	return responses, incident.db.DB.Save(&models.IncidentRequest{
		KodeKejadian:      request.KodeKejadian,
		PenyebabKejadian1: request.PenyebabKejadian1,
		CreatedAt:         &timeNow,
	}).Error
}

// Update implements IncidentDefinition
func (incident IncidentRepository) Update(request *models.IncidentRequest) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")
	return true, incident.db.DB.Save(&models.IncidentRequest{
		ID:                request.ID,
		KodeKejadian:      request.KodeKejadian,
		PenyebabKejadian1: request.PenyebabKejadian1,
		CreatedAt:         request.CreatedAt,
		UpdatedAt:         &timeNow,
	}).Error
}

// WithTrx implements IncidentDefinition
func (incident IncidentRepository) WithTrx(trxHandle *gorm.DB) IncidentRepository {
	if trxHandle == nil {
		incident.logger.Zap.Error("transaction Database not found in gin context")
		return incident
	}

	incident.db.DB = trxHandle
	return incident
}

func NewIncidentRepository(
	db lib.Database,
	dbRaw lib.Database,
	logger logger.Logger,
) IncidentDefinition {
	return IncidentRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: time.Second * 100,
	}
}
