package controller

import (
	"fmt"
	"riskmanagement/lib"
	models "riskmanagement/models/riskcontrol"
	services "riskmanagement/services/riskcontrol"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type RiskControlController struct {
	logger  logger.Logger
	service services.RiskControlDefinition
}

func NewRiskControlController(
	RiskControlService services.RiskControlDefinition,
	logger logger.Logger,
) RiskControlController {
	return RiskControlController{
		logger:  logger,
		service: RiskControlService,
	}
}

func (riskControl RiskControlController) GetAll(c *gin.Context) {
	datas, err := riskControl.service.GetAll()
	if err != nil {
		riskControl.logger.Zap.Error(err)
	}
	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", datas)
}
func (riskControl RiskControlController) GetOne(c *gin.Context) {
	paramID := c.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		riskControl.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	data, err := riskControl.service.GetOne(int64(id))
	if err != nil {
		riskControl.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", data)
}
func (riskControl RiskControlController) Store(c *gin.Context) {
	data := models.RiskControlRequest{}

	if err := c.Bind(&data); err != nil {
		riskControl.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	fmt.Println(data)
	if err := riskControl.service.Store(&data); err != nil {
		riskControl.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "200", "Internal error", data)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Input data berhasil", data)
}
func (riskControl RiskControlController) Update(c *gin.Context) {
	data := models.RiskControlRequest{}

	if err := c.Bind(&data); err != nil {
		riskControl.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	if err := riskControl.service.Update(&data); err != nil {
		riskControl.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "200", "Internal error", data)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Update data berhasil", data)
}
func (riskControl RiskControlController) Delete(c *gin.Context) {
	paramID := c.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		riskControl.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	if err := riskControl.service.Delete(int64(id)); err != nil {
		riskControl.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Data berhasil dihapus", true)
}
