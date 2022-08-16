package controllers

import (
	"fmt"
	"riskmanagement/lib"
	models "riskmanagement/models/subactivity"
	services "riskmanagement/services/subactivity"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type SubActivityController struct {
	logger  logger.Logger
	service services.SubAvtivityDefinition
}

func NewSubActivityController(SubActivityService services.SubAvtivityDefinition, logger logger.Logger) SubActivityController {
	return SubActivityController{
		service: SubActivityService,
		logger:  logger,
	}
}

func (subactivity SubActivityController) GetAll(c *gin.Context) {
	datas, err := subactivity.service.GetAll()
	if err != nil {
		subactivity.logger.Zap.Error(err)
	}
	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", datas)
}

func (subactivity SubActivityController) GetOne(c *gin.Context) {
	paramID := c.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		subactivity.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "200", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := subactivity.service.GetOne(int64(id))
	if err != nil {
		subactivity.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhsil", data)
}

func (subactivity SubActivityController) Store(c *gin.Context) {
	data := models.SubActivityRequest{}

	if err := c.Bind(&data); err != nil {
		subactivity.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak sesuai : "+err.Error(), "")
		return
	}

	fmt.Println(data)
	if err := subactivity.service.Store(&data); err != nil {
		subactivity.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "200", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Inputs Data Berhasil", data)
}

func (subactivity SubActivityController) Update(c *gin.Context) {
	data := models.SubActivityRequest{}

	if err := c.Bind(&data); err != nil {
		subactivity.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak sesuai : "+err.Error(), "")
		return
	}

	if err := subactivity.service.Update(&data); err != nil {
		subactivity.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", data)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Update Data Berhasil", data)
}

func (subactivity SubActivityController) Delete(c *gin.Context) {
	paramID := c.Param("id")
	id, err := strconv.Atoi(paramID)

	if err != nil {
		subactivity.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidka sesuai", "")
		return
	}

	if err := subactivity.service.Delete(int64(id)); err != nil {
		subactivity.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Data berhasil dihapus", true)
}
