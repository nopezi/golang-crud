package access_places

import (
	"fmt"
	"infolelang/lib"
	models "infolelang/models/vehicle_category"
	"time"

	elastic "gitlab.com/golang-package-library/elasticsearch"
	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type VehicleCategoryDefinition interface {
	GetAll() (responses []models.VehicleCategoryResponse, err error)
	GetOne(id int64) (responses models.VehicleCategoryResponse, err error)
	Store(request *models.VehicleCategoryRequest) (responses bool, err error)
	Update(request *models.VehicleCategoryRequest) (responses bool, err error)
	Delete(id int64) (err error)
	WithTrx(trxHandle *gorm.DB) VehicleCategoryRepository
}
type VehicleCategoryRepository struct {
	db      lib.Database
	db2     lib.Databases
	elastic elastic.Elasticsearch
	logger  logger.Logger
	timeout time.Duration
}

func NewVehicleCategoryReporitory(
	db lib.Database,
	db2 lib.Databases,
	elastic elastic.Elasticsearch,
	logger logger.Logger) VehicleCategoryDefinition {
	return VehicleCategoryRepository{
		db:      db,
		db2:     db2,
		elastic: elastic,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// WithTrx implements VehicleCategoryDefinition
func (VehicleCategory VehicleCategoryRepository) WithTrx(trxHandle *gorm.DB) VehicleCategoryRepository {
	if trxHandle == nil {
		VehicleCategory.logger.Zap.Error("transaction Database not found in gin context. ")
		return VehicleCategory
	}
	VehicleCategory.db.DB = trxHandle
	return VehicleCategory
}

// GetAll implements VehicleCategoryDefinition
func (VehicleCategory VehicleCategoryRepository) GetAll() (responses []models.VehicleCategoryResponse, err error) {
	return responses, VehicleCategory.db.DB.Where("status = 1").Find(&responses).Error
}

// GetOne implements VehicleCategoryDefinition
func (VehicleCategory VehicleCategoryRepository) GetOne(id int64) (responses models.VehicleCategoryResponse, err error) {
	return responses, VehicleCategory.db.DB.Where("id = ?", id).Find(&responses).Error
}

// Store implements VehicleCategoryDefinition
func (VehicleCategory VehicleCategoryRepository) Store(request *models.VehicleCategoryRequest) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")
	fmt.Println("repo =", models.VehicleCategoryRequest{
		Name:      request.Name,
		CreatedAt: &timeNow,
	})
	err = VehicleCategory.db.DB.Save(&models.VehicleCategory{
		Name:      request.Name,
		CreatedAt: &timeNow,
	}).Error
	fmt.Println(err)
	return true, err
}

// Update implements VehicleCategoryDefinition
func (VehicleCategory VehicleCategoryRepository) Update(request *models.VehicleCategoryRequest) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")
	return true, VehicleCategory.db.DB.Save(&models.VehicleCategoryRequest{
		ID:        request.ID,
		Name:      request.Name,
		CreatedAt: request.CreatedAt,
		UpdatedAt: &timeNow,
	}).Error
}

// Delete implements VehicleCategoryDefinition
func (VehicleCategory VehicleCategoryRepository) Delete(id int64) (err error) {
	return VehicleCategory.db.DB.Where("id = ?", id).Delete(&models.VehicleCategoryResponse{}).Error
}
