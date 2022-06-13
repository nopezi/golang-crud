package access_places

import (
	"fmt"
	"infolelang/lib"
	models "infolelang/models/categories"
	"time"

	elastic "gitlab.com/golang-package-library/elasticsearch"
	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type CategoryDefinition interface {
	GetAll() (responses []models.CategoryResponse, err error)
	GetOne(id int64) (responses models.CategoryResponse, err error)
	Store(request *models.CategoryRequest) (responses bool, err error)
	Update(request *models.CategoryRequest) (responses bool, err error)
	Delete(id int64) (err error)
	WithTrx(trxHandle *gorm.DB) CategoryRepository
}
type CategoryRepository struct {
	db      lib.Database
	db2     lib.Databases
	elastic elastic.Elasticsearch
	logger  logger.Logger
	timeout time.Duration
}

func NewCategoryReporitory(
	db lib.Database,
	db2 lib.Databases,
	elastic elastic.Elasticsearch,
	logger logger.Logger) CategoryDefinition {
	return CategoryRepository{
		db:      db,
		db2:     db2,
		elastic: elastic,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// WithTrx implements CategoryDefinition
func (category CategoryRepository) WithTrx(trxHandle *gorm.DB) CategoryRepository {
	if trxHandle == nil {
		category.logger.Zap.Error("transaction Database not found in gin context. ")
		return category
	}
	category.db.DB = trxHandle
	return category
}

// GetAll implements CategoryDefinition
func (category CategoryRepository) GetAll() (responses []models.CategoryResponse, err error) {
	return responses, category.db.DB.Where("status = 1").Find(&responses).Error
}

// GetOne implements CategoryDefinition
func (category CategoryRepository) GetOne(id int64) (responses models.CategoryResponse, err error) {
	return responses, category.db.DB.Where("id = ?", id).Find(&responses).Error
}

// Store implements CategoryDefinition
func (category CategoryRepository) Store(request *models.CategoryRequest) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")
	fmt.Println("repo =", models.CategoryRequest{
		Name:      request.Name,
		CreatedAt: &timeNow,
	})
	err = category.db.DB.Save(&models.Categories{
		Name:      request.Name,
		CreatedAt: &timeNow,
	}).Error
	fmt.Println(err)
	return true, err
}

// Update implements CategoryDefinition
func (category CategoryRepository) Update(request *models.CategoryRequest) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")
	return true, category.db.DB.Save(&models.CategoryRequest{
		ID:        request.ID,
		Name:      request.Name,
		CreatedAt: request.CreatedAt,
		UpdatedAt: &timeNow,
	}).Error
}

// Delete implements CategoryDefinition
func (category CategoryRepository) Delete(id int64) (err error) {
	return category.db.DB.Where("id = ?", id).Delete(&models.CategoryResponse{}).Error
}
