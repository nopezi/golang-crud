package controllers

import (
	"infolelang/lib"
	models "infolelang/models/sub_categories"

	services "infolelang/services/sub_category"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SubCategoryController struct {
	logger  lib.Logger
	service services.SubCategoryDefinition
}

func NewSubCategoryController(SubCategoryService services.SubCategoryDefinition, logger lib.Logger) SubCategoryController {
	return SubCategoryController{
		service: SubCategoryService,
		logger:  logger,
	}
}

func (subCategory SubCategoryController) GetAll(c *gin.Context) {
	datas, err := subCategory.service.GetAll()
	if err != nil {
		subCategory.logger.Zap.Error(err)
	}
	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", datas)
}

func (subCategory SubCategoryController) GetOne(c *gin.Context) {
	paramId := c.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		subCategory.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
	}

	data, err := subCategory.service.GetOne(int64(id))
	if err != nil {
		subCategory.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", data)
	}
	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", data)
}

func (subCategory SubCategoryController) Store(c *gin.Context) {
	data := models.SubCategoriesRequest{}
	if err := c.Bind(&data); err != nil {
		subCategory.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	if err := subCategory.service.Store(&data); err != nil {
		subCategory.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", data)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", data)
}

func (subCategory SubCategoryController) Update(c *gin.Context) {
	data := models.SubCategoriesRequest{}
	// paramID := c.Param("id")

	if err := c.Bind(&data); err != nil {
		subCategory.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	// id, err := strconv.Atoi(paramID)
	// if err != nil {
	// 	SubCategory.logger.Zap.Error(err)
	// 	lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
	// 	return
	// }

	if err := subCategory.service.Update(&data); err != nil {
		subCategory.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", data)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", data)

}

func (subCategory SubCategoryController) Delete(c *gin.Context) {
	paramID := c.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		subCategory.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	if err := subCategory.service.Delete(int64(id)); err != nil {
		subCategory.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}
	lib.ReturnToJson(c, 200, "200", "data deleted", "")
}
