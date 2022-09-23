package controllers

import (
	"fmt"
	"riskmanagement/lib"
	models "riskmanagement/models/incident"
	services "riskmanagement/services/incident"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type IncidentController struct {
	logger  logger.Logger
	service services.IncidentDefinition
}

func NewIncidentController(
	IncidentService services.IncidentDefinition,
	logger logger.Logger,
) IncidentController {
	return IncidentController{
		service: IncidentService,
		logger:  logger,
	}
}

func (incident IncidentController) GetAll(c *gin.Context) {
	datas, err := incident.service.GetAll()
	if err != nil {
		incident.logger.Zap.Error(err)
	}
	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", datas)
}

func (incident IncidentController) GetOne(c *gin.Context) {
	paramID := c.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		incident.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	data, err := incident.service.GetOne(int64(id))
	if err != nil {
		incident.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", data)
}

func (incident IncidentController) Store(c *gin.Context) {
	data := models.IncidentRequest{}

	if err := c.Bind(&data); err != nil {
		incident.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	fmt.Println(data)
	if err := incident.service.Store(&data); err != nil {
		incident.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "200", "Internal error", data)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Input data berhasil", data)
}

func (incident IncidentController) Update(c *gin.Context) {
	data := models.IncidentRequest{}

	if err := c.Bind(&data); err != nil {
		incident.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	if err := incident.service.Update(&data); err != nil {
		incident.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "200", "Internal error", data)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Update data berhasil", data)
}

func (incident IncidentController) Delete(c *gin.Context) {
	paramID := c.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		incident.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	if err := incident.service.Delete(int64(id)); err != nil {
		incident.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Data berhasil dihapus", true)
}
