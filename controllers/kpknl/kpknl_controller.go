package controllers

import (
	"infolelang/lib"
	models "infolelang/models/kpknl"

	services "infolelang/services/kpknl"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type KpknlController struct {
	logger  logger.Logger
	service services.KpknlDefinition
}

func NewKpknlController(KpknlService services.KpknlDefinition, logger logger.Logger) KpknlController {
	return KpknlController{
		service: KpknlService,
		logger:  logger,
	}
}

func (kpknl KpknlController) GetAll(c *gin.Context) {
	datas, err := kpknl.service.GetAll()
	if err != nil {
		kpknl.logger.Zap.Error(err)
	}
	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", datas)
}

func (kpknl KpknlController) GetOne(c *gin.Context) {
	paramId := c.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		kpknl.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
	}

	data, err := kpknl.service.GetOne(int64(id))
	if err != nil {
		kpknl.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", data)
	}
	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", data)
}

func (kpknl KpknlController) Store(c *gin.Context) {
	data := models.KpknlRequest{}
	if err := c.Bind(&data); err != nil {
		kpknl.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	if err := kpknl.service.Store(&data); err != nil {
		kpknl.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", data)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", data)
}

func (kpknl KpknlController) Update(c *gin.Context) {
	data := models.KpknlRequest{}
	// paramID := c.Param("id")

	if err := c.Bind(&data); err != nil {
		kpknl.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	// id, err := strconv.Atoi(paramID)
	// if err != nil {
	// 	kpknl.logger.Zap.Error(err)
	// 	lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
	// 	return
	// }

	if err := kpknl.service.Update(&data); err != nil {
		kpknl.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", data)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", data)

}

func (kpknl KpknlController) Delete(c *gin.Context) {
	paramID := c.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		kpknl.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	if err := kpknl.service.Delete(int64(id)); err != nil {
		kpknl.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}
	lib.ReturnToJson(c, 200, "200", "data deleted", "")
}
