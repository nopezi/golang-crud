package controllers

import (
	"fmt"
	"riskmanagement/lib"
	models "riskmanagement/models/activity"
	services "riskmanagement/services/activity"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type ActivityController struct {
	logger  logger.Logger
	service services.ActivityDefinition
}

func NewActivityController(ActivityService services.ActivityDefinition, logger logger.Logger) ActivityController {
	return ActivityController{
		service: ActivityService,
		logger:  logger,
	}
}

func (activity ActivityController) GetAll(c *gin.Context) {
	datas, err := activity.service.GetAll()
	if err != nil {
		activity.logger.Zap.Error(err)
	}
	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", datas)
}

func (activity ActivityController) GetOne(c *gin.Context) {
	paramID := c.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		activity.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := activity.service.GetOne(int64(id))
	if err != nil {
		activity.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (activity ActivityController) Store(c *gin.Context) {
	data := models.ActivityRequest{}

	if err := c.Bind(&data); err != nil {
		activity.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}
	fmt.Println(data)
	if err := activity.service.Store(&data); err != nil {
		activity.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", true)
}

func (activity ActivityController) Update(c *gin.Context) {
	data := models.ActivityRequest{}

	if err := c.Bind(&data); err != nil {
		activity.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	if err := activity.service.Update(&data); err != nil {
		activity.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", data)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", data)
}

func (activity ActivityController) Delete(c *gin.Context) {
	paramID := c.Param("id")
	id, err := strconv.Atoi(paramID)

	if err != nil {
		activity.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	if err := activity.service.Delete(int64(id)); err != nil {
		activity.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Data berhasil dihapus", true)
}
