package controllers

import (
	"infolelang/lib"
	models "infolelang/models/banner"

	services "infolelang/services/banner"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type BannerController struct {
	logger  logger.Logger
	service services.BannerDefinition
}

func NewBannerController(BannerService services.BannerDefinition, logger logger.Logger) BannerController {
	return BannerController{
		service: BannerService,
		logger:  logger,
	}
}

func (banner BannerController) GetAll(c *gin.Context) {
	datas, err := banner.service.GetAll()
	if err != nil {
		banner.logger.Zap.Error(err)
	}
	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", datas)
}

func (banner BannerController) Store(c *gin.Context) {
	data := models.BannerRequest{}

	if err := c.Bind(&data); err != nil {
		banner.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	status, err := banner.service.Store(&data)
	if err != nil || !status {
		banner.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", true)
}

func (banner BannerController) Delete(c *gin.Context) {
	data := models.BannerImageRequest{}

	if err := c.Bind(&data); err != nil {
		banner.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	status, err := banner.service.Delete(data)
	if err != nil || !status {
		banner.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Data berhasil dihapus", true)
}
