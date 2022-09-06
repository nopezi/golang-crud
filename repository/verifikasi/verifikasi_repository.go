package verifikasi

import (
	"riskmanagement/lib"
	models "riskmanagement/models/verifikasi"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type VerifikasiDefinition interface {
	WithTrx(trxHandle *gorm.DB) VerifikasiRepository
	GetAll() (responses []models.VerifikasiResponse, err error)
	GetOne(id int64) (responses models.VerifikasiResponse, err error)
	Store(request *models.Verifikasi, tx *gorm.DB) (responses *models.Verifikasi, err error)
	Delete(request *models.VerifikasiUpdateDelete, include []string, tx *gorm.DB) (responses bool, err error)
	DeleteAnomaliData(id int64, tx *gorm.DB) (err error)
	DeleteLampiranVerifikasi(id int64, tx *gorm.DB) (err error)
	KonfirmSave(request *models.VerifikasiUpdateMaintain, include []string, tx *gorm.DB) (response bool, err error)
	UpdateAllVerifikasi(request *models.VerifikasiUpdateAll, include []string, tx *gorm.DB) (response bool, err error)
}

type VerifikasiRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

func NewVerfikasiRepository(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger,
) VerifikasiDefinition {
	return VerifikasiRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// Delete implements VerifikasiDefinition
func (verifikasi VerifikasiRepository) Delete(request *models.VerifikasiUpdateDelete, include []string, tx *gorm.DB) (responses bool, err error) {
	return true, tx.Save(&request).Error
}

// DeleteAnomaliData implements VerifikasiDefinition
func (verifikasi VerifikasiRepository) DeleteAnomaliData(id int64, tx *gorm.DB) (err error) {
	return tx.Where("id = ?", id).Delete(&models.VerifikasiAnomaliDataRequest{}).Error
}

// GetAll implements VerifikasiDefinition
func (verifikasi VerifikasiRepository) GetAll() (responses []models.VerifikasiResponse, err error) {
	return responses, verifikasi.db.DB.Find(&responses).Error
}

// GetOne implements VerifikasiDefinition
func (verifikasi VerifikasiRepository) GetOne(id int64) (responses models.VerifikasiResponse, err error) {
	err = verifikasi.db.DB.Raw(`
		SELECT 
			verif.*
		FROM verifikasi verif 
		WHERE verif.id = ?`, id).Find(&responses).Error

	if err != nil {
		verifikasi.logger.Zap.Error(err)
		return responses, err
	}
	return responses, err
}

// Store implements VerifikasiDefinition
func (verifikasi VerifikasiRepository) Store(request *models.Verifikasi, tx *gorm.DB) (responses *models.Verifikasi, err error) {
	return request, tx.Save(&request).Error
}

// WithTrx implements VerifikasiDefinition
func (verifikasi VerifikasiRepository) WithTrx(trxHandle *gorm.DB) VerifikasiRepository {
	if trxHandle == nil {
		verifikasi.logger.Zap.Error("transaction Database not found in gin context.")
		return verifikasi
	}

	verifikasi.db.DB = trxHandle
	return verifikasi
}

// DeleteLampiranVerifikasi implements VerifikasiDefinition
func (VerifikasiRepository) DeleteLampiranVerifikasi(id int64, tx *gorm.DB) (err error) {
	return tx.Where("id = ?", id).Delete(&models.VerifikasiFilesRequest{}).Error
}

// KonfirmSave implements VerifikasiDefinition
func (VerifikasiRepository) KonfirmSave(request *models.VerifikasiUpdateMaintain, include []string, tx *gorm.DB) (response bool, err error) {
	return true, tx.Save(&request).Error
}

// UpdateAllVerifikasi implements VerifikasiDefinition
func (VerifikasiRepository) UpdateAllVerifikasi(request *models.VerifikasiUpdateAll, include []string, tx *gorm.DB) (response bool, err error) {
	return true, tx.Save(&request).Error
}
