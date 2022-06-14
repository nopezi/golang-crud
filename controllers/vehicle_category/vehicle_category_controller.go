package controllers

import (
	"fmt"
	"infolelang/lib"
	models "infolelang/models/vehicle_category"

	services "infolelang/services/vehicle_category"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type VehicleCategoryController struct {
	logger  logger.Logger
	service services.VehicleCategoryDefinition
}

func NewVehicleCategoryController(VehicleCategoryService services.VehicleCategoryDefinition, logger logger.Logger) VehicleCategoryController {
	return VehicleCategoryController{
		service: VehicleCategoryService,
		logger:  logger,
	}
}

func (VehicleCategory VehicleCategoryController) GetAll(c *gin.Context) {
	datas, err := VehicleCategory.service.GetAll()
	if err != nil {
		VehicleCategory.logger.Zap.Error(err)
	}
	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", datas)
}

func (VehicleCategory VehicleCategoryController) GetOne(c *gin.Context) {
	paramId := c.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		VehicleCategory.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	_, err = VehicleCategory.service.GetOne(int64(id))
	if err != nil {
		VehicleCategory.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", true)
}

func (VehicleCategory VehicleCategoryController) Store(c *gin.Context) {
	data := models.VehicleCategoryRequest{}

	if err := c.Bind(&data); err != nil {
		VehicleCategory.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}
	fmt.Println(data)
	if err := VehicleCategory.service.Store(&data); err != nil {
		VehicleCategory.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", true)
}

func (VehicleCategory VehicleCategoryController) Update(c *gin.Context) {
	data := models.VehicleCategoryRequest{}
	// paramID := c.Param("id")

	if err := c.Bind(&data); err != nil {
		VehicleCategory.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	// id, err := strconv.Atoi(paramID)
	// if err != nil {
	// 	VehicleCategory.logger.Zap.Error(err)
	// 	lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
	// 	return
	// }

	if err := VehicleCategory.service.Update(&data); err != nil {
		VehicleCategory.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", data)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", data)

}

func (VehicleCategory VehicleCategoryController) Delete(c *gin.Context) {
	paramID := c.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		VehicleCategory.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	if err := VehicleCategory.service.Delete(int64(id)); err != nil {
		VehicleCategory.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Data berhasil dihapus", true)
}
