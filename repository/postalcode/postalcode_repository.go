package postalcode

import (
	"infolelang/lib"
	models "infolelang/models/postalcode"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type PostalcodeDefinition interface {
	GetAll() (responses []models.PostalcodeResponse, err error)
	GetOne(id int64) (responses models.PostalcodeResponse, err error)
	FindPostalCode(postalcode string) (responses models.PostalcodeResponse, err error)
	Store(request *models.PostalcodeRequest) (responses bool, err error)
	Update(request *models.PostalcodeRequest) (responses bool, err error)
	Delete(id int64) (err error)
	WithTrx(trxHandle *gorm.DB) PostalcodeRepository
}
type PostalcodeRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

func NewPostalcodeReporitory(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger) PostalcodeDefinition {
	return PostalcodeRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// WithTrx implements PostalcodeDefinition
func (Postalcode PostalcodeRepository) WithTrx(trxHandle *gorm.DB) PostalcodeRepository {
	if trxHandle == nil {
		Postalcode.logger.Zap.Error("transaction Database not found in gin context. ")
		return Postalcode
	}
	Postalcode.db.DB = trxHandle
	return Postalcode
}

// GetAll implements PostalcodeDefinition
func (Postalcode PostalcodeRepository) GetAll() (responses []models.PostalcodeResponse, err error) {
	return responses, Postalcode.db.DB.Find(&responses).Error
}

// GetOne implements PostalcodeDefinition
func (Postalcode PostalcodeRepository) GetOne(id int64) (responses models.PostalcodeResponse, err error) {
	return responses, Postalcode.db.DB.Where("id = ?", id).Find(&responses).Error
}

// GetOne implements PostalcodeDefinition
func (Postalcode PostalcodeRepository) FindPostalCode(postalcode string) (responses models.PostalcodeResponse, err error) {
	find := Postalcode.db.DB.Where("postal_code = ? ", postalcode).Find(&responses).Error

	if find == nil {
		return responses, err
	}
	return responses, err
}

// Store implements PostalcodeDefinition
func (Postalcode PostalcodeRepository) Store(request *models.PostalcodeRequest) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")

	return true, Postalcode.db.DB.Save(&models.PostalcodeRequest{
		PostalCode: request.PostalCode,
		Region:     request.Region,
		District:   request.District,
		City:       request.City,
		Province:   request.Province,
		Enabled:    request.Enabled,
		CreatedAt:  &timeNow,
	}).Error
}

// Update implements PostalcodeDefinition
func (Postalcode PostalcodeRepository) Update(request *models.PostalcodeRequest) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")

	return true, Postalcode.db.DB.Save(&models.PostalcodeRequest{
		ID:         request.ID,
		PostalCode: request.PostalCode,
		Region:     request.Region,
		District:   request.District,
		City:       request.City,
		Province:   request.Province,
		Enabled:    request.Enabled,
		CreatedAt:  request.CreatedAt,
		UpdatedAt:  &timeNow,
	}).Error
}

// Delete implements PostalcodeDefinition
func (Postalcode PostalcodeRepository) Delete(id int64) (err error) {
	return Postalcode.db.DB.Where("id = ?", id).Delete(&models.PostalcodeResponse{}).Error
}
