package controllers

import (
	"riskmanagement/lib"
	models "riskmanagement/models/materi"
	services "riskmanagement/services/materi"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type MateriController struct {
	logger  logger.Logger
	service services.MateriDefinition
}

func NewMateriController(MateriService services.MateriDefinition, logger logger.Logger) MateriController {
	return MateriController{
		service: MateriService,
		logger:  logger,
	}
}

func (materi MateriController) GetAll(c *gin.Context) {
	datas, err := materi.service.GetAll()
	if err != nil {
		materi.logger.Zap.Error(err)
	}

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", datas)
}

func (materi MateriController) Store(c *gin.Context) {
	data := models.MateriRequest{}
	if err := c.Bind(&data); err != nil {
		materi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), data)
		return
	}

	status, err := materi.service.Store(&data)
	if err != nil {
		materi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	if !status {
		materi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Input Data Berhasil", true)
}
