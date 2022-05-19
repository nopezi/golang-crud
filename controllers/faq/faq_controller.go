package controllers

import (
	"infolelang/lib"
	models "infolelang/models/faq"

	services "infolelang/services/faq"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FaqController struct {
	logger  lib.Logger
	service services.FaqDefinition
}

func NewFaqController(FaqService services.FaqDefinition, logger lib.Logger) FaqController {
	return FaqController{
		service: FaqService,
		logger:  logger,
	}
}

func (ap FaqController) GetAll(c *gin.Context) {
	datas, err := ap.service.GetAll()
	if err != nil {
		ap.logger.Zap.Error(err)
	}
	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", datas)
}

func (ap FaqController) GetOne(c *gin.Context) {
	paramId := c.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		ap.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
	}

	data, err := ap.service.GetOne(int64(id))
	if err != nil {
		ap.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", data)
	}
	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", data)
}

func (ap FaqController) Store(c *gin.Context) {
	data := models.FaqRequest{}
	if err := c.Bind(&data); err != nil {
		ap.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	if err := ap.service.Store(&data); err != nil {
		ap.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", data)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", data)
}

func (ap FaqController) Update(c *gin.Context) {
	data := models.FaqRequest{}
	// paramID := c.Param("id")

	if err := c.Bind(&data); err != nil {
		ap.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	// id, err := strconv.Atoi(paramID)
	// if err != nil {
	// 	ap.logger.Zap.Error(err)
	// 	lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
	// 	return
	// }

	if err := ap.service.Update(&data); err != nil {
		ap.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", data)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", data)

}

func (ap FaqController) Delete(c *gin.Context) {
	paramID := c.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		ap.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	if err := ap.service.Delete(int64(id)); err != nil {
		ap.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}
	lib.ReturnToJson(c, 200, "200", "data deleted", "")
}
