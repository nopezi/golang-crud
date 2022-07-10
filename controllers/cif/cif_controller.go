package cif

import (
	"infolelang/lib"

	services "infolelang/services/cif"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type CifController struct {
	logger  logger.Logger
	service services.CifDefinition
}

func NewCifController(CifService services.CifDefinition, logger logger.Logger) CifController {
	return CifController{
		service: CifService,
		logger:  logger,
	}
}

func (cif CifController) InquiryCif(c *gin.Context) {
	paramId := c.Param("id")
	// id, err := strconv.Atoi(paramId)
	// if err != nil {
	// 	cif.logger.Zap.Error(err)
	// 	lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
	// 	return
	// }

	data, err := cif.service.InquiryCif(paramId)
	if err != nil {
		cif.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", data)
}
