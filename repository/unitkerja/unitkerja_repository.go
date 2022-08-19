package unitkerja

import (
	"fmt"
	"riskmanagement/lib"
	models "riskmanagement/models/unitkerja"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type UnitKerjaDefinition interface {
	GetAll() (responses []models.UnitKerjaResponse, err error)
	GetOne(id int64) (responses models.UnitKerjaResponse, err error)
	Store(request *models.UnitKerjaRequest) (responses bool, err error)
	Update(request *models.UnitKerjaRequest) (responses bool, err error)
	Delete(id int64) (err error)
	WithTrx(trxHandle *gorm.DB) UnitKerjaRepository
}

type UnitKerjaRepository struct {
	db      lib.Database
	dbRaw   lib.Database
	logger  logger.Logger
	timeout time.Duration
}

func NewUnitKerjaRepository(db lib.Database, dbRaw lib.Database, logger logger.Logger) UnitKerjaDefinition {
	return UnitKerjaRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// Delete implements ActicityDefinition
func (unitKerja UnitKerjaRepository) Delete(id int64) (err error) {
	return unitKerja.db.DB.Where("id = ?", id).Delete(&models.UnitKerjaResponse{}).Error
}

// GetAll implements ActicityDefinition
func (unitKerja UnitKerjaRepository) GetAll() (responses []models.UnitKerjaResponse, err error) {
	return responses, unitKerja.db.DB.Find(&responses).Error
}

// GetOne implements ActicityDefinition
func (unitKerja UnitKerjaRepository) GetOne(id int64) (responses models.UnitKerjaResponse, err error) {
	return responses, unitKerja.db.DB.Where("id = ?", id).Find(&responses).Error
}

// Store implements ActicityDefinition
func (unitKerja UnitKerjaRepository) Store(request *models.UnitKerjaRequest) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")
	fmt.Println("repo = ", models.UnitKerjaRequest{
		NamaUker:  request.NamaUker,
		Status:    request.Status,
		CreatedAt: &timeNow,
	})
	err = unitKerja.db.DB.Save(&models.UnitKerjaRequest{
		NamaUker:  request.NamaUker,
		Status:    request.Status,
		CreatedAt: &timeNow,
	}).Error

	fmt.Println(err)
	return true, err
}

// Update implements ActicityDefinition
func (unitKerja UnitKerjaRepository) Update(request *models.UnitKerjaRequest) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")
	return true, unitKerja.db.DB.Save(&models.UnitKerjaRequest{
		ID:        request.ID,
		NamaUker:  request.NamaUker,
		Status:    request.Status,
		CreatedAt: request.CreatedAt,
		UpdatedAt: &timeNow,
	}).Error
}

// WithTrx implements ActicityDefinition
func (unitKerja UnitKerjaRepository) WithTrx(trxHandle *gorm.DB) UnitKerjaRepository {
	if trxHandle == nil {
		unitKerja.logger.Zap.Error("transaction Database not found in gin context")
		return unitKerja
	}
	unitKerja.db.DB = trxHandle
	return unitKerja
}
