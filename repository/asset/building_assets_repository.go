package asset

import (
	"infolelang/lib"
	models "infolelang/models/building_assets"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type BuildingAssetDefinition interface {
	GetAll() (responses []models.BuildingAssetsResponse, err error)
	GetOne(id int64) (responses models.BuildingAssetsResponse, err error)
	GetOneAsset(id int64) (responses models.BuildingAssetsResponse, err error)
	Store(request *models.BuildingAssets, tx *gorm.DB) (responses *models.BuildingAssets, err error)
	Update(request *models.BuildingAssetsRequest) (responses bool, err error)
	Delete(id int64) (err error)
	WithTrx(trxHandle *gorm.DB) BuildingAssetRepository
}
type BuildingAssetRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

func NewBuildingAssetReporitory(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger) BuildingAssetDefinition {
	return BuildingAssetRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// WithTrx implements BuildingAssetDefinition
func (buildingAsset BuildingAssetRepository) WithTrx(trxHandle *gorm.DB) BuildingAssetRepository {
	if trxHandle == nil {
		buildingAsset.logger.Zap.Error("transaction Database not found in gin context. ")
		return buildingAsset
	}
	buildingAsset.db.DB = trxHandle
	return buildingAsset
}

// GetAll implements BuildingAssetDefinition
func (buildingAsset BuildingAssetRepository) GetAll() (responses []models.BuildingAssetsResponse, err error) {
	return responses, buildingAsset.db.DB.Find(&responses).Error
}

// GetOne implements BuildingAssetDefinition
func (buildingAsset BuildingAssetRepository) GetOne(id int64) (responses models.BuildingAssetsResponse, err error) {
	return responses, buildingAsset.db.DB.Raw(`
		SELECT
		ba.id,
		ba.asset_id,
		ba.certificate_type_id,
		ct.name certificate_type_name,
		ba.certificate_number,
		ba.build_year,
		ba.surface_area,
		ba.building_area,
		ba.direction,
		ba.number_of_floors,
		ba.number_of_bedrooms,
		ba.number_of_bathrooms,
		ba.electrical_power,
		ba.carport,
		ba.created_at,
		ba.updated_at
		FROM building_assets ba 
		JOIN certificate_type ct on ba.certificate_type_id  = ct.id 
		WHERE  ba.asset_id = ? `, id).Find(&responses).Error
}

// GetOneAsset implements BuildingAssetDefinition
func (buildingAsset BuildingAssetRepository) GetOneAsset(id int64) (responses models.BuildingAssetsResponse, err error) {
	return responses, buildingAsset.db.DB.Raw(`
	SELECT
	ba.id,
	ba.asset_id,
	ba.certificate_type_id,
	ct.name certificate_type_name,
	ba.certificate_number,
	ba.build_year,
	ba.surface_area,
	ba.building_area,
	ba.direction,
	ba.number_of_floors,
	ba.number_of_bedrooms,
	ba.number_of_bathrooms,
	ba.electrical_power,
	ba.carport,
	ba.created_at,
	ba.updated_at
	FROM building_assets ba 
	JOIN certificate_type ct on ba.certificate_type_id  = ct.id 
	WHERE  ba.asset_id = ? `, id).Find(&responses).Error
}

// Store implements BuildingAssetDefinition
func (buildingAsset BuildingAssetRepository) Store(request *models.BuildingAssets, tx *gorm.DB) (responses *models.BuildingAssets, err error) {
	return request, tx.Save(&request).Error
}

// Update implements BuildingAssetDefinition
func (buildingAsset BuildingAssetRepository) Update(request *models.BuildingAssetsRequest) (responses bool, err error) {
	return true, buildingAsset.db.DB.Save(&responses).Error
}

// Delete implements BuildingAssetDefinition
func (buildingAsset BuildingAssetRepository) Delete(id int64) (err error) {
	return buildingAsset.db.DB.Where("id = ?", id).Delete(&models.BuildingAssetsResponse{}).Error
}
