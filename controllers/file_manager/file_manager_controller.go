package file_manager

import (
	"infolelang/lib"
	models "infolelang/models/file_manager"
	services "infolelang/services/file_manager"

	minio "gitlab.com/golang-package-library/minio"

	"github.com/gin-gonic/gin"
)

type FileManagerController struct {
	minio   minio.Minio
	logger  lib.Logger
	service services.FileManagerDefinition
}

func NewFileManagerController(
	FileManagerService services.FileManagerDefinition,
	logger lib.Logger,
	minio minio.Minio) FileManagerController {
	return FileManagerController{
		minio:   minio,
		service: FileManagerService,
		logger:  logger,
	}
}

func (fm FileManagerController) MakeUpload(c *gin.Context) {
	request := models.FileManagerRequest{}
	file, err := c.FormFile("file")
	subdir := c.PostForm("subdir")
	request.File = file
	request.Subdir = subdir
	if err != nil {
		fm.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error : "+err.Error(), nil)
		return
	}

	// if err := c.Bind(&request); err != nil {
	// 	fm.logger.Zap.Error(err)
	// 	lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), nil)
	// 	return
	// }

	datas, err := fm.service.MakeUpload(request)
	if err != nil {
		fm.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error : "+err.Error(), datas)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", datas)
}

func (fm FileManagerController) GetFile(c *gin.Context) {
	request := models.FileManagerRequest{}
	if err := c.Bind(&request); err != nil {
		fm.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), nil)
		return
	}

	datas, err := fm.service.GetFile(request)
	if err != nil {
		fm.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error : "+err.Error(), datas)
	}

	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", datas)
}

func (fm FileManagerController) RemoveObject(c *gin.Context) {
	request := models.FileManagerRequest{}
	if err := c.Bind(&request); err != nil {
		fm.logger.Zap.Error(err)

		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), nil)
	}

	datas, err := fm.service.RemoveObject(request)
	if err != nil {
		fm.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error : "+err.Error(), datas)
	}

	if !datas {
		lib.ReturnToJson(c, 200, "200", "Remove Gagal : "+err.Error(), datas)
	}

	lib.ReturnToJson(c, 200, "200", "Remove Success", datas)
}
