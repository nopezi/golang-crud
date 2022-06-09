package access_places

import (
	"infolelang/lib"
	models "infolelang/models/sub_categories"
	"time"

	elastic "gitlab.com/golang-package-library/elasticsearch"
	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type SubCategoryDefinition interface {
	GetAll() (responses []models.SubCategoriesResponse, err error)
	GetOne(id int64) (responses models.SubCategoriesResponse, err error)
	Store(request *models.SubCategoriesRequest) (responses bool, err error)
	Update(request *models.SubCategoriesRequest) (responses bool, err error)
	Delete(id int64) (err error)
	WithTrx(trxHandle *gorm.DB) SubCategoryRepository
}
type SubCategoryRepository struct {
	db      lib.Database
	db2     lib.Databases
	elastic elastic.Elasticsearch
	logger  logger.Logger
	timeout time.Duration
}

func NewSubCategoryReporitory(
	db lib.Database,
	db2 lib.Databases,
	elastic elastic.Elasticsearch,
	logger logger.Logger) SubCategoryDefinition {
	return SubCategoryRepository{
		db:      db,
		db2:     db2,
		elastic: elastic,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// WithTrx implements SubCategoryDefinition
func (subCategory SubCategoryRepository) WithTrx(trxHandle *gorm.DB) SubCategoryRepository {
	if trxHandle == nil {
		subCategory.logger.Zap.Error("transaction Database not found in gin context. ")
		return subCategory
	}
	subCategory.db.DB = trxHandle
	return subCategory
}

// GetAll implements SubCategoryDefinition
func (subCategory SubCategoryRepository) GetAll() (responses []models.SubCategoriesResponse, err error) {
	return responses, subCategory.db.DB.Find(&responses).Error
}

// GetOne implements SubCategoryDefinition
func (subCategory SubCategoryRepository) GetOne(id int64) (responses models.SubCategoriesResponse, err error) {
	return responses, subCategory.db.DB.Where("id = ?", id).Find(&responses).Error
}

// Store implements SubCategoryDefinition
func (subCategory SubCategoryRepository) Store(request *models.SubCategoriesRequest) (responses bool, err error) {
	return responses, subCategory.db.DB.Save(&responses).Error
}

// Update implements SubCategoryDefinition
func (subCategory SubCategoryRepository) Update(request *models.SubCategoriesRequest) (responses bool, err error) {
	return true, subCategory.db.DB.Save(&responses).Error
}

// Delete implements SubCategoryDefinition
func (subCategory SubCategoryRepository) Delete(id int64) (err error) {
	return subCategory.db.DB.Where("id = ?", id).Delete(&models.SubCategoriesResponse{}).Error
}
