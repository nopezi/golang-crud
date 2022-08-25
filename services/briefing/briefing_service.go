package briefing

import (
	"fmt"
	"riskmanagement/lib"
	models "riskmanagement/models/briefing"

	"github.com/google/uuid"

	briefingRepo "riskmanagement/repository/briefing"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

var (
	timeNow = lib.GetTimeNow("timestime")
	UUID    = uuid.NewString()
)

type BriefingDefinition interface {
	WithTrx(trxHandle *gorm.DB) BriefingService
	GetAll() (responses []models.BriefingResponse, err error)
	GetOne(id int64) (responses models.BriefingResponseGetOneString, status bool, err error)
	Store(request models.BriefingRequest) (status bool, err error)
	Delete(request *models.BriefingRequestUpdate) (responses bool, err error)
	DeleteBriefingMateri(request *models.BriefMateriRequest) (status bool, err error)
}

type BriefingService struct {
	db             lib.Database
	logger         logger.Logger
	briefingRepo   briefingRepo.BriefingDefinition
	briefingMateri briefingRepo.BriefingMateriDefinition
}

func NewBriefingService(
	db lib.Database,
	logger logger.Logger,
	briefingRepo briefingRepo.BriefingDefinition,
	briefingMateri briefingRepo.BriefingMateriDefinition,
) BriefingDefinition {
	return BriefingService{
		db:             db,
		logger:         logger,
		briefingRepo:   briefingRepo,
		briefingMateri: briefingMateri,
	}
}

// Delete implements BriefingDefinition
func (briefing BriefingService) Delete(request *models.BriefingRequestUpdate) (responses bool, err error) {
	tx := briefing.db.DB.Begin()

	getOneBriefing, exist, err := briefing.GetOne(request.ID)
	if err != nil {
		briefing.logger.Zap.Error(err)
		tx.Rollback()
		return false, err
	}

	updateDataBriefing := &models.BriefingUpdateDelete{
		ID:            request.ID,
		LastMakerID:   request.LastMakerID,
		LastMakerDesc: request.LastMakerDesc,
		LastMakerDate: request.LastMakerDate,
		Deleted:       true,
		Action:        "updateDelete",
		Status:        "02b", //selesai
		UpdatedAt:     &timeNow,
	}

	_, err = briefing.briefingRepo.Delete(updateDataBriefing,
		[]string{
			"last_maker_id",
			"last_maker_desc",
			"last_maker_date",
			"deleted",
			"action",
			"status",
			"updated_at",
		}, tx)

	if err != nil {
		tx.Rollback()
		briefing.logger.Zap.Error(err)
		return false, err
	}

	if exist {
		fmt.Println("getOneBriefing", getOneBriefing)
		tx.Commit()
		return true, err
	}
	return false, err
}

// DeleteBriefingMateri implements BriefingDefinition
func (briefing BriefingService) DeleteBriefingMateri(request *models.BriefMateriRequest) (status bool, err error) {
	tx := briefing.db.DB.Begin()
	err = briefing.briefingRepo.DeleteBriefingMateri(request.ID, tx)

	if err != nil {
		tx.Rollback()
		briefing.logger.Zap.Error(err)
		return false, err
	}
	tx.Commit()
	return true, err
}

// GetAll implements BriefingDefinition
func (briefing BriefingService) GetAll() (responses []models.BriefingResponse, err error) {
	return briefing.briefingRepo.GetAll()
}

// GetOne implements BriefingDefinition
func (briefing BriefingService) GetOne(id int64) (responses models.BriefingResponseGetOneString, status bool, err error) {
	dataBriefing, err := briefing.briefingRepo.GetOne(id)
	fmt.Println(dataBriefing)
	if dataBriefing.ID != 0 {
		fmt.Println("Bukan 0")

		materi, err := briefing.briefingMateri.GetOneBriefing(dataBriefing.ID)

		responses = models.BriefingResponseGetOneString{
			ID:            dataBriefing.ID,
			NoPelaporan:   dataBriefing.NoPelaporan,
			UnitKerja:     dataBriefing.UnitKerja,
			Peserta:       dataBriefing.Peserta,
			JumlahPeserta: dataBriefing.JumlahPeserta,
			MakerID:       dataBriefing.MakerID,
			MakerDesc:     dataBriefing.MakerDesc,
			MakerDate:     dataBriefing.MakerDate,
			LastMakerID:   dataBriefing.LastMakerID,
			LastMakerDesc: dataBriefing.LastMakerDesc,
			LastMakerDate: dataBriefing.LastMakerDate,
			Status:        dataBriefing.Status,
			Action:        dataBriefing.Action,
			Deleted:       dataBriefing.Deleted,
			CreatedAt:     dataBriefing.CreatedAt,
			UpdatedAt:     dataBriefing.UpdatedAt,
			Materi:        materi,
		}
		return responses, true, err
	}
	return responses, false, err
}

// Store implements BriefingDefinition
func (briefing BriefingService) Store(request models.BriefingRequest) (status bool, err error) {
	tx := briefing.db.DB.Begin()
	//Input Briefing
	reqBriefing := &models.Briefing{
		NoPelaporan:   request.NoPelaporan,
		UnitKerja:     request.UnitKerja,
		Peserta:       request.Peserta,
		JumlahPeserta: request.JumlahPeserta,
		MakerID:       request.MakerID,
		MakerDesc:     request.MakerDesc,
		MakerDate:     &timeNow,
		LastMakerID:   request.LastMakerID,
		LastMakerDesc: request.LastMakerDesc,
		LastMakerDate: &timeNow,
		Status:        "01a",
		Action:        "Create",
		CreatedAt:     &timeNow,
	}

	dataBriefing, err := briefing.briefingRepo.Store(reqBriefing, tx)
	if err != nil {
		tx.Rollback()
		briefing.logger.Zap.Error(err)
		return false, err
	}

	fmt.Println("dataBriefing", dataBriefing)

	//Input Briefing Materi
	if len(request.Materi) != 0 {
		for _, value := range request.Materi {
			_, err = briefing.briefingMateri.Store(&models.BriefingMateri{
				BriefingID:        dataBriefing.ID,
				ActivityID:        value.ActivityID,
				SubActivityID:     value.SubActivityID,
				ProductID:         value.ProductID,
				JudulMateri:       value.JudulMateri,
				RekomendasiMateri: value.RekomendasiMateri,
				MateriTambahan:    value.MateriTambahan,
				CreatedAt:         &timeNow,
			}, tx)

			if err != nil {
				tx.Rollback()
				briefing.logger.Zap.Error(err)
				return false, err
			}
		}
	} else {
		tx.Rollback()
		briefing.logger.Zap.Error(err)
		return false, err
	}

	tx.Commit()
	return true, err
}

// WithTrx implements BriefingDefinition
func (briefing BriefingService) WithTrx(trxHandle *gorm.DB) BriefingService {
	briefing.briefingRepo = briefing.briefingRepo.WithTrx(trxHandle)
	return briefing
}
