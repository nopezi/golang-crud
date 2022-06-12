package facility

import (
	"fmt"
	"infolelang/lib"
	models "infolelang/models/facilities"

	services "infolelang/services/facility"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type FacilityController struct {
	logger  logger.Logger
	service services.FacilitiesDefinition
}

func NewFacilityController(FacilityService services.FacilitiesDefinition, logger logger.Logger) FacilityController {
	return FacilityController{
		service: FacilityService,
		logger:  logger,
	}
}

func (Facility FacilityController) GetAll(c *gin.Context) {
	datas, err := Facility.service.GetAll()
	if err != nil {
		Facility.logger.Zap.Error(err)
	}
	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", datas)
}

func (Facility FacilityController) GetOne(c *gin.Context) {
	paramId := c.Param("id")
	fmt.Println(paramId)
	id, err := strconv.Atoi(paramId)
	if err != nil {
		Facility.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	data, err := Facility.service.GetOne(int64(id))
	if err != nil {
		Facility.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", data)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", data)
}

func (Facility FacilityController) Store(c *gin.Context) {
	data := models.FacilitiesRequest{}
	if err := c.Bind(&data); err != nil {
		Facility.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	status, err := Facility.service.Store(&data)

	if err != nil {
		Facility.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", status)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Data berhasil disimpan", status)
}

func (Facility FacilityController) Update(c *gin.Context) {
	data := models.FacilitiesRequest{}

	if err := c.Bind(&data); err != nil {
		Facility.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	_, err := Facility.service.Update(&data)
	if err != nil {
		Facility.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Data berhasil diupdate", true)

}

func (Facility FacilityController) Delete(c *gin.Context) {
	paramID := c.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		Facility.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	if err := Facility.service.Delete(int64(id)); err != nil {
		Facility.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Data berhasil dihapus", true)
}
