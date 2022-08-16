package product

import (
	"riskmanagement/lib"
	models "riskmanagement/models/product"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type ProductDefinition interface {
	GetAll() (responses []models.ProductResponse, err error)
	GetOne(id int64) (responses models.ProductResponse, err error)
	Store(request *models.ProductRequest) (responses bool, err error)
	Update(request *models.ProductRequest) (responese bool, err error)
	Delete(id int64) (err error)
	WithTrx(trxHandle *gorm.DB) ProductRepository
}

type ProductRepository struct {
	db      lib.Database
	dbRaw   lib.Database
	logger  logger.Logger
	timeout time.Duration
}

// Delete implements ProductDefinition
func (product ProductRepository) Delete(id int64) (err error) {
	return product.db.DB.Where("id = ?", id).Delete(&models.ProductResponse{}).Error
}

// GetAll implements ProductDefinition
func (product ProductRepository) GetAll() (responses []models.ProductResponse, err error) {
	return responses, product.db.DB.Find(&responses).Error
}

// GetOne implements ProductDefinition
func (product ProductRepository) GetOne(id int64) (responses models.ProductResponse, err error) {
	return responses, product.db.DB.Where("id = ?", id).Find(&responses).Error
}

// Store implements ProductDefinition
func (product ProductRepository) Store(request *models.ProductRequest) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")
	return responses, product.db.DB.Save(&models.ProductRequest{
		Name:      request.Name,
		CreatedAt: &timeNow,
	}).Error
}

// Update implements ProductDefinition
func (product ProductRepository) Update(request *models.ProductRequest) (responese bool, err error) {
	timeNow := lib.GetTimeNow("timestime")
	return true, product.db.DB.Save(&models.ProductRequest{
		ID:        request.ID,
		Name:      request.Name,
		CreatedAt: request.CreatedAt,
		UpdatedAt: &timeNow,
	}).Error
}

// WithTrx implements ProductDefinition
func (product ProductRepository) WithTrx(trxHandle *gorm.DB) ProductRepository {
	if trxHandle == nil {
		product.logger.Zap.Error("transaction Database not found in gin context")
		return product
	}

	product.db.DB = trxHandle
	return product
}

func NewProductRepository(
	db lib.Database,
	dbRaw lib.Database,
	logger logger.Logger,
) ProductDefinition {
	return ProductRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: time.Second * 100,
	}
}
