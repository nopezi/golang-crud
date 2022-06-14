package access_places

import (
	"fmt"
	"infolelang/lib"
	models "infolelang/models/vehicle_capacity"
	"time"

	elastic "gitlab.com/golang-package-library/elasticsearch"
	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type VehicleCapacityDefinition interface {
	GetAll() (responses []models.VehicleCapacityResponse, err error)
	GetOne(id int64) (responses models.VehicleCapacityResponse, err error)
	Store(request *models.VehicleCapacityRequest) (responses bool, err error)
	Update(request *models.VehicleCapacityRequest) (responses bool, err error)
	Delete(id int64) (err error)
	WithTrx(trxHandle *gorm.DB) VehicleCapacityRepository
}
type VehicleCapacityRepository struct {
	db      lib.Database
	db2     lib.Databases
	elastic elastic.Elasticsearch
	logger  logger.Logger
	timeout time.Duration
}

func NewVehicleCapacityReporitory(
	db lib.Database,
	db2 lib.Databases,
	elastic elastic.Elasticsearch,
	logger logger.Logger) VehicleCapacityDefinition {
	return VehicleCapacityRepository{
		db:      db,
		db2:     db2,
		elastic: elastic,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// WithTrx implements VehicleCapacityDefinition
func (VehicleCapacity VehicleCapacityRepository) WithTrx(trxHandle *gorm.DB) VehicleCapacityRepository {
	if trxHandle == nil {
		VehicleCapacity.logger.Zap.Error("transaction Database not found in gin context. ")
		return VehicleCapacity
	}
	VehicleCapacity.db.DB = trxHandle
	return VehicleCapacity
}

// GetAll implements VehicleCapacityDefinition
func (VehicleCapacity VehicleCapacityRepository) GetAll() (responses []models.VehicleCapacityResponse, err error) {
	return responses, VehicleCapacity.db.DB.Where("status = 1").Find(&responses).Error
}

// GetOne implements VehicleCapacityDefinition
func (VehicleCapacity VehicleCapacityRepository) GetOne(id int64) (responses models.VehicleCapacityResponse, err error) {
	return responses, VehicleCapacity.db.DB.Where("id = ?", id).Find(&responses).Error
}

// Store implements VehicleCapacityDefinition
func (VehicleCapacity VehicleCapacityRepository) Store(request *models.VehicleCapacityRequest) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")
	fmt.Println("repo =", models.VehicleCapacityRequest{
		Name:      request.Name,
		CreatedAt: &timeNow,
	})
	err = VehicleCapacity.db.DB.Save(&models.VehicleCapacity{
		Name:      request.Name,
		CreatedAt: &timeNow,
	}).Error
	fmt.Println(err)
	return true, err
}

// Update implements VehicleCapacityDefinition
func (VehicleCapacity VehicleCapacityRepository) Update(request *models.VehicleCapacityRequest) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")
	return true, VehicleCapacity.db.DB.Save(&models.VehicleCapacityRequest{
		ID:        request.ID,
		Name:      request.Name,
		CreatedAt: request.CreatedAt,
		UpdatedAt: &timeNow,
	}).Error
}

// Delete implements VehicleCapacityDefinition
func (VehicleCapacity VehicleCapacityRepository) Delete(id int64) (err error) {
	return VehicleCapacity.db.DB.Where("id = ?", id).Delete(&models.VehicleCapacityResponse{}).Error
}
