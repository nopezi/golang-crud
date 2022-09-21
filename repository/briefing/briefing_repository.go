package briefing

import (
	"riskmanagement/lib"
	models "riskmanagement/models/briefing"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type BriefingDefinition interface {
	WithTrx(trxHandle *gorm.DB) BriefingRepository
	GetAll() (responses []models.BriefingResponse, err error)
	GetData() (responses []models.BriefingResponseData, err error)
	GetOne(id int64) (responses models.BriefingResponse, err error)
	Store(request *models.Briefing, tx *gorm.DB) (responses *models.Briefing, err error)
	Delete(request *models.BriefingUpdateDelete, include []string, tx *gorm.DB) (responses bool, err error)
	DeleteBriefingMateri(id int64, tx *gorm.DB) (err error)
	UpdateAllBrief(request *models.BriefingUpdateMateri, include []string, tx *gorm.DB) (responses bool, err error)
}

type BriefingRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

// GetData implements BriefingDefinition
func (briefing BriefingRepository) GetData() (responses []models.BriefingResponseData, err error) {
	rows, err := briefing.db.DB.Raw(`
		SELECT 
			brf.id 'id',
			brf.no_pelaporan 'no_pelaporan',
			brf.unit_kerja 'unit_kerja',
			bm.judul_materi 'judul_materi',
			act.name 'aktifitas',
			CASE
				WHEN brf.status = "01a" && brf.action = "Draft" THEN "Draft"
				WHEN brf.status = "02b" && (brf.action = "Update" || brf.action ="Selesai")   THEN "Selesai"
				ELSE "Delete"
			END 'status_brf'
		FROM briefing brf 
		JOIN briefing_materis bm ON bm.briefing_id = brf.id
		JOIN activity act ON bm.activity_id = act.id
		WHERE brf.deleted != 1
		GROUP BY brf.id
	`).Rows()

	defer rows.Close()

	var brief models.BriefingResponseData
	for rows.Next() {
		briefing.db.DB.ScanRows(rows, &brief)
		responses = append(responses, brief)
	}
	return responses, err
}

// UpdateAllBrief implements BriefingDefinition
func (briefing BriefingRepository) UpdateAllBrief(request *models.BriefingUpdateMateri, include []string, tx *gorm.DB) (responses bool, err error) {
	return true, tx.Save(&request).Error
}

func NewBriefingRepository(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger,
) BriefingDefinition {
	return BriefingRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// Delete implements BriefingDefinition
func (briefing BriefingRepository) Delete(request *models.BriefingUpdateDelete, include []string, tx *gorm.DB) (responses bool, err error) {
	return true, tx.Save(&request).Error
}

// DeleteBriefingMateri implements BriefingDefinition
func (briefing BriefingRepository) DeleteBriefingMateri(id int64, tx *gorm.DB) (err error) {
	return tx.Where("id = ?", id).Delete(&models.BriefingMateriRequest{}).Error
}

// GetAll implements BriefingDefinition
func (briefing BriefingRepository) GetAll() (responses []models.BriefingResponse, err error) {
	return responses, briefing.db.DB.Find(&responses).Error
}

// GetOne implements BriefingDefinition
func (briefing BriefingRepository) GetOne(id int64) (responses models.BriefingResponse, err error) {
	err = briefing.db.DB.Raw(`
	SELECT 
		brf.id,
		brf.no_pelaporan,
		brf.unit_kerja,
		brf.peserta,
		brf.jumlah_peserta,
		brf.maker_id,
		brf.maker_desc,
		brf.maker_date,
		brf.last_maker_id,
		brf.last_maker_desc,
		brf.last_maker_date,
		brf.status,
		brf.action,
		brf.deleted,
		brf.created_at,
		brf.updated_at
	FROM briefing brf
	WHERE brf.id = ?`, id).Find(&responses).Error

	if err != nil {
		briefing.logger.Zap.Error(err)
		return responses, err
	}
	return responses, err
}

// Store implements BriefingDefinition
func (briefing BriefingRepository) Store(request *models.Briefing, tx *gorm.DB) (responses *models.Briefing, err error) {
	return request, tx.Save(&request).Error
}

// WithTrx implements BriefingDefinition
func (briefing BriefingRepository) WithTrx(trxHandle *gorm.DB) BriefingRepository {
	if trxHandle == nil {
		briefing.logger.Zap.Error("transaction Database not found in gin context")
		return briefing
	}
	briefing.db.DB = trxHandle
	return briefing
}
