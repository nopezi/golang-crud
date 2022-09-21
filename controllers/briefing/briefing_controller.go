package controllers

import (
	"fmt"
	"riskmanagement/lib"
	models "riskmanagement/models/briefing"
	services "riskmanagement/services/briefing"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type BriefingController struct {
	logger  logger.Logger
	service services.BriefingDefinition
}

func NewBriefingController(
	BriefingService services.BriefingDefinition,
	logger logger.Logger,
) BriefingController {
	return BriefingController{
		service: BriefingService,
		logger:  logger,
	}
}

func (briefing BriefingController) GetAll(c *gin.Context) {
	datas, err := briefing.service.GetAll()
	if err != nil {
		briefing.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	if len(datas) == 0 {
		briefing.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "404", "Data Tidak ditemukan !", "")
		return
	}
	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", datas)
}

func (briefing BriefingController) GetData(c *gin.Context) {
	datas, err := briefing.service.GetData()
	if err != nil {
		briefing.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	if len(datas) == 0 {
		briefing.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "404", "Data Tidak ditemukan !", "")
		return
	}
	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", datas)
}

func (briefing BriefingController) GetOne(c *gin.Context) {
	paramID := c.Param("id")
	id, err := strconv.Atoi(paramID)

	if err != nil {
		briefing.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "200", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	data, status, err := briefing.service.GetOne(int64(id))
	if err != nil {
		briefing.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", data)
		return
	}

	if !status {
		briefing.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "404", "Data Tidak Ditemukan", nil)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", data)
}

func (briefing BriefingController) Store(c *gin.Context) {
	data := models.BriefingRequest{}
	if err := c.Bind(&data); err != nil {
		briefing.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), data)
		return
	}
	// fmt.Println(data.)
	status, err := briefing.service.Store(data)
	if err != nil {
		briefing.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	if !status {
		briefing.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Input Data Berhasil", true)
}

func (briefing BriefingController) Delete(c *gin.Context) {
	data := models.BriefingRequestUpdate{}

	if err := c.Bind(&data); err != nil {
		briefing.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	status, err := briefing.service.Delete(&data)
	if err != nil {
		briefing.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	if !status {
		briefing.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Data Gagal Dihapus", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Hapus data berhasil", true)
}

func (briefing BriefingController) DeleteBriefingMateri(c *gin.Context) {
	data := models.BriefMateriRequest{}

	if err := c.Bind(&data); err != nil {
		briefing.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	status, err := briefing.service.DeleteBriefingMateri(&data)
	if err != nil {
		briefing.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	if !status {
		briefing.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Data Gagal Disimpan", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Delete data berhasil", true)
}

func (briefing BriefingController) UpdateAllBrief(c *gin.Context) {
	data := models.BriefingResponseMaintain{}

	if err := c.Bind(&data); err != nil {
		briefing.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	fmt.Println(data)

	status, err := briefing.service.UpdateAllBrief(&data)
	if err != nil {
		briefing.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	if !status {
		briefing.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Data Gagal diupdate", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Update data berhasil", true)
}
