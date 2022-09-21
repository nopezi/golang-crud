package riskcontrol

import (
	models "riskmanagement/models/riskcontrol"
	repository "riskmanagement/repository/riskcontrol"

	"gitlab.com/golang-package-library/logger"
)

type RiskControlDefinition interface {
	GetAll() (responses []models.RiskControlResponse, err error)
	GetOne(id int64) (responses models.RiskControlResponse, err error)
	Store(request *models.RiskControlRequest) (err error)
	Update(request *models.RiskControlRequest) (err error)
	Delete(id int64) (err error)
}

type RiskControlService struct {
	logger     logger.Logger
	repository repository.RiskControlDefinition
}

func NewRiskControService(
	logger logger.Logger,
	repository repository.RiskControlDefinition,
) RiskControlDefinition {
	return RiskControlService{
		logger:     logger,
		repository: repository,
	}
}

// Delete implements RiskControlDefinition
func (riskControl RiskControlService) Delete(id int64) (err error) {
	return riskControl.repository.Delete(id)
}

// GetAll implements RiskControlDefinition
func (riskControl RiskControlService) GetAll() (responses []models.RiskControlResponse, err error) {
	return riskControl.repository.GetAll()
}

// GetOne implements RiskControlDefinition
func (riskControl RiskControlService) GetOne(id int64) (responses models.RiskControlResponse, err error) {
	return riskControl.repository.GetOne(id)
}

// Store implements RiskControlDefinition
func (riskControl RiskControlService) Store(request *models.RiskControlRequest) (err error) {
	_, err = riskControl.repository.Store(request)
	return err
}

// Update implements RiskControlDefinition
func (riskControl RiskControlService) Update(request *models.RiskControlRequest) (err error) {
	_, err = riskControl.repository.Update(request)
	return err
}
