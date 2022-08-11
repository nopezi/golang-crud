package access_places

import (
	"fmt"
	"infolelang/lib"
	models "infolelang/models/vehicle_color"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type VehicleColorDefinition interface {
	GetAll() (responses []models.VehicleColorResponse, err error)
	GetOne(id int64) (responses models.VehicleColorResponse, err error)
	Store(request *models.VehicleColorRequest) (responses bool, err error)
	Update(request *models.VehicleColorRequest) (responses bool, err error)
	Delete(id int64) (err error)
	WithTrx(trxHandle *gorm.DB) VehicleColorRepository
}
type VehicleColorRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

func NewVehicleColorReporitory(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger) VehicleColorDefinition {
	return VehicleColorRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// WithTrx implements VehicleColorDefinition
func (VehicleColor VehicleColorRepository) WithTrx(trxHandle *gorm.DB) VehicleColorRepository {
	if trxHandle == nil {
		VehicleColor.logger.Zap.Error("transaction Database not found in gin context. ")
		return VehicleColor
	}
	VehicleColor.db.DB = trxHandle
	return VehicleColor
}

// GetAll implements VehicleColorDefinition
func (VehicleColor VehicleColorRepository) GetAll() (responses []models.VehicleColorResponse, err error) {
	return responses, VehicleColor.db.DB.Where("status = 1").Find(&responses).Error
}

// GetOne implements VehicleColorDefinition
func (VehicleColor VehicleColorRepository) GetOne(id int64) (responses models.VehicleColorResponse, err error) {
	return responses, VehicleColor.db.DB.Where("id = ?", id).Find(&responses).Error
}

// Store implements VehicleColorDefinition
func (VehicleColor VehicleColorRepository) Store(request *models.VehicleColorRequest) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")
	fmt.Println("repo =", models.VehicleColorRequest{
		Name:      request.Name,
		CreatedAt: &timeNow,
	})
	err = VehicleColor.db.DB.Save(&models.VehicleColor{
		Name:      request.Name,
		CreatedAt: &timeNow,
	}).Error
	fmt.Println(err)
	return true, err
}

// Update implements VehicleColorDefinition
func (VehicleColor VehicleColorRepository) Update(request *models.VehicleColorRequest) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")
	return true, VehicleColor.db.DB.Save(&models.VehicleColorRequest{
		ID:        request.ID,
		Name:      request.Name,
		CreatedAt: request.CreatedAt,
		UpdatedAt: &timeNow,
	}).Error
}

// Delete implements VehicleColorDefinition
func (VehicleColor VehicleColorRepository) Delete(id int64) (err error) {
	return VehicleColor.db.DB.Where("id = ?", id).Delete(&models.VehicleColorResponse{}).Error
}
