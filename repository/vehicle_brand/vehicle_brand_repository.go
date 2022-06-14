package access_places

import (
	"fmt"
	"infolelang/lib"
	models "infolelang/models/vehicle_brand"
	"time"

	elastic "gitlab.com/golang-package-library/elasticsearch"
	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type VehicleBrandDefinition interface {
	GetAll() (responses []models.VehicleBrandResponse, err error)
	GetOne(id int64) (responses models.VehicleBrandResponse, err error)
	Store(request *models.VehicleBrandRequest) (responses bool, err error)
	Update(request *models.VehicleBrandRequest) (responses bool, err error)
	Delete(id int64) (err error)
	WithTrx(trxHandle *gorm.DB) VehicleBrandRepository
}
type VehicleBrandRepository struct {
	db      lib.Database
	db2     lib.Databases
	elastic elastic.Elasticsearch
	logger  logger.Logger
	timeout time.Duration
}

func NewVehicleBrandReporitory(
	db lib.Database,
	db2 lib.Databases,
	elastic elastic.Elasticsearch,
	logger logger.Logger) VehicleBrandDefinition {
	return VehicleBrandRepository{
		db:      db,
		db2:     db2,
		elastic: elastic,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// WithTrx implements VehicleBrandDefinition
func (VehicleBrand VehicleBrandRepository) WithTrx(trxHandle *gorm.DB) VehicleBrandRepository {
	if trxHandle == nil {
		VehicleBrand.logger.Zap.Error("transaction Database not found in gin context. ")
		return VehicleBrand
	}
	VehicleBrand.db.DB = trxHandle
	return VehicleBrand
}

// GetAll implements VehicleBrandDefinition
func (VehicleBrand VehicleBrandRepository) GetAll() (responses []models.VehicleBrandResponse, err error) {
	return responses, VehicleBrand.db.DB.Where("status = 1").Find(&responses).Error
}

// GetOne implements VehicleBrandDefinition
func (VehicleBrand VehicleBrandRepository) GetOne(id int64) (responses models.VehicleBrandResponse, err error) {
	return responses, VehicleBrand.db.DB.Where("id = ?", id).Find(&responses).Error
}

// Store implements VehicleBrandDefinition
func (VehicleBrand VehicleBrandRepository) Store(request *models.VehicleBrandRequest) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")
	fmt.Println("repo =", models.VehicleBrandRequest{
		Name:      request.Name,
		CreatedAt: &timeNow,
	})
	err = VehicleBrand.db.DB.Save(&models.VehicleBrand{
		Name:      request.Name,
		CreatedAt: &timeNow,
	}).Error
	fmt.Println(err)
	return true, err
}

// Update implements VehicleBrandDefinition
func (VehicleBrand VehicleBrandRepository) Update(request *models.VehicleBrandRequest) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")
	return true, VehicleBrand.db.DB.Save(&models.VehicleBrandRequest{
		ID:        request.ID,
		Name:      request.Name,
		CreatedAt: request.CreatedAt,
		UpdatedAt: &timeNow,
	}).Error
}

// Delete implements VehicleBrandDefinition
func (VehicleBrand VehicleBrandRepository) Delete(id int64) (err error) {
	return VehicleBrand.db.DB.Where("id = ?", id).Delete(&models.VehicleBrandResponse{}).Error
}
