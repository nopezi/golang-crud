package CertificateType

import (
	"infolelang/lib"
	models "infolelang/models/certificate_type"
	"time"

	elastic "gitlab.com/golang-package-library/elasticsearch"
	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type CertificateTypeDefinition interface {
	GetAll() (responses []models.CertificateTypeResponse, err error)
	GetOne(id int64) (responses models.CertificateTypeResponse, err error)
	Store(request *models.CertificateTypeRequest) (responses bool, err error)
	Update(request *models.CertificateTypeRequest) (responses bool, err error)
	Delete(id int64) (err error)
	WithTrx(trxHandle *gorm.DB) CertificateTypeRepository
}
type CertificateTypeRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	elastic elastic.Elasticsearch
	logger  logger.Logger
	timeout time.Duration
}

func NewCertificateTypeReporitory(
	db lib.Database,
	dbRaw lib.Databases,
	elastic elastic.Elasticsearch,
	logger logger.Logger) CertificateTypeDefinition {
	return CertificateTypeRepository{
		db:      db,
		dbRaw:   dbRaw,
		elastic: elastic,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// WithTrx implements CertificateTypeDefinition
func (CertificateType CertificateTypeRepository) WithTrx(trxHandle *gorm.DB) CertificateTypeRepository {
	if trxHandle == nil {
		CertificateType.logger.Zap.Error("transaction Database not found in gin context. ")
		return CertificateType
	}
	CertificateType.db.DB = trxHandle
	return CertificateType
}

// GetAll implements CertificateTypeDefinition
func (CertificateType CertificateTypeRepository) GetAll() (responses []models.CertificateTypeResponse, err error) {
	return responses, CertificateType.db.DB.Find(&responses).Error
}

// GetOne implements CertificateTypeDefinition
func (CertificateType CertificateTypeRepository) GetOne(id int64) (responses models.CertificateTypeResponse, err error) {
	return responses, CertificateType.db.DB.Where("id = ?", id).Find(&responses).Error
}

// Store implements CertificateTypeDefinition
func (CertificateType CertificateTypeRepository) Store(request *models.CertificateTypeRequest) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")
	return responses, CertificateType.db.DB.Save(&models.CertificateTypeRequest{
		Name:        request.Name,
		Icon:        request.Icon,
		Description: request.Description,
		// Status:      request.Status,
		CreatedAt: &timeNow,
	}).Error
}

// Update implements CertificateTypeDefinition
func (CertificateType CertificateTypeRepository) Update(request *models.CertificateTypeRequest) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")
	return true, CertificateType.db.DB.Save(&models.CertificateTypeRequest{
		ID:          request.ID,
		Name:        request.Name,
		Icon:        request.Icon,
		Description: request.Description,
		// Status:      request.Status,
		CreatedAt: request.CreatedAt,
		UpdatedAt: &timeNow,
	}).Error
}

// Delete implements CertificateTypeDefinition
func (CertificateType CertificateTypeRepository) Delete(id int64) (err error) {
	return CertificateType.db.DB.Where("id = ?", id).Delete(&models.CertificateTypeResponse{}).Error
}
