package coaching

import (
	"fmt"
	"riskmanagement/lib"
	models "riskmanagement/models/coaching"

	"gorm.io/gorm"

	coachingRepo "riskmanagement/repository/coaching"

	"github.com/google/uuid"
	"gitlab.com/golang-package-library/logger"
)

var (
	timeNow = lib.GetTimeNow("timestime")
	UUID    = uuid.NewString()
)

type CoachingDefinition interface {
	WithTrx(trxHandle *gorm.DB) CoachingService
	GetAll() (responses []models.CoachingResponse, err error)
	GetOne(id int64) (responses models.CoachingResponsesGetOneString, status bool, err error)
	Store(request models.CoachingRequest) (status bool, err error)
	Delete(request *models.CoachingRequestUpdate) (responses bool, err error)
	DeleteCoachingActivity(request *models.CoachingActRequest) (status bool, err error)
	UpdateAllCoaching(request *models.CoachingResponseMaintain) (status bool, err error)
}

type CoachingService struct {
	db               lib.Database
	logger           logger.Logger
	coachingRepo     coachingRepo.CoachingDefinition
	coachingActivity coachingRepo.CoachingActivityDefinition
}

func NewCoachingService(
	db lib.Database,
	logger logger.Logger,
	coachingRepo coachingRepo.CoachingDefinition,
	coachingActivity coachingRepo.CoachingActivityDefinition,
) CoachingDefinition {
	return CoachingService{
		db:               db,
		logger:           logger,
		coachingRepo:     coachingRepo,
		coachingActivity: coachingActivity,
	}
}

// Delete implements CoachingDefinition
func (coaching CoachingService) Delete(request *models.CoachingRequestUpdate) (responses bool, err error) {
	tx := coaching.db.DB.Begin()

	getOneCoaching, exist, err := coaching.GetOne(request.ID)
	if err != nil {
		coaching.logger.Zap.Error(err)
		tx.Rollback()
		return false, err
	}

	updateDataCoaching := &models.CoachingUpdateDelete{
		ID:            request.ID,
		LastMakerID:   request.LastMakerID,
		LastMakerDesc: request.LastMakerDesc,
		LastMakerDate: &timeNow,
		Status:        "02b", //selesai
		Action:        "updateDelete",
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
	_, err = coaching.coachingRepo.Delete(updateDataCoaching, include, tx)

	if err != nil {
		tx.Rollback()
		coaching.logger.Zap.Error(err)
		return false, err
	}

	if exist {
		fmt.Println("getOneCoaching", getOneCoaching)
		tx.Commit()
		return true, err
	}

	return false, err
}

// DeleteCoachingActivity implements CoachingDefinition
func (coaching CoachingService) DeleteCoachingActivity(request *models.CoachingActRequest) (status bool, err error) {
	tx := coaching.db.DB.Begin()
	err = coaching.coachingRepo.DeleteCoachingActivity(request.ID, tx)

	if err != nil {
		tx.Rollback()
		coaching.logger.Zap.Error(err)
		return false, err
	}

	tx.Commit()
	return true, err
}

// GetAll implements CoachingDefinition
func (coaching CoachingService) GetAll() (responses []models.CoachingResponse, err error) {
	return coaching.coachingRepo.GetAll()
}

// GetOne implements CoachingDefinition
func (coaching CoachingService) GetOne(id int64) (responses models.CoachingResponsesGetOneString, status bool, err error) {
	dataCoaching, err := coaching.coachingRepo.GetOne(id)
	fmt.Println(dataCoaching)

	if dataCoaching.ID != 0 {
		fmt.Println("Bukan 0")

		activity, err := coaching.coachingActivity.GetOneActivity(dataCoaching.ID)

		responses = models.CoachingResponsesGetOneString{
			ID:            dataCoaching.ID,
			NoPelaporan:   dataCoaching.NoPelaporan,
			UnitKerja:     dataCoaching.UnitKerja,
			Peserta:       dataCoaching.Peserta,
			JumlahPeserta: dataCoaching.JumlahPeserta,
			ActivityID:    dataCoaching.ActivityID,
			SubActivityID: dataCoaching.SubActivityID,
			MakerID:       dataCoaching.MakerID,
			MakerDesc:     dataCoaching.MakerDesc,
			MakerDate:     dataCoaching.MakerDate,
			LastMakerID:   dataCoaching.LastMakerID,
			LastMakerDesc: dataCoaching.LastMakerDesc,
			LastMakerDate: dataCoaching.LastMakerDate,
			Status:        dataCoaching.Status,
			Action:        dataCoaching.Action,
			Deleted:       dataCoaching.Deleted,
			Activity:      activity,
			UpdatedAt:     dataCoaching.UpdatedAt,
			CreatedAt:     dataCoaching.CreatedAt,
		}

		return responses, true, err
	}
	return responses, false, err
}

// Store implements CoachingDefinition
func (coaching CoachingService) Store(request models.CoachingRequest) (status bool, err error) {
	tx := coaching.db.DB.Begin()

	reqCoaching := &models.Coaching{
		NoPelaporan:   request.NoPelaporan,
		UnitKerja:     request.UnitKerja,
		Peserta:       request.Peserta,
		JumlahPeserta: request.JumlahPeserta,
		ActivityID:    request.ActivityID,
		SubActivityID: request.SubActivityID,
		MakerID:       request.MakerID,
		MakerDesc:     request.MakerDesc,
		MakerDate:     &timeNow,
		LastMakerID:   request.LastMakerID,
		LastMakerDesc: request.LastMakerDesc,
		LastMakerDate: request.LastMakerDate,
		Status:        "01a",
		Action:        "Create",
		CreatedAt:     &timeNow,
	}

	dataCoaching, err := coaching.coachingRepo.Store(reqCoaching, tx)
	if err != nil {
		tx.Rollback()
		coaching.logger.Zap.Error(err)
		return false, err
	}

	fmt.Println("dataCoaching : ", dataCoaching)

	if len(request.Activity) != 0 {
		for _, value := range request.Activity {
			_, err = coaching.coachingActivity.Store(&models.CoachingActivity{
				CoachingID:        dataCoaching.ID,
				RiskIssueID:       value.RiskIssueID,
				JudulMateri:       value.JudulMateri,
				RekomendasiMateri: value.RekomendasiMateri,
				MateriTambahan:    value.MateriTambahan,
				CreatedAt:         &timeNow,
			}, tx)

			if err != nil {
				tx.Rollback()
				coaching.logger.Zap.Error(err)
				return false, err
			}
		}
	} else {
		tx.Rollback()
		coaching.logger.Zap.Error(err)
		return false, err
	}

	tx.Commit()
	return true, err
}

// UpdateAllCoaching implements CoachingDefinition
func (coaching CoachingService) UpdateAllCoaching(request *models.CoachingResponseMaintain) (status bool, err error) {
	tx := coaching.db.DB.Begin()

	updateCoaching := &models.CoachingUpdateActivity{
		ID:            request.ID,
		UnitKerja:     request.UnitKerja,
		Peserta:       request.Peserta,
		JumlahPeserta: request.JumlahPeserta,
		ActivityID:    request.ActivityID,
		SubActivityID: request.SubActivityID,
		LastMakerID:   request.LastMakerID,
		LastMakerDesc: request.LastMakerDesc,
		LastMakerDate: &timeNow,
		Deleted:       false,
		Action:        "Update",
		Status:        "02b",
		UpdatedAt:     &timeNow,
	}

	include := []string{
		"unit_kerja",
		"peserta",
		"jumlah_peserta",
		"activity_id",
		"sub_activity_id",
		"last_maker_id",
		"last_maker_desc",
		"last_maker_date",
		"deleted",
		"action",
		"status",
		"updated_at",
	}

	_, err = coaching.coachingRepo.UpdateAllCoaching(updateCoaching, include, tx)

	if err != nil {
		tx.Rollback()
		coaching.logger.Zap.Error(err)
		return false, err
	}

	if len(request.Activity) != 0 {
		for _, value := range request.Activity {
			updateCoachinAct := &models.CoachingActivity{
				ID:                value.ID,
				CoachingID:        request.ID,
				RiskIssueID:       value.RiskIssueID,
				JudulMateri:       value.JudulMateri,
				RekomendasiMateri: value.RekomendasiMateri,
				MateriTambahan:    value.MateriTambahan,
				CreatedAt:         value.CreatedAt,
				UpdatedAt:         &timeNow,
			}

			_, err = coaching.coachingActivity.Store(updateCoachinAct, tx)

			if err != nil {
				tx.Rollback()
				coaching.logger.Zap.Error(err)
				return false, err
			}
		}
	} else {
		if err != nil {
			tx.Rollback()
			coaching.logger.Zap.Error(err)
			return false, err
		}
	}

	tx.Commit()
	return true, err
}

// WithTrx implements CoachingDefinition
func (coaching CoachingService) WithTrx(trxHandle *gorm.DB) CoachingService {
	coaching.coachingRepo = coaching.coachingRepo.WithTrx(trxHandle)
	return coaching
}
