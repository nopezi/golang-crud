package CertificateType

import (
	models "infolelang/models/certificate_type"
	repository "infolelang/repository/certificate_type"

	"gitlab.com/golang-package-library/logger"
)

type CertificateTypeDefinition interface {
	GetAll() (responses []models.CertificateTypeResponse, err error)
	GetOne(id int64) (responses models.CertificateTypeResponse, err error)
	Store(request *models.CertificateTypeRequest) (status bool, err error)
	Update(request *models.CertificateTypeRequest) (status bool, err error)
	Delete(id int64) (err error)
}
type CertificateTypeService struct {
	logger     logger.Logger
	repository repository.CertificateTypeDefinition
}

func NewCertificateTypeService(logger logger.Logger, repository repository.CertificateTypeDefinition) CertificateTypeDefinition {
	return CertificateTypeService{
		logger:     logger,
		repository: repository,
	}
}

// GetAll implements CertificateTypeDefinition
func (CertificateType CertificateTypeService) GetAll() (responses []models.CertificateTypeResponse, err error) {
	return CertificateType.repository.GetAll()
}

// GetOne implements CertificateTypeDefinition
func (CertificateType CertificateTypeService) GetOne(id int64) (responses models.CertificateTypeResponse, err error) {
	return CertificateType.repository.GetOne(id)
}

// Store implements CertificateTypeDefinition
func (CertificateType CertificateTypeService) Store(request *models.CertificateTypeRequest) (status bool, err error) {
	status, err = CertificateType.repository.Store(request)
	return status, err
}

// Update implements CertificateTypeDefinition
func (CertificateType CertificateTypeService) Update(request *models.CertificateTypeRequest) (status bool, err error) {
	status, err = CertificateType.repository.Update(request)
	return status, err
}

// Delete implements CertificateTypeDefinition
func (CertificateType CertificateTypeService) Delete(id int64) (err error) {
	return CertificateType.repository.Delete(id)
}
