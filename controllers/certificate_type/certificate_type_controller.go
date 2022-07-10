package CertificateType

import (
	"fmt"
	"infolelang/lib"
	models "infolelang/models/certificate_type"

	services "infolelang/services/certificate_type"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type CertificateTypeController struct {
	logger  logger.Logger
	service services.CertificateTypeDefinition
}

func NewCertificateTypeController(CertificateTypeService services.CertificateTypeDefinition, logger logger.Logger) CertificateTypeController {
	return CertificateTypeController{
		service: CertificateTypeService,
		logger:  logger,
	}
}

func (CertificateType CertificateTypeController) GetAll(c *gin.Context) {
	datas, err := CertificateType.service.GetAll()
	if err != nil {
		CertificateType.logger.Zap.Error(err)
	}
	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", datas)
}

func (CertificateType CertificateTypeController) GetOne(c *gin.Context) {
	paramId := c.Param("id")
	fmt.Println(paramId)
	id, err := strconv.Atoi(paramId)
	if err != nil {
		CertificateType.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	data, err := CertificateType.service.GetOne(int64(id))
	if err != nil {
		CertificateType.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", data)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", data)
}

func (CertificateType CertificateTypeController) Store(c *gin.Context) {
	data := models.CertificateTypeRequest{}
	if err := c.Bind(&data); err != nil {
		CertificateType.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	status, err := CertificateType.service.Store(&data)

	if err != nil {
		CertificateType.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", status)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Data berhasil disimpan", status)
}

func (CertificateType CertificateTypeController) Update(c *gin.Context) {
	data := models.CertificateTypeRequest{}

	if err := c.Bind(&data); err != nil {
		CertificateType.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	_, err := CertificateType.service.Update(&data)
	if err != nil {
		CertificateType.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Data berhasil diupdate", true)

}

func (CertificateType CertificateTypeController) Delete(c *gin.Context) {
	paramID := c.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		CertificateType.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	if err := CertificateType.service.Delete(int64(id)); err != nil {
		CertificateType.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Data berhasil dihapus", true)
}
