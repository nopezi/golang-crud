package controllers

import (
	"fmt"
	"infolelang/lib"
	models "infolelang/models/vehicle_transmission"

	services "infolelang/services/vehicle_transmission"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type VehicleTransmissionController struct {
	logger  logger.Logger
	service services.VehicleTransmissionDefinition
}

func NewVehicleTransmissionController(VehicleTransmissionService services.VehicleTransmissionDefinition, logger logger.Logger) VehicleTransmissionController {
	return VehicleTransmissionController{
		service: VehicleTransmissionService,
		logger:  logger,
	}
}

func (VehicleTransmission VehicleTransmissionController) GetAll(c *gin.Context) {
	datas, err := VehicleTransmission.service.GetAll()
	if err != nil {
		VehicleTransmission.logger.Zap.Error(err)
	}
	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", datas)
}

func (VehicleTransmission VehicleTransmissionController) GetOne(c *gin.Context) {
	paramId := c.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		VehicleTransmission.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	_, err = VehicleTransmission.service.GetOne(int64(id))
	if err != nil {
		VehicleTransmission.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", true)
}

func (VehicleTransmission VehicleTransmissionController) Store(c *gin.Context) {
	data := models.VehicleTransmissionRequest{}

	if err := c.Bind(&data); err != nil {
		VehicleTransmission.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}
	fmt.Println(data)
	if err := VehicleTransmission.service.Store(&data); err != nil {
		VehicleTransmission.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", true)
}

func (VehicleTransmission VehicleTransmissionController) Update(c *gin.Context) {
	data := models.VehicleTransmissionRequest{}
	// paramID := c.Param("id")

	if err := c.Bind(&data); err != nil {
		VehicleTransmission.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	// id, err := strconv.Atoi(paramID)
	// if err != nil {
	// 	VehicleTransmission.logger.Zap.Error(err)
	// 	lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
	// 	return
	// }

	if err := VehicleTransmission.service.Update(&data); err != nil {
		VehicleTransmission.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", data)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", data)

}

func (VehicleTransmission VehicleTransmissionController) Delete(c *gin.Context) {
	paramID := c.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		VehicleTransmission.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	if err := VehicleTransmission.service.Delete(int64(id)); err != nil {
		VehicleTransmission.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Data berhasil dihapus", true)
}
