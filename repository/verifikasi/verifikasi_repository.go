package verifikasi

import (
	"math"
	"riskmanagement/lib"
	models "riskmanagement/models/verifikasi"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type VerifikasiDefinition interface {
	WithTrx(trxHandle *gorm.DB) VerifikasiRepository
	GetAll() (responses []models.VerifikasiResponse, err error)
	GetListData() (responses []models.VerifikasiList, err error)
	GetOne(id int64) (responses models.VerifikasiResponse, err error)
	FilterVerifikasi(request *models.VerifikasiFilterRequest) (responses []models.VerifikasiListFilter, totalRows int, totalData int, err error)
	Store(request *models.Verifikasi, tx *gorm.DB) (responses *models.Verifikasi, err error)
	Delete(request *models.VerifikasiUpdateDelete, include []string, tx *gorm.DB) (responses bool, err error)
	DeleteAnomaliData(id int64, tx *gorm.DB) (err error)
	DeleteLampiranVerifikasi(id int64, tx *gorm.DB) (err error)
	KonfirmSave(request *models.VerifikasiUpdateMaintain, include []string, tx *gorm.DB) (response bool, err error)
	UpdateAllVerifikasi(request *models.VerifikasiUpdateAll, include []string, tx *gorm.DB) (response bool, err error)
	GetNoPelaporan(request *models.NoPalaporanRequest) (responses []models.NoPelaporanNullResponse, err error)
	GetLastID() (responses []models.VerifikasiLastID, err error)
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

// GetListData implements VerifikasiDefinition
func (verifikasi VerifikasiRepository) GetListData() (responses []models.VerifikasiList, err error) {
	rows, err := verifikasi.db.DB.Raw(`
		SELECT
			verif.id 'id',
			verif.no_pelaporan 'no_pelaporan',
			verif.unit_kerja 'unit_kerja',
			act.name 'aktifitas',
			CASE
				WHEN verif.status = "01a" && verif.action = "Draft" THEN "Draft"
				WHEN verif.status = "02b" && (verif.action = "Update" || verif.action ="Selesai")   THEN "Selesai"
				ELSE "Delete"
			END 'status_verif'
		FROM verifikasi verif
		JOIN activity act on verif.activity_id = act.id
		WHERE verif.deleted != 1
		GROUP BY verif.id
	`).Rows()

	defer rows.Close()
	var verif models.VerifikasiList
	for rows.Next() {
		verifikasi.db.DB.ScanRows(rows, &verif)
		responses = append(responses, verif)
	}

	return responses, err
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
func (verifikasi VerifikasiRepository) DeleteLampiranVerifikasi(id int64, tx *gorm.DB) (err error) {
	return tx.Where("id = ?", id).Delete(&models.VerifikasiFilesRequest{}).Error
}

// KonfirmSave implements VerifikasiDefinition
func (verifikasi VerifikasiRepository) KonfirmSave(request *models.VerifikasiUpdateMaintain, include []string, tx *gorm.DB) (response bool, err error) {
	return true, tx.Save(&request).Error
}

// UpdateAllVerifikasi implements VerifikasiDefinition
func (verifikasi VerifikasiRepository) UpdateAllVerifikasi(request *models.VerifikasiUpdateAll, include []string, tx *gorm.DB) (response bool, err error) {
	return true, tx.Save(&request).Error
}

// FilterVerifikasi implements VerifikasiDefinition
func (verifikasi VerifikasiRepository) FilterVerifikasi(request *models.VerifikasiFilterRequest) (responses []models.VerifikasiListFilter, totalRows int, totalData int, err error) {
	where := " WHERE verif.deleted != 1"
	whereCount := " WHERE verif.deleted != 1"

	if request.NoPelaporan != "" {
		where += " AND verif.no_pelaporan = '" + request.NoPelaporan + "'"
		whereCount += " AND verif.no_pelaporan = '" + request.NoPelaporan + "'"
	}

	if request.UnitKerja != "" {
		where += " AND verif.unit_kerja = '" + request.UnitKerja + "'"
		whereCount += " AND verif.unit_kerja = '" + request.UnitKerja + "'"
	}

	if request.ActivityID != "" {
		where += " AND verif.activity_id = '" + request.ActivityID + "'"
		whereCount += " AND verif.activity_id = '" + request.ActivityID + "'"
	}

	if request.RiskIssueID != "" {
		where += " AND verif.risk_issue_id = '" + request.RiskIssueID + "'"
		whereCount += " AND verif.risk_issue_id = '" + request.RiskIssueID + "'"
	}

	if request.Status != "" && request.Status != "Semua" && request.Status != "Selesai" {
		where += " AND verif.action = '" + request.Status + "'"
		whereCount += " AND verif.action = '" + request.Status + "'"
	}

	if request.Status == "Selesai" {
		where += " AND verif.status = '02b' AND (verif.action = 'Update' || verif.action = 'Selesai')"
	}

	if request.TglAwal != "" && request.TglAkhir != "" {
		where += " AND CAST(created_at as date) BETWEEN '" + request.TglAwal + "' AND '" + request.TglAkhir + "'"
		whereCount += " AND CAST(created_at as date) BETWEEN '" + request.TglAwal + "' AND '" + request.TglAkhir + "'"

		// where += " AND CAST(created_at as date) >= '" + request.TglAwal + "' AND AND CAST(created_at as date) <= '" + request.TglAkhir + "'"
		// whereCount += " AND CAST(created_at as date) >= '" + request.TglAwal + "' AND AND CAST(created_at as date) <= '" + request.TglAkhir + "'"
	}

	query := `SELECT
				verif.id 'id',
				verif.no_pelaporan 'no_pelaporan',
				verif.unit_kerja 'unit_kerja',
				act.name 'aktifitas',
				CASE
					WHEN verif.status = "01a" && verif.action = "Draft" THEN "Draft"
					WHEN verif.status = "02b" && (verif.action = "Update" || verif.action ="Selesai")   THEN "Selesai"
					ELSE "Delete"
				END 'status_verif'
			FROM verifikasi verif
			JOIN activity act on verif.activity_id = act.id
			` + where + ` ORDER BY id desc LIMIT ? OFFSET ?`

	verifikasi.logger.Zap.Info(query)
	rows, err := verifikasi.dbRaw.DB.Query(query, request.Limit, request.Offset)

	verifikasi.logger.Zap.Info("rows ", rows)
	if err != nil {
		return responses, totalRows, totalData, err
	}

	response := models.VerifikasiListFilter{}
	for rows.Next() {
		_ = rows.Scan(
			&response.ID,
			&response.NoPelaporan,
			&response.UnitKerja,
			&response.Aktifitas,
			&response.StatusVerif,
		)
		responses = append(responses, response)
	}

	if err = rows.Err(); err != nil {
		return responses, totalRows, totalData, err
	}

	paginateQuery := `SELECT count(*) FROM verifikasi verif
					JOIN activity act on verif.activity_id = act.id` + whereCount + ` GROUP BY verif.id`

	err = verifikasi.dbRaw.DB.QueryRow(paginateQuery).Scan(&totalRows)

	result := float64(totalRows) / float64(request.Limit)
	resultFinal := int(math.Ceil(result))

	return responses, resultFinal, totalRows, err

}

// GetNoPelaporan implements VerifikasiDefinition
func (verifikasi VerifikasiRepository) GetNoPelaporan(request *models.NoPalaporanRequest) (responses []models.NoPelaporanNullResponse, err error) {
	kode := "VER-"
	today := lib.GetTimeNow("date2")

	if request.ORGEH != "" {
		kode += request.ORGEH + "-" + today
	}

	query := `SELECT RIGHT(CONCAT("0000",(count(*) + 1)), 4) 'no_pelaporan' FROM verifikasi WHERE no_pelaporan like '%` + kode + `%'`

	verifikasi.logger.Zap.Info(query)
	rows, err := verifikasi.dbRaw.DB.Query(query)

	verifikasi.logger.Zap.Info("rows ", rows)
	if err != nil {
		return responses, err
	}

	response := models.NoPelaporanNullResponse{}
	for rows.Next() {
		_ = rows.Scan(
			&response.NoPelaporan,
		)
		responses = append(responses, response)
	}

	if err = rows.Err(); err != nil {
		return responses, err
	}

	return responses, err
}

// GetLastID implements VerifikasiDefinition
func (verifikasi VerifikasiRepository) GetLastID() (responses []models.VerifikasiLastID, err error) {
	query := "SELECT id FROM verifikasi ORDER BY id DESC LIMIT 1"
	verifikasi.logger.Zap.Info(query)
	rows, err := verifikasi.dbRaw.DB.Query(query)

	verifikasi.logger.Zap.Info("rows ", rows)
	if err != nil {
		return responses, err
	}

	response := models.VerifikasiLastID{}
	for rows.Next() {
		_ = rows.Scan(
			&response.ID,
		)

		responses = append(responses, response)
	}

	if err = rows.Err(); err != nil {
		return responses, err
	}

	return responses, err
}
