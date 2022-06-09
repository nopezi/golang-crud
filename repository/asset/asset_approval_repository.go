package asset

import (
	"infolelang/lib"
	models "infolelang/models/assets"
	"time"

	elastic "gitlab.com/golang-package-library/elasticsearch"
	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type AssetApprovalDefinition interface {
	GetAll() (responses []models.AssetApprovalsResponse, err error)
	GetOne(id int64) (responses models.AssetApprovalsResponse, err error)
	Store(request *models.AssetApprovalsRequest) (responses bool, err error)
	Update(request *models.AssetApprovalsRequest) (responses bool, err error)
	Delete(id int64) (err error)
	WithTrx(trxHandle *gorm.DB) AssetApprovalRepository
}
type AssetApprovalRepository struct {
	db      lib.Database
	db2     lib.Databases
	elastic elastic.Elasticsearch
	logger  logger.Logger
	timeout time.Duration
}

func NewAssetApprovalReporitory(
	db lib.Database,
	db2 lib.Databases,
	elastic elastic.Elasticsearch,
	logger logger.Logger) AssetApprovalDefinition {
	return AssetApprovalRepository{
		db:      db,
		db2:     db2,
		elastic: elastic,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// WithTrx implements AssetApprovalDefinition
func (AssetApproval AssetApprovalRepository) WithTrx(trxHandle *gorm.DB) AssetApprovalRepository {
	if trxHandle == nil {
		AssetApproval.logger.Zap.Error("transaction Database not found in gin context. ")
		return AssetApproval
	}
	AssetApproval.db.DB = trxHandle
	return AssetApproval
}

// GetAll implements AssetApprovalDefinition
func (assetApproval AssetApprovalRepository) GetAll() (responses []models.AssetApprovalsResponse, err error) {
	return responses, assetApproval.db.DB.Find(&responses).Error
}

// GetOne implements AssetApprovalDefinition
func (assetApproval AssetApprovalRepository) GetOne(id int64) (responses models.AssetApprovalsResponse, err error) {
	return responses, assetApproval.db.DB.Where("id = ?", id).Find(&responses).Error
}

// Store implements AssetApprovalDefinition
func (assetApproval AssetApprovalRepository) Store(request *models.AssetApprovalsRequest) (responses bool, err error) {
	return responses, assetApproval.db.DB.Save(&responses).Error
}

// Update implements AssetApprovalDefinition
func (assetApproval AssetApprovalRepository) Update(request *models.AssetApprovalsRequest) (responses bool, err error) {
	return true, assetApproval.db.DB.Save(&responses).Error
}

// Delete implements AssetApprovalDefinition
func (assetApproval AssetApprovalRepository) Delete(id int64) (err error) {
	return assetApproval.db.DB.Where("id = ?", id).Delete(&models.AssetApprovalsResponse{}).Error
}
