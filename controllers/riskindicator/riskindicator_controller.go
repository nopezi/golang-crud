package controllers

import (
	"fmt"
	"riskmanagement/lib"
	models "riskmanagement/models/riskindicator"
	services "riskmanagement/services/riskindicator"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type RiskIndicatorController struct {
	logger  logger.Logger
	service services.RiskIndicatorDefinition
}

func NewRiskIndicatorController(
	RiskIndicatorService services.RiskIndicatorDefinition,
	logger logger.Logger,
) RiskIndicatorController {
	return RiskIndicatorController{
		service: RiskIndicatorService,
		logger:  logger,
	}
}

func (riskIndicator RiskIndicatorController) GetAll(c *gin.Context) {
	datas, err := riskIndicator.service.GetAll()
	if err != nil {
		riskIndicator.logger.Zap.Error(err)
	}
	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", datas)
}

func (riskIndicator RiskIndicatorController) GetOne(c *gin.Context) {
	paramID := c.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		riskIndicator.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	data, err := riskIndicator.service.GetOne(int64(id))
	if err != nil {
		riskIndicator.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", data)
}
func (riskIndicator RiskIndicatorController) Store(c *gin.Context) {
	data := models.RiskIndicatorRequest{}

	if err := c.Bind(&data); err != nil {
		riskIndicator.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	fmt.Println(data)
	if err := riskIndicator.service.Store(&data); err != nil {
		riskIndicator.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "200", "Internal error", data)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Input data berhasil", data)
}
func (riskIndicator RiskIndicatorController) Update(c *gin.Context) {
	data := models.RiskIndicatorRequest{}

	if err := c.Bind(&data); err != nil {
		riskIndicator.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	if err := riskIndicator.service.Update(&data); err != nil {
		riskIndicator.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "200", "Internal error", data)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Update data berhasil", data)
}
func (riskIndicator RiskIndicatorController) Delete(c *gin.Context) {
	paramID := c.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		riskIndicator.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	if err := riskIndicator.service.Delete(int64(id)); err != nil {
		riskIndicator.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Data berhasil dihapus", true)
}
