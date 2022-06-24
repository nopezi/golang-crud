package asset

import (
	"infolelang/lib"
	models "infolelang/models/assets"
	facilitiesModel "infolelang/models/facilities"
	"time"

	elastic "gitlab.com/golang-package-library/elasticsearch"
	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type AssetFacilityDefinition interface {
	GetAll() (responses []models.AssetFacilitiesResponse, err error)
	GetOne(id int64) (responses models.AssetFacilitiesResponse, err error)
	GetOneAsset(id int64) (responses []facilitiesModel.FacilitiesResponse, err error)
	Store(request *models.AssetFacilities) (responses *models.AssetFacilities, err error)
	Update(request *models.AssetFacilitiesRequest) (responses bool, err error)
	Delete(id int64) (err error)
	WithTrx(trxHandle *gorm.DB) AssetFacilityRepository
}
type AssetFacilityRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	elastic elastic.Elasticsearch
	logger  logger.Logger
	timeout time.Duration
}

func NewAssetFacilityReporitory(
	db lib.Database,
	dbRaw lib.Databases,
	elastic elastic.Elasticsearch,
	logger logger.Logger) AssetFacilityDefinition {
	return AssetFacilityRepository{
		db:      db,
		dbRaw:   dbRaw,
		elastic: elastic,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// WithTrx implements AssetFacilityDefinition
func (AssetFacility AssetFacilityRepository) WithTrx(trxHandle *gorm.DB) AssetFacilityRepository {
	if trxHandle == nil {
		AssetFacility.logger.Zap.Error("transaction Database not found in gin context. ")
		return AssetFacility
	}
	AssetFacility.db.DB = trxHandle
	return AssetFacility
}

// GetAll implements AssetFacilityDefinition
func (AssetFacility AssetFacilityRepository) GetAll() (responses []models.AssetFacilitiesResponse, err error) {
	return responses, AssetFacility.db.DB.Find(&responses).Error
}

// GetOne implements AssetFacilityDefinition
func (AssetFacility AssetFacilityRepository) GetOne(id int64) (responses models.AssetFacilitiesResponse, err error) {
	return responses, AssetFacility.db.DB.Where("id = ?", id).Find(&responses).Error
}

// GetOneAsset implements AssetFacilityDefinition
func (AssetFacility AssetFacilityRepository) GetOneAsset(id int64) (responses []facilitiesModel.FacilitiesResponse, err error) {
	// return responses, AssetFacility.db.DB.Where("asset_id = ?", id).Find(&responses).Error
	rows, err := AssetFacility.db.DB.Raw(`
	select f.id,f.name, 
	f.icon , af.status,
	f.description  from 
	asset_facilities af 
	left join facilities f 
	on af.facility_id  = f.id 
	where af.asset_id = ? 
	order by f.id  asc`, id).Rows()

	defer rows.Close()

	var facilitiy facilitiesModel.FacilitiesResponse
	for rows.Next() {
		// ScanRows scan a row into user
		AssetFacility.db.DB.ScanRows(rows, &facilitiy)
		responses = append(responses, facilitiy)
		// do something
	}
	return responses, err
}

// Store implements AssetFacilityDefinition
func (AssetFacility AssetFacilityRepository) Store(request *models.AssetFacilities) (responses *models.AssetFacilities, err error) {
	return request, AssetFacility.db.DB.Save(&request).Error
}

// Update implements AssetFacilityDefinition
func (AssetFacility AssetFacilityRepository) Update(request *models.AssetFacilitiesRequest) (responses bool, err error) {
	return true, AssetFacility.db.DB.Save(&request).Error
}

// Delete implements AssetFacilityDefinition
func (AssetFacility AssetFacilityRepository) Delete(id int64) (err error) {
	return AssetFacility.db.DB.Where("id = ?", id).Delete(&models.AssetFacilitiesResponse{}).Error
}
