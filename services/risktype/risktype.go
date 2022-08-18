package risktype

import (
	"fmt"
	models "riskmanagement/models/risktype"
	repository "riskmanagement/repository/risktype"

	"gitlab.com/golang-package-library/logger"
)

type RiskTypeDefinition interface {
	GetAll() (responses []models.RiskTypeResponse, err error)
	GetOne(id int64) (responses models.RiskTypeResponse, err error)
	Store(request *models.RiskTypeRequest) (err error)
	Update(request *models.RiskTypeRequest) (err error)
	Delete(id int64) (err error)
}

type RiskTypeService struct {
	logger     logger.Logger
	repository repository.RiskTypeDefinition
}

// Delete implements RiskIssueDefinition
func (riskType RiskTypeService) Delete(id int64) (err error) {
	return riskType.repository.Delete(id)
}

// GetAll implements RiskIssueDefinition
func (riskType RiskTypeService) GetAll() (responses []models.RiskTypeResponse, err error) {
	return riskType.repository.GetAll()
}

// GetOne implements RiskIssueDefinition
func (riskType RiskTypeService) GetOne(id int64) (responses models.RiskTypeResponse, err error) {
	return riskType.repository.GetOne(id)
}

// Store implements RiskIssueDefinition
func (riskType RiskTypeService) Store(request *models.RiskTypeRequest) (err error) {
	fmt.Println("service = ", request)
	_, err = riskType.repository.Store(request)
	return err
}

// Update implements RiskIssueDefinition
func (riskType RiskTypeService) Update(request *models.RiskTypeRequest) (err error) {
	_, err = riskType.repository.Update(request)
	return err
}

func NewRiskTypeService(
	logger logger.Logger,
	repository repository.RiskTypeDefinition,
) RiskTypeDefinition {
	return RiskTypeService{
		logger:     logger,
		repository: repository,
	}
}
