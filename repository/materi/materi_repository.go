package materi

import (
	"fmt"
	"riskmanagement/lib"
	models "riskmanagement/models/materi"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type MateriDefinition interface {
	// GetAll() (responses []models.Materi, err error)
	GetAll() (responses []models.MateriAllResponse, err error)
	GetAllMateriFiles(materiID int64) (responses []models.MateriFilesResponse, err error)
	Store(request *models.Materi, tx *gorm.DB) (responses *models.Materi, err error)
	StoreMateriFiles(request *models.MateriRequest, tx *gorm.DB) (responses bool, err error)
	Delete(request models.MateriFilesRequest) (status bool, err error)
	WithTrx(trxHandle *gorm.DB) MateriRepository
}

type MateriRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

func NewMateriRepository(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger,
) MateriDefinition {
	return MateriRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// Delete implements MaterDefinition
func (materi MateriRepository) Delete(request models.MateriFilesRequest) (status bool, err error) {
	err = materi.db.DB.Where("files_id = ?", request.FilesID).Delete(&models.MateriFiles{}).Error
	if err != nil {
		materi.logger.Zap.Error(err)
		return false, err
	}
	return true, err
}

// GetAll implements MaterDefinition
// func (materi MateriRepository) GetAll() (responses []models.Materi, err error) {
// 	return responses, materi.db.DB.Find(&responses).Error
// }

func (materi MateriRepository) GetAll() (responses []models.MateriAllResponse, err error) {
	return responses, materi.db.DB.Find(&responses).Error
}

// GetAllMateriFiles implements MaterDefinition
func (materi MateriRepository) GetAllMateriFiles(materiID int64) (responses []models.MateriFilesResponse, err error) {
	rows, err := materi.db.DB.Raw(`
				SELECT 
					mf.id materi_files_id,
					mf.materi_id, 
					mf.files_id, f.filename, 
					f.path ,f.extension,f.size  
				FROM materi_files mf
				LEFT JOIN files f on mf.files_id = f.id 
				LEFT JOIN materi m on mf.materi_id = m.id
				WHERE m.id = ? `, materiID).Rows()

	defer rows.Close()
	var materiFiles models.MateriFilesResponse
	for rows.Next() {
		materi.db.DB.ScanRows(rows, &materiFiles)
		responses = append(responses, materiFiles)
	}

	return responses, err
}

// Store implements MaterDefinition
func (materi MateriRepository) Store(request *models.Materi, tx *gorm.DB) (responses *models.Materi, err error) {
	return request, tx.Save(&request).Error
}

// StoreMateriFiles implements MaterDefinition
func (materi MateriRepository) StoreMateriFiles(request *models.MateriRequest, tx *gorm.DB) (responses bool, err error) {
	err = tx.Save(&models.MateriFiles{
		MateriID: request.MateriID,
		FilesID:  request.FilesID,
	}).Error
	fmt.Println(err)
	return true, err
}

// WithTrx implements MaterDefinition
func (materi MateriRepository) WithTrx(trxHandle *gorm.DB) MateriRepository {
	if trxHandle == nil {
		materi.logger.Zap.Error("transaction Database not found in gin context")
		return materi
	}

	materi.db.DB = trxHandle
	return materi
}
