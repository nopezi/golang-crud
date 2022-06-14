package controllers

import (
	"fmt"
	"infolelang/lib"
	models "infolelang/models/vehicle_color"

	services "infolelang/services/vehicle_color"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type VehicleColorController struct {
	logger  logger.Logger
	service services.VehicleColorDefinition
}

func NewVehicleColorController(VehicleColorService services.VehicleColorDefinition, logger logger.Logger) VehicleColorController {
	return VehicleColorController{
		service: VehicleColorService,
		logger:  logger,
	}
}

func (VehicleColor VehicleColorController) GetAll(c *gin.Context) {
	datas, err := VehicleColor.service.GetAll()
	if err != nil {
		VehicleColor.logger.Zap.Error(err)
	}
	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", datas)
}

func (VehicleColor VehicleColorController) GetOne(c *gin.Context) {
	paramId := c.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		VehicleColor.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	_, err = VehicleColor.service.GetOne(int64(id))
	if err != nil {
		VehicleColor.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", true)
}

func (VehicleColor VehicleColorController) Store(c *gin.Context) {
	data := models.VehicleColorRequest{}

	if err := c.Bind(&data); err != nil {
		VehicleColor.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}
	fmt.Println(data)
	if err := VehicleColor.service.Store(&data); err != nil {
		VehicleColor.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", true)
}

func (VehicleColor VehicleColorController) Update(c *gin.Context) {
	data := models.VehicleColorRequest{}
	// paramID := c.Param("id")

	if err := c.Bind(&data); err != nil {
		VehicleColor.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	// id, err := strconv.Atoi(paramID)
	// if err != nil {
	// 	VehicleColor.logger.Zap.Error(err)
	// 	lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
	// 	return
	// }

	if err := VehicleColor.service.Update(&data); err != nil {
		VehicleColor.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", data)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", data)

}

func (VehicleColor VehicleColorController) Delete(c *gin.Context) {
	paramID := c.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		VehicleColor.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	if err := VehicleColor.service.Delete(int64(id)); err != nil {
		VehicleColor.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Data berhasil dihapus", true)
}
