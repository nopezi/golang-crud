package unitkerja

import (
	"fmt"
	models "riskmanagement/models/unitkerja"
	repository "riskmanagement/repository/unitkerja"

	"gitlab.com/golang-package-library/logger"
)

type UnitKerjaDefinition interface {
	GetAll() (responses []models.UnitKerjaResponse, err error)
	GetOne(id int64) (responses models.UnitKerjaResponse, err error)
	Store(request *models.UnitKerjaRequest) (err error)
	Update(request *models.UnitKerjaRequest) (err error)
	Delete(id int64) (err error)
}

type UnitKerjaService struct {
	logger     logger.Logger
	repository repository.UnitKerjaDefinition
}

func NewUnitKerjaService(
	logger logger.Logger,
	repository repository.UnitKerjaDefinition,
) UnitKerjaDefinition {
	return UnitKerjaService{
		logger:     logger,
		repository: repository,
	}
}

// Delete implements UnitKerjaDefinition
func (unitKerja UnitKerjaService) Delete(id int64) (err error) {
	return unitKerja.repository.Delete(id)
}

// GetAll implements UnitKerjaDefinition
func (unitKerja UnitKerjaService) GetAll() (responses []models.UnitKerjaResponse, err error) {
	return unitKerja.repository.GetAll()
}

// GetOne implements UnitKerjaDefinition
func (unitKerja UnitKerjaService) GetOne(id int64) (responses models.UnitKerjaResponse, err error) {
	return unitKerja.repository.GetOne(id)
}

// Store implements UnitKerjaDefinition
func (unitKerja UnitKerjaService) Store(request *models.UnitKerjaRequest) (err error) {
	fmt.Println("service =", request)
	_, err = unitKerja.repository.Store(request)
	return err
}

// Update implements UnitKerjaDefinition
func (unitKerja UnitKerjaService) Update(request *models.UnitKerjaRequest) (err error) {
	_, err = unitKerja.repository.Update(request)
	return err
}
