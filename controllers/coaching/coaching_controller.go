package controllers

import (
	"fmt"
	"riskmanagement/lib"
	models "riskmanagement/models/coaching"
	services "riskmanagement/services/coaching"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type CoachingController struct {
	logger  logger.Logger
	service services.CoachingDefinition
}

func NewCoachingController(
	CoachingService services.CoachingDefinition,
	logger logger.Logger,
) CoachingController {
	return CoachingController{
		service: CoachingService,
		logger:  logger,
	}
}

func (coaching CoachingController) GetAll(c *gin.Context) {
	datas, err := coaching.service.GetAll()

	if err != nil {
		coaching.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	if len(datas) == 0 {
		coaching.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "404", "Data tidak ditemukan", "")
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", datas)
}

func (coaching CoachingController) Store(c *gin.Context) {
	data := models.CoachingRequest{}

	if err := c.Bind(&data); err != nil {
		coaching.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), data)
		return
	}

	status, err := coaching.service.Store(data)
	if err != nil {
		coaching.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	if !status {
		coaching.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error status", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Input data berhasil", true)
}

func (coaching CoachingController) GetOne(c *gin.Context) {
	paramID := c.Param("id")
	id, err := strconv.Atoi(paramID)

	if err != nil {
		coaching.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	data, status, err := coaching.service.GetOne(int64(id))
	if err != nil {
		coaching.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	if !status {
		coaching.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "404", "Data tidak ditemukan", nil)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", data)
}

func (coaching CoachingController) DeleteCoachingActivity(c *gin.Context) {
	data := models.CoachingActRequest{}

	if err := c.Bind(&data); err != nil {
		coaching.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	status, err := coaching.service.DeleteCoachingActivity(&data)
	if err != nil {
		coaching.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	if !status {
		coaching.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Data gagal disimpan", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Delete data berhasil", true)
}

func (coaching CoachingController) Delete(c *gin.Context) {
	data := models.CoachingRequestUpdate{}

	if err := c.Bind(&data); err != nil {
		coaching.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	status, err := coaching.service.Delete(&data)
	if err != nil {
		coaching.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	if !status {
		coaching.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Data Gagal Dihapus", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Hapus data berhasil", true)
}

func (coaching CoachingController) UpdateAllCoaching(c *gin.Context) {
	data := models.CoachingResponseMaintain{}

	if err := c.Bind(&data); err != nil {
		coaching.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	fmt.Println(data)

	status, err := coaching.service.UpdateAllCoaching(&data)
	if err != nil {
		coaching.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	if !status {
		coaching.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Data Gagal Diupdate", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Update data berhasil", true)
}
