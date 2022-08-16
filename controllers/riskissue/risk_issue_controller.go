package controller

import (
	"fmt"
	"riskmanagement/lib"
	models "riskmanagement/models/riskissue"
	services "riskmanagement/services/riskissue"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type RiskIssueController struct {
	logger  logger.Logger
	service services.RiskIssueDefinition
}

func NewRiskIssueController(
	RiskIssueService services.RiskIssueDefinition,
	logger logger.Logger,
) RiskIssueController {
	return RiskIssueController{
		service: RiskIssueService,
		logger:  logger,
	}
}

func (riskIssue RiskIssueController) GetAll(c *gin.Context) {
	datas, err := riskIssue.service.GetAll()
	if err != nil {
		riskIssue.logger.Zap.Error(err)
	}
	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", datas)
}
func (riskIssue RiskIssueController) GetOne(c *gin.Context) {
	paramID := c.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	data, err := riskIssue.service.GetOne(int64(id))
	if err != nil {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", data)
}
func (riskIssue RiskIssueController) Store(c *gin.Context) {
	data := models.RiskIssueRequest{}

	if err := c.Bind(&data); err != nil {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	fmt.Println(data)
	if err := riskIssue.service.Store(&data); err != nil {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "200", "Internal error", data)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Input data berhasil", data)
}
func (riskIssue RiskIssueController) Update(c *gin.Context) {
	data := models.RiskIssueRequest{}

	if err := c.Bind(&data); err != nil {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	if err := riskIssue.service.Update(&data); err != nil {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "200", "Internal error", data)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Update data berhasil", data)
}
func (riskIssue RiskIssueController) Delete(c *gin.Context) {
	paramID := c.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	if err := riskIssue.service.Delete(int64(id)); err != nil {
		riskIssue.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Data berhasil dihapus", true)
}
