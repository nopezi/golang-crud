package controllers

import (
	"infolelang/lib"
	models "infolelang/models/mcs"

	services "infolelang/services/mcs"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type McsController struct {
	logger  logger.Logger
	service services.McsDefinition
}

func NewMcsController(McsService services.McsDefinition, logger logger.Logger) McsController {
	return McsController{
		service: McsService,
		logger:  logger,
	}
}

func (Mcs McsController) GetMcs(c *gin.Context) {
	data := models.McsRequest{}

	if err := c.Bind(&data); err != nil {
		Mcs.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}
	response, err := Mcs.service.GetMcs(&data)
	if err != nil {
		Mcs.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", response)
}
