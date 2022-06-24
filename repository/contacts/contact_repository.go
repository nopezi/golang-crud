package contacts

import (
	"infolelang/lib"
	models "infolelang/models/contacts"
	"time"

	elastic "gitlab.com/golang-package-library/elasticsearch"
	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type ContactDefinition interface {
	GetAll() (responses []models.ContactsResponse, err error)
	GetOne(id int64) (responses models.ContactsResponse, err error)
	GetOneAsset(id int64) (responses models.ContactsResponse, err error)
	Store(request *models.Contacts) (responses *models.Contacts, err error)
	Update(request *models.ContactsRequest) (responses bool, err error)
	Delete(id int64) (err error)
	WithTrx(trxHandle *gorm.DB) ContactRepository
}
type ContactRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	elastic elastic.Elasticsearch
	logger  logger.Logger
	timeout time.Duration
}

func NewContactReporitory(
	db lib.Database,
	dbRaw lib.Databases,
	elastic elastic.Elasticsearch,
	logger logger.Logger) ContactDefinition {
	return ContactRepository{
		db:      db,
		dbRaw:   dbRaw,
		elastic: elastic,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// WithTrx implements ContactDefinition
func (contact ContactRepository) WithTrx(trxHandle *gorm.DB) ContactRepository {
	if trxHandle == nil {
		contact.logger.Zap.Error("transaction Database not found in gin context. ")
		return contact
	}
	contact.db.DB = trxHandle
	return contact
}

// GetAll implements ContactDefinition
func (contact ContactRepository) GetAll() (responses []models.ContactsResponse, err error) {
	return responses, contact.db.DB.Find(&responses).Error
}

// GetOne implements ContactDefinition
func (contact ContactRepository) GetOne(id int64) (responses models.ContactsResponse, err error) {
	return responses, contact.db.DB.Where("id = ?", id).Find(&responses).Error
}

// GetOneAsset implements ContactDefinition
func (contact ContactRepository) GetOneAsset(id int64) (responses models.ContactsResponse, err error) {
	return responses, contact.db.DB.Where("asset_id = ?", id).Find(&responses).Error
}

// Store implements ContactDefinition
func (contact ContactRepository) Store(request *models.Contacts) (responses *models.Contacts, err error) {
	return request, contact.db.DB.Save(&request).Error
}

// Update implements ContactDefinition
func (Contact ContactRepository) Update(request *models.ContactsRequest) (responses bool, err error) {
	return true, Contact.db.DB.Save(&request).Error
}

// Delete implements ContactDefinition
func (Contact ContactRepository) Delete(id int64) (err error) {
	return Contact.db.DB.Where("id = ?", id).Delete(&models.ContactsResponse{}).Error
}
