package address

import (
	"infolelang/lib"
	models "infolelang/models/addresses"
	"time"

	elastic "gitlab.com/golang-package-library/elasticsearch"
	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type AddressDefinition interface {
	GetAll() (responses []models.AddressesResponse, err error)
	GetOne(id int64) (responses models.AddressesResponse, err error)
	GetOneAsset(id int64) (responses models.AddressesResponse, err error)
	Store(request *models.Addresses) (responses *models.Addresses, err error)
	Update(request *models.AddressesRequest) (responses bool, err error)
	Delete(id int64) (err error)
	WithTrx(trxHandle *gorm.DB) AddressRepository
}
type AddressRepository struct {
	db      lib.Database
	db2     lib.Databases
	elastic elastic.Elasticsearch
	logger  logger.Logger
	timeout time.Duration
}

func NewAddressReporitory(
	db lib.Database,
	db2 lib.Databases,
	elastic elastic.Elasticsearch,
	logger logger.Logger) AddressDefinition {
	return AddressRepository{
		db:      db,
		db2:     db2,
		elastic: elastic,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// WithTrx implements AddressDefinition
func (address AddressRepository) WithTrx(trxHandle *gorm.DB) AddressRepository {
	if trxHandle == nil {
		address.logger.Zap.Error("transaction Database not found in gin context. ")
		return address
	}
	address.db.DB = trxHandle
	return address
}

// GetAll implements AddressDefinition
func (address AddressRepository) GetAll() (responses []models.AddressesResponse, err error) {
	return responses, address.db.DB.Find(&responses).Error
}

// GetOne implements AddressDefinition
func (address AddressRepository) GetOne(id int64) (responses models.AddressesResponse, err error) {
	return responses, address.db.DB.Where("id = ?", id).Find(&responses).Error
}

// GetOneAsset implements AddressDefinition
func (address AddressRepository) GetOneAsset(id int64) (responses models.AddressesResponse, err error) {

	return responses, address.db.DB.Raw(`
		select a.id, 
		a.asset_id,
		a.address, 
		a.longitude, 
		a.langitude, 
		CONCAT(a.longitude,',',a.langitude)  longlat, 
		rpc.postal_code,
		rpc.region, 
		rpc.district, 
		rpc.city, 
		rpc.province,
		a.created_at,
		a.updated_at
		from addresses a 
		left join ref_postal_code rpc on a.postalcode_id = rpc.postal_code
		where asset_id = ?`, id).Find(&responses).Error
}

// Store implements AddressDefinition
func (address AddressRepository) Store(request *models.Addresses) (responses *models.Addresses, err error) {
	return request, address.db.DB.Save(&request).Error
}

// Update implements AddressDefinition
func (address AddressRepository) Update(request *models.AddressesRequest) (responses bool, err error) {
	return true, address.db.DB.Save(&request).Error
}

// Delete implements AddressDefinition
func (address AddressRepository) Delete(id int64) (err error) {
	return address.db.DB.Where("id = ?", id).Delete(&models.AddressesResponse{}).Error
}
