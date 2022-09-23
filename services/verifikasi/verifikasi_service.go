package verifikasi

import (
	"fmt"
	"os"
	"riskmanagement/lib"
	models "riskmanagement/models/verifikasi"
	verifikasi "riskmanagement/repository/verifikasi"

	requestFile "riskmanagement/models/files"
	fileRepo "riskmanagement/repository/files"

	"github.com/google/uuid"
	"gitlab.com/golang-package-library/logger"
	minio "gitlab.com/golang-package-library/minio"
	"gorm.io/gorm"
)

var (
	timeNow = lib.GetTimeNow("timestime")
	UUID    = uuid.NewString()
)

type VerifikasiDefinition interface {
	WithTrx(trxHandle *gorm.DB) VerifikasiService
	GetAll() (responses []models.VerifikasiResponse, err error)
	GetListData() (responses []models.VerifikasiList, err error)
	GetOne(id int64) (responses models.VerifikasiResponseGetOne, status bool, err error)
	FilterVerifikasi(request models.VerifikasiFilterRequest) (responses []models.VerifikasiList, pangination lib.Pagination, err error)
	Store(request models.VerifikasiRequest) (status bool, err error)
	Delete(request *models.VerifikasiRequestUpdateMaintain) (response bool, err error)
	KonfirmSave(request *models.VerifikasiUpdateMaintain) (response bool, err error)
	DeleteLampiranVerifikasi(request *models.VerifikasiFileRequest) (status bool, err error)
	UpdateAllVerifikasi(request *models.VerifikasiRequestMaintain) (status bool, err error)
}

type VerifikasiService struct {
	db                    lib.Database
	minio                 minio.Minio
	logger                logger.Logger
	verifikasiRepo        verifikasi.VerifikasiDefinition
	verifikasiAnomali     verifikasi.VerifikasiAnomaliDefinition
	verifikasiFile        verifikasi.VerifikasiFilesDefinition
	verifikasiPIC         verifikasi.VerifikasiPICDefinition
	verifikasiRiskControl verifikasi.VerifikasiRiskControlDefinition
	fileRepo              fileRepo.FilesDefinition
}

func NewVerifikasiService(
	db lib.Database,
	minio minio.Minio,
	logger logger.Logger,
	verifikasiRepo verifikasi.VerifikasiDefinition,
	verifikasiAnomali verifikasi.VerifikasiAnomaliDefinition,
	verifikasiFile verifikasi.VerifikasiFilesDefinition,
	verifikasiPIC verifikasi.VerifikasiPICDefinition,
	verifikasiRiskControl verifikasi.VerifikasiRiskControlDefinition,
	fileRepo fileRepo.FilesDefinition,
) VerifikasiDefinition {
	return VerifikasiService{
		db:                    db,
		minio:                 minio,
		logger:                logger,
		verifikasiRepo:        verifikasiRepo,
		verifikasiAnomali:     verifikasiAnomali,
		verifikasiFile:        verifikasiFile,
		verifikasiPIC:         verifikasiPIC,
		verifikasiRiskControl: verifikasiRiskControl,
		fileRepo:              fileRepo,
	}
}

// Delete implements VerifikasiDefinition
func (verifikasi VerifikasiService) Delete(request *models.VerifikasiRequestUpdateMaintain) (response bool, err error) {
	tx := verifikasi.db.DB.Begin()

	getOneVerifikasi, exist, err := verifikasi.GetOne(request.ID)
	if err != nil {
		verifikasi.logger.Zap.Error(err)
		tx.Rollback()
		return false, err
	}

	UpdateDataVerifikasi := &models.VerifikasiUpdateDelete{
		ID:            request.ID,
		LastMakerID:   request.LastMakerID,
		LastMakerDesc: request.LastMakerDesc,
		LastMakerDate: &timeNow,
		Status:        "02b", //selesai
		Action:        "UpdateDelete",
		Deleted:       true,
		UpdatedAt:     &timeNow,
	}

	include := []string{
		"last_maker_id",
		"last_maker_desc",
		"last_maker_date",
		"deleted",
		"status",
		"action",
		"updated_at",
	}

	_, err = verifikasi.verifikasiRepo.Delete(UpdateDataVerifikasi, include, tx)
	if err != nil {
		tx.Rollback()
		verifikasi.logger.Zap.Error(err)
		return false, err
	}

	if exist {
		fmt.Println("getOneVerif", getOneVerifikasi)
		tx.Commit()
		return true, err
	}

	return false, err
}

// GetAll implements VerifikasiDefinition
func (verifikasi VerifikasiService) GetAll() (responses []models.VerifikasiResponse, err error) {
	return verifikasi.verifikasiRepo.GetAll()
}

// GetListData implements VerifikasiDefinition
func (verifikasi VerifikasiService) GetListData() (responses []models.VerifikasiList, err error) {
	return verifikasi.verifikasiRepo.GetListData()
}

// GetOne implements VerifikasiDefinition
func (verifikasi VerifikasiService) GetOne(id int64) (responses models.VerifikasiResponseGetOne, status bool, err error) {
	dataVerif, err := verifikasi.verifikasiRepo.GetOne(id)
	fmt.Println(dataVerif)

	if dataVerif.ID != 0 {
		fmt.Println("bukan 0")

		data_anomali, err := verifikasi.verifikasiAnomali.GetOneByVerifikasi(dataVerif.ID)
		files, err := verifikasi.verifikasiFile.GetOneFileByID(dataVerif.ID)
		pic_tindak_lanjut, err := verifikasi.verifikasiPIC.GetOneByPIC(dataVerif.ID)
		risk_control, err := verifikasi.verifikasiRiskControl.GetOneDataByID(dataVerif.ID)

		responses = models.VerifikasiResponseGetOne{
			ID:                        dataVerif.ID,
			NoPelaporan:               dataVerif.NoPelaporan,
			UnitKerja:                 dataVerif.UnitKerja,
			ActivityID:                dataVerif.ActivityID,
			SubActivityID:             dataVerif.SubActivityID,
			ProductID:                 dataVerif.ProductID,
			RiskIssueID:               dataVerif.RiskIssueID,
			RiskIndicatorID:           dataVerif.RiskIndicatorID,
			IncidentCauseID:           dataVerif.IncidentCauseID,
			SubIncidentCauseID:        dataVerif.SubIncidentCauseID,
			ApplicationID:             dataVerif.ApplicationID,
			HasilVerifikasi:           dataVerif.HasilVerifikasi,
			KunjunganNasabah:          dataVerif.KunjunganNasabah,
			IndikasiFraud:             dataVerif.IndikasiFraud,
			JenisKerugianFinansial:    dataVerif.JenisKerugianFinansial,
			JumlahPerkiraanKerugian:   dataVerif.JumlahPerkiraanKerugian,
			JenisKerugianNonFinansial: dataVerif.JenisKerugianNonFinansial,
			RekomendasiTindakLanjut:   dataVerif.RekomendasiTindakLanjut,
			RencanaTindakLanjut:       dataVerif.RencanaTindakLanjut,
			RiskTypeID:                dataVerif.RiskTypeID,
			TanggalDitemukan:          dataVerif.TanggalDitemukan,
			TanggalMulaiRTL:           dataVerif.TanggalMulaiRTL,
			TanggalTargetSelesai:      dataVerif.TanggalTargetSelesai,
			MakerID:                   dataVerif.MakerID,
			MakerDesc:                 dataVerif.MakerDesc,
			MakerDate:                 dataVerif.MakerDate,
			LastMakerID:               dataVerif.LastMakerID,
			LastMakerDesc:             dataVerif.LastMakerDesc,
			LastMakerDate:             dataVerif.LastMakerDate,
			Status:                    dataVerif.Status,
			Action:                    dataVerif.Action,
			Deleted:                   dataVerif.Deleted,
			DataAnomali:               data_anomali,
			PICTindakLanjut:           pic_tindak_lanjut,
			Files:                     files,
			RiskControl:               risk_control,
			UpdatedAt:                 dataVerif.UpdatedAt,
			CreatedAt:                 dataVerif.CreatedAt,
		}

		return responses, true, err
	}

	return responses, false, err

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
		Status:                    "01a",
		Action:                    "Draft",
		Deleted:                   false,
		CreatedAt:                 &timeNow,
	}

	dataVerif, err := verifikasi.verifikasiRepo.Store(reqVerif, tx)

	if err != nil {
		tx.Rollback()
		verifikasi.logger.Zap.Error(err)
		return false, err
	}
	// fmt.Println("data verifikasi : ", dataVerif)
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
				// CreatedAt:       &timeNow,
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

	//Begin Input Kelemahan Kontrol
	if len(request.RiskControl) != 0 {
		for _, value := range request.RiskControl {
			_, err = verifikasi.verifikasiRiskControl.Store(&models.VerifikasiRiskControl{
				VerifikasiId:  dataVerif.ID,
				RiskControlID: value.RiskControlID,
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
	//End Input Kelemahan Kontrol

	//Begin Input data PIC
	if len(request.PICTindakLanjut) != 0 {
		for _, value := range request.PICTindakLanjut {
			_, err = verifikasi.verifikasiPIC.Store(&models.VerifikasiPICTindakLanjut{
				VerifikasiID:          dataVerif.ID,
				PICID:                 value.PICID,
				TanggalTindakLanjut:   value.TanggalTindakLanjut,
				DeskripsiTindakLanjut: value.DeskripsiTindakLanjut,
				Status:                "01a",
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

	//Begin Input Lampiran
	bucket := os.Getenv("BUCKET_NAME")

	if len(request.Files) != 0 {

		for _, value := range request.Files {
			var destinationPath string
			bucketExist := verifikasi.minio.BucketExist(verifikasi.minio.Client(), bucket)

			sourcePath := fmt.Sprint(value.Path)
			newPath := "verifikasi/" +
				lib.GetTimeNow("year") + "/" +
				lib.GetTimeNow("month") + "/" +
				lib.GetTimeNow("day")

			destinationPath = newPath + "/" + value.Filename

			if bucketExist {
				fmt.Println("Exist")
				fmt.Println(bucket)
				fmt.Println(sourcePath)
				fmt.Println(destinationPath)
				uploaded := verifikasi.minio.PutObject(verifikasi.minio.MinioClient, bucket, destinationPath, sourcePath)

				fmt.Println(uploaded)
			} else {
				fmt.Println("Not Exist")
				fmt.Println(bucket)
				fmt.Println(sourcePath)
				fmt.Println(destinationPath)
				verifikasi.minio.MakeBucket(verifikasi.minio.Client(), bucket, "")
				// uploaded := materi.minio.CopyObject(materi.minio.Client(), bucket, sourcePath, bucket, destinationPath)
				uploaded := verifikasi.minio.PutObject(verifikasi.minio.MinioClient, bucket, destinationPath, sourcePath)
				fmt.Println(uploaded)
			}

			files, err := verifikasi.fileRepo.Store(&requestFile.Files{
				Filename:  value.Filename,
				Path:      destinationPath,
				Extension: value.Extension,
				Size:      value.Size,
				CreatedAt: &timeNow,
			}, tx)

			if err != nil {
				tx.Rollback()
				verifikasi.logger.Zap.Error(err)
				return false, err
			}

			_, err = verifikasi.verifikasiFile.Store(&models.VerifikasiFiles{
				VerifikasiID: dataVerif.ID,
				FilesID:      files.ID,
				// CreatedAt:    &timeNow,
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

	//End Input Lampiran

	tx.Commit()
	return true, err
}

// WithTrx implements VerifikasiDefinition
func (verifikasi VerifikasiService) WithTrx(trxHandle *gorm.DB) VerifikasiService {
	verifikasi.verifikasiRepo = verifikasi.verifikasiRepo.WithTrx(trxHandle)
	return verifikasi
}

// DeleteLampiranVerifikasi implements VerifikasiDefinition
func (verifikasi VerifikasiService) DeleteLampiranVerifikasi(request *models.VerifikasiFileRequest) (status bool, err error) {
	bucket := os.Getenv("BUCKET_NAME")

	ok := verifikasi.minio.RemoveObject(verifikasi.minio.Client(), bucket, request.Path)

	if !ok {
		return false, err
	} else {
		tx := verifikasi.db.DB.Begin()
		err = verifikasi.fileRepo.Delete(request.FilesID, tx)
		if err != nil {
			tx.Rollback()
			verifikasi.logger.Zap.Error(err)
			return false, err
		}

		err = verifikasi.verifikasiRepo.DeleteLampiranVerifikasi(request.VerifikasiLampiranID, tx)
		if err != nil {
			tx.Rollback()
			verifikasi.logger.Zap.Error(err)
			return false, err
		}

		tx.Commit()

	}

	return true, err

}

// KonfirmSave implements VerifikasiDefinition
func (verifikasi VerifikasiService) KonfirmSave(request *models.VerifikasiUpdateMaintain) (response bool, err error) {
	tx := verifikasi.db.DB.Begin()

	getOneVerifikasi, exist, err := verifikasi.GetOne(request.ID)
	if err != nil {
		verifikasi.logger.Zap.Error(err)
		tx.Rollback()
		return false, err
	}

	UpdateDataVerifikasi := &models.VerifikasiUpdateMaintain{
		ID:            request.ID,
		LastMakerID:   request.LastMakerID,
		LastMakerDesc: request.LastMakerDesc,
		LastMakerDate: &timeNow,
		Status:        "02b", //selesai
		Action:        "Selesai",
		UpdatedAt:     &timeNow,
	}

	include := []string{
		"last_maker_id",
		"last_maker_desc",
		"last_maker_date",
		"status",
		"action",
		"updated_at",
	}

	_, err = verifikasi.verifikasiRepo.KonfirmSave(UpdateDataVerifikasi, include, tx)
	if err != nil {
		tx.Rollback()
		verifikasi.logger.Zap.Error(err)
		return false, err
	}

	if exist {
		fmt.Println("getOneVerif", getOneVerifikasi)
		tx.Commit()
		return true, err
	}
	return false, err

}

// UpdateAllVerifikasi implements VerifikasiDefinition
func (verifikasi VerifikasiService) UpdateAllVerifikasi(request *models.VerifikasiRequestMaintain) (status bool, err error) {
	tx := verifikasi.db.DB.Begin()

	updateVerifikasi := &models.VerifikasiUpdateAll{
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
		LastMakerID:               request.LastMakerID,
		LastMakerDesc:             request.LastMakerDesc,
		LastMakerDate:             &timeNow,
		Status:                    "02b",
		Action:                    "Update",
		UpdatedAt:                 &timeNow,
	}

	include := []string{
		"no_pelaporan",
		"unit_kerja",
		"activity_id",
		"sub_activity_id",
		"product_id",
		"risk_issue_id",
		"risk_indicator_id",
		"incident_cause_id",
		"sub_incident_cause_id",
		"application_id",
		"hasil_verifikasi",
		"kunjungan_nasabah",
		"indikasi_fraud",
		"jenis_kerugian_finansial",
		"rekomendasi_tindak_lanjut",
		"rencana_tindak_lanjut",
		"risk_type_id",
		"tanggal_ditemukan",
		"tanggal_mulai_rtl",
		"tanggal_target_selesai",
		"last_maker_id",
		"last_maker_desc",
		"last_maker_date",
		"status",
		"action",
		"updated_at",
	}

	_, err = verifikasi.verifikasiRepo.UpdateAllVerifikasi(updateVerifikasi, include, tx)

	if err != nil {
		tx.Rollback()
		verifikasi.logger.Zap.Error(err)
		return false, err
	}

	//Update & add Data Anomali
	if len(request.DataAnomali) != 0 {
		for _, value := range request.DataAnomali {
			updateAnomali := &models.VerifikasiAnomaliData{
				ID:              value.ID,
				VerifikasiID:    request.ID,
				TanggalKejadian: value.TanggalKejadian,
				NomorRekening:   value.NomorRekening,
				Nominal:         value.Nominal,
				Keterangan:      value.Keterangan,
			}

			_, err = verifikasi.verifikasiAnomali.Store(updateAnomali, tx)

			if err != nil {
				tx.Rollback()
				verifikasi.logger.Zap.Error(err)
				return false, err
			}
		}
	} else {
		if err != nil {
			tx.Rollback()
			verifikasi.logger.Zap.Error(err)
			return false, err
		}
	}
	//#update & add Anomali

	//Update & add Risk Control
	if len(request.RiskControl) != 0 {
		for _, value := range request.RiskControl {
			updateRiskControl := &models.VerifikasiRiskControl{
				ID:            value.ID,
				VerifikasiId:  request.ID,
				RiskControlID: value.RiskControlID,
			}
			_, err = verifikasi.verifikasiRiskControl.Store(updateRiskControl, tx)

			if err != nil {
				tx.Rollback()
				verifikasi.logger.Zap.Error(err)
				return false, err
			}

		}
	} else {
		if err != nil {
			tx.Rollback()
			verifikasi.logger.Zap.Error(err)
			return false, err
		}
	}
	//#Update & add Risk Control

	//Update & add Data PIC Tindak Lanjut
	if len(request.PICTindakLanjut) != 0 {
		for _, value := range request.PICTindakLanjut {
			updatePIC := &models.VerifikasiPICTindakLanjut{
				ID:                    value.ID,
				VerifikasiID:          request.ID,
				PICID:                 value.PICID,
				TanggalTindakLanjut:   value.TanggalTindakLanjut,
				DeskripsiTindakLanjut: value.DeskripsiTindakLanjut,
			}

			_, err = verifikasi.verifikasiPIC.Store(updatePIC, tx)

			if err != nil {
				tx.Rollback()
				verifikasi.logger.Zap.Error(err)
				return false, err
			}

		}
	} else {
		if err != nil {
			tx.Rollback()
			verifikasi.logger.Zap.Error(err)
			return false, err
		}
	}
	//#Update & add Data PIC Tindak Lanjut

	//#Update Lampiran
	bucket := os.Getenv("BUCKET_NAME")

	if len(request.Files) != 0 {

		err := verifikasi.verifikasiFile.DeleteFilesByID(request.ID, tx)
		if err != nil {
			tx.Rollback()
			verifikasi.logger.Zap.Error(err)
			return false, err
		}

		for _, value := range request.Files {
			var destinationPath string
			bucketExist := verifikasi.minio.BucketExist(verifikasi.minio.Client(), bucket)

			sourcePath := fmt.Sprintf(value.Path)
			newPath := "verifikasi/" +
				lib.GetTimeNow("year") + "/" +
				lib.GetTimeNow("month") + "/" +
				lib.GetTimeNow("day")

			destinationPath = newPath + "/" + value.Filename

			if bucketExist {
				fmt.Println("Exist")
				fmt.Println(bucket)
				fmt.Println(sourcePath)
				fmt.Println(destinationPath)
				uploaded := verifikasi.minio.PutObject(verifikasi.minio.MinioClient, bucket, destinationPath, sourcePath)

				fmt.Println(uploaded)
			} else {
				fmt.Println("Not Exist")
				fmt.Println(bucket)
				fmt.Println(sourcePath)
				fmt.Println(destinationPath)
				verifikasi.minio.MakeBucket(verifikasi.minio.Client(), bucket, "")
				// uploaded := materi.minio.CopyObject(materi.minio.Client(), bucket, sourcePath, bucket, destinationPath)
				uploaded := verifikasi.minio.PutObject(verifikasi.minio.MinioClient, bucket, destinationPath, sourcePath)
				fmt.Println(uploaded)
			}

			files, err := verifikasi.fileRepo.Store(&requestFile.Files{
				ID:        value.ID,
				Filename:  value.Filename,
				Path:      destinationPath,
				Extension: value.Extension,
				Size:      value.Size,
				UpdatedAt: &timeNow,
			}, tx)

			if err != nil {
				tx.Rollback()
				verifikasi.logger.Zap.Error(err)
				return false, err
			}

			_, err = verifikasi.verifikasiFile.Store(&models.VerifikasiFiles{
				VerifikasiID: request.ID,
				FilesID:      files.ID,
			}, tx)
		}
	} else {
		tx.Rollback()
		verifikasi.logger.Zap.Error(err)
		return false, err
	}
	//#UpdateLampiran

	tx.Commit()
	return true, err
}

// FilterVerifikasi implements VerifikasiDefinition
func (verifikasi VerifikasiService) FilterVerifikasi(request models.VerifikasiFilterRequest) (responses []models.VerifikasiList, pangination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort
	dataVerif, totalRows, totalData, err := verifikasi.verifikasiRepo.FilterVerifikasi(&request)
	if err != nil {
		verifikasi.logger.Zap.Error(err)
		return responses, pangination, err
	}

	for _, response := range dataVerif {
		responses = append(responses, models.VerifikasiList{
			ID:          response.ID.Int64,
			NoPelaporan: response.NoPelaporan.String,
			UnitKerja:   response.UnitKerja.String,
			Aktifitas:   response.Aktifitas.String,
			StatusVerif: response.StatusVerif.String,
		})
	}

	pangination = lib.SetPaginationResponse(page, limit, totalRows, totalData)

	return responses, pangination, err
}
