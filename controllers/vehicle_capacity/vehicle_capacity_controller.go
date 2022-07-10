package controllers

import (
	"fmt"
	"infolelang/lib"
	models "infolelang/models/vehicle_capacity"

	services "infolelang/services/vehicle_capacity"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type VehicleCapacityController struct {
	logger  logger.Logger
	service services.VehicleCapacityDefinition
}

func NewVehicleCapacityController(VehicleCapacityService services.VehicleCapacityDefinition, logger logger.Logger) VehicleCapacityController {
	return VehicleCapacityController{
		service: VehicleCapacityService,
		logger:  logger,
	}
}

func (VehicleCapacity VehicleCapacityController) GetAll(c *gin.Context) {
	datas, err := VehicleCapacity.service.GetAll()
	if err != nil {
		VehicleCapacity.logger.Zap.Error(err)
	}
	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", datas)
}

func (VehicleCapacity VehicleCapacityController) GetOne(c *gin.Context) {
	paramId := c.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		VehicleCapacity.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	_, err = VehicleCapacity.service.GetOne(int64(id))
	if err != nil {
		VehicleCapacity.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", true)
}

func (VehicleCapacity VehicleCapacityController) Store(c *gin.Context) {
	data := models.VehicleCapacityRequest{}

	if err := c.Bind(&data); err != nil {
		VehicleCapacity.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}
	fmt.Println(data)
	if err := VehicleCapacity.service.Store(&data); err != nil {
		VehicleCapacity.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", true)
}

func (VehicleCapacity VehicleCapacityController) Update(c *gin.Context) {
	data := models.VehicleCapacityRequest{}
	// paramID := c.Param("id")

	if err := c.Bind(&data); err != nil {
		VehicleCapacity.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	// id, err := strconv.Atoi(paramID)
	// if err != nil {
	// 	VehicleCapacity.logger.Zap.Error(err)
	// 	lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
	// 	return
	// }

	if err := VehicleCapacity.service.Update(&data); err != nil {
		VehicleCapacity.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", data)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", data)

}

func (VehicleCapacity VehicleCapacityController) Delete(c *gin.Context) {
	paramID := c.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		VehicleCapacity.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	if err := VehicleCapacity.service.Delete(int64(id)); err != nil {
		VehicleCapacity.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Data berhasil dihapus", true)
}
