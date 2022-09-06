package controller

import (
	"riskmanagement/lib"
	models "riskmanagement/models/verifikasi"
	services "riskmanagement/services/verifikasi"
	"strconv"

	"github.com/gin-gonic/gin"

	"gitlab.com/golang-package-library/logger"
)

type VerifikasiController struct {
	logger  logger.Logger
	service services.VerifikasiDefinition
}

func NewVerifikasiController(
	verifikasiService services.VerifikasiDefinition,
	logger logger.Logger,
) VerifikasiController {
	return VerifikasiController{
		service: verifikasiService,
		logger:  logger,
	}
}

func (verifikasi VerifikasiController) GetAll(c *gin.Context) {
	datas, err := verifikasi.service.GetAll()

	if err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "200", "Internal Error", "")
		return
	}

	if len(datas) == 0 {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Data tidak ditemukan", "")
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", datas)
}

func (verifikasi VerifikasiController) Store(c *gin.Context) {
	data := models.VerifikasiRequest{}

	if err := c.Bind(&data); err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), data)
		return
	}

	status, err := verifikasi.service.Store(data)
	if err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	if !status {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error status", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Input data berhasil", true)
}

func (verifikasi VerifikasiController) GetOne(c *gin.Context) {
	paramID := c.Param("id")
	id, err := strconv.Atoi(paramID)

	if err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	data, status, err := verifikasi.service.GetOne(int64(id))
	if err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	if !status {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "404", "Data tidak ditemukan", nil)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inqueri data berhasil", data)

}

func (verifikasi VerifikasiController) DeleteLampiranVerifikasi(c *gin.Context) {
	data := models.VerifikasiFileRequest{}

	if err := c.Bind(&data); err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	status, err := verifikasi.service.DeleteLampiranVerifikasi(&data)
	if err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	if !status {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Data Gagal dihapus", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Delete data berhasil", true)
}

func (verifikasi VerifikasiController) Delete(c *gin.Context) {
	data := models.VerifikasiRequestUpdateMaintain{}

	if err := c.Bind(&data); err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	status, err := verifikasi.service.Delete(&data)
	if err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	if !status {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Data Gagal dihapus", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Delete data berhasil", true)
}

func (verifikasi VerifikasiController) KonfirmSave(c *gin.Context) {
	data := models.VerifikasiUpdateMaintain{}

	if err := c.Bind(&data); err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	status, err := verifikasi.service.KonfirmSave(&data)
	if err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	if !status {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Data Gagal disimpan", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Berhasil menyimpan data", true)
}

func (verifikasi VerifikasiController) UpdateAllVerifikasi(c *gin.Context) {
	data := models.VerifikasiRequestMaintain{}

	if err := c.Bind(&data); err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	status, err := verifikasi.service.UpdateAllVerifikasi(&data)
	if err != nil {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error : ", "")
		return
	}

	if !status {
		verifikasi.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Data Gagal Diupdate : ", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Update data berhasil", true)
}
