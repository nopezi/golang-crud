package controller

import (
	"riskmanagement/lib"
	models "riskmanagement/models/verifikasi"
	services "riskmanagement/services/verifikasi"

	"github.com/gin-gonic/gin"

	"gitlab.com/golang-package-library/logger"
)

type VerifikasiController struct {
	logger  logger.Logger
	service services.VerifikasiDefinition
}

func NewVerifikasiController(
	verifikasiService services.VerifikasiDefinition,
	logger logger.Logger,
) VerifikasiController {
	return VerifikasiController{
		service: verifikasiService,
		logger:  logger,
	}
}

func (verifikasi VerifikasiController) GetAll(c *gin.Context) {
	datas, err := verifikasi.service.GetAll()

	if err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "200", "Internal Error", "")
		return
	}

	if len(datas) == 0 {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Data tidak ditemukan", "")
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", datas)
}

func (verifikasi VerifikasiController) Store(c *gin.Context) {
	data := models.VerifikasiRequest{}

	if err := c.Bind(&data); err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), data)
		return
	}

	status, err := verifikasi.service.Store(data)
	if err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	if !status {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error status", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Input data berhasil", true)
}
