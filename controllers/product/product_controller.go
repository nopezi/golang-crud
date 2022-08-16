package controller

import (
	"fmt"
	"riskmanagement/lib"
	models "riskmanagement/models/product"
	services "riskmanagement/services/product"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type ProductController struct {
	logger  logger.Logger
	service services.ProductDefinition
}

func NewProductController(
	ProductService services.ProductDefinition,
	logger logger.Logger,
) ProductController {
	return ProductController{
		service: ProductService,
		logger:  logger,
	}
}

func (product ProductController) GetAll(c *gin.Context) {
	datas, err := product.service.GetAll()
	if err != nil {
		product.logger.Zap.Error(err)
	}
	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", datas)
}

func (product ProductController) GetOne(c *gin.Context) {
	paramID := c.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		product.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak sesuai : "+err.Error(), "")
		return
	}

	data, err := product.service.GetOne(int64(id))
	if err != nil {
		product.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", data)
}

func (product ProductController) Store(c *gin.Context) {
	data := models.ProductRequest{}

	if err := c.Bind(&data); err != nil {
		product.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	fmt.Println(data)
	if err := product.service.Store(&data); err != nil {
		product.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Input data berhasil", true)
}

func (product ProductController) Update(c *gin.Context) {
	data := models.ProductRequest{}

	if err := c.Bind(&data); err != nil {
		product.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	if err := product.service.Update(&data); err != nil {
		product.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", data)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Update data berhasil", data)
}

func (product ProductController) Delete(c *gin.Context) {
	paramID := c.Param("id")
	id, err := strconv.Atoi(paramID)

	if err != nil {
		product.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	if err := product.service.Delete(int64(id)); err != nil {
		product.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Data berhasil dihapus", true)
}
