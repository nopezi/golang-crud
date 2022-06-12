package controllers

import (
	"fmt"
	"infolelang/lib"
	models "infolelang/models/categories"

	services "infolelang/services/category"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type CategoryController struct {
	logger  logger.Logger
	service services.CategoryDefinition
}

func NewCategoryController(CategoryService services.CategoryDefinition, logger logger.Logger) CategoryController {
	return CategoryController{
		service: CategoryService,
		logger:  logger,
	}
}

func (category CategoryController) GetAll(c *gin.Context) {
	datas, err := category.service.GetAll()
	if err != nil {
		category.logger.Zap.Error(err)
	}
	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", datas)
}

func (category CategoryController) GetOne(c *gin.Context) {
	paramId := c.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		category.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	_, err = category.service.GetOne(int64(id))
	if err != nil {
		category.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", true)
}

func (category CategoryController) Store(c *gin.Context) {
	data := models.CategoryRequest{}

	if err := c.Bind(&data); err != nil {
		category.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}
	fmt.Println(data)
	if err := category.service.Store(&data); err != nil {
		category.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", true)
}

func (category CategoryController) Update(c *gin.Context) {
	data := models.CategoryRequest{}
	// paramID := c.Param("id")

	if err := c.Bind(&data); err != nil {
		category.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	// id, err := strconv.Atoi(paramID)
	// if err != nil {
	// 	Category.logger.Zap.Error(err)
	// 	lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
	// 	return
	// }

	if err := category.service.Update(&data); err != nil {
		category.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", data)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", data)

}

func (category CategoryController) Delete(c *gin.Context) {
	paramID := c.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		category.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	if err := category.service.Delete(int64(id)); err != nil {
		category.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Data berhasil dihapus", true)
}
