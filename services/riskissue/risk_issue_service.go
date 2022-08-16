package riskissue

import (
	"fmt"
	models "riskmanagement/models/riskissue"
	repository "riskmanagement/repository/riskissue"

	"gitlab.com/golang-package-library/logger"
)

type RiskIssueDefinition interface {
	GetAll() (responses []models.RiskIssueResponse, err error)
	GetOne(id int64) (responses models.RiskIssueResponse, err error)
	Store(request *models.RiskIssueRequest) (err error)
	Update(request *models.RiskIssueRequest) (err error)
	Delete(id int64) (err error)
}

type RiskIssueService struct {
	logger     logger.Logger
	repository repository.RiskIssueDefinition
}

// Delete implements RiskIssueDefinition
func (riskIssue RiskIssueService) Delete(id int64) (err error) {
	return riskIssue.repository.Delete(id)
}

// GetAll implements RiskIssueDefinition
func (riskIssue RiskIssueService) GetAll() (responses []models.RiskIssueResponse, err error) {
	return riskIssue.repository.GetAll()
}

// GetOne implements RiskIssueDefinition
func (riskIssue RiskIssueService) GetOne(id int64) (responses models.RiskIssueResponse, err error) {
	return riskIssue.repository.GetOne(id)
}

// Store implements RiskIssueDefinition
func (riskIssue RiskIssueService) Store(request *models.RiskIssueRequest) (err error) {
	fmt.Println("service = ", request)
	_, err = riskIssue.repository.Store(request)
	return err
}

// Update implements RiskIssueDefinition
func (riskIssue RiskIssueService) Update(request *models.RiskIssueRequest) (err error) {
	_, err = riskIssue.repository.Update(request)
	return err
}

func NewRiskIssueService(
	logger logger.Logger,
	repository repository.RiskIssueDefinition,
) RiskIssueDefinition {
	return RiskIssueService{
		logger:     logger,
		repository: repository,
	}
}
