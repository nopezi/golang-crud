package riskindicator

import (
	"fmt"
	models "riskmanagement/models/riskindicator"
	repository "riskmanagement/repository/riskindicator"

	"gitlab.com/golang-package-library/logger"
)

type RiskIndicatorDefinition interface {
	GetAll() (responses []models.RiskIndicatorResponse, err error)
	GetOne(id int64) (responses models.RiskIndicatorResponse, err error)
	Store(request *models.RiskIndicatorRequest) (err error)
	Update(request *models.RiskIndicatorRequest) (err error)
	Delete(id int64) (err error)
}

type RiskIndicatorService struct {
	logger     logger.Logger
	repository repository.RiskIndicatorDefinition
}

// Delete implements RiskIndicatorDefinition
func (riskIndicator RiskIndicatorService) Delete(id int64) (err error) {
	return riskIndicator.repository.Delete(id)
}

// GetAll implements RiskIndicatorDefinition
func (riskIndicator RiskIndicatorService) GetAll() (responses []models.RiskIndicatorResponse, err error) {
	return riskIndicator.repository.GetAll()
}

// GetOne implements RiskIndicatorDefinition
func (riskIndicator RiskIndicatorService) GetOne(id int64) (responses models.RiskIndicatorResponse, err error) {
	return riskIndicator.repository.GetOne(id)
}

// Store implements RiskIndicatorDefinition
func (riskIndicator RiskIndicatorService) Store(request *models.RiskIndicatorRequest) (err error) {
	fmt.Println("service = ", request)
	_, err = riskIndicator.repository.Store(request)
	return err
}

// Update implements RiskIndicatorDefinition
func (riskIndicator RiskIndicatorService) Update(request *models.RiskIndicatorRequest) (err error) {
	_, err = riskIndicator.repository.Update(request)
	return err
}

func NewRiskIndicatorService(
	logger logger.Logger,
	repository repository.RiskIndicatorDefinition,
) RiskIndicatorDefinition {
	return RiskIndicatorService{
		logger:     logger,
		repository: repository,
	}
}
