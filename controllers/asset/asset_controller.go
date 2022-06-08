package controllers

import (
	"infolelang/lib"
	models "infolelang/models/assets"

	services "infolelang/services/asset"
	"strconv"

	"github.com/gin-gonic/gin"
	// "gitlab.com/golang-package-library/minio"
	// minio "gitlab.com/golang-package-library/minio"
)

type AssetController struct {
	// minio   minio.Minio
	logger  lib.Logger
	service services.AssetDefinition
}

func NewAssetController(
	AssetService services.AssetDefinition,
	logger lib.Logger,
	// minio minio.Minio,
) AssetController {
	return AssetController{
		service: AssetService,
		logger:  logger,
		// minio:   minio,
	}
}

func (asset AssetController) GetAll(c *gin.Context) {
	datas, err := asset.service.GetAll()
	if err != nil {
		asset.logger.Zap.Error(err)
	}
	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", datas)
}

func (asset AssetController) GetOne(c *gin.Context) {
	paramId := c.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		asset.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
	}

	data, err := asset.service.GetOne(int64(id))
	if err != nil {
		asset.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", data)
	}
	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", data)
}

func (asset AssetController) Store(c *gin.Context) {
	data := models.AssetsRequest{}
	if err := c.Bind(&data); err != nil {
		asset.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	if err := asset.service.Store(&data); err != nil {
		asset.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}
	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", data)
}

func (asset AssetController) Update(c *gin.Context) {
	data := models.AssetsRequest{}
	// paramID := c.Param("id")

	if err := c.BindJSON(&data); err != nil {
		asset.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	// id, err := strconv.Atoi(paramID)
	// if err != nil {
	// 	Asset.logger.Zap.Error(err)
	// 	lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
	// 	return
	// }

	if err := asset.service.Update(&data); err != nil {
		asset.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", data)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", data)

}

func (asset AssetController) Delete(c *gin.Context) {
	paramID := c.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		asset.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	if err := asset.service.Delete(int64(id)); err != nil {
		asset.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}
	lib.ReturnToJson(c, 200, "200", "data deleted", "")
}
