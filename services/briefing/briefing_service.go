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
	DeleteBriefingMateri(request *models.BriefingMateriRequest) (status bool, err error)
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
	panic("unimplemented")
}

// DeleteBriefingMateri implements BriefingDefinition
func (briefing BriefingService) DeleteBriefingMateri(request *models.BriefingMateriRequest) (status bool, err error) {
	panic("unimplemented")
}

// GetAll implements BriefingDefinition
func (briefing BriefingService) GetAll() (responses []models.BriefingResponse, err error) {
	return briefing.briefingRepo.GetAll()
}

// GetOne implements BriefingDefinition
func (briefing BriefingService) GetOne(id int64) (responses models.BriefingResponseGetOneString, status bool, err error) {
	panic("unimplemented")
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
		for range request.Materi {
			_, err = briefing.briefingMateri.Store(&models.BriefingMateri{
				BriefingID:        dataBriefing.ID,
				ActivityID:        0,
				SubActivityID:     0,
				ProductID:         0,
				JudulMateri:       "",
				RekomendasiMateri: "",
				MateriTambahan:    "",
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
