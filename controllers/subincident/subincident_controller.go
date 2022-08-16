package controllers

import (
	"fmt"
	"riskmanagement/lib"
	models "riskmanagement/models/subincident"
	services "riskmanagement/services/subincident"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type SubIncidentController struct {
	logger  logger.Logger
	service services.SubIncidentDefinition
}

func NewSubIncidentController(
	SubIncidentService services.SubIncidentDefinition,
	logger logger.Logger,
) SubIncidentController {
	return SubIncidentController{
		service: SubIncidentService,
		logger:  logger,
	}
}

func (subIncident SubIncidentController) GetAll(c *gin.Context) {
	datas, err := subIncident.service.GetAll()
	if err != nil {
		subIncident.logger.Zap.Error(err)
	}
	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", datas)
}

func (subIncident SubIncidentController) GetOne(c *gin.Context) {
	paramID := c.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		subIncident.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "200", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := subIncident.service.GetOne(int64(id))
	if err != nil {
		subIncident.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhsil", data)
}

func (subIncident SubIncidentController) Store(c *gin.Context) {
	data := models.SubIncidentRequest{}

	if err := c.Bind(&data); err != nil {
		subIncident.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak sesuai : "+err.Error(), "")
		return
	}

	fmt.Println(data)
	if err := subIncident.service.Store(&data); err != nil {
		subIncident.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "200", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Inputs Data Berhasil", data)
}

func (subIncident SubIncidentController) Update(c *gin.Context) {
	data := models.SubIncidentRequest{}

	if err := c.Bind(&data); err != nil {
		subIncident.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak sesuai : "+err.Error(), "")
		return
	}

	if err := subIncident.service.Update(&data); err != nil {
		subIncident.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", data)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Update Data Berhasil", data)
}

func (subIncident SubIncidentController) Delete(c *gin.Context) {
	paramID := c.Param("id")
	id, err := strconv.Atoi(paramID)

	if err != nil {
		subIncident.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidka sesuai", "")
		return
	}

	if err := subIncident.service.Delete(int64(id)); err != nil {
		subIncident.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Data berhasil dihapus", true)
}
