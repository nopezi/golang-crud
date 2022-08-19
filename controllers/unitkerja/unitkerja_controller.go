package controller

import (
	"fmt"
	"riskmanagement/lib"
	models "riskmanagement/models/unitkerja"
	services "riskmanagement/services/unitkerja"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type UnitKerjaController struct {
	logger  logger.Logger
	service services.UnitKerjaDefinition
}

func NewUnitKerjaController(
	UnitKerjaService services.UnitKerjaDefinition,
	logger logger.Logger,
) UnitKerjaController {
	return UnitKerjaController{
		service: UnitKerjaService,
		logger:  logger,
	}
}

func (unitKerja UnitKerjaController) GetAll(c *gin.Context) {
	datas, err := unitKerja.service.GetAll()
	if err != nil {
		unitKerja.logger.Zap.Error(err)
	}
	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", datas)
}

func (unitKerja UnitKerjaController) GetOne(c *gin.Context) {
	paramID := c.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		unitKerja.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak sesuai : "+err.Error(), "")
		return
	}

	data, err := unitKerja.service.GetOne(int64(id))
	if err != nil {
		unitKerja.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", data)
}

func (unitKerja UnitKerjaController) Store(c *gin.Context) {
	data := models.UnitKerjaRequest{}

	if err := c.Bind(&data); err != nil {
		unitKerja.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	fmt.Println(data)
	if err := unitKerja.service.Store(&data); err != nil {
		unitKerja.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Input data berhasil", true)
}

func (unitKerja UnitKerjaController) Update(c *gin.Context) {
	data := models.UnitKerjaRequest{}

	if err := c.Bind(&data); err != nil {
		unitKerja.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	if err := unitKerja.service.Update(&data); err != nil {
		unitKerja.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", data)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Update data berhasil", data)
}

func (unitKerja UnitKerjaController) Delete(c *gin.Context) {
	paramID := c.Param("id")
	id, err := strconv.Atoi(paramID)

	if err != nil {
		unitKerja.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	if err := unitKerja.service.Delete(int64(id)); err != nil {
		unitKerja.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Data berhasil dihapus", true)
}
