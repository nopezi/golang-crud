package subincident

import (
	"riskmanagement/lib"
	models "riskmanagement/models/subincident"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type SubIncidentDefinition interface {
	GetAll() (responses []models.SubIncidentResponse, err error)
	GetOne(id int64) (responses models.SubIncidentResponse, err error)
	Store(request *models.SubIncidentRequest) (responses bool, err error)
	Update(request *models.SubIncidentRequest) (responses bool, err error)
	Delete(id int64) (err error)
	WithTrx(trxHandle *gorm.DB) SubIncidentRepository
}

type SubIncidentRepository struct {
	db      lib.Database
	dbRaw   lib.Database
	logger  logger.Logger
	timeout time.Duration
}

// Delete implements SubIncidentDefinition
func (subIncident SubIncidentRepository) Delete(id int64) (err error) {
	return subIncident.db.DB.Where("id = ?", id).Delete(&models.SubIncidentResponse{}).Error
}

// GetAll implements SubIncidentDefinition
func (subIncident SubIncidentRepository) GetAll() (responses []models.SubIncidentResponse, err error) {
	return responses, subIncident.db.DB.Find(&responses).Error

}

// GetOne implements SubIncidentDefinition
func (subIncident SubIncidentRepository) GetOne(id int64) (responses models.SubIncidentResponse, err error) {
	return responses, subIncident.db.DB.Where("id = ?", id).Find(&responses).Error
}

// Store implements SubIncidentDefinition
func (subIncident SubIncidentRepository) Store(request *models.SubIncidentRequest) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")
	return responses, subIncident.db.DB.Save(&models.SubIncidentRequest{
		KodeKejadian:             request.KodeKejadian,
		KodeSubKejadian:          request.KodeSubKejadian,
		KriteriaPenyebabKejadian: request.KriteriaPenyebabKejadian,
		CreatedAt:                &timeNow,
	}).Error
}

// Update implements SubIncidentDefinition
func (subIncident SubIncidentRepository) Update(request *models.SubIncidentRequest) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")
	return responses, subIncident.db.DB.Save(&models.SubIncidentRequest{
		ID:                       request.ID,
		KodeKejadian:             request.KodeKejadian,
		KodeSubKejadian:          request.KodeSubKejadian,
		KriteriaPenyebabKejadian: request.KriteriaPenyebabKejadian,
		CreatedAt:                request.CreatedAt,
		UpdatedAt:                &timeNow,
	}).Error
}

// WithTrx implements SubIncidentDefinition
func (subIncident SubIncidentRepository) WithTrx(trxHandle *gorm.DB) SubIncidentRepository {
	if trxHandle == nil {
		subIncident.logger.Zap.Error("transaction Database not found in gin context")
		return subIncident
	}

	subIncident.db.DB = trxHandle
	return subIncident
}

func NewSubIncidentRepository(
	db lib.Database,
	dbRaw lib.Database,
	logger logger.Logger,
) SubIncidentDefinition {
	return SubIncidentRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: time.Second * 100,
	}
}
