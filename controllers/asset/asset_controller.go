package controllers

import (
	"database/sql"
	"fmt"
	"infolelang/lib"
	models "infolelang/models/assets"

	services "infolelang/services/asset"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
	// minio "gitlab.com/golang-package-library/minio"
)

type AssetController struct {
	// minio   minio.Minio
	logger  logger.Logger
	service services.AssetDefinition
}

func NewAssetController(
	AssetService services.AssetDefinition,
	logger logger.Logger,
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
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	if len(datas) == 0 {
		asset.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "404", "Data Tidak Ditemukan", "")
		return
	}
	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", datas)
}

func (asset AssetController) GetOne(c *gin.Context) {
	paramId := c.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		asset.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	data, err := asset.service.GetOne(int64(id))
	if err != nil {
		asset.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", data)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", data)
}

func (asset AssetController) Store(c *gin.Context) {
	data := models.AssetsRequest{}
	// trxHandle := c.MustGet(constants.DBTransaction).(*gorm.DB)

	if err := c.Bind(&data); err != nil {
		asset.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	status, err := asset.service.Store(&data)
	if err != nil {
		asset.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	if !status {
		asset.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Create data berhasil", true)
}

func (asset AssetController) UpdatePublish(c *gin.Context) {
	data := models.AssetsRequestUpdate{}

	if err := c.BindJSON(&data); err != nil {
		asset.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	status, err := asset.service.UpdatePublish(&data)
	if err != nil {
		asset.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}

	if !status {
		asset.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Update data Gagal", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Update data berhasil", status)
}

func (asset AssetController) UpdateApproval(c *gin.Context) {
	data := models.AssetsRequestUpdate{}

	if err := c.Bind(&data); err != nil {
		asset.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	status, err := asset.service.UpdateApproval(&data)
	if err != nil {
		asset.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}

	if !status {
		asset.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Update data Gagal", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Update data berhasil", status)
}

func (asset AssetController) UpdateMaintain(c *gin.Context) {
	data := models.AssetsRequest{}

	if err := c.BindJSON(&data); err != nil {
		asset.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	status, err := asset.service.UpdateMaintain(&data)
	if err != nil {
		asset.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", data)
		return
	}

	if !status {
		asset.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Update data Gagal", data)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Update data berhasil", data)
}

func (asset AssetController) Delete(c *gin.Context) {
	data := models.AssetsRequestUpdate{}

	if err := c.Bind(&data); err != nil {
		asset.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	status, err := asset.service.Delete(&data)
	if err != nil {
		asset.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	if !status {
		asset.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Data Gagal disimpan", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Update data berhasil", true)
}

func (asset AssetController) GetApproval(c *gin.Context) {
	request := models.AssetsRequestMaintain{}
	if err := c.Bind(&request); err != nil {
		asset.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	datas, pagination, err := asset.service.GetApproval(request)
	if err != nil {
		asset.logger.Zap.Error(err)
	}

	if pagination.Total == 0 {
		lib.ReturnToJson(c, 200, "404", "Data Kosong", datas)
		return
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", datas)
		return
	}
	fmt.Println("Data Approvals controller=>", datas)
	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquiry data berhasil", datas, pagination)
}

func (asset AssetController) GetMaintain(c *gin.Context) {
	request := models.AssetsRequestMaintain{}

	if err := c.Bind(&request); err != nil {
		asset.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	datas, pagination, err := asset.service.GetMaintain(request)
	if err != nil {
		asset.logger.Zap.Error(err)
	}

	if pagination.Total == 0 {
		lib.ReturnToJson(c, 200, "404", "Data Kosong", datas)
		return
	}

	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquiry data berhasil", datas, pagination)
}
