package access_places

import (
	"fmt"
	"infolelang/lib"
	models "infolelang/models/vehicle_transmission"
	"time"

	elastic "gitlab.com/golang-package-library/elasticsearch"
	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type VehicleTransmissionDefinition interface {
	GetAll() (responses []models.VehicleTransmissionResponse, err error)
	GetOne(id int64) (responses models.VehicleTransmissionResponse, err error)
	Store(request *models.VehicleTransmissionRequest) (responses bool, err error)
	Update(request *models.VehicleTransmissionRequest) (responses bool, err error)
	Delete(id int64) (err error)
	WithTrx(trxHandle *gorm.DB) VehicleTransmissionRepository
}
type VehicleTransmissionRepository struct {
	db      lib.Database
	db2     lib.Databases
	elastic elastic.Elasticsearch
	logger  logger.Logger
	timeout time.Duration
}

func NewVehicleTransmissionReporitory(
	db lib.Database,
	db2 lib.Databases,
	elastic elastic.Elasticsearch,
	logger logger.Logger) VehicleTransmissionDefinition {
	return VehicleTransmissionRepository{
		db:      db,
		db2:     db2,
		elastic: elastic,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// WithTrx implements VehicleTransmissionDefinition
func (VehicleTransmission VehicleTransmissionRepository) WithTrx(trxHandle *gorm.DB) VehicleTransmissionRepository {
	if trxHandle == nil {
		VehicleTransmission.logger.Zap.Error("transaction Database not found in gin context. ")
		return VehicleTransmission
	}
	VehicleTransmission.db.DB = trxHandle
	return VehicleTransmission
}

// GetAll implements VehicleTransmissionDefinition
func (VehicleTransmission VehicleTransmissionRepository) GetAll() (responses []models.VehicleTransmissionResponse, err error) {
	return responses, VehicleTransmission.db.DB.Where("status = 1").Find(&responses).Error
}

// GetOne implements VehicleTransmissionDefinition
func (VehicleTransmission VehicleTransmissionRepository) GetOne(id int64) (responses models.VehicleTransmissionResponse, err error) {
	return responses, VehicleTransmission.db.DB.Where("id = ?", id).Find(&responses).Error
}

// Store implements VehicleTransmissionDefinition
func (VehicleTransmission VehicleTransmissionRepository) Store(request *models.VehicleTransmissionRequest) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")
	fmt.Println("repo =", models.VehicleTransmissionRequest{
		Name:      request.Name,
		CreatedAt: &timeNow,
	})
	err = VehicleTransmission.db.DB.Save(&models.VehicleTransmission{
		Name:      request.Name,
		CreatedAt: &timeNow,
	}).Error
	fmt.Println(err)
	return true, err
}

// Update implements VehicleTransmissionDefinition
func (VehicleTransmission VehicleTransmissionRepository) Update(request *models.VehicleTransmissionRequest) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")
	return true, VehicleTransmission.db.DB.Save(&models.VehicleTransmissionRequest{
		ID:        request.ID,
		Name:      request.Name,
		CreatedAt: request.CreatedAt,
		UpdatedAt: &timeNow,
	}).Error
}

// Delete implements VehicleTransmissionDefinition
func (VehicleTransmission VehicleTransmissionRepository) Delete(id int64) (err error) {
	return VehicleTransmission.db.DB.Where("id = ?", id).Delete(&models.VehicleTransmissionResponse{}).Error
}
