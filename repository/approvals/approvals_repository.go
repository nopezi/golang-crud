package approvals

import (
	"infolelang/lib"
	models "infolelang/models/approvals"
	"time"

	elastic "gitlab.com/golang-package-library/elasticsearch"
	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type ApprovalDefinition interface {
	GetAll() (responses []models.ApprovalsResponse, err error)
	GetOne(id int64) (responses models.ApprovalsResponse, err error)
	GetOneAsset(id int64) (responses models.ApprovalsResponse, err error)
	Store(request *models.Approvals) (responses *models.Approvals, err error)
	Update(request *models.ApprovalsRequest) (responses bool, err error)
	Delete(id int64) (err error)
	DeleteApprovals(id int64) (err error)
	WithTrx(trxHandle *gorm.DB) ApprovalRepository
}
type ApprovalRepository struct {
	db      lib.Database
	db2     lib.Databases
	elastic elastic.Elasticsearch
	logger  logger.Logger
	timeout time.Duration
}

func NewApprovalReporitory(
	db lib.Database,
	db2 lib.Databases,
	elastic elastic.Elasticsearch,
	logger logger.Logger) ApprovalDefinition {
	return ApprovalRepository{
		db:      db,
		db2:     db2,
		elastic: elastic,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// WithTrx implements ApprovalDefinition
func (Approval ApprovalRepository) WithTrx(trxHandle *gorm.DB) ApprovalRepository {
	if trxHandle == nil {
		Approval.logger.Zap.Error("transaction Database not found in gin context. ")
		return Approval
	}
	Approval.db.DB = trxHandle
	return Approval
}

// GetAll implements ApprovalDefinition
func (approval ApprovalRepository) GetAll() (responses []models.ApprovalsResponse, err error) {
	return responses, approval.db.DB.Find(&responses).Error
}

// GetOne implements ApprovalDefinition
func (approval ApprovalRepository) GetOne(id int64) (responses models.ApprovalsResponse, err error) {
	return responses, approval.db.DB.Where("id = ?", id).Find(&responses).Error
}

// GetOneAsset implements ApprovalDefinition
func (approval ApprovalRepository) GetOneAsset(id int64) (responses models.ApprovalsResponse, err error) {
	return responses, approval.db.DB.Where("asset_id = ?", id).Find(&responses).Error
}

// Store implements ApprovalDefinition
func (approval ApprovalRepository) Store(request *models.Approvals) (responses *models.Approvals, err error) {
	return request, approval.db.DB.Save(&request).Error
}

// Update implements ApprovalDefinition
func (approval ApprovalRepository) Update(request *models.ApprovalsRequest) (responses bool, err error) {
	return true, approval.db.DB.Save(&request).Error
}

// Delete implements ApprovalDefinition
func (approval ApprovalRepository) Delete(id int64) (err error) {
	return approval.db.DB.Where("id = ?", id).Delete(&models.ApprovalsResponse{}).Error
}

// Delete implements ApprovalDefinition
func (approval ApprovalRepository) DeleteApprovals(id int64) (err error) {
	return approval.db.DB.Delete(&models.ApprovalsResponse{}).Error
}
