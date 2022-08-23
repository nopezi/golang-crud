package materi

import (
	"riskmanagement/lib"
	models "riskmanagement/models/materi"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type MateriDefinition interface {
	GetAll() (responses []models.MateriResponses, err error)
	GetOne(id int64) (responses models.MateriResponses, err error)
	GetOneMateri(id int64) (responses []models.MateriResponses, err error)
	Store(request *models.Materi, tx *gorm.DB) (reponses *models.Materi, err error)
	Update(request *models.MateriRequest) (responses bool, err error)
	Delete(id int64, tx *gorm.DB) (err error)
	WithTrx(trxHandle *gorm.DB) MateriRepository
}

type MateriRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

// Delete implements MateriDefinition
func (MateriRepository) Delete(id int64, tx *gorm.DB) (err error) {
	return tx.Where("id = ?", id).Delete(&models.MateriResponses{}).Error
}

// GetAll implements MateriDefinition
func (materi MateriRepository) GetAll() (responses []models.MateriResponses, err error) {
	return responses, materi.db.DB.Find(&responses).Error
}

// GetOne implements MateriDefinition
func (materi MateriRepository) GetOne(id int64) (responses models.MateriResponses, err error) {
	return responses, materi.db.DB.Raw("SELECT * FROM rekomendasi_materi where id = ?", id).Find(&responses).Error
}

// GetOneMateri implements MateriDefinition
func (materi MateriRepository) GetOneMateri(id int64) (responses []models.MateriResponses, err error) {
	panic("unimplemented")
}

// Store implements MateriDefinition
func (materi MateriRepository) Store(request *models.Materi, tx *gorm.DB) (reponses *models.Materi, err error) {
	return request, tx.Save(&request).Error
}

// Update implements MateriDefinition
func (materi MateriRepository) Update(request *models.MateriRequest) (responses bool, err error) {
	return true, materi.db.DB.Save(&request).Error
}

// WithTrx implements MateriDefinition
func (materi MateriRepository) WithTrx(trxHandle *gorm.DB) MateriRepository {
	panic("unimplemented")
}

func NewMateriRepository(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger) MateriDefinition {
	return MateriRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: time.Second * 100,
	}
}
