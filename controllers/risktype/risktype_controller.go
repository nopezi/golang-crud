package controller

import (
	"fmt"
	"riskmanagement/lib"
	models "riskmanagement/models/risktype"
	services "riskmanagement/services/risktype"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type RiskTypeController struct {
	logger  logger.Logger
	service services.RiskTypeDefinition
}

func NewRiskTypeController(
	RiskTypeService services.RiskTypeDefinition,
	logger logger.Logger,
) RiskTypeController {
	return RiskTypeController{
		service: RiskTypeService,
		logger:  logger,
	}
}

func (riskType RiskTypeController) GetAll(c *gin.Context) {
	datas, err := riskType.service.GetAll()
	if err != nil {
		riskType.logger.Zap.Error(err)
	}
	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", datas)
}
func (riskType RiskTypeController) GetOne(c *gin.Context) {
	paramID := c.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		riskType.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	data, err := riskType.service.GetOne(int64(id))
	if err != nil {
		riskType.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", data)
}
func (riskType RiskTypeController) Store(c *gin.Context) {
	data := models.RiskTypeRequest{}

	if err := c.Bind(&data); err != nil {
		riskType.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	fmt.Println(data)
	if err := riskType.service.Store(&data); err != nil {
		riskType.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "200", "Internal error", data)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Input data berhasil", data)
}
func (riskType RiskTypeController) Update(c *gin.Context) {
	data := models.RiskTypeRequest{}

	if err := c.Bind(&data); err != nil {
		riskType.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	if err := riskType.service.Update(&data); err != nil {
		riskType.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "200", "Internal error", data)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Update data berhasil", data)
}
func (riskType RiskTypeController) Delete(c *gin.Context) {
	paramID := c.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		riskType.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	if err := riskType.service.Delete(int64(id)); err != nil {
		riskType.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Data berhasil dihapus", true)
}
