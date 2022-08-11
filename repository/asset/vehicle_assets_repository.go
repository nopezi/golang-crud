package asset

import (
	"infolelang/lib"
	models "infolelang/models/vehicle_assets"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type VehicleAssetDefinition interface {
	GetAll() (responses []models.VehicleAssetsResponse, err error)
	GetOne(id int64) (responses models.VehicleAssetsResponse, err error)
	GetOneAsset(id int64) (responses models.VehicleAssetsResponse, err error)
	Store(request *models.VehicleAssets, tx *gorm.DB) (responses *models.VehicleAssets, err error)
	Update(request *models.VehicleAssetsRequest) (responses bool, err error)
	Delete(id int64) (err error)
	WithTrx(trxHandle *gorm.DB) VehicleAssetRepository
}
type VehicleAssetRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

func NewVehicleAssetReporitory(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger) VehicleAssetDefinition {
	return VehicleAssetRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// WithTrx implements VehicleAssetDefinition
func (vehicleAsset VehicleAssetRepository) WithTrx(trxHandle *gorm.DB) VehicleAssetRepository {
	if trxHandle == nil {
		vehicleAsset.logger.Zap.Error("transaction Database not found in gin context. ")
		return vehicleAsset
	}
	vehicleAsset.db.DB = trxHandle
	return vehicleAsset
}

// GetAll implements VehicleAssetDefinition
func (vehicleAsset VehicleAssetRepository) GetAll() (responses []models.VehicleAssetsResponse, err error) {
	return responses, vehicleAsset.db.DB.Find(&responses).Error
}

// GetOne implements VehicleAssetDefinition
func (vehicleAsset VehicleAssetRepository) GetOne(id int64) (responses models.VehicleAssetsResponse, err error) {
	return responses, vehicleAsset.db.DB.Raw(`
		SELECT
		va.id,
		va.asset_id,
		va.vehicle_type,
		va.certificate_type_id,
		va.certificate_number,
		va.series,
		va.brand_id,
		va.type,
		va.production_year,
		va.transmission_id,
		va.machine_capacity_id,
		va.color_id,
		va.number_of_seat,
		va.number_of_usage,
		va.machine_number,
		va.body_number,
		va.licence_date,
		va.created_at,
		va.updated_at,
		ct.name certificate_type_name,
		vb.name brand_name,
		vt.name transmission_name,
		vc.name machine_capacity_name,
		vc2.name color_name
		from vehicle_assets va 
		left join certificate_type ct on va.certificate_type_id = ct.id 
		left join vehicle_brand vb on va.brand_id = vb.id 
		left join vehicle_transmission vt on va.transmission_id = vt.id 
		left join vehicle_capacity vc on va.machine_capacity_id = vc.id 
		left join vehicle_color vc2 on va.color_id = vc2.id 
		WHERE va.id = ?`, id).Find(&responses).Error
}

// GetOneAsset implements VehicleAssetDefinition
func (vehicleAsset VehicleAssetRepository) GetOneAsset(id int64) (responses models.VehicleAssetsResponse, err error) {
	err = vehicleAsset.db.DB.Raw(`
	select 
	va.id,
	va.asset_id,
	va.vehicle_type,
	va.certificate_type_id,
	va.certificate_number,
	va.series,
	va.brand_id,
	va.type,
	va.production_year,
	va.transmission_id,
	va.machine_capacity_id,
	va.color_id,
	va.number_of_seat,
	va.number_of_usage,
	va.machine_number,
	va.body_number,
	va.licence_date,
	va.created_at,
	va.updated_at,
	ct.name certificate_type_name,
	vb.name brand_name,
	vt.name transmission_name,
	vc.name machine_capacity_name,
	vc2.name color_name
	from vehicle_assets va 
	left join certificate_type ct on va.certificate_type_id = ct.id 
	left join vehicle_brand vb on va.brand_id = vb.id 
	left join vehicle_transmission vt on va.transmission_id = vt.id 
	left join vehicle_capacity vc on va.machine_capacity_id = vc.id 
	left join vehicle_color vc2 on va.color_id = vc2.id 
	where va.asset_id = ?`, id).Find(&responses).Error

	if err != nil {
		vehicleAsset.logger.Zap.Error(err)
		return responses, err
	}
	return responses, err
}

// Store implements VehicleAssetDefinition
func (vehicleAsset VehicleAssetRepository) Store(request *models.VehicleAssets, tx *gorm.DB) (responses *models.VehicleAssets, err error) {
	return request, tx.Save(&request).Error
}

// Update implements VehicleAssetDefinition
func (vehicleAsset VehicleAssetRepository) Update(request *models.VehicleAssetsRequest) (responses bool, err error) {
	return true, vehicleAsset.db.DB.Save(&responses).Error
}

// Delete implements VehicleAssetDefinition
func (vehicleAsset VehicleAssetRepository) Delete(id int64) (err error) {
	return vehicleAsset.db.DB.Where("id = ?", id).Delete(&models.VehicleAssetsResponse{}).Error
}
