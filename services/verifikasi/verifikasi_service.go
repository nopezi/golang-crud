package verifikasi

import (
	"fmt"
	"riskmanagement/lib"
	models "riskmanagement/models/verifikasi"
	verifikasi "riskmanagement/repository/verifikasi"

	"github.com/google/uuid"
	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

var (
	timeNow = lib.GetTimeNow("timestime")
	UUID    = uuid.NewString()
)

type VerifikasiDefinition interface {
	WithTrx(trxHandle *gorm.DB) VerifikasiService
	GetAll() (responses []models.VerifikasiResponse, err error)
	GetOne(id int64) (responses models.VerifikasiResponseGetOne, status bool, err error)
	Store(request models.VerifikasiRequest) (status bool, err error)
	Delete(request *models.VerifikasiRequestUpdateMaintain) (response bool, err error)
}

type VerifikasiService struct {
	db                lib.Database
	logger            logger.Logger
	verifikasiRepo    verifikasi.VerifikasiDefinition
	verifikasiAnomali verifikasi.VerifikasiAnomaliDefinition
	verifikasiFile    verifikasi.VerifikasiFilesDefinition
	verifikasiPIC     verifikasi.VerifikasiPICDefinition
}

func NewVerifikasiService(
	db lib.Database,
	logger logger.Logger,
	verifikasiRepo verifikasi.VerifikasiDefinition,
	verifikasiAnomali verifikasi.VerifikasiAnomaliDefinition,
	verifikasiFile verifikasi.VerifikasiFilesDefinition,
	verifikasiPIC verifikasi.VerifikasiPICDefinition,
) VerifikasiDefinition {
	return VerifikasiService{
		db:                db,
		logger:            logger,
		verifikasiRepo:    verifikasiRepo,
		verifikasiAnomali: verifikasiAnomali,
		verifikasiFile:    verifikasiFile,
		verifikasiPIC:     verifikasiPIC,
	}
}

// Delete implements VerifikasiDefinition
func (verifikasi VerifikasiService) Delete(request *models.VerifikasiRequestUpdateMaintain) (response bool, err error) {
	panic("unimplemented")
}

// GetAll implements VerifikasiDefinition
func (verifikasi VerifikasiService) GetAll() (responses []models.VerifikasiResponse, err error) {
	return verifikasi.verifikasiRepo.GetAll()
}

// GetOne implements VerifikasiDefinition
func (verifikasi VerifikasiService) GetOne(id int64) (responses models.VerifikasiResponseGetOne, status bool, err error) {
	panic("unimplemented")
}

// Store implements VerifikasiDefinition
func (verifikasi VerifikasiService) Store(request models.VerifikasiRequest) (status bool, err error) {
	tx := verifikasi.db.DB.Begin()

	//input data verifikasi
	reqVerif := &models.Verifikasi{
		ID:                        request.ID,
		NoPelaporan:               request.NoPelaporan,
		UnitKerja:                 request.UnitKerja,
		ActivityID:                request.ActivityID,
		SubActivityID:             request.SubActivityID,
		ProductID:                 request.ProductID,
		RiskIssueID:               request.RiskIssueID,
		RiskIndicatorID:           request.RiskIndicatorID,
		IncidentCauseID:           request.IncidentCauseID,
		SubIncidentCauseID:        request.SubIncidentCauseID,
		ApplicationID:             request.ApplicationID,
		HasilVerifikasi:           request.HasilVerifikasi,
		KunjunganNasabah:          request.KunjunganNasabah,
		IndikasiFraud:             request.IndikasiFraud,
		JenisKerugianFinansial:    request.JenisKerugianFinansial,
		JumlahPerkiraanKerugian:   request.JumlahPerkiraanKerugian,
		JenisKerugianNonFinansial: request.JenisKerugianNonFinansial,
		RekomendasiTindakLanjut:   request.RekomendasiTindakLanjut,
		RencanaTindakLanjut:       request.RencanaTindakLanjut,
		RiskTypeID:                request.RiskTypeID,
		TanggalDitemukan:          request.TanggalDitemukan,
		TanggalMulaiRTL:           request.TanggalMulaiRTL,
		TanggalTargetSelesai:      request.TanggalTargetSelesai,
		MakerID:                   request.MakerID,
		MakerDesc:                 request.MakerDesc,
		MakerDate:                 &timeNow,
		LastMakerID:               request.LastMakerID,
		LastMakerDesc:             request.LastMakerDesc,
		LastMakerDate:             &timeNow,
		Status:                    "Draft",
		Action:                    "01a",
		Deleted:                   false,
		CreatedAt:                 &timeNow,
	}

	dataVerif, err := verifikasi.verifikasiRepo.Store(reqVerif, tx)

	if err != nil {
		tx.Rollback()
		verifikasi.logger.Zap.Error(err)
		return false, err
	}
	fmt.Println("data verifikasi : ", dataVerif)
	//end data verifikasi

	//Begin Input data anomali
	if len(request.DataAnomali) != 0 {
		for _, value := range request.DataAnomali {
			_, err = verifikasi.verifikasiAnomali.Store(&models.VerifikasiAnomaliData{
				VerifikasiID:    dataVerif.ID,
				TanggalKejadian: value.TanggalKejadian,
				NomorRekening:   value.NomorRekening,
				Nominal:         value.Nominal,
				Keterangan:      value.Keterangan,
				CreatedAt:       &timeNow,
			}, tx)

			if err != nil {
				tx.Rollback()
				verifikasi.logger.Zap.Error(err)
				return false, err
			}
		}
	} else {
		tx.Rollback()
		verifikasi.logger.Zap.Error(err)
		return false, err
	}
	//End Input data anomali

	//Begin Input data PIC
	if len(request.PICTindakLanjut) != 0 {
		for _, value := range request.PICTindakLanjut {
			_, err = verifikasi.verifikasiPIC.Store(&models.VerifikasiPICTindakLanjut{
				VerifikasiID:          dataVerif.ID,
				PICID:                 value.PICID,
				TanggalTindakLanjut:   value.TanggalTindakLanjut,
				DeskripsiTindakLanjut: value.DeskripsiTindakLanjut,
				CreatedAt:             &timeNow,
			}, tx)

			if err != nil {
				tx.Rollback()
				verifikasi.logger.Zap.Error(err)
				return false, err
			}
		}
	} else {
		tx.Rollback()
		verifikasi.logger.Zap.Error(err)
		return false, err
	}
	//End Input data PIC

	tx.Commit()
	return true, err
}

// WithTrx implements VerifikasiDefinition
func (verifikasi VerifikasiService) WithTrx(trxHandle *gorm.DB) VerifikasiService {
	panic("unimplemented")
}
