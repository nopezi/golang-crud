package controllers

import (
	"fmt"
	"infolelang/lib"
	models "infolelang/models/vehicle_brand"

	services "infolelang/services/vehicle_brand"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type VehicleBrandController struct {
	logger  logger.Logger
	service services.VehicleBrandDefinition
}

func NewVehicleBrandController(VehicleBrandService services.VehicleBrandDefinition, logger logger.Logger) VehicleBrandController {
	return VehicleBrandController{
		service: VehicleBrandService,
		logger:  logger,
	}
}

func (VehicleBrand VehicleBrandController) GetAll(c *gin.Context) {
	datas, err := VehicleBrand.service.GetAll()
	if err != nil {
		VehicleBrand.logger.Zap.Error(err)
	}
	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", datas)
}

func (VehicleBrand VehicleBrandController) GetOne(c *gin.Context) {
	paramId := c.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		VehicleBrand.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	_, err = VehicleBrand.service.GetOne(int64(id))
	if err != nil {
		VehicleBrand.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", true)
}

func (VehicleBrand VehicleBrandController) Store(c *gin.Context) {
	data := models.VehicleBrandRequest{}

	if err := c.Bind(&data); err != nil {
		VehicleBrand.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}
	fmt.Println(data)
	if err := VehicleBrand.service.Store(&data); err != nil {
		VehicleBrand.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", true)
}

func (VehicleBrand VehicleBrandController) Update(c *gin.Context) {
	data := models.VehicleBrandRequest{}
	// paramID := c.Param("id")

	if err := c.Bind(&data); err != nil {
		VehicleBrand.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	// id, err := strconv.Atoi(paramID)
	// if err != nil {
	// 	VehicleBrand.logger.Zap.Error(err)
	// 	lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
	// 	return
	// }

	if err := VehicleBrand.service.Update(&data); err != nil {
		VehicleBrand.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", data)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", data)

}

func (VehicleBrand VehicleBrandController) Delete(c *gin.Context) {
	paramID := c.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		VehicleBrand.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	if err := VehicleBrand.service.Delete(int64(id)); err != nil {
		VehicleBrand.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Data berhasil dihapus", true)
}
