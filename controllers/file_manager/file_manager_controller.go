package file_manager

import (
	"infolelang/lib"
	models "infolelang/models/file_manager"
	services "infolelang/services/file_manager"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FileManagerController struct {
	logger  lib.Logger
	service services.FileManagerDefinition
}

func NewFileManagerController(
	FileManagerService services.FileManagerService,
	logger lib.Logger) FileManagerController {
	return FileManagerController{
		service: FileManagerService,
		logger:  logger,
	}
}

func (fm FileManagerController) MakeUpload(c *gin.Context) {
	request := models.FileManagerRequest{}
	if err := c.Bind(&request); err != nil {
		fm.logger.Zap.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	datas, err := fm.service.MakeUpload(request)
	if err != nil {
		fm.logger.Zap.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", datas)
}

func (fm FileManagerController) GetFile(c *gin.Context) {
	request := models.FileManagerRequest{}
	if err := c.Bind(&request); err != nil {
		fm.logger.Zap.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	datas, err := fm.service.GetFile(request)
	if err != nil {
		fm.logger.Zap.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", datas)
}

func (fm FileManagerController) RemoveObject(c *gin.Context) {
	request := models.FileManagerRequest{}
	if err := c.Bind(&request); err != nil {
		fm.logger.Zap.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	datas, err := fm.service.RemoveObject(request)
	if err != nil {
		fm.logger.Zap.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", datas)
}
